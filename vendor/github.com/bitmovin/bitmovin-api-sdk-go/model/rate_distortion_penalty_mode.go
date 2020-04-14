package model
// RateDistortionPenaltyMode : Penalty for 32x32 intra transfer units in non-I slices.
type RateDistortionPenaltyMode string

// List of RateDistortionPenaltyMode
const (
	RateDistortionPenaltyMode_DISABLED RateDistortionPenaltyMode = "DISABLED"
	RateDistortionPenaltyMode_NORMAL RateDistortionPenaltyMode = "NORMAL"
	RateDistortionPenaltyMode_MAXIMUM RateDistortionPenaltyMode = "MAXIMUM"
)