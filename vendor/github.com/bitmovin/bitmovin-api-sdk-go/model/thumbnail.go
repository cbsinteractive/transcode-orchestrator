package model
import (
	"time"
)

type Thumbnail struct {
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
	// Height of the thumbnail. (required)
	Height *int32 `json:"height,omitempty"`
	//  Pattern which describes the thumbnail filenames. For example with thumbnail-%number%.png as pattern and 3 positions: thumbnail-3_0.png, thumbnail-5_0.png and thumbnail-25_5.png. (The number represents the position in the source video in seconds, in the previous example the first filename represents the thumbnail at 3s, the second one at 5s and the third one at 25.5s).
	Pattern string `json:"pattern,omitempty"`
	// The interval in which to create thumbnails. In seconds (E.g. a value of 4 means create a thumbnail every 4 seconds). Mutually exclusive with positions/unit. Has to be equal to or greater than 1.
	Interval *float64 `json:"interval,omitempty"`
	// Position in the unit where the thumbnail should be created from. Mutually exclusive with interval.
	Positions []float64 `json:"positions,omitempty"`
	Outputs []EncodingOutput `json:"outputs,omitempty"`
	// Unit of the values in the positions array.
	Unit ThumbnailUnit `json:"unit,omitempty"`
}

