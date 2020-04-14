package model
import (
	"time"
)

type CustomPlayerBuildDetails struct {
	// Name of the resource. Can be freely chosen by the user.
	Name string `json:"name,omitempty"`
	// Description of the resource. Can be freely chosen by the user.
	Description string `json:"description,omitempty"`
	// Creation timestamp formatted in UTC: YYYY-MM-DDThh:mm:ssZ (required)
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Modified timestamp formatted in UTC: YYYY-MM-DDThh:mm:ssZ (required)
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
	// User-specific meta data. This can hold anything.
	CustomData *map[string]map[string]interface{} `json:"customData,omitempty"`
	// Id of the resource (required)
	Id string `json:"id,omitempty"`
	// The player version that should be used for the custom player build. If not set the 'latest' version is used. (required)
	PlayerVersion string `json:"playerVersion,omitempty"`
	// The domains that the player is locked to. If not set the player will only work with 'localhost'. Not more than 49 additional domains can be added. (required)
	Domains []string `json:"domains,omitempty"`
}

