package model
import (
	"time"
)

// An external WebVTT file that is added to an encoding. The size limit for a sidecar file is 10 MB
type WebVttSidecarFile struct {
	// Id of input (required)
	InputId string `json:"inputId,omitempty"`
	// Path to sidecar file (required)
	InputPath string `json:"inputPath,omitempty"`
	Outputs []EncodingOutput `json:"outputs,omitempty"`
	ErrorMode SidecarErrorMode `json:"errorMode,omitempty"`
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
	Segmentation *WebVttSidecarFileSegmentation `json:"segmentation,omitempty"`
}

