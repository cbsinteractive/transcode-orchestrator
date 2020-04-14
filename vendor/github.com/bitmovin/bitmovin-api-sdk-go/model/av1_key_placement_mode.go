package model
// Av1KeyPlacementMode : Specifies whether keyframes should be placed at fixed intervals or the encoder may determine optimal placement automatically
type Av1KeyPlacementMode string

// List of Av1KeyPlacementMode
const (
	Av1KeyPlacementMode_AUTO Av1KeyPlacementMode = "AUTO"
	Av1KeyPlacementMode_FIXED Av1KeyPlacementMode = "FIXED"
	Av1KeyPlacementMode_DISABLED Av1KeyPlacementMode = "DISABLED"
)