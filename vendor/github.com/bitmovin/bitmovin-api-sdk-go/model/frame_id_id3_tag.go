package model
import (
	"time"
)

type FrameIdId3Tag struct {
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
	PositionMode Id3TagPositionMode `json:"positionMode,omitempty"`
	// Frame number at which the Tag should be inserted
	Frame *int64 `json:"frame,omitempty"`
	// Time in seconds where the Tag should be inserted
	Time *float64 `json:"time,omitempty"`
	// Base64 Encoded Data (required)
	Bytes string `json:"bytes,omitempty"`
	// 4 character long Frame ID (required)
	FrameId string `json:"frameId,omitempty"`
}
func (o FrameIdId3Tag) Id3TagType() Id3TagType {
    return Id3TagType_FRAME_ID
}

