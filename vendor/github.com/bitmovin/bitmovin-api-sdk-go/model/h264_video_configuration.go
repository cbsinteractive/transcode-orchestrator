package model
import (
	"time"
)

type H264VideoConfiguration struct {
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
	Crf *float64 `json:"crf,omitempty"`
	// When setting a profile, all other settings must not exceed the limits which are defined in the profile. Otherwise, a higher profile may be automatically chosen. (required)
	Profile ProfileH264 `json:"profile,omitempty"`
	// Sets the amount of b frames.
	Bframes *int32 `json:"bframes,omitempty"`
	// Sets the amount of reference frames.
	RefFrames *int32 `json:"refFrames,omitempty"`
	// Sets the minimum of quantization factor.
	QpMin *int32 `json:"qpMin,omitempty"`
	// Sets the maximum of quantization factor.
	QpMax *int32 `json:"qpMax,omitempty"`
	MvPredictionMode MvPredictionMode `json:"mvPredictionMode,omitempty"`
	// Sets the maximum Motion-Vector-Search-Range
	MvSearchRangeMax *int32 `json:"mvSearchRangeMax,omitempty"`
	// Enable or disable CABAC
	Cabac *bool `json:"cabac,omitempty"`
	// Maximum Bitrate
	MaxBitrate *int64 `json:"maxBitrate,omitempty"`
	// Minimum Bitrate
	MinBitrate *int64 `json:"minBitrate,omitempty"`
	// Playback device buffer size
	Bufsize *int64 `json:"bufsize,omitempty"`
	// Minimum GOP length, the minimum distance between I-frames
	MinGop *int32 `json:"minGop,omitempty"`
	// Maximum GOP length, the maximum distance between I-frames
	MaxGop *int32 `json:"maxGop,omitempty"`
	// Enable open-gop, allows referencing frames from a previous gop
	OpenGop *bool `json:"openGop,omitempty"`
	// Minimum interval in seconds between key frames
	MinKeyframeInterval *float64 `json:"minKeyframeInterval,omitempty"`
	// Maximum interval in seconds between key frames
	MaxKeyframeInterval *float64 `json:"maxKeyframeInterval,omitempty"`
	// If three-pass encoding is used and a level is set for the encoder, the bitrate for some segments may exceed the bitrate limit which is defined by the level.
	Level LevelH264 `json:"level,omitempty"`
	BAdaptiveStrategy BAdapt `json:"bAdaptiveStrategy,omitempty"`
	MotionEstimationMethod H264MotionEstimationMethod `json:"motionEstimationMethod,omitempty"`
	// Number of frames for frame-type decision lookahead
	RcLookahead *int32 `json:"rcLookahead,omitempty"`
	// Subpixel motion estimation and mode decision
	SubMe H264SubMe `json:"subMe,omitempty"`
	// Enables or disables Trellis quantization. NOTE: This requires cabac
	Trellis H264Trellis `json:"trellis,omitempty"`
	// Partitions to consider. Analyzing more partition options improves quality at the cost of speed.
	Partitions []H264Partition `json:"partitions,omitempty"`
	// Number of slices per frame.
	Slices *int32 `json:"slices,omitempty"`
	// Using TOP_FIELD_FIRST or BOTTOM_FIELD_FIRST will output interlaced video
	InterlaceMode H264InterlaceMode `json:"interlaceMode,omitempty"`
	// Scene change sensitivity. The higher the value, the more likely an I-frame will be inserted. Set to 0 to disable it which is advised for scenarios where fixed GOP is required, e.g., adaptive streaming outputs like DASH, HLS and Smooth. Having this setting enabled can improve quality for progressive output with an increased internal chunk length (see `internalChunkLength` of muxings).
	SceneCutThreshold *int32 `json:"sceneCutThreshold,omitempty"`
	// Signal hypothetical reference decoder (HRD) information (requires bufsize to be set)
	NalHrd H264NalHrd `json:"nalHrd,omitempty"`
	// Keep some B-frames as references
	BPyramid H264BPyramid `json:"bPyramid,omitempty"`
	// Defines whether CEA 608/708 subtitles are copied from the input video stream
	Cea608708SubtitleConfig *Cea608708SubtitleConfiguration `json:"cea608708SubtitleConfig,omitempty"`
	// Strength of the in-loop deblocking filter. Higher values deblock more effectively but also soften the image
	DeblockAlpha *int32 `json:"deblockAlpha,omitempty"`
	// Threshold of the in-loop deblocking filter. Higher values apply deblocking stronger on non flat blocks, lower values on flat blocks
	DeblockBeta *int32 `json:"deblockBeta,omitempty"`
	// Controls the adaptive quantization algorithm
	AdaptiveQuantizationMode AdaptiveQuantMode `json:"adaptiveQuantizationMode,omitempty"`
	// Values greater than 1 reduce blocking and blurring in flat and textured areas. Values less than 1 reduces ringing artifacts at the cost of more banding artifacts. Negative values are not allowed
	AdaptiveQuantizationStrength *float64 `json:"adaptiveQuantizationStrength,omitempty"`
	// Allow references on a per partition basis, rather than per-macroblock basis
	MixedReferences *bool `json:"mixedReferences,omitempty"`
	// Enables adaptive spatial transform (high profile 8x8 transform)
	AdaptiveSpatialTransform *bool `json:"adaptiveSpatialTransform,omitempty"`
	// Enables fast skip detection on P-frames. Disabling this very slightly increases quality but at a large speed loss
	FastSkipDetectionPFrames *bool `json:"fastSkipDetectionPFrames,omitempty"`
	// Enable open-gop, allows referencing frames from a previous gop
	WeightedPredictionBFrames *bool `json:"weightedPredictionBFrames,omitempty"`
	// Defines the mode for weighted prediction for P-frames
	WeightedPredictionPFrames WeightedPredictionPFrames `json:"weightedPredictionPFrames,omitempty"`
	// Enable macroblock tree ratecontrol. Macroblock tree rate control tracks how often blocks of the frame are used for prediciting future frames
	MacroblockTreeRatecontrol *bool `json:"macroblockTreeRatecontrol,omitempty"`
	// Ratio between constant bitrate (0.0) and constant quantizer (1.0). Valid range 0.0 - 1.0
	QuantizerCurveCompression *float64 `json:"quantizerCurveCompression,omitempty"`
	// Psychovisual Rate Distortion retains fine details like film grain at the expense of more blocking artifacts. Higher values make the video appear sharper and more detailed but with a higher risk of blocking artifacts. Needs to have subMe with RD_IP, RD_ALL, RD_REF_IP or RD_REF_ALL
	PsyRateDistortionOptimization *float64 `json:"psyRateDistortionOptimization,omitempty"`
	// Higher values will improve sharpness and detail retention but might come at costs of artifacts. Needs to have trellis enabled
	PsyTrellis *float64 `json:"psyTrellis,omitempty"`
}
func (o H264VideoConfiguration) CodecConfigType() CodecConfigType {
    return CodecConfigType_H264
}

