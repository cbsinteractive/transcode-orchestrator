package model
import (
	"time"
)

type Vp8VideoConfiguration struct {
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
	// Sets the constant rate factor for quality-based variable bitrate. Either bitrate or crf is required.
	Crf *int32 `json:"crf,omitempty"`
	// Number of frames to look ahead for alternate reference frame selection.
	LagInFrames *int32 `json:"lagInFrames,omitempty"`
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
	CpuUsed *int32 `json:"cpuUsed,omitempty"`
	NoiseSensitivity Vp8NoiseSensitivity `json:"noiseSensitivity,omitempty"`
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
	Quality Vp8Quality `json:"quality,omitempty"`
	// A change threshold on blocks below which they will be skipped by the encoder.
	StaticThresh *int32 `json:"staticThresh,omitempty"`
	// altref noise reduction max frame count.
	ArnrMaxFrames *int32 `json:"arnrMaxFrames,omitempty"`
	// altref noise reduction filter strength.
	ArnrStrength *int32 `json:"arnrStrength,omitempty"`
	ArnrType Vp8ArnrType `json:"arnrType,omitempty"`
}
func (o Vp8VideoConfiguration) CodecConfigType() CodecConfigType {
    return CodecConfigType_VP8
}

