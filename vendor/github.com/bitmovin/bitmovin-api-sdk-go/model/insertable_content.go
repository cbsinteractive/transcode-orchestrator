package model
import (
	"time"
)

type InsertableContent struct {
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
	// List of video files to be inserted in the live stream. These have to match the codec, aspect ration and frame rate of the live stream.
	Inputs []InsertableContentInput `json:"inputs,omitempty"`
	// Status of the insertable content.
	Status InsertableContentStatus `json:"status,omitempty"`
}

