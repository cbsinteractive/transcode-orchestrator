package model
import (
	"time"
)

type Vp9VideoConfiguration struct {
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
	// Use a set of well defined configurations preset to support certain use cases. Can be overwritten with more specific values.
	PresetConfiguration PresetConfiguration `json:"presetConfiguration,omitempty"`
	// Sets the constant rate factor for quality-based variable bitrate. Either bitrate or crf is required.
	Crf *int32 `json:"crf,omitempty"`
	// Number of frames to look ahead for alternate reference frame selection.
	LagInFrames *int32 `json:"lagInFrames,omitempty"`
	// Enables error resiliency feature
	ErrorResiliencyEnabled *bool `json:"errorResiliencyEnabled,omitempty"`
	// Number of tile columns to use, log2. Depending on the encoding width there are limitations on this value. The minimum values are 2 for width >= 1920 and 1 for width >= 1280. The minimum width of each tile is 256 pixels so the maximum values are 0 for width < 256, 1 for width < 512, 2 for width < 1024, 3 for width < 2048, 4 for width < 4096, 5 for width < 8192. If the value is too high or too low it will be overridden.
	TileColumns *int32 `json:"tileColumns,omitempty"`
	// Number of tile rows to use, log2.
	TileRows *int32 `json:"tileRows,omitempty"`
	// Enable frame parallel decodability features
	FrameParallel *bool `json:"frameParallel,omitempty"`
	// Maximum I-frame bitrate (percentage) 0=unlimited
	MaxIntraRate *int64 `json:"maxIntraRate,omitempty"`
	// Sets the minimum of quantization factor.
	QpMin *int32 `json:"qpMin,omitempty"`
	// Sets the maximum of quantization factor.
	QpMax *int32 `json:"qpMax,omitempty"`
	// Datarate undershoot (min) target (percentage).
	RateUndershootPct *int32 `json:"rateUndershootPct,omitempty"`
	// Datarate overshoot (max) target (percentage).
	RateOvershootPct *int32 `json:"rateOvershootPct,omitempty"`
	// Client buffer size (ms)
	ClientBufferSize *int64 `json:"clientBufferSize,omitempty"`
	// Client initial buffer size (ms)
	ClientInitialBufferSize *int64 `json:"clientInitialBufferSize,omitempty"`
	// CBR/VBR bias (0=CBR, 100=VBR)
	BiasPct *int32 `json:"biasPct,omitempty"`
	// Enable noise sensitivity on Y channel
	NoiseSensitivity *bool `json:"noiseSensitivity,omitempty"`
	// Controls the tradeoff between compression efficiency and encoding speed. Higher values indicate a faster encoding. The minimum value for width * height >= 1280 * 720 is 2. If the value is too low it will be overridden.
	CpuUsed *int32 `json:"cpuUsed,omitempty"`
	// Enable automatic alternate reference frames (2pass only)
	AutomaticAltRefFramesEnabled *bool `json:"automaticAltRefFramesEnabled,omitempty"`
	// Target level (255: off, 0: only keep level stats; 10: level 1.0; 11: level 1.1; ... 62: level 6.2)
	TargetLevel *int32 `json:"targetLevel,omitempty"`
	// Enable row based non-deterministic multi-threading
	RowMultiThreadingEnabled *bool `json:"rowMultiThreadingEnabled,omitempty"`
	// Loop filter sharpness.
	Sharpness *int32 `json:"sharpness,omitempty"`
	// Minimum GOP length, the minimum distance between I-frames.
	MinGop *int32 `json:"minGop,omitempty"`
	// Maximum GOP length, the maximum distance between I-frames
	MaxGop *int32 `json:"maxGop,omitempty"`
	// Minimum interval in seconds between key frames
	MinKeyframeInterval *float64 `json:"minKeyframeInterval,omitempty"`
	// Maximum interval in seconds between key frames
	MaxKeyframeInterval *float64 `json:"maxKeyframeInterval,omitempty"`
	Quality Vp9Quality `json:"quality,omitempty"`
	// Lossless mode
	Lossless *bool `json:"lossless,omitempty"`
	// A change threshold on blocks below which they will be skipped by the encoder.
	StaticThresh *int32 `json:"staticThresh,omitempty"`
	AqMode Vp9AqMode `json:"aqMode,omitempty"`
	// altref noise reduction max frame count.
	ArnrMaxFrames *int32 `json:"arnrMaxFrames,omitempty"`
	// altref noise reduction filter strength.
	ArnrStrength *int32 `json:"arnrStrength,omitempty"`
	ArnrType Vp9ArnrType `json:"arnrType,omitempty"`
}
func (o Vp9VideoConfiguration) CodecConfigType() CodecConfigType {
    return CodecConfigType_VP9
}

