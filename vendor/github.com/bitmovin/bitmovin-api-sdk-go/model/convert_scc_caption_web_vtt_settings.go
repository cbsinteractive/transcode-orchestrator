package model

type ConvertSccCaptionWebVttSettings struct {
	PositionMode ConvertSccPositionMode `json:"positionMode,omitempty"`
	// Remove flash (blinking) information when converting SCC to WebVTT
	RemoveFlash *bool `json:"removeFlash,omitempty"`
	// Remove color information when converting SCC to WebVTT
	RemoveColor *bool `json:"removeColor,omitempty"`
}

