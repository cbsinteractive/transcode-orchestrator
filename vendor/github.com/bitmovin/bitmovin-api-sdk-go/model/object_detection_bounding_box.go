package model

type ObjectDetectionBoundingBox struct {
	// Horizontal position of the top left corner, given as a fraction of the frame width
	TopLeftX *float64 `json:"topLeftX,omitempty"`
	// Vertical position of the top left corner, given as a fraction of the frame height
	TopLeftY *float64 `json:"topLeftY,omitempty"`
	// Horizontal position of the bottom right corner, given as a fraction of the frame width
	BottomRightX *float64 `json:"bottomRightX,omitempty"`
	// Vertical position of the bottom right corner, given as a fraction of the frame height
	BottomRightY *float64 `json:"bottomRightY,omitempty"`
}

