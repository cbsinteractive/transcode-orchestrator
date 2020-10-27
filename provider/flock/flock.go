package flock

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/cbsinteractive/transcode-orchestrator/config"
	"github.com/cbsinteractive/transcode-orchestrator/db"
	"github.com/cbsinteractive/transcode-orchestrator/provider"
)

const (
	// Name identifies the Flock provider by name
	Name = "flock"
)

func init() {
	err := provider.Register(Name, flockFactory)
	if err != nil {
		fmt.Printf("registering flock factory: %v", err)
	}
}

type flock struct {
	cfg        *config.Flock
	repository db.Repository
	client     *http.Client
}

type JobRequest struct {
	Job JobSpec `json:"job"`
}

type JobSpec struct {
	Source  string      `json:"source"`
	Outputs []JobOutput `json:"outputs"`
	Labels  []string    `json:"labels,omitempty"`
}

type JobOutput struct {
	AudioChannels      int    `json:"audio_channels"`
	AudioBitrateKbps   int    `json:"audio_bitrate_kbps"`
	VideoBitrateKbps   int    `json:"video_bitrate_kbps"`
	Width              int    `json:"width,omitempty"`
	Height             int    `json:"height,omitempty"`
	KeyframesPerSecond int    `json:"keyframes_sec,omitempty"`
	KeyframeInterval   int    `json:"keyframe_interval,omitempty"`
	MultiPass          bool   `json:"multipass"`
	Preset             string `json:"preset,omitempty"`
	Tag                string `json:"tag,omitempty"`
	Destination        string `json:"destination"`
	AudioCodec         string `json:"acodec,omitempty"`
	VideoCodec         string `json:"vcodec,omitempty"`
}

type NewJobResponse struct {
	JobID  int    `json:"job_id"`
	Status string `json:"status"`
}

type JobResponse struct {
	CreateTime      string              `json:"create_time"`
	CreateTimestamp float64             `json:"create_timestamp"`
	Duration        string              `json:"duration"`
	ID              int                 `json:"id"`
	Progress        float64             `json:"progress_pct"`
	Status          string              `json:"status"`
	UpdateTime      string              `json:"update_time"`
	UpdateTimestamp float64             `json:"update_timestamp"`
	Outputs         []JobResponseOutput `json:"outputs"`
}

type JobResponseOutput struct {
	CreateTime      string  `json:"create_time"`
	CreateTimestamp float64 `json:"create_timestamp"`
	Destination     string  `json:"destination"`
	Duration        string  `json:"duration"`
	Encoder         string  `json:"encoder"`
	ID              int     `json:"id"`
	Progress        float64 `json:"progress_pct"`
	Request         map[string]interface{}
	Info            string  `json:"info"`
	Status          string  `json:"status"`
	UpdateTime      string  `json:"update_time"`
	UpdateTimestamp float64 `json:"update_timestamp"`
}

func (p *flock) Transcode(ctx context.Context, job *db.Job) (*provider.JobStatus, error) {
	jobReq, err := p.flockJobRequestFrom(ctx, job)
	if err != nil {
		return nil, fmt.Errorf("generating flock job request: %w", err)
	}

	jsonValue, err := json.Marshal(jobReq)
	if err != nil {
		return nil, fmt.Errorf("marshaling job request %+v to json: %w", jobReq, err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		fmt.Sprintf("%s/api/v1/jobs", p.cfg.Endpoint), bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, fmt.Errorf("creating job request: %w", err)
	}
	req.Header.Set("Authorization", p.cfg.Credential)
	req.Header.Set("Content-Type", "application/json")

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("submitting new job: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %w", err)
	}

	var newJob NewJobResponse
	err = json.Unmarshal(body, &newJob)
	if err != nil {
		return nil, fmt.Errorf("parsing flock response: %w", err)
	}

	return &provider.JobStatus{
		ProviderName:  Name,
		ProviderJobID: fmt.Sprintf("%d", newJob.JobID),
		Status:        provider.StatusQueued,
	}, nil
}

func (p *flock) flockJobRequestFrom(ctx context.Context, job *db.Job) (*JobRequest, error) {
	presets := []db.Preset{}
	for _, output := range job.Outputs {
		presetName := output.Preset.Name
		presetResponse, err := p.GetPreset(ctx, presetName)
		if err != nil {
			return nil, err
		}

		localPreset, ok := presetResponse.(*db.LocalPreset)
		if !ok {
			return nil, fmt.Errorf("could not convert preset response into a db.LocalPreset")
		}

		presets = append(presets, localPreset.Preset)
	}

	var jobReq JobRequest
	jobReq.Job.Source = job.SourceMedia

	for _, label := range job.Labels {
		jobReq.Job.Labels = append(jobReq.Job.Labels, label)
	}

	jobOuts := make([]JobOutput, 0, len(job.Outputs))
	for i, output := range job.Outputs {
		var jobOut JobOutput
		jobOut.Destination = joinBaseAndParts(job.DestinationBasePath, job.RootFolder(), output.FileName)
		jobOut.AudioCodec = presets[i].Audio.Codec
		jobOut.VideoCodec = presets[i].Video.Codec
		jobOut.Preset = "slow"
		jobOut.MultiPass = presets[i].TwoPass

		if jobOut.VideoCodec == "hevc" {
			jobOut.Tag = "hvc1"
		}
		if presets[i].Video.GopUnit == db.GopUnitSeconds {
			jobOut.KeyframesPerSecond = int(presets[i].Video.GopSize)
		} else {
			jobOut.KeyframeInterval = int(presets[i].Video.GopSize)
		}
		if br := presets[i].Video.Bitrate; br > 0 {
			jobOut.VideoBitrateKbps = br / 1000
		}
		if br := presets[i].Audio.Bitrate; br > 0 {
			jobOut.AudioBitrateKbps = br / 1000
			jobOut.AudioChannels = 2
		}
		if w := presets[i].Video.Width; w > 0 {
			jobOut.Width = presets[i].Video.Width
		}
		if h := presets[i].Video.Height; h > 0 {
			jobOut.Height = h
		}
		jobOuts = append(jobOuts, jobOut)
	}

	jobReq.Job.Outputs = jobOuts
	return &jobReq, nil
}

