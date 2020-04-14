package model
type LimitReferences string

// List of LimitReferences
const (
	LimitReferences_DISABLED LimitReferences = "DISABLED"
	LimitReferences_DEPTH LimitReferences = "DEPTH"
	LimitReferences_CU LimitReferences = "CU"
	LimitReferences_DEPTH_AND_CU LimitReferences = "DEPTH_AND_CU"
)