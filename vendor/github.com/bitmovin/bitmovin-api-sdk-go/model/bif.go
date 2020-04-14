package model
import (
	"time"
)

// Either height or width is required. It is also possible to set both properties.
type Bif struct {
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
	// Height of one thumbnail
	Height *int32 `json:"height,omitempty"`
	// Width of one thumbnail. Roku recommends a width of 240 for SD and 320 for HD.
	Width *int32 `json:"width,omitempty"`
	// Distance in seconds between a screenshot
	Distance *float64 `json:"distance,omitempty"`
	// Filename of the Bif image. (required)
	Filename string `json:"filename,omitempty"`
	Outputs []EncodingOutput `json:"outputs,omitempty"`
}

