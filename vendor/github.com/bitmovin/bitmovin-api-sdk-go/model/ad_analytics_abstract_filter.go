package model

type AdAnalyticsAbstractFilter interface {
    AnalyticsQueryOperator() AnalyticsQueryOperator
}

type BaseAdAnalyticsAbstractFilter struct {
    Name AdAnalyticsAttribute `json:"name"`
    Operator AnalyticsQueryOperator `json:"operator"`
}

func (b BaseAdAnalyticsAbstractFilter) AnalyticsQueryOperator() AnalyticsQueryOperator {
    return b.Operator
}

