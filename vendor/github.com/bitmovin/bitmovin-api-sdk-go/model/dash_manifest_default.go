package model
import (
	"time"
)

type DashManifestDefault struct {
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
	Profile DashProfile `json:"profile,omitempty"`
	// The filename of your manifest
	ManifestName string `json:"manifestName,omitempty"`
	// List of additional XML namespaces to add to the DASH Manifest
	Namespaces []XmlNamespace `json:"namespaces,omitempty"`
	// List of UTC Timings to use for live streaming
	UtcTimings []UtcTiming `json:"utcTimings,omitempty"`
	// The id of the encoding to create a default manifest from (required)
	EncodingId string `json:"encodingId,omitempty"`
	// The version of the default manifest generator
	Version DashManifestDefaultVersion `json:"version,omitempty"`
}

