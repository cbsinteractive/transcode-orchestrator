package hybrik

import (
	"context"
	"encoding/json"
	"fmt"
	"path"
	"strings"

	hy "github.com/cbsinteractive/hybrik-sdk-go"
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
	segmentedRendering      *hy.SegmentedRendering
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

type driver struct {
	c      hy.ClientInterface
	config *config.Hybrik
}

func (p driver) String() string {
	return "Hybrik"
}

func hybrikTranscoderFactory(cfg *config.Config) (provider.Provider, error) {
	api, err := hy.NewClient(hy.Config{
		URL:            cfg.Hybrik.URL,
		ComplianceDate: cfg.Hybrik.ComplianceDate,
		OAPIKey:        cfg.Hybrik.OAPIKey,
		OAPISecret:     cfg.Hybrik.OAPISecret,
		AuthKey:        cfg.Hybrik.AuthKey,
		AuthSecret:     cfg.Hybrik.AuthSecret,
	})
	if err != nil {
		return &driver{}, err
	}

	return &driver{
		c:      api,
		config: cfg.Hybrik,
	}, nil
}

func (p *driver) Create(ctx context.Context, j *Job) (*Status, error) {
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

func (p *driver) create(j *Job) ([]byte, error) {
	c := p.jobRequest(j)
	return json.MarshalIndent(c, "", "\t")
}

func (p *driver) Valid(j *Job) error {
	// provider supported
	// valid credentials
	// valid codecs and features
	return nil
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
	return []string{v}
}

// jobRequest assumes j was already validated, error conditions
// are impossible by design as they are overridden by defaults
func (p *driver) jobRequest(j *Job) hy.CreateJob {
	eg, err := p.assemble(j)
	if err != nil {
		println(err.Error()) // TODO(as): fix return
	}
	conn := []hy.Connection{}
	task := []hy.Element{p.srcFrom(j)}
	prev := task
	for _, eg := range eg {
		src := []hy.ConnectionFrom{}
		dst := []hy.ToSuccess{}
		for _, e := range prev {
			src = append(src, hy.ConnectionFrom{Element: e.UID})
		}
		for _, e := range eg {
			dst = append(dst, hy.ToSuccess{Element: e.UID})
			task = append(task, e)
		}
		conn = append(conn, hy.Connection{
			From: src,
			To:   hy.ConnectionTo{Success: dst},
		})
		prev = eg
	}

	return hy.CreateJob{
		Name: fmt.Sprintf("Job %s [%s]", j.ID, path.Base(j.Input.Name)),
		Payload: hy.CreateJobPayload{
			Elements:    task,
			Connections: conn,
		},
	}
}

func (p *driver) StorageFallback() (path string) {
	return p.config.Destination
}

func (p *driver) Status(_ context.Context, j *Job) (*Status, error) {
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

	var output job.Dir
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
				output.File = append(output.File, files...)
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

func features(job *Job) *hy.SegmentedRendering {
	v, has := job.ExecutionFeatures["segmentedRendering"]
	if !has || !canSegment(j.Input) {
		return feat, nil
	}
	data, _ := json.Marshal(v)
	sr := hy.SegmentedRendering{}
	if err := json.Unmarshal(data, &sr); err != nil {
		return nil
	}
	return &sr

}

func (p *driver) Cancel(_ context.Context, id string) error {
	return p.c.StopJob(id)
}

func videoTarget(v job.Video) *hy.VideoTarget {
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
	return &hy.VideoTarget{
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

func audioTarget(a job.Audio) []hy.AudioTarget {
	if (a == job.Audio{}) {
		return []hy.AudioTarget{}
	}
	return []hy.AudioTarget{
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
func (p *driver) Healthcheck() error {
	// For now, just call list jobs. If this errors, then we can consider the service unhealthy
	_, err := p.c.CallAPI("GET", "/jobs/info", nil, nil)
	return err
}

// Capabilities describes the capabilities of the provider.
func (p *driver) Capabilities() provider.Capabilities {
	// we can support quite a bit more format wise, but unsure of schema so limiting to known supported video-transcoding-api formats for now...
	return provider.Capabilities{
		InputFormats:  []string{"prores", "h264", "h265"},
		OutputFormats: []string{"mp4", "hls", "webm", "mov"},
		Destinations:  []string{storageProviderS3.string(), storageProviderGCS.string()},
	}
}
