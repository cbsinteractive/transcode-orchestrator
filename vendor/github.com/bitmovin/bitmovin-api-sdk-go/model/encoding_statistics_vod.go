package model
import (
	"time"
)

type EncodingStatisticsVod struct {
	// Date, format. yyyy-MM-dd (required)
	Date *time.Time `json:"date,omitempty"`
	// Bytes encoded for this encoding. (required)
	BytesEncoded *int64 `json:"bytesEncoded,omitempty"`
	// Time in seconds encoded for this encoding. (required)
	TimeEncoded *int64 `json:"timeEncoded,omitempty"`
	// Time in seconds encoded for this encoding. (required)
	TimeEnqueued *int64 `json:"timeEnqueued,omitempty"`
	// The realtime factor. (required)
	RealTimeFactor *float64 `json:"realTimeFactor,omitempty"`
}

