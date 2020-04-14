package model
import (
	"time"
)

type Drm interface {
    DrmType() DrmType
}

type BaseDrm struct {
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
    Outputs []EncodingOutput `json:"outputs,omitempty"`
    Type DrmType `json:"type"`
}

func (b BaseDrm) DrmType() DrmType {
    return b.Type
}

