package model
import (
	"time"
)

type PrewarmEncoderSettings struct {
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
	// Encoder Version to be prewarmed. Only one encoder of this version can be prewarmed per cluster. (required)
	EncoderVersion string `json:"encoderVersion,omitempty"`
	// The minimum number of prewarmed encoders of this Version (required)
	MinPrewarmed *int32 `json:"minPrewarmed,omitempty"`
	// The maximum number of concurrent prewarmed encoders of this Version
	MaxPrewarmed *int32 `json:"maxPrewarmed,omitempty"`
	LogLevel LogLevel `json:"logLevel,omitempty"`
}

