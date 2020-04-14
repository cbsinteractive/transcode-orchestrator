package model
// Vp9AqMode : Adaptive quantization mode.
type Vp9AqMode string

// List of Vp9AqMode
const (
	Vp9AqMode_NONE Vp9AqMode = "NONE"
	Vp9AqMode_VARIANCE Vp9AqMode = "VARIANCE"
	Vp9AqMode_COMPLEXITY Vp9AqMode = "COMPLEXITY"
	Vp9AqMode_CYCLIC Vp9AqMode = "CYCLIC"
)