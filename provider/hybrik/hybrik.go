package hybrik

import (
	"encoding/json"
	"fmt"
	"path"
	"strconv"
	"strings"

	hwrapper "github.com/cbsinteractive/hybrik-sdk-go"
	"github.com/cbsinteractive/video-transcoding-api/config"
	"github.com/cbsinteractive/video-transcoding-api/db"
	"github.com/cbsinteractive/video-transcoding-api/db/redis"
	"github.com/cbsinteractive/video-transcoding-api/provider"
	"github.com/pkg/errors"
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
	// ErrBitrateNan is an error returned when the bitrate field of db.Preset is not a valid number
	ErrBitrateNan = errors.New("bitrate not a number")

	// ErrVideoWidthNan is an error returned when the preset video width of db.Preset is not a valid number
	ErrVideoWidthNan = errors.New("preset video width not a number")

	// ErrVideoHeightNan is an error returned when the preset video height of db.Preset is not a valid number
	ErrVideoHeightNan = errors.New("preset video height not a number")

	// ErrUnsupportedContainer is returned when the container format is not present in the provider's capabilities list
	ErrUnsupportedContainer = errors.New("container format unsupported. Hybrik provider capabilities may need to be updated")
)

func init() {
	provider.Register(Name, hybrikTranscoderFactory)
}

type hybrikProvider struct {
	c          hwrapper.ClientInterface
	config     *config.Hybrik
	repository db.Repository
}

func (p hybrikProvider) String() string {
	return "Hybrik"
}

func hybrikTranscoderFactory(cfg *config.Config) (provider.TranscodingProvider, error) {
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

	dbRepo, err := redis.NewRepository(cfg)
	if err != nil {
		return nil, fmt.Errorf("error initializing hybrik wrapper: %s", err)
	}

	return &hybrikProvider{
		c:          api,
		config:     cfg.Hybrik,
		repository: dbRepo,
	}, nil
}

func (p *hybrikProvider) Transcode(job *db.Job) (*provider.JobStatus, error) {
	cj, err := p.createJobReqBodyFrom(job)
	if err != nil {
		return &provider.JobStatus{}, err
	}

	id, err := p.c.QueueJob(cj)
	if err != nil {
		return &provider.JobStatus{}, err
	}

	return &provider.JobStatus{
		ProviderName:  Name,
		ProviderJobID: id,
		Status:        provider.StatusQueued,
	}, nil
}

func (p *hybrikProvider) createJobReqBodyFrom(job *db.Job) (string, error) {
	cj, err := p.createJobReqFrom(job)
	if err != nil {
		return "", errors.Wrap(err, "generating create job request from db.Job")
	}

	resp, err := json.MarshalIndent(cj, "", "\t")
	if err != nil {
		return "", errors.Wrapf(err, "marshalling create job %v into json", cj)
	}

	return string(resp), nil
}

