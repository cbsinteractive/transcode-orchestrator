package model

type AbstractCondition interface {
    ConditionType() ConditionType
}

type BaseAbstractCondition struct {
    Type ConditionType `json:"type"`
}

func (b BaseAbstractCondition) ConditionType() ConditionType {
    return b.Type
}

