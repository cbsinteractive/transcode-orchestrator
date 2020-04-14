package model
import (
	"time"
)

type Id3Tag interface {
    Id3TagType() Id3TagType
}

type BaseId3Tag struct {
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
    Id string `json:"id"`
    PositionMode Id3TagPositionMode `json:"positionMode,omitempty"`
    // Frame number at which the Tag should be inserted
    Frame *int64 `json:"frame,omitempty"`
    // Time in seconds where the Tag should be inserted
    Time *float64 `json:"time,omitempty"`
    Type Id3TagType `json:"type"`
}

func (b BaseId3Tag) Id3TagType() Id3TagType {
    return b.Type
}

