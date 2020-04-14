package model

type ProgressiveWebmMuxingInformation struct {
	// The mime type of the muxing
	MimeType string `json:"mimeType,omitempty"`
	// The file size of the muxing in bytes
	FileSize *int64 `json:"fileSize,omitempty"`
	// The container format used
	ContainerFormat string `json:"containerFormat,omitempty"`
	// The bitrate of the container if available (tracks + container overhead)
	ContainerBitrate *int64 `json:"containerBitrate,omitempty"`
	// The duration of the container in seconds
	Duration *float64 `json:"duration,omitempty"`
	// Information about the video tracks in the container
	VideoTracks []MuxingInformationVideoTrack `json:"videoTracks,omitempty"`
	// Information about the audio tracks in the container
	AudioTracks []MuxingInformationAudioTrack `json:"audioTracks,omitempty"`
}

