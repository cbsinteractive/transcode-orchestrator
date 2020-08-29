package transcoding

import "github.com/cbsinteractive/pkg/video"

// PresetName is a custom string type with the name of the preset
type PresetName string

// CreatePresetResponse contains the results of requests to generate new presets
type CreatePresetResponse struct {
	Results   map[string]NewPresetSummary
	PresetMap string
}

// NewPresetSummary contains preset creation results for a single provider
type NewPresetSummary struct {
	PresetID string
	Error    string
}

// DeletePresetResponse contains the results of requests to remove a preset
type DeletePresetResponse struct {
	Results   map[string]DeletePresetSummary `json:"results"`
	PresetMap string                         `json:"presetMap"`
}

// DeletePresetSummary contains preset deletion results for a single provider
type DeletePresetSummary struct {
	PresetID string `json:"presetId"`
	Error    string `json:"error,omitempty"`
}

// CreatePresetRequest represents the request body structure when creating a preset
type CreatePresetRequest struct {
	Providers     []string      `json:"providers"`
	Preset        Preset        `json:"preset"`
	OutputOptions OutputOptions `json:"outputOptions"`
}

// Preset defines the set of parameters of a given preset
type Preset struct {
	Name            PresetName  `json:"name,omitempty"`
	Description     string      `json:"description,omitempty"`
	SourceContainer string      `json:"sourceContainer,omitempty"`
	Container       string      `json:"container,omitempty"`
	RateControl     string      `json:"rateControl,omitempty"`
	TwoPass         bool        `json:"twoPass"`
	Video           VideoPreset `json:"video"`
	Audio           AudioPreset `json:"audio"`
}

// OutputOptions is the set of options for the output file.
type OutputOptions struct {
	Extension string `json:"extension"`
}

// VideoPreset defines the set of parameters for video on a given preset
type VideoPreset struct {
	Profile             string              `json:"profile,omitempty"`
	ProfileLevel        string              `json:"profileLevel,omitempty"`
	Width               string              `json:"width,omitempty"`
	Height              string              `json:"height,omitempty"`
	Codec               string              `json:"codec,omitempty"`
	Bitrate             string              `json:"bitrate,omitempty"`
	GopSize             string              `json:"gopSize,omitempty"`
	GopUnit             string              `json:"gopUnit,omitempty"`
	GopMode             string              `json:"gopMode,omitempty"`
	InterlaceMode       string              `json:"interlaceMode,omitempty"`
	HDR10Settings       HDR10Settings       `json:"hdr10"`
	DolbyVisionSettings DolbyVisionSettings `json:"dolbyVision"`
	Overlays            *Overlays           `json:"overlays,omitempty"`

	// Crop contains offsets for top, bottom, left and right src cropping
	Crop *video.Crop `json:"crop,omitempty" redis-hash:"crop,expand,omitempty"`
}

//Overlays defines all the overlay settings for a Video preset
type Overlays struct {
	Images         []Image         `json:"images,omitempty"`
	TimecodeBurnin *TimecodeBurnin `json:"timecodeBurnin,omitempty"`
}

//Image defines the image overlay settings
type Image struct {
	URL string `json:"url"`
}

//TimecodeBurnin defines the timecode burnin settings
type TimecodeBurnin struct {
	Enabled  bool   `json:"enabled"`
	FontSize int    `json:"fontSize,omitempty"`
	Position int    `json:"position,omitempty"`
	Prefix   string `json:"prefix,omitempty"`
}

// HDR10Settings defines a set of configurations for defining HDR10 metadata
type HDR10Settings struct {
	Enabled       bool   `json:"enabled"`
	MaxCLL        uint   `json:"maxCLL,omitempty"`
	MaxFALL       uint   `json:"maxFALL,omitempty"`
	MasterDisplay string `json:"masterDisplay,omitempty"`
}

// DolbyVisionSettings defines a set of configurations for setting DolbyVision metadata
type DolbyVisionSettings struct {
	Enabled bool `json:"enabled" redis-hash:"enabled"`
}

// AudioPreset defines the set of parameters for audio on a given preset
type AudioPreset struct {
	Codec   string `json:"codec,omitempty"`
	Bitrate string `json:"bitrate,omitempty"`
}
