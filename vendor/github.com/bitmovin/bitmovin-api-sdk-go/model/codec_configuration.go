package model
import (
	"time"
)

type CodecConfiguration interface {
    CodecConfigType() CodecConfigType
}

type BaseCodecConfiguration struct {
    // Name of the resource. Can be freely chosen by the user. (required)
    Name string `json:"name"`
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
    Type CodecConfigType `json:"type"`
}

func (b BaseCodecConfiguration) CodecConfigType() CodecConfigType {
    return b.Type
}

