package model

type AutoRepresentation struct {
	// This is the threshold that determines whether the settings of the lower or the upper template representation (codec configuration) should be used, when representations are added automatically. The value must be between 0 and 1. Values nearer 0 will favour the higher representation, values nearer 1 will favour the lower representation.
	AdoptConfigurationThreshold *float64 `json:"adoptConfigurationThreshold,omitempty"`
}

