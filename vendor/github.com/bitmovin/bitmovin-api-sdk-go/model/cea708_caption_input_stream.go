package model
import (
	"time"
)

type Cea708CaptionInputStream struct {
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
	// Id of the Input (required)
	InputId string `json:"inputId,omitempty"`
	// Path to media file (required)
	InputPath string `json:"inputPath,omitempty"`
	// The channel number of the subtitle on the respective stream position. Must not be smaller than 1 (required)
	Channel *int32 `json:"channel,omitempty"`
}
func (o Cea708CaptionInputStream) InputStreamType() InputStreamType {
    return InputStreamType_CAPTION_CEA708
}

