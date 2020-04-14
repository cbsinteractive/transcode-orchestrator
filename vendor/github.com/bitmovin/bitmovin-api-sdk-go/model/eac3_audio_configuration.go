package model
import (
	"time"
)

type Eac3AudioConfiguration struct {
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
	// Target bitrate for the encoded audio in bps (required)
	Bitrate *int64 `json:"bitrate,omitempty"`
	// Audio sampling rate Hz
	Rate *float64 `json:"rate,omitempty"`
	// Channel layout of the audio codec configuration
	ChannelLayout Ac3ChannelLayout `json:"channelLayout,omitempty"`
	// The highest frequency that will pass the audio encoder. This value is optional.
	CutoffFrequency *int32 `json:"cutoffFrequency,omitempty"`
}
func (o Eac3AudioConfiguration) CodecConfigType() CodecConfigType {
    return CodecConfigType_EAC3
}

