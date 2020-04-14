package model
import (
	"time"
)

type DailyStatistics struct {
	// Date for the shown data. Format: yyyy-MM-dd (required)
	Date *time.Time `json:"date,omitempty"`
	// Bytes encoded. (required)
	BytesEncoded *int64 `json:"bytesEncoded,omitempty"`
	// Time in seconds encoded for this day. (required)
	TimeEncoded *int64 `json:"timeEncoded,omitempty"`
	// The billable minutes.
	BillableMinutes *float64 `json:"billableMinutes,omitempty"`
	// Label identifier.
	Label string `json:"label,omitempty"`
	// Billable minutes for each encoding configuration.
	BillableEncodingMinutes []BillableEncodingMinutes `json:"billableEncodingMinutes,omitempty"`
	// Billable minutes for muxings.
	BillableTransmuxingMinutes *float64 `json:"billableTransmuxingMinutes,omitempty"`
	// Billable minutes for features
	BillableFeatureMinutes []BillableEncodingFeatureMinutes `json:"billableFeatureMinutes,omitempty"`
}

