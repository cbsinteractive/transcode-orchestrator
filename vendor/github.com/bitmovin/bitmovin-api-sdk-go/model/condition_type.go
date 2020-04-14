package model
// ConditionType : Type of the condition
type ConditionType string

// List of ConditionType
const (
	ConditionType_CONDITION ConditionType = "CONDITION"
	ConditionType_AND ConditionType = "AND"
	ConditionType_OR ConditionType = "OR"
)