package model
import (
	"time"
)

type RedundantRtmpInput struct {
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
	// When there is no input signal present and this threshold in seconds is reached it will switch to another ingest point
	DelayThreshold *int32 `json:"delayThreshold,omitempty"`
	IngestPoints []RtmpIngestPoint `json:"ingestPoints,omitempty"`
}
func (o RedundantRtmpInput) InputType() InputType {
    return InputType_REDUNDANT_RTMP
}

