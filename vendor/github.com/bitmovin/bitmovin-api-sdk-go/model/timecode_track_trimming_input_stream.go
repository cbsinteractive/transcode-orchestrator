package model
import (
	"time"
)

type TimecodeTrackTrimmingInputStream struct {
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
	// The id of the ingest input stream that should be trimmed
	InputStreamId string `json:"inputStreamId,omitempty"`
	// Defines the timecode, in SMPTE-12M format, of the frame from which the encoding should start. The frame indicated by this value will be included in the encoding
	StartTimeCode string `json:"startTimeCode,omitempty"`
	// Defines the timecode, in SMPTE-12M format, of the frame at which the encoding should stop. The frame indicated by this value will be included in the encoding
	EndTimeCode string `json:"endTimeCode,omitempty"`
}
func (o TimecodeTrackTrimmingInputStream) InputStreamType() InputStreamType {
    return InputStreamType_TRIMMING_TIME_CODE_TRACK
}

