package db

import (
	"errors"
	"time"
)

// Job represents the job that is persisted in the repository of the Transcoding
// API.
//
// swagger:model
type Job struct {
	// id of the job. It's automatically generated by the API when creating
	// a new Job.
	//
	// unique: true
	ID string `redis-hash:"jobID" json:"jobId"`

	// name of the provider
	//
	// required: true
	ProviderName string `redis-hash:"providerName" json:"providerName"`

	// id of the job on the provider
	//
	// required: true
	ProviderJobID string `redis-hash:"providerJobID" json:"providerJobId"`

	// configuration for adaptive streaming jobs
	// Defaults to false.
	//
	// required: false
	StreamingParams StreamingParams `redis-hash:"streamingparams,expand" json:"streamingParams,omitempty"`

	// ExecutionEnv contains configurations for the environment used while transcoding
	//
	// required: false
	ExecutionEnv ExecutionEnvironment `redis-hash:"executionenvironment,expand" json:"executionEnv,omitempty"`

	// configuration for execution features for the selected provider
	//
	// required: false
	ExecutionFeatures ExecutionFeatures `redis-hash:"-" json:"executionFeatures,omitempty"`

	// string value of the execution config for auditing jobs after the fact
	//
	// required: false
	ExecutionCfgReport string `redis-hash:"execution-cfg,omitempty" json:"executionCfgReport,omitempty"`

	// Time of the creation of the job in the API
	//
	// required: true
	CreationTime time.Time `redis-hash:"creationTime" json:"creationTime"`

	// Source of the job
	//
	// required: true
	SourceMedia string `redis-hash:"source" json:"source"`

	// Base Destination of the job
	//
	// required: false
	DestinationBasePath string `redis-hash:"destbasepath,omitempty" json:"destinationBasePath,omitempty"`

	// SidecarAssets contain a map of string keys to file locations
	//
	// required: false
	SidecarAssets map[SidecarAssetKind]string `redis-hash:"sidecarassets,omitempty,expand" json:"sidecarAssets,omitempty"`

	// Output list of the given job
	//
	// required: true
	Outputs []TranscodeOutput `redis-hash:"-" json:"outputs"`
}

type SidecarAssetKind = string

const SidecarAssetKindDolbyVisionMetadata SidecarAssetKind = "dolbyVisionMetadata"

// ExecutionEnvironment contains configurations for the environment used while transcoding
type ExecutionEnvironment struct {
	Cloud       string                  `redis-hash:"cloud,omitempty" json:"cloud"`
	Region      string                  `redis-hash:"region,omitempty" json:"region"`
	ComputeTags map[ComputeClass]string `redis-hash:"computetags,omitempty,expand" json:"computeTags,omitempty"`
	InputAlias  string                  `redis-hash:"inputalias,omitempty" json:"inputAlias,omitempty"`
	OutputAlias string                  `redis-hash:"outputalias,omitempty" json:"outputAlias,omitempty"`
}

// ComputeClass represents a group of resources with similar capability
type ComputeClass = string

const (
	// ComputeClassTranscodeDefault runs any default transcodes
	ComputeClassTranscodeDefault ComputeClass = "transcodeDefault"

	// ComputeClassDolbyVisionTranscode runs Dolby Vision transcodes
	ComputeClassDolbyVisionTranscode ComputeClass = "doViTranscode"

	// ComputeClassDolbyVisionPreprocess runs Dolby Vision pre-processing
	ComputeClassDolbyVisionPreprocess ComputeClass = "doViPreprocess"

	// ComputeClassDolbyVisionPreprocess runs Dolby Vision pre-processing
	ComputeClassDolbyVisionMezzQC ComputeClass = "doViMezzQC"
)

// TranscodeOutput represents a transcoding output. It's a combination of the
// preset and the output file name.
type TranscodeOutput struct {
	// Presetmap for the output
	//
	// required: true
	Preset PresetMap `redis-hash:"presetmap,expand" json:"presetmap"`

	// Filename for the output
	//
	// required: true
	FileName string `redis-hash:"filename" json:"filename"`
}

// StreamingParams represents the params necessary to create Adaptive Streaming jobs
//
// swagger:model
type StreamingParams struct {
	// duration of the segment
	//
	// required: true
	SegmentDuration uint `redis-hash:"segmentDuration" json:"segmentDuration"`

	// the protocol name (hls or dash)
	//
	// required: true
	Protocol string `redis-hash:"protocol" json:"protocol"`

	// the playlist file name
	// required: true
	PlaylistFileName string `redis-hash:"playlistFileName" json:"playlistFileName,omitempty"`
}

// ExecutionFeatures is a map whose key is a custom feature name and value is a json string
// representing the corresponding custom feature definition
type ExecutionFeatures map[string]interface{}

// PresetSummary holds references to external resources that represent the configurations
// of audio and video streams and their containers
//
// swagger:model
type PresetSummary struct {
	Name          string `redis-hash:"-"`
	Container     string `redis-hash:"container"`
	VideoCodec    string `redis-hash:"videocodec,omitempty"`
	VideoConfigID string `redis-hash:"videoconfigid,omitempty"`
	AudioCodec    string `redis-hash:"audiocodec,omitempty"`
	AudioConfigID string `redis-hash:"audioconfigid,omitempty"`
}

