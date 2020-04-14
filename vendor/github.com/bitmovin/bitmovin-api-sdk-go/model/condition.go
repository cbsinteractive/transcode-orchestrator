package model

type Condition struct {
	// The attribute that should be used for the evaluation. Valid values include, depending on the context: - HEIGHT - WIDTH - BITRATE - FPS - ASPECTRATIO - INPUTSTREAM - LANGUAGE - CHANNELFORMAT - CHANNELLAYOUT - STREAMCOUNT - AUDIOSTREAMCOUNT - VIDEOSTREAMCOUNT - DURATION - CONNECTION_STATUS - CONNECTION_STATUS_JUST_CHANGED (required)
	Attribute string `json:"attribute,omitempty"`
	// The operator that should be used for the evaluation (required)
	Operator ConditionOperator `json:"operator,omitempty"`
	// The value that should be used for comparison (required)
	Value string `json:"value,omitempty"`
}
func (o Condition) ConditionType() ConditionType {
    return ConditionType_CONDITION
}

