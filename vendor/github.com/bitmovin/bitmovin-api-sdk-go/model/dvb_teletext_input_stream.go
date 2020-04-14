package model
import (
	"time"
)

type DvbTeletextInputStream struct {
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
	// Id of input
	InputId string `json:"inputId,omitempty"`
	// Path to media file
	InputPath string `json:"inputPath,omitempty"`
	// Specifies the algorithm how the stream in the input file will be selected
	SelectionMode StreamSelectionMode `json:"selectionMode,omitempty"`
	// Position of the stream
	Position *int32 `json:"position,omitempty"`
	// Page number of the subtitles
	Page *int32 `json:"page,omitempty"`
}
func (o DvbTeletextInputStream) InputStreamType() InputStreamType {
    return InputStreamType_DVB_TELETEXT
}

