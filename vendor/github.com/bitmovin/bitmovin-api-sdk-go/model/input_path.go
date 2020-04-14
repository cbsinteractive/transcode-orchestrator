package model

type InputPath struct {
	// Id of input (required)
	InputId string `json:"inputId,omitempty"`
	// Path to media file (required)
	InputPath string `json:"inputPath,omitempty"`
}

