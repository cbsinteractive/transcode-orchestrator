package model

type StatisticsPerLabel struct {
	// Bytes encoded total. (required)
	BytesEncodedTotal *int64 `json:"bytesEncodedTotal,omitempty"`
	// Time in seconds encoded for all contained daily statistics. (required)
	TimeEncodedTotal *int64 `json:"timeEncodedTotal,omitempty"`
	// An optional error message, when the event is in error state (occurs at event: ERROR) (required)
	Label string `json:"label,omitempty"`
	// The billable minutes.
	BillableMinutes *float64 `json:"billableMinutes,omitempty"`
	// Billable minutes for each encoding configuration
	BillableEncodingMinutes []BillableEncodingMinutes `json:"billableEncodingMinutes,omitempty"`
	// Billable minutes for muxings.
	BillableTransmuxingMinutes *float64 `json:"billableTransmuxingMinutes,omitempty"`
	// Billable minutes for features
	BillableFeatureMinutes []BillableEncodingFeatureMinutes `json:"billableFeatureMinutes,omitempty"`
}

