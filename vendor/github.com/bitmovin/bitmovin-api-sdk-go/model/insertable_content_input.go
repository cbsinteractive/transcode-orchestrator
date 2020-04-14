package model

type InsertableContentInput struct {
	// Id of the input hosting the video file (required)
	InputId string `json:"inputId,omitempty"`
	// Path to the file on the input (required)
	InputPath string `json:"inputPath,omitempty"`
	// Description of this input
	Description string `json:"description,omitempty"`
}

