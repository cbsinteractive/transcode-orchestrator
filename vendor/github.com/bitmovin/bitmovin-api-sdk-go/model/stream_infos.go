package model
import (
	"time"
)

type StreamInfos struct {
	// Timestamp of the event formatted in UTC: YYYY-MM-DDThh:mm:ssZ (required)
	Time *time.Time `json:"time,omitempty"`
	// Details about billable minutes for each resolution category
	StreamInfos []StreamInfosDetails `json:"streamInfos,omitempty"`
}

