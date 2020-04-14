package model
import (
	"time"
)

type Stream struct {
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
	InputStreams []StreamInput `json:"inputStreams,omitempty"`
	Outputs []EncodingOutput `json:"outputs,omitempty"`
	// Set true to create quality metadata for this stream
	CreateQualityMetaData *bool `json:"createQualityMetaData,omitempty"`
	// Id of the codec configuration (required)
	CodecConfigId string `json:"codecConfigId,omitempty"`
	// Number of encoded segments. Available after encoding finishes.
	SegmentsEncoded *int32 `json:"segmentsEncoded,omitempty"`
	// Conditions to evaluate before creating the stream. If this evaluation fails, the stream won't be created. All muxings that depend on the stream will also not be created.
	Conditions *AbstractCondition `json:"conditions,omitempty"`
	// If this is set and contains objects, then this stream has been ignored during the encoding process
	IgnoredBy []Ignoring `json:"ignoredBy,omitempty"`
	// Mode of the stream
	Mode StreamMode `json:"mode,omitempty"`
	// The encoding mode of the stream which was applied by the assigned codec configuration
	SelectedEncodingMode EncodingMode `json:"selectedEncodingMode,omitempty"`
	// Settings to configure Per-Title on stream level
	PerTitleSettings *StreamPerTitleSettings `json:"perTitleSettings,omitempty"`
	Metadata *StreamMetadata `json:"metadata,omitempty"`
	// Determines how to react to errors during decoding
	DecodingErrorMode DecodingErrorMode `json:"decodingErrorMode,omitempty"`
	// Contains stream properties which may not have been defined in the configuration
	AppliedSettings *AppliedStreamSettings `json:"appliedSettings,omitempty"`
}

