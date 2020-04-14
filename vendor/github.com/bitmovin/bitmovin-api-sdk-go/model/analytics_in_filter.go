package model

type AnalyticsInFilter struct {
	Name *map[string]interface{} `json:"name,omitempty"`
	Value []map[string]interface{} `json:"value,omitempty"`
}
func (o AnalyticsInFilter) AnalyticsQueryOperator() AnalyticsQueryOperator {
    return AnalyticsQueryOperator_IN
}

