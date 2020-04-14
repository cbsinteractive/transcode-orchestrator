package model
import (
	"time"
)

type AsperaInput struct {
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
	// Minimal download bandwidth. Examples: 100k, 100m, 100g
	MinBandwidth string `json:"minBandwidth,omitempty"`
	// Maximal download bandwidth. Examples: 100k, 100m, 100g
	MaxBandwidth string `json:"maxBandwidth,omitempty"`
	// Host to use for Aspera transfers (required)
	Host string `json:"host,omitempty"`
	// Username to log into Aspera host (either password and user must be set or token)
	Username string `json:"username,omitempty"`
	// corresponding password (either password and user must be set or token)
	Password string `json:"password,omitempty"`
	// Token used for authentication (either password and user must be set or token)
	Token string `json:"token,omitempty"`
}
func (o AsperaInput) InputType() InputType {
    return InputType_ASPERA
}

