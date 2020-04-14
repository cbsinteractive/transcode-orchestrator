package model
// ScalingAlgorithm : Determines the algorithm used for scaling
type ScalingAlgorithm string

// List of ScalingAlgorithm
const (
	ScalingAlgorithm_FAST_BILINEAR ScalingAlgorithm = "FAST_BILINEAR"
	ScalingAlgorithm_BILINEAR ScalingAlgorithm = "BILINEAR"
	ScalingAlgorithm_BICUBIC ScalingAlgorithm = "BICUBIC"
	ScalingAlgorithm_EXPERIMENTAL ScalingAlgorithm = "EXPERIMENTAL"
	ScalingAlgorithm_NEAREST_NEIGHBOR ScalingAlgorithm = "NEAREST_NEIGHBOR"
	ScalingAlgorithm_AVERAGING_AREA ScalingAlgorithm = "AVERAGING_AREA"
	ScalingAlgorithm_BICUBIC_LUMA_BILINEAR_CHROMA ScalingAlgorithm = "BICUBIC_LUMA_BILINEAR_CHROMA"
	ScalingAlgorithm_GAUSS ScalingAlgorithm = "GAUSS"
	ScalingAlgorithm_SINC ScalingAlgorithm = "SINC"
	ScalingAlgorithm_LANCZOS ScalingAlgorithm = "LANCZOS"
	ScalingAlgorithm_SPLINE ScalingAlgorithm = "SPLINE"
)