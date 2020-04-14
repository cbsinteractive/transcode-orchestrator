package model

type AnalyticsLessThanOrEqualFilter struct {
	Name *map[string]interface{} `json:"name,omitempty"`
	Value *map[string]interface{} `json:"value,omitempty"`
}
func (o AnalyticsLessThanOrEqualFilter) AnalyticsQueryOperator() AnalyticsQueryOperator {
    return AnalyticsQueryOperator_LTE
}

