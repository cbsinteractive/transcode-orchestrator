package model
import (
	"time"
)

type ZixiInput struct {
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
	Host string `json:"host,omitempty"`
	Port *int32 `json:"port,omitempty"`
	Stream string `json:"stream,omitempty"`
	Password string `json:"password,omitempty"`
	Latency *int32 `json:"latency,omitempty"`
	MinBitrate *int32 `json:"minBitrate,omitempty"`
	DecryptionType string `json:"decryptionType,omitempty"`
	DecryptionKey string `json:"decryptionKey,omitempty"`
}
func (o ZixiInput) InputType() InputType {
    return InputType_ZIXI
}

