package model
// MvPredictionMode : Sets the Motion Vector Prediction Mode.
type MvPredictionMode string

// List of MvPredictionMode
const (
	MvPredictionMode_NONE MvPredictionMode = "NONE"
	MvPredictionMode_SPATIAL MvPredictionMode = "SPATIAL"
	MvPredictionMode_TEMPORAL MvPredictionMode = "TEMPORAL"
	MvPredictionMode_AUTO MvPredictionMode = "AUTO"
)