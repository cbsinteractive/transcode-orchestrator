package hybrik

import (
	"context"
	"encoding/json"
	"fmt"
	"path"
	"strings"

	hwrapper "github.com/cbsinteractive/hybrik-sdk-go"
	"github.com/cbsinteractive/transcode-orchestrator/config"
	"github.com/cbsinteractive/transcode-orchestrator/job"
	"github.com/cbsinteractive/transcode-orchestrator/provider"
	"github.com/pkg/errors"
)

type (
	Job    = job.Job
	Status = job.Status
)

type executionFeatures struct {
	segmentedRendering      *hwrapper.SegmentedRendering
	doViPreProcSegmentation doViPreProcSegmentation
}

const (
	// Name describes the name of the transcoder
	Name                       = "hybrik"
	queued                     = "queued"
	active                     = "active"
	completed                  = "completed"
	failed                     = "failed"
	activeRunning              = "running"
	activeWaiting              = "waiting"
	hls                        = "hls"
	transcodeElementIDTemplate = "transcode_task_%d"
)

var (
	ErrUnsupportedContainer = errors.New("container format unsupported. Hybrik provider capabilities may need to be updated")
)

func init() {
	provider.Register(Name, hybrikTranscoderFactory)
}

type hybrikProvider struct {
	c      hwrapper.ClientInterface
	config *config.Hybrik
}

func (p hybrikProvider) String() string {
	return "Hybrik"
}

func hybrikTranscoderFactory(cfg *config.Config) (provider.Provider, error) {
	api, err := hwrapper.NewClient(hwrapper.Config{
		URL:            cfg.Hybrik.URL,
		ComplianceDate: cfg.Hybrik.ComplianceDate,
		OAPIKey:        cfg.Hybrik.OAPIKey,
		OAPISecret:     cfg.Hybrik.OAPISecret,
		AuthKey:        cfg.Hybrik.AuthKey,
		AuthSecret:     cfg.Hybrik.AuthSecret,
	})
	if err != nil {
		return &hybrikProvider{}, err
	}

	return &hybrikProvider{
		c:      api,
		config: cfg.Hybrik,
	}, nil
}

func (p *hybrikProvider) Create(ctx context.Context, j *Job) (*Status, error) {
	c, err := p.create(j)
	if err != nil {
		return nil, err
	}

	id, err := p.c.QueueJob(string(c))
	if err != nil {
		return nil, err
	}

	return &Status{
		ProviderName:  Name,
		ProviderJobID: id,
		State:         job.StateQueued,
	}, nil
}

func (p *hybrikProvider) create(j *Job) ([]byte, error) {
	c, err := p.createJobReqFrom(job)
	if err != nil {
		return nil, fmt.Errorf("hybrik: createjob: %w", err)
	}
	return json.MarshalIndent(c, "", "\t")
}

func (p *hybrikProvider) Valid(j *Job) error {
	// provider supported
	// valid credentials
	// valid codecs and features
}

/*
	destination: storageLocation{
		provider: destStorageProvider,
		path:     fmt.Sprintf("%s/%s", destinationPath, j.RootFolder()),
	},
*/

func tag(j *Job, name string, fallback ...string) []string {
	v := j.ExecutionEnv.ComputeTags[name]
	if len(v) == 0 {
		return fallback
	}
	return v
}

// jobRequest assumes j was already validated, error conditions
// are impossible by design as they are overridden by defaults
func (p *hybrikProvider) jobRequest(jo *job.Job) hwrapper.CreateJob {
	j := Job{jo, map[string][]string{}}
	for k, v := range j.ExecutionEnv.ComputeTags {
		if len(v) == 0 {
			v = defaultTag[k]
		}
		j.tag[k] = v
	}
	feat := features(j)
	cfg.executionFeatures = execFeatures
	cfg.computeTags = j.ExecutionEnv.ComputeTags

	conn := []hwrapper.Connection{}
	task := []hwrapper.Element{cfg.source}
	prev := task
	for _, eg := range p.elementAssemblerFrom(p.outputCfgsFrom(ctx, j))(cfg) {
		src := []hwrapper.ConnectionFrom{}
		dst := []hwrapper.ToSuccess{}
		for _, e := range prev {
			src = append(src, hwrapper.ConnectionFrom{Element: e.UID})
		}
		for _, e := range eg {
			dst = append(dst, hwrapper.ToSuccess{Element: e.UID})
			task = append(task, e)
		}
		conn = append(conn, hwrapper.Connection{
			From: src,
			To:   hwrapper.ConnectionTo{Success: dst},
		})
		prev = eg
	}

	return hwrapper.CreateJob{
		Name: fmt.Sprintf("Job %s [%s]", cfg.jobID, path.Base(cfg.sourceLocation.path)),
		Payload: hwrapper.CreateJobPayload{
			Elements:    task,
			Connections: conn,
		},
	}
}

