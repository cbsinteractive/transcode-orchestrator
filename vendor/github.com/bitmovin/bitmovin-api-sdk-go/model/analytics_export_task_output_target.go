package model

type AnalyticsExportTaskOutputTarget struct {
	// Path where the export should be saved (required)
	OutputPath string `json:"outputPath,omitempty"`
	// Id of the output that should be used (required)
	OutputId string `json:"outputId,omitempty"`
}

