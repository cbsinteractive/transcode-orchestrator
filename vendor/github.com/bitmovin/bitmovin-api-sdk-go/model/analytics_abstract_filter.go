package model

type AnalyticsAbstractFilter interface {
    AnalyticsQueryOperator() AnalyticsQueryOperator
}

type BaseAnalyticsAbstractFilter struct {
    Name AnalyticsAttribute `json:"name"`
    Operator AnalyticsQueryOperator `json:"operator"`
}

func (b BaseAnalyticsAbstractFilter) AnalyticsQueryOperator() AnalyticsQueryOperator {
    return b.Operator
}

