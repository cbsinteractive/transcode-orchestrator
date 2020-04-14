package model
type PositionMode string

// List of PositionMode
const (
	PositionMode_KEYFRAME PositionMode = "KEYFRAME"
	PositionMode_TIME PositionMode = "TIME"
	PositionMode_SEGMENT PositionMode = "SEGMENT"
)