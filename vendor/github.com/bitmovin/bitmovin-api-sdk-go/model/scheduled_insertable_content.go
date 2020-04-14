package model
import (
	"time"
)

type ScheduledInsertableContent struct {
	// Name of the resource. Can be freely chosen by the user.
	Name string `json:"name,omitempty"`
	// Description of the resource. Can be freely chosen by the user.
	Description string `json:"description,omitempty"`
	// Creation timestamp formatted in UTC: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Modified timestamp formatted in UTC: YYYY-MM-DDThh:mm:ssZ
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
	// User-specific meta data. This can hold anything.
	CustomData *map[string]map[string]interface{} `json:"customData,omitempty"`
	// Id of the resource (required)
	Id string `json:"id,omitempty"`
	// Id of the insertable content to play instead of the live stream
	ContentId string `json:"contentId,omitempty"`
	// Time to to play the content in UTC: YYYY-MM-DDThh:mm:ssZ
	RunAt *time.Time `json:"runAt,omitempty"`
	// Duration for how long to play the content. Cut off if shorter, loop if longer than actual duration.
	DurationInSeconds *float64 `json:"durationInSeconds,omitempty"`
	// Status of the scheduled insertable content.
	Status ScheduledInsertableContentStatus `json:"status,omitempty"`
}

