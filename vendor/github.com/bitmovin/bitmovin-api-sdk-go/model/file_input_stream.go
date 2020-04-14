package model
import (
	"time"
)

type FileInputStream struct {
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
	// Id of input (required)
	InputId string `json:"inputId,omitempty"`
	// Path to file (required)
	InputPath string `json:"inputPath,omitempty"`
	FileType FileInputStreamType `json:"fileType,omitempty"`
}
func (o FileInputStream) InputStreamType() InputStreamType {
    return InputStreamType_FILE
}

