package model

type AnalyticsGreaterThanFilter struct {
	Name *map[string]interface{} `json:"name,omitempty"`
	Value *map[string]interface{} `json:"value,omitempty"`
}
func (o AnalyticsGreaterThanFilter) AnalyticsQueryOperator() AnalyticsQueryOperator {
    return AnalyticsQueryOperator_GT
}

