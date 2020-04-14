package model
// ForceFlushMode : Force the encoder to flush frames. Default is DISABLED.
type ForceFlushMode string

// List of ForceFlushMode
const (
	ForceFlushMode_DISABLED ForceFlushMode = "DISABLED"
	ForceFlushMode_ALL_FRAMES ForceFlushMode = "ALL_FRAMES"
	ForceFlushMode_SLICE_TYPE ForceFlushMode = "SLICE_TYPE"
)