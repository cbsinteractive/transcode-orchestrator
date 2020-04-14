package model

type Vp9PerTitleConfiguration struct {
	// The minimum bitrate that will be used by the Per-Title algorithm.
	MinBitrate *int32 `json:"minBitrate,omitempty"`
	// The maximum bitrate that will be used by the Per-Title algorithm. It will not generate any rendition with a higher bitrate.
	MaxBitrate *int32 `json:"maxBitrate,omitempty"`
	// The minimum ratio between the bitrates of generated renditions, e.g. if the first bitrate is 240,000, a minimum ratio of 1.5 will require the next higher bitrate to be at least 360,000
	MinBitrateStepSize *float64 `json:"minBitrateStepSize,omitempty"`
	// The maximum ratio between the bitrates of neighbouring renditions, e.g., if the first bitrate is 240,000, a maximum ratio of 1.5 will require the next higher bitrate to be at most 360,000
	MaxBitrateStepSize *float64 `json:"maxBitrateStepSize,omitempty"`
	AutoRepresentations *AutoRepresentation `json:"autoRepresentations,omitempty"`
	// Will modify the assumed complexity for the Per-Title algorithm (> 0.0). Values higher than 1 will increase complexity and thus select smaller resolutions for given bitrates. This will also result in a higher bitrate for the top rendition. Values lower than 1 will decrease assumed complexity and thus select higher resolutions for given bitrates and also decrease the bitrate of the top rendition
	ComplexityFactor *float64 `json:"complexityFactor,omitempty"`
	// Additional configuration for fixed resolution and bitrate templates
	FixedResolutionAndBitrateConfiguration *PerTitleFixedResolutionAndBitrateConfiguration `json:"fixedResolutionAndBitrateConfiguration,omitempty"`
	// Desired target quality of the highest representation expressed as CRF value
	TargetQualityCrf *float64 `json:"targetQualityCrf,omitempty"`
}

