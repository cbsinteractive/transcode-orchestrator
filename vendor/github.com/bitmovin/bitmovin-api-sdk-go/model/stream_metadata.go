package model

type StreamMetadata struct {
	// Language of the media contained in the stream. If the value is not set, then no metadata tag is set for the media stream.
	Language string `json:"language,omitempty"`
}

