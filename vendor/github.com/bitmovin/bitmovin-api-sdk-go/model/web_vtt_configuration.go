package model
import (
	"time"
)

type WebVttConfiguration struct {
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
	// If set to true, the hours section on webvtt timestamp values will explicitely have zeroes instead of being omitted for values where hours = 0.
	AppendOptionalZeroHour *bool `json:"appendOptionalZeroHour,omitempty"`
	// If set to true, the region information of the resulting webvtt file will be omitted. Defaults to false.
	IgnoreRegion *bool `json:"ignoreRegion,omitempty"`
}

