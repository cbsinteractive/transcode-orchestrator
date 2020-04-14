package model
import (
	"time"
)

type H265VideoConfiguration struct {
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
	Profile ProfileH265 `json:"profile,omitempty"`
	// Sets the amount of b frames
	Bframes *int32 `json:"bframes,omitempty"`
	// Sets the amount of reference frames
	RefFrames *int32 `json:"refFrames,omitempty"`
	// Sets the quantization factor
	Qp *int32 `json:"qp,omitempty"`
	// Maximum Bitrate
	MaxBitrate *int64 `json:"maxBitrate,omitempty"`
	// Minimum Bitrate
	MinBitrate *int64 `json:"minBitrate,omitempty"`
	// Specify the size of the VBV buffer (kbits)
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
	Level LevelH265 `json:"level,omitempty"`
	// Number of frames for slice-type decision lookahead
	RcLookahead *int32 `json:"rcLookahead,omitempty"`
	// Set the level of effort in determining B frame placement
	BAdapt BAdapt `json:"bAdapt,omitempty"`
	MaxCTUSize MaxCtuSize `json:"maxCTUSize,omitempty"`
	TuIntraDepth TuIntraDepth `json:"tuIntraDepth,omitempty"`
	TuInterDepth TuInterDepth `json:"tuInterDepth,omitempty"`
	MotionSearch MotionSearch `json:"motionSearch,omitempty"`
	// Set the amount of subpel refinement to perform.
	SubMe *int32 `json:"subMe,omitempty"`
	// Set the Motion search range.
	MotionSearchRange *int32 `json:"motionSearchRange,omitempty"`
	// Enable weighted prediction in P slices
	WeightPredictionOnPSlice *bool `json:"weightPredictionOnPSlice,omitempty"`
	// Enable weighted prediction in B slices
	WeightPredictionOnBSlice *bool `json:"weightPredictionOnBSlice,omitempty"`
	// Toggle sample adaptive offset loop filter
	Sao *bool `json:"sao,omitempty"`
	// Set the mastering display color volume SEI info (SMPTE ST 2086). For example `G(13250,34500)B(7500,3000)R(34000,16000)WP(15635,16450)L(10000000,1)` describes a P3D65 1000-nits monitor, where G(x=0.265, y=0.690), B(x=0.150, y=0.060), R(x=0.680, y=0.320), WP(x=0.3127, y=0.3290), L(max=1000, min=0.0001). Part of HDR-10 metadata.
	MasterDisplay string `json:"masterDisplay,omitempty"`
	// Set the max content light level (MaxCLL). Use together with maxPictureAverageLightLevel (which will be 0 if not set). Part of HDR-10 metadata.
	MaxContentLightLevel *int32 `json:"maxContentLightLevel,omitempty"`
	// Set the maximum picture average light level (MaxFALL). Use together with maxContentLightLevel (which will be 0 if not set). Part of HDR-10 metadata.
	MaxPictureAverageLightLevel *int32 `json:"maxPictureAverageLightLevel,omitempty"`
	// Force signaling of HDR parameters in SEI packets. Enabled automatically when masterDisplay or maxContentLightLevel/maxPictureAverageLightLevel are set. Useful when there is a desire to signal 0 values for maxContentLightLevel and maxPictureAverageLightLevel.
	Hdr *bool `json:"hdr,omitempty"`
	// Scene Change sensitivity. The higher the value, the more likely an I-Frame will be inserted. Set to 0 to disable it.
	SceneCutThreshold *int32 `json:"sceneCutThreshold,omitempty"`
	// Controls the adaptive quantization algorithm
	AdaptiveQuantizationMode AdaptiveQuantMode `json:"adaptiveQuantizationMode,omitempty"`
	// By enabling this video stream will be signaled as HLG
	EnableHlgSignaling *bool `json:"enableHlgSignaling,omitempty"`
	// Specifies the source format of the original analog video prior to digitizing and encoding
	VideoFormat VideoFormat `json:"videoFormat,omitempty"`
	// Psycho-visual rate-distortion retains fine details like film grain at the expense of more blocking artifacts. Higher values make the video appear sharper and more detailed but with a higher risk of blocking artifacts. Needs to have subMe with RD_IP, RD_ALL, RD_REF_IP, RD_REF_ALL, QPRD or FULL_RD
	PsyRateDistortionOptimization *float64 `json:"psyRateDistortionOptimization,omitempty"`
	// Strength of psycho-visual optimizations in quantization. Only has an effect in presets which use RDOQ (rd-levels 4 and 5). The value must be between 0 and 50, 1.0 is typical
	PsyRateDistortionOptimizedQuantization *float64 `json:"psyRateDistortionOptimizedQuantization,omitempty"`
	// Signal hypothetical reference decoder (HRD) information
	EnableHrdSignaling *bool `json:"enableHrdSignaling,omitempty"`
	// Enables the use of lookahead’s lowres motion vector fields to determine the amount of reuse of each block to tune adaptive quantization factors.
	Cutree *bool `json:"cutree,omitempty"`
	// Minimum CU size (width and height). By using 16 or 32 the encoder will not analyze the cost of CUs below that minimum threshold, saving considerable amounts of compute with a predictable increase in bitrate. This setting has a large effect on performance on the faster presets.
	MinCodingUnitSize MinCodingUnitSize `json:"minCodingUnitSize,omitempty"`
	// Use multiple worker threads to measure the estimated cost of each frame within the lookahead. The higher this parameter, the less accurate the frame costs will be which will result in less accurate B-frame and scene-cut decisions. Valid range: 0 - 16
	LookaheadSlices *int32 `json:"lookaheadSlices,omitempty"`
	// If enabled, limit references per depth, CU or both.
	LimitReferences LimitReferences `json:"limitReferences,omitempty"`
	// Enable analysis of rectangular motion partitions Nx2N and 2NxN.
	RectangularMotionPartitionsAnalysis *bool `json:"rectangularMotionPartitionsAnalysis,omitempty"`
	// Enable analysis of asymmetric motion partitions.
	AsymetricMotionPartitionsAnalysis *bool `json:"asymetricMotionPartitionsAnalysis,omitempty"`
	// When enabled, will limit modes analyzed for each CU using cost metrics from the 4 sub-CUs. This can significantly improve performance when `rectangularMotionPartitionsAnalysis` and/or `asymetricMotionPartitionsAnalysis` are enabled at minimal compression efficiency loss.
	LimitModes *bool `json:"limitModes,omitempty"`
	// Maximum number of neighbor (spatial and temporal) candidate blocks that the encoder may consider for merging motion predictions. Valid range: 1 - 5
	MaxMerge *int32 `json:"maxMerge,omitempty"`
	// Measure 2Nx2N merge candidates first; if no residual is found, additional modes at that depth are not analysed.
	EarlySkip *bool `json:"earlySkip,omitempty"`
	// If enabled exits early from CU depth recursion. When a skip CU is found, additional heuristics are used to decide whether to terminate recursion.
	RecursionSkip *bool `json:"recursionSkip,omitempty"`
	// Enable faster search method for angular intra predictions.
	FastSearchForAngularIntraPredictions *bool `json:"fastSearchForAngularIntraPredictions,omitempty"`
	// Enables the evaluation of intra modes in B slices.
	EvaluationOfIntraModesInBSlices *bool `json:"evaluationOfIntraModesInBSlices,omitempty"`
	// Hide sign bit of one coefficient per coding tree unit.
	SignHide *bool `json:"signHide,omitempty"`
	// Level of rate-distortion optimization in mode decision. The lower the value the faster the encode, the higher the value higher the compression efficiency. Valid range: 0 - 4
	RateDistortionLevelForModeDecision *int32 `json:"rateDistortionLevelForModeDecision,omitempty"`
	// Specify the amount of rate-distortion analysis to use within quantization.
	RateDistortionLevelForQuantization RateDistortionLevelForQuantization `json:"rateDistortionLevelForQuantization,omitempty"`
	// Sets the minimum of quantization factor. Valid value range: 0 - 69
	QpMin *int32 `json:"qpMin,omitempty"`
	// Sets the maximum of quantization factor. Valid value range: 0 - 69
	QpMax *int32 `json:"qpMax,omitempty"`
	// The encoder may begin encoding a row as soon as the row above it is at least two CTUs ahead in the encode process. Default is enabled.
	WavefrontParallelProcessing *bool `json:"wavefrontParallelProcessing,omitempty"`
	// Encode each incoming frame as multiple parallel slices that may be decoded independently. Default is 1.
	Slices *int32 `json:"slices,omitempty"`
	// Copy buffers of input picture in frame. Default is enabled.
	CopyPicture *bool `json:"copyPicture,omitempty"`
	// If high tier is disabled the encoder will attempt to encode only at the main tier. Default is enabled.
	LevelHighTier *bool `json:"levelHighTier,omitempty"`
	// Enable skipping split rate distortion analysis when sum of split CU RD cost larger than one split CU RD cost for intra CU. Default disabled.
	SkipSplitRateDistortionAnalysis *bool `json:"skipSplitRateDistortionAnalysis,omitempty"`
	// If enabled, consider lossless mode in CU RDO decisions. Default is disabled.
	CodingUnitLossless *bool `json:"codingUnitLossless,omitempty"`
	// Enable evaluation of transform skip (bypass DCT but still use quantization) coding for 4x4 TU coded blocks. Default is NONE.
	TransformSkip TransformSkipMode `json:"transformSkip,omitempty"`
	// Enable QP based rate distortion refinement. Default is disabled.
	RefineRateDistortionCost *bool `json:"refineRateDistortionCost,omitempty"`
	// Enables early exit from transform unit depth recursion, for inter coded blocks. Default is DISABLED.
	LimitTransformUnitDepthRecursion LimitTransformUnitDepthRecursionMode `json:"limitTransformUnitDepthRecursion,omitempty"`
	// An integer value, which denotes strength of noise reduction in intra CUs. Default 0.
	NoiseReductionIntra *int32 `json:"noiseReductionIntra,omitempty"`
	// An integer value, which denotes strength of noise reduction in inter CUs. Default 0.
	NoiseReductionInter *int32 `json:"noiseReductionInter,omitempty"`
	// Penalty for 32x32 intra transform units in non-I slices. Default DISABLED.
	RateDistortionPenalty RateDistortionPenaltyMode `json:"rateDistortionPenalty,omitempty"`
	// Penalty for 32x32 intra transform units in non-I slices. Default DISABLED.
	MaximumTransformUnitSize MaxTransformUnitSize `json:"maximumTransformUnitSize,omitempty"`
	// Increases the RD level at points where quality drops due to VBV rate control enforcement. Default 0.
	DynamicRateDistortionStrength *int32 `json:"dynamicRateDistortionStrength,omitempty"`
	// It is used for mode selection during analysis of CTUs and can achieve significant gain in terms of objective quality metrics SSIM and PSNR. Default false.
	SsimRateDistortionOptimization *bool `json:"ssimRateDistortionOptimization,omitempty"`
	// Enable temporal motion vector predictors in P and B slices. Default true.
	TemporalMotionVectorPredictors *bool `json:"temporalMotionVectorPredictors,omitempty"`
	// Enable motion estimation with source frame pixels, in this mode, motion estimation can be computed independently. Default false.
	AnalyzeSourceFramePixels *bool `json:"analyzeSourceFramePixels,omitempty"`
	// Enable strong intra smoothing for 32x32 intra blocks. Default true.
	StrongIntraSmoothing *bool `json:"strongIntraSmoothing,omitempty"`
	// When generating intra predictions for blocks in inter slices, only intra-coded reference pixels are used. Default false.
	ConstrainedIntraPrediction *bool `json:"constrainedIntraPrediction,omitempty"`
	// This value represents the percentage difference between the inter cost and intra cost of a frame used in scenecut detection. Default 5.0.
	ScenecutBias *float64 `json:"scenecutBias,omitempty"`
	// Number of RADL pictures allowed infront of IDR. Requires fixed keyframe interval. Valid values 0 - `bframes`. Default 0.
	AllowedRADLBeforeIDR *int32 `json:"allowedRADLBeforeIDR,omitempty"`
	// Number of frames for GOP boundary decision lookahead. Valid values 0 - `rcLookahead`. Default 0
	GopLookahead *int32 `json:"gopLookahead,omitempty"`
	// Bias towards B frames in slicetype decision. The higher the bias the more likely the encoder is to use B frames. Default 0
	BframeBias *int32 `json:"bframeBias,omitempty"`
	// Force the encoder to flush frames. Default is DISABLED.
	ForceFlush ForceFlushMode `json:"forceFlush,omitempty"`
	// Adjust the strength of the adaptive quantization offsets. Default 1.0.
	AdaptiveQuantizationStrength *float64 `json:"adaptiveQuantizationStrength,omitempty"`
	// Adjust the AQ offsets based on the relative motion of each block with respect to the motion of the frame. Default false.
	AdaptiveQuantizationMotion *bool `json:"adaptiveQuantizationMotion,omitempty"`
	// Enable adaptive quantization for sub-CTUs. This parameter specifies the minimum CU size at which QP can be adjusted. Default is same as `maxCTUSize`.
	QuantizationGroupSize QuantizationGroupSize `json:"quantizationGroupSize,omitempty"`
	// Enables stricter conditions to control bitrate deviance from the target bitrate in ABR mode. Bit rate adherence is prioritised over quality. Default false.
	StrictCbr *bool `json:"strictCbr,omitempty"`
	// Offset of Cb chroma QP from the luma QP selected by rate control. This is a general way to spend more or less bits on the chroma channel. Default 0
	QpOffsetChromaCb *int32 `json:"qpOffsetChromaCb,omitempty"`
	// Offset of Cr chroma QP from the luma QP selected by rate control. This is a general way to spend more or less bits on the chroma channel. Default 0
	QpOffsetChromaCr *int32 `json:"qpOffsetChromaCr,omitempty"`
	// QP ratio factor between I and P slices. This ratio is used in all of the rate control modes. Default 1.4
	IpRatio *float64 `json:"ipRatio,omitempty"`
	// QP ratio factor between P and B slices. This ratio is used in all of the rate control modes. Default 1.3
	PbRatio *float64 `json:"pbRatio,omitempty"`
	// Sets the quantizer curve compression factor. It weights the frame quantizer based on the complexity of residual (measured by lookahead). Default 0.6
	QuantizerCurveCompressionFactor *float64 `json:"quantizerCurveCompressionFactor,omitempty"`
	// The maximum single adjustment in QP allowed to rate control. Default 4
	QpStep *int32 `json:"qpStep,omitempty"`
	// Enables a specialised ratecontrol algorithm for film grain content. Default false.
	GrainOptimizedRateControl *bool `json:"grainOptimizedRateControl,omitempty"`
	// Temporally blur quants. Default 0.5
	BlurQuants *float64 `json:"blurQuants,omitempty"`
	// Temporally blur complexity. Default 20.0
	BlurComplexity *float64 `json:"blurComplexity,omitempty"`
	// Specify how to handle depencency between SAO and deblocking filter. When enabled, non-deblocked pixels are used for SAO analysis. When disabled, SAO analysis skips the right/bottom boundary areas. Default false.
	SaoNonDeblock *bool `json:"saoNonDeblock,omitempty"`
	// Limit SAO filter computation by early terminating SAO process based on inter prediction mode, CTU spatial-domain correlations, and relations between luma and chroma. Default false.
	LimitSao *bool `json:"limitSao,omitempty"`
	// Will use low-pass subband dct approximation instead of the standard dct for 16x16 and 32x32 blocks. Default false.
	LowpassDct *bool `json:"lowpassDct,omitempty"`
}
func (o H265VideoConfiguration) CodecConfigType() CodecConfigType {
    return CodecConfigType_H265
}

