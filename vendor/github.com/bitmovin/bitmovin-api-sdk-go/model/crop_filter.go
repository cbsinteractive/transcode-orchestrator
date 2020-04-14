package model
import (
	"time"
)

type CropFilter struct {
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
	// Amount of pixels which will be cropped of the input video from the left side.
	Left *int32 `json:"left,omitempty"`
	// Amount of pixels which will be cropped of the input video from the right side.
	Right *int32 `json:"right,omitempty"`
	// Amount of pixels which will be cropped of the input video from the top.
	Top *int32 `json:"top,omitempty"`
	// Amount of pixels which will be cropped of the input video from the bottom.
	Bottom *int32 `json:"bottom,omitempty"`
	Unit PositionUnit `json:"unit,omitempty"`
}
func (o CropFilter) FilterType() FilterType {
    return FilterType_CROP
}

