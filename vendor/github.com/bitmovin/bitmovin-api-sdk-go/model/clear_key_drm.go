package model
import (
	"time"
)

type ClearKeyDrm struct {
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
	Outputs []EncodingOutput `json:"outputs,omitempty"`
	// 16 byte encryption key, 32 hexadecimal characters (required)
	Key string `json:"key,omitempty"`
	// 16 byte key id (required)
	Kid string `json:"kid,omitempty"`
}
func (o ClearKeyDrm) DrmType() DrmType {
    return DrmType_CLEARKEY
}

