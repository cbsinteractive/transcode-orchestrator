package model
// QuantizationGroupSize : Enable adaptive quantization for sub-CTUs. This parameter specifies the minimum CU size at which QP can be adjusted.
type QuantizationGroupSize string

// List of QuantizationGroupSize
const (
	QuantizationGroupSize_QGS_8x8 QuantizationGroupSize = "QGS_8x8"
	QuantizationGroupSize_QGS_16x16 QuantizationGroupSize = "QGS_16x16"
	QuantizationGroupSize_QGS_32x32 QuantizationGroupSize = "QGS_32x32"
	QuantizationGroupSize_QGS_64x64 QuantizationGroupSize = "QGS_64x64"
)