package model
// ConditionOperator : The operator that should be used for the evaluation
type ConditionOperator string

// List of ConditionOperator
const (
	ConditionOperator_EQUAL ConditionOperator = "=="
	ConditionOperator_NOT_EQUAL ConditionOperator = "!="
	ConditionOperator_LESS_THAN_OR_EQUAL ConditionOperator = "<="
	ConditionOperator_LESS_THAN ConditionOperator = "<"
	ConditionOperator_GREATER_THAN ConditionOperator = ">"
	ConditionOperator_GREATER_THAN_OR_EQUAL ConditionOperator = ">="
)