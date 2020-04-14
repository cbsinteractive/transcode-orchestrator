package model

type AbstractConjunction struct {
	// Array to perform the AND/OR evaluation on
	Conditions []AbstractCondition `json:"conditions,omitempty"`
}

