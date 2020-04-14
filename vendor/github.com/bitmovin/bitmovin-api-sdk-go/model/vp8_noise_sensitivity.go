package model
// Vp8NoiseSensitivity : Noise sensitivity (frames to blur).
type Vp8NoiseSensitivity string

// List of Vp8NoiseSensitivity
const (
	Vp8NoiseSensitivity_OFF Vp8NoiseSensitivity = "OFF"
	Vp8NoiseSensitivity_ON_Y_ONLY Vp8NoiseSensitivity = "ON_Y_ONLY"
	Vp8NoiseSensitivity_ON_YUV Vp8NoiseSensitivity = "ON_YUV"
	Vp8NoiseSensitivity_ON_YUV_AGGRESSIVE Vp8NoiseSensitivity = "ON_YUV_AGGRESSIVE"
	Vp8NoiseSensitivity_ADAPTIVE Vp8NoiseSensitivity = "ADAPTIVE"
)