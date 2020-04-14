package model
import (
	"time"
)

type EncodingStats struct {
	// Date, format. yyyy-MM-dd
	Date *time.Time `json:"date,omitempty"`
	// The id of the encoding (required)
	EncodingId string `json:"encodingId,omitempty"`
	// Total bytes encoded
	BytesEncoded *int64 `json:"bytesEncoded,omitempty"`
	// Total time encoded
	TimeEncoded *int64 `json:"timeEncoded,omitempty"`
	// Downloaded size of the input file
	DownloadedSize *int64 `json:"downloadedSize,omitempty"`
	// Billable minutes
	BillableMinutes *float64 `json:"billableMinutes,omitempty"`
	// Detailed statistics per stream
	BillableEncodingMinutes []BillableEncodingMinutes `json:"billableEncodingMinutes,omitempty"`
	// Billable transmuxing minutes (required)
	BillableTransmuxingMinutes *float64 `json:"billableTransmuxingMinutes,omitempty"`
	// Billable feature minutes
	BillableFeatureMinutes *float64 `json:"billableFeatureMinutes,omitempty"`
	// Detailed statistics per stream (required)
	Streams []StatisticsPerStream `json:"streams,omitempty"`
	// Detailed statistics per muxing (required)
	Muxings []StatisticsPerMuxing `json:"muxings,omitempty"`
	// Detailed statistics per feature
	Features []BillableEncodingFeatureMinutes `json:"features,omitempty"`
}

