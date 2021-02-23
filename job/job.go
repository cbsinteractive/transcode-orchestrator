package job

import (
	"errors"
	"time"

	"github.com/cbsinteractive/pkg/timecode"
	"github.com/cbsinteractive/pkg/video"
	"github.com/gofrs/uuid"
)

// Job is a transcoding job
type Job struct {
	ID string `json:"jobId"`

	Name         string
	CreationTime time.Time
	Labels []string

	SourceMedia string `json:"source"`
	SourceInfo  File
	// SourceSplice is a set of second ranges to excise from the input and catenate
	// together before processing the source. For example, [[0,1],[8,9]], will cut out
	// a two-second clip, from the first and last second of a 10s video.
	SourceSplice timecode.Splice `json:"splice,omitempty"`
	ProviderName string          `json:"provider"`

	ProviderJobID string

	// configuration for adaptive streaming jobs
	StreamingParams StreamingParams

	// ExecutionEnv contains configurations for the environment used while transcoding
	ExecutionFeatures  ExecutionFeatures
	ExecutionEnv       ExecutionEnvironment
	ExecutionCfgReport string

	DestinationBasePath string

	// SidecarAssets contain a map of string keys to file locations
	SidecarAssets map[SidecarAssetKind]string

	// Output list of the given job
	Outputs []TranscodeOutput

	// AudioDownmix holds source and output channels for configuring downmixing
	AudioDownmix *AudioDownmix

	// ExplicitKeyframeOffsets define offsets from the beginning of the media to insert keyframes when encoding
	ExplicitKeyframeOffsets []float64
}

// State is the state of a transcoding job.
type State string

const (
	StateUnknown  = State("unknown")
	StateQueued   = State("queued")
	StateStarted  = State("started")
	StateFinished = State("finished")
	StateFailed   = State("failed")
	StateCanceled = State("canceled")
)

// Status is the representation of the status
type Status struct {
	ID      string   `json:"jobID,omitempty"`
	Labels  []string `json:"labels,omitempty"`
	State   State    `json:"status,omitempty"`
	Progress float64 `json:"progress"`
	Message string   `json:"statusMessage,omitempty"`

	Input    File    `json:"sourceInfo,omitempty"`
	Output   Output  `json:"output"`

	ProviderName   string                 `json:"providerName,omitempty"`
	ProviderJobID  string                 `json:"providerJobId,omitempty"`
	ProviderStatus map[string]interface{} `json:"providerStatus,omitempty"`
}

// Output represents information about a job output.
type Output struct {
	Destination string `json:"destination,omitempty"`
	Files       []File `json:"files,omitempty"`
}

func (j Job) RootFolder() string {
	if j.Name != "" {
		if _, err := uuid.FromString(j.Name); err == nil {
			return j.Name
		}
	}

	return j.ID
}

type SidecarAssetKind = string

const SidecarAssetKindDolbyVisionMetadata SidecarAssetKind = "dolbyVisionMetadata"

// ExecutionEnvironment contains configurations for the environment used while transcoding
type ExecutionEnvironment struct {
	Cloud       string
	Region      string
	ComputeTags map[ComputeClass]string
	InputAlias  string
	OutputAlias string
}

// ComputeClass represents a group of resources with similar capability
type ComputeClass = string

// ComputeClassTranscodeDefault runs any default transcodes
// ComputeClassDolbyVisionTranscode runs Dolby Vision transcodes
// ComputeClassDolbyVisionPreprocess runs Dolby Vision pre-processing
// ComputeClassDolbyVisionMezzQC runs QC check on the mezzanine
const (
	ComputeClassTranscodeDefault      ComputeClass = "transcodeDefault"
	ComputeClassDolbyVisionTranscode  ComputeClass = "doViTranscode"
	ComputeClassDolbyVisionPreprocess ComputeClass = "doViPreprocess"
	ComputeClassDolbyVisionMezzQC     ComputeClass = "doViMezzQC"
)

// TranscodeOutput represents a transcoding output. It's a combination of the
// preset and the output file name.
type TranscodeOutput struct {
	Preset   Preset
	FileName string
}

// StreamingParams represents the params necessary to create Adaptive Streaming jobs
type StreamingParams struct {
	SegmentDuration  uint
	Protocol         string
	PlaylistFileName string
}

// ScanType is a string that represents the scan type of the content.
type ScanType string

// ScanTypeProgressive and other supported types
const (
	ScanTypeProgressive ScanType = "progressive"
	ScanTypeInterlaced  ScanType = "interlaced"
	ScanTypeUnknown     ScanType = "unknown"
)

//ChannelLayout describes layout of an audio channel
type ChannelLayout string

