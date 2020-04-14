package model
type StreamSelectionMode string

// List of StreamSelectionMode
const (
	StreamSelectionMode_AUTO StreamSelectionMode = "AUTO"
	StreamSelectionMode_POSITION_ABSOLUTE StreamSelectionMode = "POSITION_ABSOLUTE"
	StreamSelectionMode_VIDEO_RELATIVE StreamSelectionMode = "VIDEO_RELATIVE"
	StreamSelectionMode_AUDIO_RELATIVE StreamSelectionMode = "AUDIO_RELATIVE"
	StreamSelectionMode_SUBTITLE_RELATIVE StreamSelectionMode = "SUBTITLE_RELATIVE"
)