package model

type AudioGroupConfiguration struct {
	// Dropping mode (required)
	DroppingMode VariantStreamDroppingMode `json:"droppingMode,omitempty"`
	// Audio groups (required)
	Groups []AudioGroup `json:"groups,omitempty"`
}