func (p *flock) JobStatus(ctx context.Context, job *db.Job) (*provider.JobStatus, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		fmt.Sprintf("%s/api/v1/jobs/%s", p.cfg.Endpoint, job.ProviderJobID), nil)
	if err != nil {
		return nil, fmt.Errorf("creating status request: %w", err)
	}
	req.Header.Set("Authorization", p.cfg.Credential)

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("querying for provider job %s: %w", job.ProviderJobID, err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %w", err)
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("job not found for provider id %s, response: %s",
			job.ProviderJobID, string(body))
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("querying for provider job %s, status: %d response: %s",
			job.ProviderJobID, resp.StatusCode, string(body))
	}

	var jobResp JobResponse
	err = json.Unmarshal(body, &jobResp)
	if err != nil {
		return nil, fmt.Errorf("parsing flock response: %w", err)
	}

	return p.jobStatusFrom(job, &jobResp), nil
}

func (p *flock) jobStatusFrom(job *db.Job, jobResp *JobResponse) *provider.JobStatus {
	status := &provider.JobStatus{
		ProviderJobID: job.ProviderJobID,
		ProviderName:  Name,
		ProviderStatus: map[string]interface{}{
			"create_timestamp": jobResp.CreateTimestamp,
			"update_timestamp": jobResp.UpdateTimestamp,
			"duration":         jobResp.Duration,
			"progress":         jobResp.Progress,
			"status":           jobResp.Status,
		},
		Status: statusFrom(jobResp),
		Output: provider.JobOutput{
			Destination: joinBaseAndParts(job.DestinationBasePath, job.RootFolder()),
		},
		Labels: job.Labels,
	}

	outputsStatus := make([]map[string]interface{}, 0, len(jobResp.Outputs))
	outputFiles := make([]provider.OutputFile, 0, len(jobResp.Outputs))

	for _, output := range jobResp.Outputs {
		outputFiles = append(outputFiles, provider.OutputFile{Path: output.Destination})
		outputsStatus = append(outputsStatus, map[string]interface{}{
			"duration":         output.Duration,
			"destination":      output.Destination,
			"encoder":          output.Encoder,
			"info":             output.Info,
			"status":           output.Status,
			"progress":         output.Progress,
			"update_timestamp": output.UpdateTimestamp,
		})
	}

	status.Output.Files = outputFiles
	status.ProviderStatus["outputs"] = outputsStatus
	status.Progress = jobResp.Progress
	return status
}

func joinBaseAndParts(base string, elem ...string) string {
	parts := []string{strings.TrimRight(base, "/")}
	parts = append(parts, elem...)
	return strings.Join(parts, "/")
}

func statusFrom(job *JobResponse) provider.Status {
	switch job.Status {
	case "submitted":
		return provider.StatusQueued
	case "assigned", "transcoding":
		return provider.StatusStarted
	case "complete":
		return provider.StatusFinished
	case "cancelled":
		return provider.StatusCanceled
	case "error":
		return provider.StatusFailed
	}
	return provider.StatusUnknown
}

func (p *flock) CreatePreset(_ context.Context, preset db.Preset) (string, error) {
	err := p.repository.CreateLocalPreset(&db.LocalPreset{
		Name:   preset.Name,
		Preset: preset,
	})
	if err != nil {
		return "", err
	}

	return preset.Name, nil
}

func (p *flock) GetPreset(_ context.Context, presetID string) (interface{}, error) {
	return p.repository.GetLocalPreset(presetID)
}

func (p *flock) DeletePreset(ctx context.Context, presetID string) error {
	preset, err := p.GetPreset(ctx, presetID)
	if err != nil {
		return err
	}

	return p.repository.DeleteLocalPreset(preset.(*db.LocalPreset))
}

func (p *flock) CancelJob(ctx context.Context, providerID string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, fmt.Sprintf("%s/api/v1/jobs/%s", p.cfg.Endpoint, providerID), nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", p.cfg.Credential)

	resp, err := p.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading resp body: %w", err)
	}

	if c := resp.StatusCode; c/100 > 3 {
		return fmt.Errorf("received non 2xx status code, got %d with body: %s", c, string(body))
	}

	return nil
}

func (p *flock) Healthcheck() error {
	resp, err := p.client.Get(fmt.Sprintf("%s/healthcheck", p.cfg.Endpoint))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return nil
}

func (*flock) Capabilities() provider.Capabilities {
	return provider.Capabilities{
		InputFormats:  []string{"h264", "h265"},
		OutputFormats: []string{"mp4"},
		Destinations:  []string{"s3", "gs"},
	}
}

func flockFactory(cfg *config.Config) (provider.TranscodingProvider, error) {
	if cfg.Flock.Endpoint == "" || cfg.Flock.Credential == "" {
		return nil, errors.New("incomplete Flock config")
	}

	return &flock{
		cfg:    cfg.Flock,
		client: &http.Client{Timeout: time.Second * 30},
	}, nil
}
