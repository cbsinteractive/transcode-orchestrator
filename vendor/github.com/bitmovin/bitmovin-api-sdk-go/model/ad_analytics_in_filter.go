package model

type AdAnalyticsInFilter struct {
	Name *map[string]interface{} `json:"name,omitempty"`
	Value []map[string]interface{} `json:"value,omitempty"`
}
func (o AdAnalyticsInFilter) AnalyticsQueryOperator() AnalyticsQueryOperator {
    return AnalyticsQueryOperator_IN
}

