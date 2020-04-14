package model

type AdAnalyticsNotEqualFilter struct {
	Name *map[string]interface{} `json:"name,omitempty"`
	Value *map[string]interface{} `json:"value,omitempty"`
}
func (o AdAnalyticsNotEqualFilter) AnalyticsQueryOperator() AnalyticsQueryOperator {
    return AnalyticsQueryOperator_NE
}

