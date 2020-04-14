package model

type AndConjunction struct {
	// Array to perform the AND/OR evaluation on
	Conditions []AbstractCondition `json:"conditions,omitempty"`
}
func (o AndConjunction) ConditionType() ConditionType {
    return ConditionType_AND
}

