package model

type Tweaks struct {
	// Defines special audio video sync handling
	AudioVideoSyncMode AudioVideoSyncMode `json:"audioVideoSyncMode,omitempty"`
}

