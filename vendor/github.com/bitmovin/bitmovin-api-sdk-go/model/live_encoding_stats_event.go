package model
import (
	"time"
)

type LiveEncodingStatsEvent struct {
	// Timestamp of the event formatted in UTC: YYYY-MM-DDThh:mm:ssZ (required)
	Time *time.Time `json:"time,omitempty"`
	Details *LiveEncodingStatsEventDetails `json:"details,omitempty"`
}

