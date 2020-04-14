package model
// RaiUnit : Set the random access indicator (RAI) on appropriate access units (AUs).
type RaiUnit string

// List of RaiUnit
const (
	RaiUnit_NONE RaiUnit = "NONE"
	RaiUnit_ALL_PES_PACKETS RaiUnit = "ALL_PES_PACKETS"
	RaiUnit_ACQUISITION_POINT_PACKETS RaiUnit = "ACQUISITION_POINT_PACKETS"
	RaiUnit_ACCORDING_TO_INPUT RaiUnit = "ACCORDING_TO_INPUT"
)