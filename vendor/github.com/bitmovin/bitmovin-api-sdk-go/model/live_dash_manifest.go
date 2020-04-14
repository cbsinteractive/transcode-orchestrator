package model

type LiveDashManifest struct {
	// Dash manifest ids (required)
	ManifestId string `json:"manifestId,omitempty"`
	// Timeshift in seconds
	Timeshift *float64 `json:"timeshift,omitempty"`
	// Live edge offset in seconds
	LiveEdgeOffset *float64 `json:"liveEdgeOffset,omitempty"`
}

