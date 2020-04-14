package model
import (
	"time"
)

type DeinterlaceFilter struct {
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
	Parity PictureFieldParity `json:"parity,omitempty"`
	Mode DeinterlaceMode `json:"mode,omitempty"`
	FrameSelectionMode DeinterlaceFrameSelectionMode `json:"frameSelectionMode,omitempty"`
	AutoEnable DeinterlaceAutoEnable `json:"autoEnable,omitempty"`
}
func (o DeinterlaceFilter) FilterType() FilterType {
    return FilterType_DEINTERLACE
}