func (p *hybrikProvider) StorageFallback() (path string) {
	return p.config.Destination
}

func (p *hybrikProvider) Status(_ context.Context, j *Job) (*Status, error) {
	ji, err := p.c.GetJobInfo(j.ProviderJobID)
	if err != nil {
		return &Status{}, err
	}

	var status job.State
	switch ji.Status {
	case active:
		fallthrough
	case activeRunning:
		fallthrough
	case activeWaiting:
		status = job.StateStarted
	case queued:
		status = job.StateQueued
	case completed:
		status = job.StateFinished
	case failed:
		status = job.StateFailed
	}

	var output job.Output
	if status == job.StateFailed || status == job.StateFinished {
		result, err := p.c.GetJobResult(j.ProviderJobID)
		if err != nil {
			return &Status{}, err
		}

		output = job.Output{}
		for _, task := range result.Tasks {
			files, found, err := filesFrom(task)
			if err != nil {
				return &Status{}, err
			}
			if found {
				output.Files = append(output.Files, files...)
			}
		}
	}

	return &Status{
		ProviderJobID: j.ProviderJobID,
		ProviderName:  p.String(),
		Progress:      float64(ji.Progress),
		State:         status,
		Output:        output,
	}, nil
}

func features(job *Job) *hwrapper.SegmentedRendering {
	v, has := job.ExecutionFeatures["segmentedRendering"]
	if !has || !canSegment(j.Input) {
		return feat, nil
	}
	data, _ := json.Marshal(v)
	sr := hwrapper.SegmentedRendering{}
	if err := json.Unmarshal(data, &sr); err != nil {
		return nil
	}
	return &sr

}

func (p *hybrikProvider) Cancel(_ context.Context, id string) error {
	return p.c.StopJob(id)
}

func videoTarget(v job.Video) *hwrapper.VideoTarget {
	if (v == job.Video{}) {
		return nil, nil
	}

	var frames, seconds int
	if v.Gop.Seconds() {
		seconds = int(v.Gop.Size)
	} else {
		frames = int(v.Gop.Size)
	}

	profile := strings.ToLower(preset.Profile)
	level := preset.Level

	// TODO: Understand video-transcoding-api profile + level settings in relation to vp8
	// For now, we will omit and leave to encoder defaults
	if v.Codec == "vp8" {
		profile = ""
		level = ""
	}

	w, h := &preset.Width, &preset.Height
	if *w == 0 {
		w = nil
	}
	if *h == 0 {
		h = nil
	}
	return &hwrapper.VideoTarget{
		Width:             w,
		Height:            h,
		BitrateMode:       strings.ToLower(preset.Bitrate.Mode),
		BitrateKb:         v.Bitrate.Kbps(),
		Preset:            "slow",
		Codec:             preset.Codec,
		ChromaFormat:      chromaFormatYUV420P,
		Profile:           profile,
		Level:             level,
		ExactGOPFrames:    frames,
		ExactKeyFrame:     seconds,
		InterlaceMode:     preset.InterlaceMode,
		UseSceneDetection: false,
	}, nil
}

func audioTarget(a job.Audio) []hwrapper.AudioTarget {
	if (a == job.Audio{}) {
		return []hwrapper.AudioTarget{}
	}
	return []hwrapper.AudioTarget{
		{
			Codec:     a.Codec,
			Channels:  2,
			BitrateKb: a.Bitrate / 1000,
		},
	}
}

// Healthcheck should return nil if the provider is currently available
// for transcoding videos, otherwise it should return an error
// explaining what's going on.
func (p *hybrikProvider) Healthcheck() error {
	// For now, just call list jobs. If this errors, then we can consider the service unhealthy
	_, err := p.c.CallAPI("GET", "/jobs/info", nil, nil)
	return err
}

// Capabilities describes the capabilities of the provider.
func (p *hybrikProvider) Capabilities() provider.Capabilities {
	// we can support quite a bit more format wise, but unsure of schema so limiting to known supported video-transcoding-api formats for now...
	return provider.Capabilities{
		InputFormats:  []string{"prores", "h264", "h265"},
		OutputFormats: []string{"mp4", "hls", "webm", "mov"},
		Destinations:  []string{storageProviderS3.string(), storageProviderGCS.string()},
	}
}
