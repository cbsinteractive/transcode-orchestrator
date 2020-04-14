package model

type BillableEncodingMinutesDetails struct {
	// Only set if resolution information is not present.
	UNKNOWN *float64 `json:"UNKNOWN,omitempty"`
	// Billable minutes for audio. Available if stream is an audio stream.
	AUDIO *float64 `json:"AUDIO,omitempty"`
	// Billable minutes for SD resolutions.
	SD *float64 `json:"SD,omitempty"`
	// Billable minutes for HD resolutions.
	HD *float64 `json:"HD,omitempty"`
	// Billable minutes for UHD resolutions.
	UHD *float64 `json:"UHD,omitempty"`
}