func (p *hybrikProvider) createJobReqFrom(job *db.Job) (hwrapper.CreateJob, error) {
	srcStorageProvider, err := storageProviderFrom(job.SourceMedia)
	if err != nil {
		return hwrapper.CreateJob{}, errors.Wrap(err, "parsing source storage provider")
	}
	srcLocation := storageLocation{
		provider: srcStorageProvider,
		path:     job.SourceMedia,
	}

	destinationPath := p.destinationForJob(job)
	destStorageProvider, err := storageProviderFrom(destinationPath)
	if err != nil {
		return hwrapper.CreateJob{}, errors.Wrap(err, "parsing destination storage provider")
	}

	srcElement, err := p.srcFrom(job, srcLocation)
	if err != nil {
		return hwrapper.CreateJob{}, errors.Wrap(err, "creating the hybrik source element")
	}

	cfg := jobCfg{
		jobID:          job.ID,
		sourceLocation: srcLocation,
		destination: storageLocation{
			provider: destStorageProvider,
			path:     fmt.Sprintf("%s/%s", destinationPath, job.ID),
		},
		streamingParams:      job.StreamingParams,
		executionEnvironment: job.ExecutionEnv,
		source:               srcElement,
	}

	execFeatures, err := executionFeaturesFrom(job, srcLocation.provider)
	if err != nil {
		return hwrapper.CreateJob{}, err
	}
	cfg.executionFeatures = execFeatures

	if job.ExecutionEnv.ComputeTags != nil {
		cfg.computeTags = job.ExecutionEnv.ComputeTags
	} else {
		cfg.computeTags = map[db.ComputeClass]string{}
	}

	outputCfgs, err := p.outputCfgsFrom(job)
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

func (p *hybrikProvider) destinationForJob(job *db.Job) string {
	if path := job.DestinationBasePath; path != "" {
		return path
	}

	return p.config.Destination
}

func (p *hybrikProvider) JobStatus(job *db.Job) (*provider.JobStatus, error) {
	ji, err := p.c.GetJobInfo(job.ProviderJobID)
	if err != nil {
		return &provider.JobStatus{}, err
	}

	var status provider.Status
	switch ji.Status {
	case active:
		fallthrough
	case activeRunning:
		fallthrough
	case activeWaiting:
		status = provider.StatusStarted
	case queued:
		status = provider.StatusQueued
	case completed:
		status = provider.StatusFinished
	case failed:
		status = provider.StatusFailed
	}

	var output provider.JobOutput
	if status == provider.StatusFailed || status == provider.StatusFinished {
		result, err := p.c.GetJobResult(job.ProviderJobID)
		if err != nil {
			return &provider.JobStatus{}, err
		}

		output = provider.JobOutput{}
		for _, task := range result.Tasks {
			files, found, err := filesFrom(task)
			if err != nil {
				return &provider.JobStatus{}, err
			}
			if found {
				output.Files = append(output.Files, files...)
			}
		}
	}

	return &provider.JobStatus{
		ProviderJobID: job.ProviderJobID,
		ProviderName:  p.String(),
		Progress:      float64(ji.Progress),
		Status:        status,
		Output:        output,
	}, nil
}

func executionFeaturesFrom(job *db.Job, storageProvider storageProvider) (executionFeatures, error) {
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

func (p *hybrikProvider) CancelJob(id string) error {
	return p.c.StopJob(id)
}

func (p *hybrikProvider) CreatePreset(preset db.Preset) (string, error) {
	err := p.repository.CreateLocalPreset(&db.LocalPreset{
		Name:   preset.Name,
		Preset: preset,
	})
	if err != nil {
		return "", err
	}

	return preset.Name, nil
}

func videoTargetFrom(preset db.VideoPreset, rateControl string) (*hwrapper.VideoTarget, error) {
	if (preset == db.VideoPreset{}) {
		return nil, nil
	}

	var minGOPFrames, maxGOPFrames, gopSize int

	gopSize, err := strconv.Atoi(preset.GopSize)
	if err != nil {
		return &hwrapper.VideoTarget{}, err
	}

	minGOPFrames = gopSize
	maxGOPFrames = gopSize

	bitrate, err := strconv.Atoi(preset.Bitrate)
	if err != nil {
		return &hwrapper.VideoTarget{}, ErrBitrateNan
	}

	var videoWidth *int
	var videoHeight *int

	if preset.Width != "" {
		var presetWidth int
		presetWidth, err = strconv.Atoi(preset.Width)
		if err != nil {
			return &hwrapper.VideoTarget{}, ErrVideoWidthNan
		}
		videoWidth = &presetWidth
	}

	if preset.Height != "" {
		var presetHeight int
		presetHeight, err = strconv.Atoi(preset.Height)
		if err != nil {
			return &hwrapper.VideoTarget{}, ErrVideoHeightNan
		}
		videoHeight = &presetHeight
	}

	videoProfile := strings.ToLower(preset.Profile)
	videoLevel := preset.ProfileLevel

	// TODO: Understand video-transcoding-api profile + level settings in relation to vp8
	// For now, we will omit and leave to encoder defaults
	if preset.Codec == "vp8" {
		videoProfile = ""
		videoLevel = ""
	}

	return &hwrapper.VideoTarget{
		Width:             videoWidth,
		Height:            videoHeight,
		BitrateMode:       strings.ToLower(rateControl),
		BitrateKb:         bitrate / 1000,
		Preset:            presetSlow,
		Codec:             preset.Codec,
		ChromaFormat:      chromaFormatYUV420P,
		Profile:           videoProfile,
		Level:             videoLevel,
		MinGOPFrames:      minGOPFrames,
		MaxGOPFrames:      maxGOPFrames,
		ExactGOPFrames:    maxGOPFrames,
		InterlaceMode:     preset.InterlaceMode,
		UseSceneDetection: false,
	}, nil
}

func audioTargetFrom(preset db.AudioPreset) ([]hwrapper.AudioTarget, error) {
	if (preset == db.AudioPreset{}) {
		return []hwrapper.AudioTarget{}, nil
	}

	audioBitrate, err := strconv.Atoi(preset.Bitrate)
	if err != nil {
		return nil, ErrBitrateNan
	}
	return []hwrapper.AudioTarget{
		{
			Codec:     preset.Codec,
			Channels:  2,
			BitrateKb: audioBitrate / 1000,
		},
	}, nil
}

func (p *hybrikProvider) DeletePreset(presetID string) error {
	preset, err := p.GetPreset(presetID)
	if err != nil {
		return err
	}

	return p.repository.DeleteLocalPreset(preset.(*db.LocalPreset))
}

func (p *hybrikProvider) GetPreset(presetID string) (interface{}, error) {
	return p.repository.GetLocalPreset(presetID)
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
