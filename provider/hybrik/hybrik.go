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

	featureSegmentedRendering = "segmentedRendering"
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
	cj, err := p.createJobReqBodyFrom(ctx, j)
	if err != nil {
		return nil, err
	}

	id, err := p.c.QueueJob(cj)
	if err != nil {
		return nil, err
	}

	return &Status{
		ProviderName:  Name,
		ProviderJobID: id,
		State:         job.StateQueued,
	}, nil
}

func (p *hybrikProvider) createJobReqBodyFrom(ctx context.Context, job *Job) (string, error) {
	cj, err := p.createJobReqFrom(ctx, job)
	if err != nil {
		return "", errors.Wrap(err, "generating create job request from Job")
	}

	resp, err := json.MarshalIndent(cj, "", "\t")
	if err != nil {
		return "", errors.Wrapf(err, "marshalling create job %v into json", cj)
	}

	return string(resp), nil
}

func (p *hybrikProvider) createJobReqFrom(ctx context.Context, j *Job) (hwrapper.CreateJob, error) {
	srcStorageProvider, err := storageProviderFrom(j.SourceMedia)
	if err != nil {
		return hwrapper.CreateJob{}, errors.Wrap(err, "parsing source storage provider")
	}
	srcLocation := storageLocation{
		provider: srcStorageProvider,
		path:     j.SourceMedia,
	}

	destinationPath := p.destinationForJob(j)
	destStorageProvider, err := storageProviderFrom(destinationPath)
	if err != nil {
		return hwrapper.CreateJob{}, errors.Wrap(err, "parsing destination storage provider")
	}

	srcElement, err := p.srcFrom(j, srcLocation)
	if err != nil {
		return hwrapper.CreateJob{}, errors.Wrap(err, "creating the hybrik source element")
	}

	cfg := jobCfg{
		jobID:          j.ID,
		sourceLocation: srcLocation,
		destination: storageLocation{
			provider: destStorageProvider,
			path:     fmt.Sprintf("%s/%s", destinationPath, j.RootFolder()),
		},
		streamingParams:      j.StreamingParams,
		executionEnvironment: j.ExecutionEnv,
		source:               srcElement,
	}

	execFeatures, err := executionFeaturesFrom(j, srcLocation.provider)
	if err != nil {
		return hwrapper.CreateJob{}, err
	}
	cfg.executionFeatures = execFeatures

	if j.ExecutionEnv.ComputeTags != nil {
		cfg.computeTags = j.ExecutionEnv.ComputeTags
	} else {
		cfg.computeTags = map[job.ComputeClass]string{}
	}

	outputCfgs, err := p.outputCfgsFrom(ctx, j)
	if err != nil {
		return hwrapper.CreateJob{}, err
	}
	cfg.outputCfgs = outputCfgs

	elmAssembler, err := p.elementAssemblerFrom(cfg.outputCfgs)
	if err != nil {
		return hwrapper.CreateJob{}, err
	}

	elementGroups, err := elmAssembler(cfg)
	if err != nil {
		return hwrapper.CreateJob{}, err
	}
	cfg.elementGroups = elementGroups

	connections := []hwrapper.Connection{}
	prevElements := []hwrapper.Element{cfg.source}
	allTaskElements := []hwrapper.Element{}

	for _, elementGroup := range elementGroups {
		fromConnections := []hwrapper.ConnectionFrom{}
		for _, prevElement := range prevElements {
			fromConnections = append(fromConnections, hwrapper.ConnectionFrom{Element: prevElement.UID})
		}

		toSuccessElements := []hwrapper.ToSuccess{}
		for _, element := range elementGroup {
			allTaskElements = append(allTaskElements, element)
			toSuccessElements = append(toSuccessElements, hwrapper.ToSuccess{Element: element.UID})
		}

		toConnections := hwrapper.ConnectionTo{Success: toSuccessElements}
		connections = append(connections, hwrapper.Connection{
			From: fromConnections,
			To:   toConnections,
		})
		prevElements = elementGroup
	}

	// create the full job structure
	cj := hwrapper.CreateJob{
		Name: fmt.Sprintf("Job %s [%s]", cfg.jobID, path.Base(cfg.sourceLocation.path)),
		Payload: hwrapper.CreateJobPayload{
			Elements:    append([]hwrapper.Element{cfg.source}, allTaskElements...),
			Connections: connections,
		},
	}

	if _, found := supportedPackagingProtocols[strings.ToLower(cfg.streamingParams.Protocol)]; found {
		cj, err = p.enrichCreateJobWithPackagingCfg(cj, cfg, prevElements)
		if err != nil {
			return hwrapper.CreateJob{}, errors.Wrap(err, "creating packaging config")
		}
	}

	return cj, nil
}