// LocalPreset is a struct to persist encoding configurations. Some providers don't have
// the ability to store presets on it's side so we persist locally.
//
// swagger:model
type LocalPreset struct {
	// name of the local preset
	//
	// unique: true
	// required: true
	Name string `redis-hash:"-" json:"name"`

	// the preset structure
	// required: true
	Preset Preset `redis-hash:"preset,expand" json:"preset"`
}

// Preset defines the set of parameters of a given preset
type Preset struct {
	Name            string      `json:"name,omitempty" redis-hash:"name"`
	Description     string      `json:"description,omitempty" redis-hash:"description,omitempty"`
	SourceContainer string      `json:"sourceContainer,omitempty" redis-hash:"sourcecontainer,omitempty"`
	Container       string      `json:"container,omitempty" redis-hash:"container,omitempty"`
	RateControl     string      `json:"rateControl,omitempty" redis-hash:"ratecontrol,omitempty"`
	TwoPass         bool        `json:"twoPass" redis-hash:"twopass"`
	Video           VideoPreset `json:"video" redis-hash:"video,expand"`
	Audio           AudioPreset `json:"audio" redis-hash:"audio,expand"`
}

// VideoPreset defines the set of parameters for video on a given preset
type VideoPreset struct {
	Profile             string              `json:"profile,omitempty" redis-hash:"profile,omitempty"`
	ProfileLevel        string              `json:"profileLevel,omitempty" redis-hash:"profilelevel,omitempty"`
	Width               string              `json:"width,omitempty" redis-hash:"width,omitempty"`
	Height              string              `json:"height,omitempty" redis-hash:"height,omitempty"`
	Codec               string              `json:"codec,omitempty" redis-hash:"codec,omitempty"`
	Bitrate             string              `json:"bitrate,omitempty" redis-hash:"bitrate,omitempty"`
	GopSize             string              `json:"gopSize,omitempty" redis-hash:"gopsize,omitempty"`
	GopUnit             string              `json:"gopUnit,omitempty" redis-hash:"gopunit,omitempty"`
	GopMode             string              `json:"gopMode,omitempty" redis-hash:"gopmode,omitempty"`
	InterlaceMode       string              `json:"interlaceMode,omitempty" redis-hash:"interlacemode,omitempty"`
	HDR10Settings       HDR10Settings       `json:"hdr10" redis-hash:"hdr10,expand,omitempty"`
	DolbyVisionSettings DolbyVisionSettings `json:"dolbyVision" redis-hash:"dolbyvision,expand,omitempty"`
}

// GopUnit defines the unit used to measure gops
type GopUnit = string

const (
	// GopUnitFrames uses Gop Frames in transcode job
	GopUnitFrames GopUnit = "frames"

	// GopUnitSeconds uses Key Intervals in transcode job
	GopUnitSeconds GopUnit = "seconds"
)

// HDR10Settings defines a set of configurations for defining HDR10 metadata
type HDR10Settings struct {
	Enabled       bool   `json:"enabled" redis-hash:"enabled"`
	MaxCLL        uint   `json:"maxCLL,omitempty" redis-hash:"maxcll,omitempty"`
	MaxFALL       uint   `json:"maxFALL,omitempty" redis-hash:"maxfll,omitempty"`
	MasterDisplay string `json:"masterDisplay,omitempty" redis-hash:"masterdisplay,omitempty"`
}

// DolbyVisionSettings defines a set of configurations for setting DolbyVision metadata
type DolbyVisionSettings struct {
	Enabled bool `json:"enabled" redis-hash:"enabled"`
}

// AudioPreset defines the set of parameters for audio on a given preset
type AudioPreset struct {
	Codec   string `json:"codec,omitempty" redis-hash:"codec,omitempty"`
	Bitrate string `json:"bitrate,omitempty" redis-hash:"bitrate,omitempty"`
}

// PresetMap represents the preset that is persisted in the repository of the
// Transcoding API
//
// Each presetmap is just an aggregator of provider presets, where each preset in
// the API maps to a preset on each provider
//
// swagger:model
type PresetMap struct {
	// name of the presetmap
	//
	// unique: true
	// required: true
	Name string `redis-hash:"presetmap_name" json:"name"`

	// mapping of provider name to provider's internal preset id.
	//
	// required: true
	ProviderMapping map[string]string `redis-hash:"pmapping,expand" json:"providerMapping"`

	// set of options in the output file for this preset.
	//
	// required: true
	OutputOpts OutputOptions `redis-hash:"output,expand" json:"output"`
}

// OutputOptions is the set of options for the output file.
//
// This type includes only configuration parameters that are not defined in
// providers (like the extension of the output file).
//
// swagger:model
type OutputOptions struct {
	// extension for the output file, it's usually attached to the
	// container (for example, webm for VP, mp4 for MPEG-4 and ts for HLS).
	//
	// The dot should not be part of the extension, i.e. use "webm" instead
	// of ".webm".
	//
	// required: true
	Extension string `redis-hash:"extension" json:"extension"`
}

// Validate checks that the OutputOptions object is properly defined.
func (o *OutputOptions) Validate() error {
	if o.Extension == "" {
		return errors.New("extension is required")
	}
	return nil
}
