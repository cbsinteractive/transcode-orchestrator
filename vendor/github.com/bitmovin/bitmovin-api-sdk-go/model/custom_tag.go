package model
import (
	"time"
)

type CustomTag struct {
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
	// The positioning mode that should be used when inserting the placement opportunity (required)
	PositionMode PositionMode `json:"positionMode,omitempty"`
	// Id of keyframe where the custom tag should be inserted. Required, when KEYFRAME is selected as position mode.
	KeyframeId string `json:"keyframeId,omitempty"`
	// Time in seconds where the custom tag should be inserted. Required, when TIME is selected as position mode.
	Time *float64 `json:"time,omitempty"`
	// The custom tag will be inserted before the specified segment. Required, when SEGMENT is selected as position mode.
	Segment *int64 `json:"segment,omitempty"`
	// The data to be contained in the custom tag. (required)
	Data string `json:"data,omitempty"`
}

