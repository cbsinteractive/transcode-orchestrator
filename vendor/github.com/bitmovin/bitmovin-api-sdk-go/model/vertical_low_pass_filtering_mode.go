package model
// VerticalLowPassFilteringMode : Mode for vertical low pass filtering
type VerticalLowPassFilteringMode string

// List of VerticalLowPassFilteringMode
const (
	VerticalLowPassFilteringMode_DISABLED VerticalLowPassFilteringMode = "DISABLED"
	VerticalLowPassFilteringMode_LOW_PASS VerticalLowPassFilteringMode = "LOW_PASS"
	VerticalLowPassFilteringMode_COMPLEX VerticalLowPassFilteringMode = "COMPLEX"
)