package model

type PerTitleFixedResolutionAndBitrateConfiguration struct {
	// Number of forced renditions above the highest fixed representation (e.g. FIXED_RESOLUTION_AND_BITRATE). Forced renditions will be added, also if the Per-Title algorithm chooses the user defined force rendition to be the highest one.
	ForcedRenditionAboveHighestFixedRepresentation *int32 `json:"forcedRenditionAboveHighestFixedRepresentation,omitempty"`
	// The factor to calculate the bitrate that will be chosen based on the bitrate of the last FIXED_RESOLUTION.
	ForcedRenditionAboveHighestFixedRepresentationFactor *float64 `json:"forcedRenditionAboveHighestFixedRepresentationFactor,omitempty"`
	// Mode to calculate the bitrate of the next representation
	ForcedRenditionAboveHighestFixedRepresentationCalculationMode PerTitleFixedResolutionAndBitrateConfigurationMode `json:"forcedRenditionAboveHighestFixedRepresentationCalculationMode,omitempty"`
}

