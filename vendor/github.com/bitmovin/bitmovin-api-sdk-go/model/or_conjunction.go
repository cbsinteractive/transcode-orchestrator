package model

type OrConjunction struct {
	// Array to perform the AND/OR evaluation on
	Conditions []AbstractCondition `json:"conditions,omitempty"`
}
func (o OrConjunction) ConditionType() ConditionType {
    return ConditionType_OR
}

