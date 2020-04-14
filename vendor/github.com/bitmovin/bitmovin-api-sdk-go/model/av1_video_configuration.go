package model
import (
	"time"
)

type Av1VideoConfiguration struct {
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
	// Width of the encoded video in pixels
	Width *int32 `json:"width,omitempty"`
	// Height of the encoded video in pixels
	Height *int32 `json:"height,omitempty"`
	// Target bitrate for the encoded video in bps. Either bitrate or crf is required.
	Bitrate *int64 `json:"bitrate,omitempty"`
	// Target frame rate of the encoded video. Must be set for live encodings
	Rate *float64 `json:"rate,omitempty"`
	PixelFormat PixelFormat `json:"pixelFormat,omitempty"`
	ColorConfig *ColorConfig `json:"colorConfig,omitempty"`
	// The numerator of the sample aspect ratio (also known as pixel aspect ratio). Must be set if sampleAspectRatioDenominator is set.
	SampleAspectRatioNumerator *int32 `json:"sampleAspectRatioNumerator,omitempty"`
	// The denominator of the sample aspect ratio (also known as pixel aspect ratio). Must be set if sampleAspectRatioNumerator is set.
	SampleAspectRatioDenominator *int32 `json:"sampleAspectRatioDenominator,omitempty"`
	// The mode of the encoding
	EncodingMode EncodingMode `json:"encodingMode,omitempty"`
	KeyPlacementMode Av1KeyPlacementMode `json:"keyPlacementMode,omitempty"`
	AdaptiveQuantMode Av1AdaptiveQuantMode `json:"adaptiveQuantMode,omitempty"`
	// Number of frames to look ahead for alternate reference frame selection
	LagInFrames *int32 `json:"lagInFrames,omitempty"`
	// Minimum (best quality) quantizer
	MinQ *int32 `json:"minQ,omitempty"`
	// Maximum (worst quality) quantizer
	MaxQ *int32 `json:"maxQ,omitempty"`
	// Rate control adaptation undershoot control
	UndershootPct *int32 `json:"undershootPct,omitempty"`
	// Rate control adaptation overshoot control
	OvershootPct *int32 `json:"overshootPct,omitempty"`
	// Decoder buffer size in milliseconds
	ClientBufferSize *int64 `json:"clientBufferSize,omitempty"`
	// Decoder buffer initial size in milliseconds
	ClientInitialBufferSize *int64 `json:"clientInitialBufferSize,omitempty"`
	// Decoder buffer optimal size in milliseconds
	ClientOptimalBufferSize *int64 `json:"clientOptimalBufferSize,omitempty"`
	// Number of tile columns to use, log2
	TileColumns *int32 `json:"tileColumns,omitempty"`
	// Number of tile rows to use, log2
	TileRows *int32 `json:"tileRows,omitempty"`
	// Enable automatic set and use alf frames
	IsAutomaticAltRefFramesEnabled *bool `json:"isAutomaticAltRefFramesEnabled,omitempty"`
	// The max number of frames to create arf
	ArnrMaxFrames *int32 `json:"arnrMaxFrames,omitempty"`
	// The filter strength for the arf
	ArnrStrength *int32 `json:"arnrStrength,omitempty"`
	// Maximum data rate for intra frames, expressed as a percentage of the average per-frame bitrate. Default value 0 meaning unlimited
	MaxIntraRate *int64 `json:"maxIntraRate,omitempty"`
	// Lossless encoding mode
	IsLossless *bool `json:"isLossless,omitempty"`
	// Enable frame parallel decoding feature
	IsFrameParallel *bool `json:"isFrameParallel,omitempty"`
	// Sets the sharpness
	Sharpness *int32 `json:"sharpness,omitempty"`
	// Enable quality boost by lowering frame level Q periodically
	IsFrameBoostEnabled *bool `json:"isFrameBoostEnabled,omitempty"`
	// Enable noise sensitivity on Y channel
	NoiseSensitivity *bool `json:"noiseSensitivity,omitempty"`
	// Minimum interval between GF/ARF frames
	MinGfInterval *int32 `json:"minGfInterval,omitempty"`
	// Maximum interval between GF/ARF frames
	MaxGfInterval *int32 `json:"maxGfInterval,omitempty"`
	// Maximum number of tile groups
	NumTileGroups *int32 `json:"numTileGroups,omitempty"`
	// Maximum number of bytes in a tile group
	MtuSize *int32 `json:"mtuSize,omitempty"`
}
func (o Av1VideoConfiguration) CodecConfigType() CodecConfigType {
    return CodecConfigType_AV1
}

