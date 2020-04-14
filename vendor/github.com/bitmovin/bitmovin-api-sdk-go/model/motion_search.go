package model
// MotionSearch : Set the Motion search method
type MotionSearch string

// List of MotionSearch
const (
	MotionSearch_DIA MotionSearch = "DIA"
	MotionSearch_HEX MotionSearch = "HEX"
	MotionSearch_UMH MotionSearch = "UMH"
	MotionSearch_STAR MotionSearch = "STAR"
	MotionSearch_FULL MotionSearch = "FULL"
)