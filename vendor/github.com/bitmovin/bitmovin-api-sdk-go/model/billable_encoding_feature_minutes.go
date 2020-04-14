package model

type BillableEncodingFeatureMinutes struct {
	// The name of the feature.
	FeatureType string `json:"featureType,omitempty"`
	// Encoded minutes related to this feature.
	EncodedMinutes *float64 `json:"encodedMinutes,omitempty"`
	// The multiplier used for this feature.
	FeatureMultiplier *float64 `json:"featureMultiplier,omitempty"`
	// The billable minutes related to this feature.
	BillableMinutes *float64 `json:"billableMinutes,omitempty"`
}

