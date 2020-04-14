package model

type StreamPerTitleSettings struct {
	// Settings for PER_TITLE_TEMPLATE_FIXED_RESOLUTION_AND_BITRATE mode
	FixedResolutionAndBitrateSettings *StreamPerTitleFixedResolutionAndBitrateSettings `json:"fixedResolutionAndBitrateSettings,omitempty"`
}

