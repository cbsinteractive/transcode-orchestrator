package model
import (
	"time"
)

type MjpegVideoConfiguration struct {
	// Name of the resource. Can be freely chosen by the user. (required)
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
	// Width of the encoded video
	Width *int32 `json:"width,omitempty"`
	// Height of the encoded video
	Height *int32 `json:"height,omitempty"`
	// Target frame rate of the encoded video! (required)
	Rate *float64 `json:"rate,omitempty"`
	// The quality scale parameter (required)
	QScale *int32 `json:"qScale,omitempty"`
	PixelFormat PixelFormat `json:"pixelFormat,omitempty"`
}
func (o MjpegVideoConfiguration) CodecConfigType() CodecConfigType {
    return CodecConfigType_MJPEG
}

