package model
import (
	"time"
)

type UnsharpFilter struct {
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
	// Must be an odd integer between 3 and 23
	LumaMatrixHorizontalSize *int32 `json:"lumaMatrixHorizontalSize,omitempty"`
	// Must be an odd integer between 3 and 23
	LumaMatrixVerticalSize *int32 `json:"lumaMatrixVerticalSize,omitempty"`
	// Negative value: blur, positive value: sharpen, floating point number, valid value range: -1.5 - 1.5
	LumaEffectStrength *float64 `json:"lumaEffectStrength,omitempty"`
	// Must be an odd integer between 3 and 23
	ChromaMatrixHorizontalSize *int32 `json:"chromaMatrixHorizontalSize,omitempty"`
	// Must be an odd integer between 3 and 23
	ChromaMatrixVerticalSize *int32 `json:"chromaMatrixVerticalSize,omitempty"`
	// Negative value: blur, positive value: sharpen, floating point number, valid value range: -1.5 - 1.5
	ChromaEffectStrength *float64 `json:"chromaEffectStrength,omitempty"`
}
func (o UnsharpFilter) FilterType() FilterType {
    return FilterType_UNSHARP
}