func (p *hybrikProvider) destinationForJob(job *Job) string {
	if path := job.DestinationBasePath; path != "" {
		return path
	}

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

func executionFeaturesFrom(job *Job, storageProvider storageProvider) (executionFeatures, error) {
	features := executionFeatures{}

	supportsSegRendering := storageProvider.supportsSegmentedRendering()
	if featureDefinition, ok := job.ExecutionFeatures[featureSegmentedRendering]; ok && supportsSegRendering {
		featureJSON, err := json.Marshal(featureDefinition)
		if err != nil {
			return executionFeatures{}, fmt.Errorf("could not marshal segmented rendering cfg to json: %v", err)
		}

		var feature SegmentedRendering
		err = json.Unmarshal(featureJSON, &feature)
		if err != nil {
			return executionFeatures{}, fmt.Errorf("could not unmarshal %q into SegmentedRendering feature: %v",
				featureDefinition, err)
		}

		features.segmentedRendering = &hwrapper.SegmentedRendering{
			Duration:                  feature.Duration,
			SceneChangeSearchDuration: feature.SceneChangeSearchDuration,
			NumTotalSegments:          feature.NumTotalSegments,
			EnableStrictCFR:           feature.EnableStrictCFR,
			MuxTimebaseOffset:         feature.MuxTimebaseOffset,
		}

		features.doViPreProcSegmentation = doViPreProcSegmentation{
			numTasks:       feature.DoViPreProcNumTasks,
			intervalLength: feature.DoViPreProcIntervalLength,
		}
	}

	return features, nil
}

func (p *hybrikProvider) Cancel(_ context.Context, id string) error {
	return p.c.StopJob(id)
}

func videoTargetFrom(preset job.Video, rateControl string) (*hwrapper.VideoTarget, error) {
	if (preset == job.Video{}) {
		return nil, nil
	}

	var exactGOPFrames, exactKeyFrames int
	switch strings.ToLower(preset.GopUnit) {
	case job.GopUnitSeconds:
		exactKeyFrames = int(preset.GopSize)
	case job.GopUnitFrames, "":
		exactGOPFrames = int(preset.GopSize)
	default:
		return &hwrapper.VideoTarget{}, fmt.Errorf("GopUnit %v not recognized", preset.GopUnit)
	}

	videoProfile := strings.ToLower(preset.Profile)
	videoLevel := preset.Level

	// TODO: Understand video-transcoding-api profile + level settings in relation to vp8
	// For now, we will omit and leave to encoder defaults
	if preset.Codec == "vp8" {
		videoProfile = ""
		videoLevel = ""
	}

	return &hwrapper.VideoTarget{
		Width:             &preset.Width,
		Height:            &preset.Height,
		BitrateMode:       strings.ToLower(rateControl),
		BitrateKb:         preset.Bitrate / 1000,
		Preset:            presetSlow,
		Codec:             preset.Codec,
		ChromaFormat:      chromaFormatYUV420P,
		Profile:           videoProfile,
		Level:             videoLevel,
		ExactGOPFrames:    exactGOPFrames,
		ExactKeyFrame:     exactKeyFrames,
		InterlaceMode:     preset.InterlaceMode,
		UseSceneDetection: false,
	}, nil
}

func audioTargetFrom(preset job.Audio) ([]hwrapper.AudioTarget, error) {
	if (preset == job.Audio{}) {
		return []hwrapper.AudioTarget{}, nil
	}
	return []hwrapper.AudioTarget{
		{
			Codec:     preset.Codec,
			Channels:  2,
			BitrateKb: preset.Bitrate / 1000,
		},
	}, nil
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
