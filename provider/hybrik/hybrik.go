package hybrik // import "github.com/NYTimes/video-transcoding-api/provider/hybrik"

import (
	"encoding/json"
	"fmt"
	"path"
	"strconv"
	"strings"

	"github.com/NYTimes/video-transcoding-api/config"
	"github.com/NYTimes/video-transcoding-api/db"
	"github.com/NYTimes/video-transcoding-api/provider"
	hwrapper "github.com/cbsinteractive/hybrik-sdk-go"
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
	transcodeElementIDTemplate = "transcode_task_%s"
	dolbyVisionElementID       = "dolby_vision_task"

	featureSegmentedRendering = "segmentedRendering"
)

var (
	// ErrBitrateNan is an error returned when the bitrate field of db.Preset is not a valid number
	ErrBitrateNan = errors.New("bitrate not a number")

	// ErrPresetOutputMatch represents an error in the hybrik encoding-wrapper provider.
	ErrPresetOutputMatch = errors.New("preset retrieved does not map to hybrik.Preset struct")

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
	c      hwrapper.ClientInterface
	config *config.Hybrik
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

	return &hybrikProvider{
		c:      api,
		config: cfg.Hybrik,
	}, nil
}

func (p *hybrikProvider) Transcode(job *db.Job) (*provider.JobStatus, error) {
	cj, err := p.presetsToTranscodeJob(job)
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

func (p *hybrikProvider) mountTranscodeElement(elementID, id, outputFilename string, destination storageLocation, segmentDuration uint,
	preset hwrapper.Preset, execFeatures executionFeatures, computeTags map[db.ComputeClass]string) hwrapper.Element {
	var e hwrapper.Element
	var subLocation *hwrapper.TranscodeLocation

	// outputFilename can be "test.mp4", or "subfolder1/subfodler2/test.mp4"
	// Handling accordingly
	subPath := path.Dir(outputFilename)
	outputFilePattern := path.Base(outputFilename)
	if subPath != "." && subPath != "/" {
		subLocation = &hwrapper.TranscodeLocation{
			StorageProvider: "relative",
			Path:            subPath,
		}
	}

	transcodePayload := hwrapper.TranscodePayload{
		LocationTargetPayload: hwrapper.LocationTargetPayload{
			Location: hwrapper.TranscodeLocation{
				StorageProvider: destination.provider,
				Path:            destination.path,
			},
			Targets: []hwrapper.TranscodeLocationTarget{
				{
					Location:    subLocation,
					FilePattern: outputFilePattern,
					Container: hwrapper.TranscodeTargetContainer{
						SegmentDuration: segmentDuration,
					},
				},
			},
		},
	}

	if execFeatures.segmentedRendering != nil {
		transcodePayload.SourcePipeline = hwrapper.TranscodeSourcePipeline{SegmentedRendering: execFeatures.segmentedRendering}
	}

	for _, modifier := range transcodePayloadModifiersFor(preset) {
		transcodePayload = modifier.runFunc(transcodePayload)
	}

	transcodeComputeTags := []string{}
	if tag, found := computeTags[db.ComputeClassTranscodeDefault]; found {
		transcodeComputeTags = append(transcodeComputeTags, tag)
	}

	// create the transcode element
	e = hwrapper.Element{
		UID:  fmt.Sprintf(transcodeElementIDTemplate, elementID),
		Kind: "transcode",
		Task: &hwrapper.ElementTaskOptions{
			Name: "Transcode - " + preset.Name,
			Tags: transcodeComputeTags,
		},
		Preset: &hwrapper.TranscodePreset{
			Key: preset.Name,
		},
		Payload: transcodePayload,
	}

	return e
}

type transcodePayloadModifier struct {
	name    string
	runFunc func(hwrapper.TranscodePayload) hwrapper.TranscodePayload
}

func transcodePayloadModifiersFor(preset hwrapper.Preset) []transcodePayloadModifier {
	modifiers := []transcodePayloadModifier{}

	for _, target := range preset.Payload.Targets {
		if hdr10 := target.Video.HDR10; hdr10 != nil && hdr10.Source != "" {
			modifiers = append(modifiers, transcodePayloadModifier{
				name:    "hdr10",
				runFunc: hdr10TranscodePayloadModifier,
			})
			break
		}
	}

	return modifiers
}

func (p *hybrikProvider) presetsToTranscodeJob(job *db.Job) (string, error) {
	srcStorageProvider, err := storageProviderFrom(job.SourceMedia)
	if err != nil {
		return "", errors.Wrap(err, "parsing source storage provider")
	}
	srcLocation := storageLocation{
		provider: srcStorageProvider,
		path:     job.SourceMedia,
	}

	destinationPath := p.destinationForJob(job)
	destStorageProvider, err := storageProviderFrom(destinationPath)
	if err != nil {
		return "", errors.Wrap(err, "parsing destination storage provider")
	}

	srcElement, err := srcFrom(job, srcLocation)
	if err != nil {
		return "", errors.Wrap(err, "creating the hybrik source element")
	}

	cfg := jobCfg{
		jobID:          job.ID,
		sourceLocation: srcLocation,
		destination: storageLocation{
			provider: destStorageProvider,
			path:     fmt.Sprintf("%s/%s", destinationPath, job.ID),
		},
		streamingParams: job.StreamingParams,
		source:          srcElement,
	}

	execFeatures, err := executionFeaturesFrom(job)
	if err != nil {
		return "", err
	}
	cfg.executionFeatures = execFeatures

	if job.ExecutionEnv.ComputeTags != nil {
		cfg.computeTags = job.ExecutionEnv.ComputeTags
	} else {
		cfg.computeTags = map[db.ComputeClass]string{}
	}

	outputCfgs, err := p.outputCfgsFrom(job)
	if err != nil {
		return "", err
	}
	cfg.outputCfgs = outputCfgs

	elmAssembler, err := p.elementAssemblerFrom(cfg.outputCfgs)
	if err != nil {
		return "", err
	}

	tasks, err := elmAssembler(cfg)
	if err != nil {
		return "", err
	}
	cfg.tasks = tasks

	transcodeSuccessConnections := []hwrapper.ToSuccess{}
	for _, task := range tasks {
		transcodeSuccessConnections = append(transcodeSuccessConnections, hwrapper.ToSuccess{Element: task.UID})
	}

	// create the full job structure
	cj := hwrapper.CreateJob{
		Name: fmt.Sprintf("Job %s [%s]", cfg.jobID, path.Base(cfg.sourceLocation.path)),
		Payload: hwrapper.CreateJobPayload{
			Elements: append([]hwrapper.Element{cfg.source}, tasks...),
			Connections: []hwrapper.Connection{{
				From: []hwrapper.ConnectionFrom{{
					Element: cfg.source.UID,
				}},
				To: hwrapper.ConnectionTo{
					Success: transcodeSuccessConnections,
				},
			}},
		},
	}

	// check if we need to add a master manifest task element
	if job.StreamingParams.Protocol == hls {
		manifestOutputDir := fmt.Sprintf("%s/%s", destinationPath, job.ID)
		manifestSubDir := path.Dir(job.StreamingParams.PlaylistFileName)
		manifestFilePattern := path.Base(job.StreamingParams.PlaylistFileName)

		if manifestSubDir != "." && manifestSubDir != "/" {
			manifestOutputDir = path.Join(manifestOutputDir, manifestSubDir)
		}

		manifestElement := hwrapper.Element{
			UID:  "manifest_creator",
			Kind: "manifest_creator",
			Payload: hwrapper.ManifestCreatorPayload{
				Location: hwrapper.TranscodeLocation{
					StorageProvider: cfg.destination.provider,
					Path:            manifestOutputDir,
				},
				FilePattern: manifestFilePattern,
				Kind:        hls,
			},
		}

		cj.Payload.Elements = append(cj.Payload.Elements, manifestElement)

		var manifestFromConnections []hwrapper.ConnectionFrom
		for _, task := range tasks {
			manifestFromConnections = append(manifestFromConnections, hwrapper.ConnectionFrom{Element: task.UID})
		}

		cj.Payload.Connections = append(cj.Payload.Connections,
			hwrapper.Connection{
				From: manifestFromConnections,
				To: hwrapper.ConnectionTo{
					Success: []hwrapper.ToSuccess{
						{Element: "manifest_creator"},
					},
				},
			},
		)

	}

	resp, err := json.Marshal(cj)
	if err != nil {
		return "", err
	}

	return string(resp), nil
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

	return &provider.JobStatus{
		ProviderJobID: job.ProviderJobID,
		ProviderName:  p.String(),
		Progress:      float64(ji.Progress),
		Status:        status,
	}, nil
}

func executionFeaturesFrom(job *db.Job) (executionFeatures, error) {
	features := executionFeatures{}
	if featureDefinition, ok := job.ExecutionFeatures[featureSegmentedRendering]; ok {
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
	hybrikPreset, err := p.hybrikPresetFrom(preset)
	if err != nil {
		return "", err
	}

	resultPreset, err := p.c.CreatePreset(hybrikPreset)
	if err != nil {
		return "", err
	}

	return resultPreset.Name, nil
}

type customPresetData struct {
	DolbyVisionEnabled bool `json:"dolbyVision"`
}

func (p *hybrikProvider) hybrikPresetFrom(preset db.Preset) (hwrapper.Preset, error) {
	var minGOPFrames, maxGOPFrames, gopSize int

	gopSize, err := strconv.Atoi(preset.Video.GopSize)
	if err != nil {
		return hwrapper.Preset{}, err
	}

	if preset.Video.GopMode == "fixed" {
		minGOPFrames = gopSize
		maxGOPFrames = gopSize
	} else {
		maxGOPFrames = gopSize
	}

	container := ""
	for _, c := range p.Capabilities().OutputFormats {
		if preset.Container == c || (preset.Container == "m3u8" && c == hls) {
			container = c
		}
	}

	if container == "" {
		return hwrapper.Preset{}, ErrUnsupportedContainer
	}

	bitrate, err := strconv.Atoi(preset.Video.Bitrate)
	if err != nil {
		return hwrapper.Preset{}, ErrBitrateNan
	}

	audioBitrate, err := strconv.Atoi(preset.Audio.Bitrate)
	if err != nil {
		return hwrapper.Preset{}, ErrBitrateNan
	}

	var videoWidth *int
	var videoHeight *int

	if preset.Video.Width != "" {
		var presetWidth int
		presetWidth, err = strconv.Atoi(preset.Video.Width)
		if err != nil {
			return hwrapper.Preset{}, ErrVideoWidthNan
		}
		videoWidth = &presetWidth
	}

	if preset.Video.Height != "" {
		var presetHeight int
		presetHeight, err = strconv.Atoi(preset.Video.Height)
		if err != nil {
			return hwrapper.Preset{}, ErrVideoHeightNan
		}
		videoHeight = &presetHeight
	}

	videoProfile := strings.ToLower(preset.Video.Profile)
	videoLevel := preset.Video.ProfileLevel

	// TODO: Understand video-transcoding-api profile + level settings in relation to vp8
	// For now, we will omit and leave to encoder defaults
	if preset.Video.Codec == "vp8" {
		videoProfile = ""
		videoLevel = ""
	}

	hybrikPreset := hwrapper.Preset{
		Key:         preset.Name,
		Name:        preset.Name,
		Description: preset.Description,
		Kind:        "transcode",
		Path:        p.config.PresetPath,
		Payload: hwrapper.PresetPayload{
			Targets: []hwrapper.PresetTarget{
				{
					FilePattern: "",
					Container: hwrapper.TranscodeContainer{
						Kind: container,
					},
					Video: hwrapper.VideoTarget{
						Width:         videoWidth,
						Height:        videoHeight,
						Codec:         preset.Video.Codec,
						BitrateKb:     bitrate / 1000,
						MinGOPFrames:  minGOPFrames,
						MaxGOPFrames:  maxGOPFrames,
						Profile:       videoProfile,
						Level:         videoLevel,
						InterlaceMode: preset.Video.InterlaceMode,
					},
					Audio: []hwrapper.AudioTarget{
						{
							Codec:     preset.Audio.Codec,
							BitrateKb: audioBitrate / 1000,
						},
					},
					ExistingFiles: "replace",
					UID:           "target",
				},
			},
		},
	}

	for _, modifier := range presetModifiersFor(preset) {
		hybrikPreset, err = modifier.runFunc(hybrikPreset, preset)
		if err != nil {
			return hwrapper.Preset{}, errors.Wrapf(err, "running %q preset modifier", modifier.name)
		}
	}

	return hybrikPreset, nil
}

type presetModifier struct {
	name    string
	runFunc func(hybrikPreset hwrapper.Preset, preset db.Preset) (hwrapper.Preset, error)
}

func presetModifiersFor(preset db.Preset) []presetModifier {
	modifiers := []presetModifier{}

	// HDR
	if _, hdrEnabled := hdrTypeFromPreset(preset); hdrEnabled {
		modifiers = append(modifiers, presetModifier{name: "hdr", runFunc: enrichPresetWithHDRMetadata})
	}

	// MXF sources
	if preset.SourceContainer == "mxf" {
		modifiers = append(modifiers, presetModifier{name: "mxf", runFunc: modifyPresetForMXFSources})
	}

	return modifiers
}

func (p *hybrikProvider) DeletePreset(presetID string) error {
	return p.c.DeletePreset(presetID)
}

func (p *hybrikProvider) GetPreset(presetID string) (interface{}, error) {
	preset, err := p.c.GetPreset(presetID)
	if err != nil {
		return nil, err
	}

	return preset, nil
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
		Destinations:  []string{storageProviderS3, storageProviderGCS},
	}
}