const (
	ChannelLayoutCenter        ChannelLayout = "C"
	ChannelLayoutLeft          ChannelLayout = "L"
	ChannelLayoutRight         ChannelLayout = "R"
	ChannelLayoutLeftSurround  ChannelLayout = "Ls"
	ChannelLayoutRightSurround ChannelLayout = "Rs"
	ChannelLayoutLeftBack      ChannelLayout = "Lb"
	ChannelLayoutRightBack     ChannelLayout = "Rb"
	ChannelLayoutLeftTotal     ChannelLayout = "Lt"
	ChannelLayoutRightTotal    ChannelLayout = "Rt"
	ChannelLayoutLFE           ChannelLayout = "LFE"
)

// AudioChannel describes the position and attributes of a
// single channel of audio inside a container
type AudioChannel struct {
	TrackIdx, ChannelIdx int
	Layout               string
}

//AudioDownmix holds source and output channels for providers
//to handle downmixing
type AudioDownmix struct {
	SrcChannels  []AudioChannel
	DestChannels []AudioChannel
}

// File
type File struct {
	Path      string `json:"path"`
	Container string `json:"container"`
	Size      int64  `json:"fileSize,omitempty"`
	Duration   time.Duration `json:"duration,omitempty"`
	Height     int         `json:"height,omitempty"`
	Width      int        `json:"width,omitempty"`
	VideoCodec string        `json:"videoCodec,omitempty"`
	FrameRate float64
	ScanType ScanType
}

// ExecutionFeatures is a map whose key is a custom feature name and value is a json string
// representing the corresponding custom feature definition
type ExecutionFeatures map[string]interface{}


// LocalPreset is a struct to persist encoding configurations. Some providers don't have
// the ability to store presets on it's side so we persist locally.
type LocalPreset struct {
	Name   string
	Preset Preset
}

// Preset defines the set of parameters of a given preset
type Preset struct {
	Name            string `json:"name,omitempty"`
	Description     string `json:"description,omitempty"`
	SourceContainer string `json:"sourceContainer,omitempty"`
	Container       string `json:"container,omitempty"`
	RateControl     string `json:"rateControl,omitempty"`
	TwoPass         bool   `json:"twoPass"`
	Video           Video  `json:"video"`
	Audio           Audio  `json:"audio"`
}

// Video transcoding parameters
type Video struct {
	Codec   string `json:"codec,omitempty"`
	Profile string `json:"profile,omitempty"`
	Level   string `json:"profileLevel,omitempty"`

	Width   int `json:"width,omitempty"`
	Height  int `json:"height,omitempty"`
	Bitrate int `json:"bitrate,omitempty"`

	GopSize       float64 `json:"gopSize,omitempty"`
	GopUnit       string  `json:"gopUnit,omitempty"`
	GopMode       string  `json:"gopMode,omitempty"`
	InterlaceMode string  `json:"interlaceMode,omitempty"`

	HDR10       HDR10       `json:"hdr10"`
	DolbyVision DolbyVision `json:"dolbyVision"`
	Overlays            *Overlays           `json:"overlays,omitempty"`

	// Crop contains offsets for top, bottom, left and right src cropping
	Crop video.Crop `json:"crop"`
}

// GopUnit defines the unit used to measure gops
type GopUnit = string

const (
	GopUnitFrames  GopUnit = "frames"
	GopUnitSeconds GopUnit = "seconds"
)

//Overlays defines all the overlay settings for a Video preset
type Overlays struct {
	Images         []Image         `json:"images,omitempty"`
	TimecodeBurnin *TimecodeBurnin `json:"timecodeBurnin,omitempty"`
}

//Image defines the image overlay settings
type Image struct {
	URL string `json:"url"`
}

// TimecodeBurnin settings
type TimecodeBurnin struct {
	Enabled  bool   `json:"enabled"`
	FontSize int    `json:"fontSize,omitempty"`
	Position int    `json:"position,omitempty"`
	Prefix   string `json:"prefix,omitempty"`
}

// HDR10 configurations and metadata
type HDR10 struct {
	Enabled       bool   `json:"enabled"`
	MaxCLL        uint   `json:"maxCLL,omitempty"`
	MaxFALL       uint   `json:"maxFALL,omitempty"`
	MasterDisplay string `json:"masterDisplay,omitempty"`
}

// DolbyVision settings
type DolbyVision struct {
	Enabled bool `json:"enabled"`
}

// Audio defines audio transcoding parameters
type Audio struct {
	Codec          string `json:"codec,omitempty"`
	Bitrate        int    `json:"bitrate,omitempty"`
	Normalization  bool   `json:"normalization,omitempty"`
	DiscreteTracks bool   `json:"discreteTracks,omitempty"`
}

type OutputOptions struct {
	Extension string
}

// Validate checks that the OutputOptions object is properly defined.
func (o *OutputOptions) Validate() error {
	if o.Extension == "" {
		return errors.New("extension is required")
	}
	return nil
}
