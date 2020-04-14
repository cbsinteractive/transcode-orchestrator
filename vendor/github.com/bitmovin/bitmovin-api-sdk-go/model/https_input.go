package model
import (
	"time"
)

type HttpsInput struct {
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
	// Host Url or IP of the HTTP server (required)
	Host string `json:"host,omitempty"`
	// Basic Auth Username, if required
	Username string `json:"username,omitempty"`
	// Basic Auth Password, if required
	Password string `json:"password,omitempty"`
	// Custom Port
	Port *int32 `json:"port,omitempty"`
}
func (o HttpsInput) InputType() InputType {
    return InputType_HTTPS
}

