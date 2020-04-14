package model
// LimitTransformUnitDepthRecursionMode : Enables early exit from transform unit depth recursion, for inter coded blocks.
type LimitTransformUnitDepthRecursionMode string

// List of LimitTransformUnitDepthRecursionMode
const (
	LimitTransformUnitDepthRecursionMode_DISABLED LimitTransformUnitDepthRecursionMode = "DISABLED"
	LimitTransformUnitDepthRecursionMode_LEVEL_1 LimitTransformUnitDepthRecursionMode = "LEVEL_1"
	LimitTransformUnitDepthRecursionMode_LEVEL_2 LimitTransformUnitDepthRecursionMode = "LEVEL_2"
	LimitTransformUnitDepthRecursionMode_LEVEL_3 LimitTransformUnitDepthRecursionMode = "LEVEL_3"
	LimitTransformUnitDepthRecursionMode_LEVEL_4 LimitTransformUnitDepthRecursionMode = "LEVEL_4"
)