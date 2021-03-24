package hybrik

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"path"

	hy "github.com/cbsinteractive/hybrik-sdk-go"
	"github.com/cbsinteractive/transcode-orchestrator/av"
	job "github.com/cbsinteractive/transcode-orchestrator/av"
	"github.com/cbsinteractive/transcode-orchestrator/config"
	"github.com/cbsinteractive/transcode-orchestrator/provider"
)

type (
	Job    = av.Job
	Status = av.Status
)

type executionFeatures struct {
	segmentedRendering      *hy.SegmentedRendering
	doViPreProcSegmentation doViPreProcSegmentation
}

const (
	// Name describes the name of the transcoder
	Name          = "hybrik"
	queued        = "queued"
	active        = "active"
	completed     = "completed"
	failed        = "failed"
	activeRunning = "running"
	activeWaiting = "waiting"
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
		Provider:      Name,
		ProviderJobID: id,
		State:         av.StateQueued,
	}, nil
}

func (p *driver) create(j *Job) ([]byte, error) {
	c, err := p.jobRequest(j)
	if err != nil {
		return nil, err
	}
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
	v := j.Env.Tags[name]
	if fallback == nil {
		// this is a code smell; just did it to make tests pass
		fallback = []string{}
	}
	if len(v) == 0 {
		return fallback
	}
	return []string{v}
}

func (p *driver) jobRequest(j *Job) (*hy.CreateJob, error) {
	if err := p.validate(j); err != nil {
		return nil, err
	}
	conn := []hy.Connection{}
	task := []hy.Element{p.srcFrom(j)}
	prev := task
	for _, eg := range p.assemble(j) {
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

	return &hy.CreateJob{
		Name: fmt.Sprintf("Job %s [%s]", j.ID, path.Base(j.Input.Name)),
		Payload: hy.CreateJobPayload{
			Elements:    task,
			Connections: conn,
		},
	}, nil
}

func (p *driver) StorageFallback() (path string) {
	return p.config.Destination
}

func (p *driver) Status(_ context.Context, j *Job) (*Status, error) {
	ji, err := p.c.GetJobInfo(j.ProviderJobID)
	if err != nil {
		return &Status{}, err
	}

	var status av.State
	switch ji.Status {
	case active:
		fallthrough
	case activeRunning:
		fallthrough
	case activeWaiting:
		status = av.StateStarted
	case queued:
		status = av.StateQueued
	case completed:
		status = av.StateFinished
	case failed:
		status = av.StateFailed
	}

	var output job.Dir
	if status == av.StateFailed || status == av.StateFinished {
		result, err := p.c.GetJobResult(j.ProviderJobID)
		if err != nil {
			return &Status{}, err
		}

		output = job.Dir{}
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
		Provider:      p.String(),
		Progress:      float64(ji.Progress),
		State:         status,
		Output:        output,
	}, nil
}

func features(j *Job) *hy.SegmentedRendering {
	s := SegmentedRendering{}
	if features0(j, &s) {
		// they're exactly the same thing... except for the json names
		return &hy.SegmentedRendering{
			Duration:                  s.Duration,
			SceneChangeSearchDuration: s.SceneChangeSearchDuration,
			NumTotalSegments:          s.NumTotalSegments,
			EnableStrictCFR:           s.EnableStrictCFR,
			MuxTimebaseOffset:         s.MuxTimebaseOffset,
		}
	}
	return nil
}

// features0; dst should point to preproc dovi
// or *hy.SegmentedRendering
func features0(j *Job, dst interface{}) bool {
	// NOTE(as): There's some strange behavior here
	// it's always looking at the segmented rendering
	// features even for the old dolby vision stuff
	v, has := j.Features["segmentedRendering"]
	if !has || j.Input.Provider() == "http" {
		// TODO(as): this check for http is a direct copy from the old
		// version, but is http the only thing that doesn't support
		// segmented rendering? what about https?
		return false
	}
	data, _ := json.Marshal(v)
	err := json.Unmarshal(data, dst)
	return err == nil
}

func (p *driver) Cancel(_ context.Context, id string) error {
	return p.c.StopJob(id)
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
