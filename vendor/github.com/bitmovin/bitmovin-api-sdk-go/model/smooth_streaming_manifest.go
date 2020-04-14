package model
import (
	"time"
)

type SmoothStreamingManifest struct {
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
	Type ManifestType `json:"type,omitempty"`
	// The outputs to store the manifest (required)
	Outputs []EncodingOutput `json:"outputs,omitempty"`
	// Filename of the server manifest
	ServerManifestName string `json:"serverManifestName,omitempty"`
	// Filename of the client manifest
	ClientManifestName string `json:"clientManifestName,omitempty"`
}

