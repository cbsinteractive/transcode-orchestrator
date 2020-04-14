package model

type AnalyticsNotContainsFilter struct {
	Name *map[string]interface{} `json:"name,omitempty"`
	Value *map[string]interface{} `json:"value,omitempty"`
}
func (o AnalyticsNotContainsFilter) AnalyticsQueryOperator() AnalyticsQueryOperator {
    return AnalyticsQueryOperator_NOTCONTAINS
}

