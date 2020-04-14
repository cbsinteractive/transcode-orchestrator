package model

type EncodingOutput struct {
	// Id of the corresponding output (required)
	OutputId string `json:"outputId,omitempty"`
	// Subdirectory where to save the files to (required)
	OutputPath string `json:"outputPath,omitempty"`
	Acl []AclEntry `json:"acl,omitempty"`
}

