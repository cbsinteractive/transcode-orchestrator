package flock

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/cbsinteractive/transcode-orchestrator/client/transcoding/job"
	"github.com/cbsinteractive/transcode-orchestrator/config"
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
	cfg    *config.Flock
	client *http.Client
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

func (p *flock) Create(ctx context.Context, j *job.Job) (*job.Status, error) {
	jr, err := p.flockJobRequestFrom(j)
	if err != nil {
		return nil, fmt.Errorf("generating flock job request: %w", err)
	}

	data, err := json.Marshal(jr)
	if err != nil {
		return nil, fmt.Errorf("marshaling job request %+v to json: %w", jr, err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		fmt.Sprintf("%s/api/v1/jobs", p.cfg.Endpoint), bytes.NewBuffer(data))
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

	return &job.Status{
		Provider:      Name,
		ProviderJobID: fmt.Sprintf("%d", newJob.JobID),
		State:         job.StateQueued,
	}, nil
}

func NewRequest(j *job.Job) (*JobRequest, error) {
	fj := &JobRequest{}
	fj.Job.Source = j.Input.Name
	fj.Job.Labels = j.Labels

	for _, orc := range j.Output.File {
		flock := JobOutput{
			Preset:           "slow",
			Destination:      j.Location(orc.Name),
			VideoCodec:       orc.Video.Codec,
			Width:            orc.Video.Width,
			Height:           orc.Video.Height,
			MultiPass:        orc.Video.Bitrate.TwoPass,
			VideoBitrateKbps: orc.Video.Bitrate.BPS / 1000,
			AudioBitrateKbps: orc.Audio.Bitrate / 1000,
			AudioCodec:       orc.Audio.Codec,
		}
		if flock.AudioBitrateKbps != 0 {
			flock.AudioChannels = 2
		}
		if flock.VideoCodec == "hevc" {
			flock.Tag = "hvc1"
		}
		if orc.Video.Gop.Seconds() {
			flock.KeyframesPerSecond = int(orc.Video.Gop.Size)
		} else {
			flock.KeyframeInterval = int(orc.Video.Gop.Size)
		}
		fj.Job.Outputs = append(fj.Job.Outputs, flock)
	}
	return fj, nil
}

func (p *flock) flockJobRequestFrom(j *job.Job) (*JobRequest, error) {
	return NewRequest(j)
}

func (p *flock) Status(ctx context.Context, job *job.Job) (*job.Status, error) {
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

func (p *flock) jobStatusFrom(j *job.Job, fj *JobResponse) *job.Status {
	status := &job.Status{
		Progress:      fj.Progress,
		ProviderJobID: j.ProviderJobID,
		Provider:      Name,
		ProviderStatus: map[string]interface{}{
			"create_timestamp": fj.CreateTimestamp,
			"update_timestamp": fj.UpdateTimestamp,
			"duration":         fj.Duration,
			"progress":         fj.Progress,
			"status":           fj.Status,
		},
		State:  state(fj),
		Output: job.Dir{Path: j.Location("")},
		Labels: j.Labels,
	}

	outstatus := []map[string]interface{}{}
	for _, o := range fj.Outputs {
		status.Output.File = append(status.Output.File, job.File{Name: o.Destination})
		outstatus = append(outstatus, map[string]interface{}{
			"duration":         o.Duration,
			"destination":      o.Destination,
			"encoder":          o.Encoder,
			"info":             o.Info,
			"status":           o.Status,
			"progress":         o.Progress,
			"update_timestamp": o.UpdateTimestamp,
		})
	}
	status.ProviderStatus["outputs"] = outstatus
	return status
}

func state(j *JobResponse) job.State {
	switch j.Status {
	case "submitted":
		return job.StateQueued
	case "assigned", "transcoding":
		return job.StateStarted
	case "complete":
		return job.StateFinished
	case "cancelled":
		return job.StateCanceled
	case "error":
		return job.StateFailed
	}
	return job.StateUnknown
}

func (p *flock) Cancel(ctx context.Context, providerID string) error {
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

func flockFactory(cfg *config.Config) (provider.Provider, error) {
	if cfg.Flock.Endpoint == "" || cfg.Flock.Credential == "" {
		return nil, errors.New("incomplete Flock config")
	}

	return &flock{
		cfg:    cfg.Flock,
		client: &http.Client{Timeout: time.Second * 30},
	}, nil
}
