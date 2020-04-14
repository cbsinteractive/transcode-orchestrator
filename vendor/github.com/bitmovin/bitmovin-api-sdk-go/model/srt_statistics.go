package model
import (
	"time"
)

type SrtStatistics struct {
	// UUID of the statistic event
	Id string `json:"id,omitempty"`
	// Timestamp when the srt statistics event was created, formatted in UTC: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// UUID of an encoding
	EncodingId string `json:"encodingId,omitempty"`
	// UUID of the SRT input used to capture this statistics
	SrtInputId string `json:"srtInputId,omitempty"`
	// Whether the SRT input that generated this statistics was selected (i.e. actively used) at the time or not
	SrtInputSelected *bool `json:"srtInputSelected,omitempty"`
	// UUID of the associated organization
	OrgId string `json:"orgId,omitempty"`
	// UUID of the associated api-user
	UserId string `json:"userId,omitempty"`
	Link *SrtStatisticLink `json:"link,omitempty"`
	Window *SrtStatisticWindow `json:"window,omitempty"`
	Recv *SrtStatisticRecv `json:"recv,omitempty"`
	Send *SrtStatisticSend `json:"send,omitempty"`
}

