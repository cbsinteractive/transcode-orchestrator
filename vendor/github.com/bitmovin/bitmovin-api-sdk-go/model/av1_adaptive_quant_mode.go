package model
// Av1AdaptiveQuantMode : AV1 has a segment based feature that allows the encoder to adaptively change quantization parameter for each segment within a frame
type Av1AdaptiveQuantMode string

// List of Av1AdaptiveQuantMode
const (
	Av1AdaptiveQuantMode_OFF Av1AdaptiveQuantMode = "OFF"
	Av1AdaptiveQuantMode_VARIANCE Av1AdaptiveQuantMode = "VARIANCE"
	Av1AdaptiveQuantMode_COMPLEXITY Av1AdaptiveQuantMode = "COMPLEXITY"
	Av1AdaptiveQuantMode_CYCLIC_REFRESH Av1AdaptiveQuantMode = "CYCLIC_REFRESH"
	Av1AdaptiveQuantMode_DELTA_QUANT Av1AdaptiveQuantMode = "DELTA_QUANT"
)