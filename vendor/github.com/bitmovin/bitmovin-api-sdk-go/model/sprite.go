package model
import (
	"time"
)

type Sprite struct {
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
	// Height of one thumbnail (required)
	Height *int32 `json:"height,omitempty"`
	// Width of one thumbnail (required)
	Width *int32 `json:"width,omitempty"`
	Unit SpriteUnit `json:"unit,omitempty"`
	// Distance in the given unit between a screenshot
	Distance *float64 `json:"distance,omitempty"`
	// Name of the sprite image. File extension \".jpg\" or \".png\" is required. (required)
	SpriteName string `json:"spriteName,omitempty"`
	// Filename of the sprite image. If not set, spriteName will be used, but without an extension.
	Filename string `json:"filename,omitempty"`
	// Filename of the vtt-file. The file-extension \".vtt\" is required. (required)
	VttName string `json:"vttName,omitempty"`
	Outputs []EncodingOutput `json:"outputs,omitempty"`
	// Number of images per file. If more images are generated than specified in this value, multiple sprites will be created. You can use the placeholder '%number%' in the spriteName to specify the naming policy.
	ImagesPerFile *int32 `json:"imagesPerFile,omitempty"`
}

