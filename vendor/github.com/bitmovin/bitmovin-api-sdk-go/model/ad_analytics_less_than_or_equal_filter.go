package model

type AdAnalyticsLessThanOrEqualFilter struct {
	Name *map[string]interface{} `json:"name,omitempty"`
	Value *map[string]interface{} `json:"value,omitempty"`
}
func (o AdAnalyticsLessThanOrEqualFilter) AnalyticsQueryOperator() AnalyticsQueryOperator {
    return AnalyticsQueryOperator_LTE
}

