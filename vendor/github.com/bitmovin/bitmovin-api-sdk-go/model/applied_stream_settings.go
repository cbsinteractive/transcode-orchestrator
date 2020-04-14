package model

type AppliedStreamSettings struct {
	// The applied width. Useful if the width in the configuration was undefined
	Width *int32 `json:"width,omitempty"`
	// The applied height. Useful if the height in the configuration was undefined
	Height *int32 `json:"height,omitempty"`
}

