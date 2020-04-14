package model
type AnalyticsQueryOperator string

// List of AnalyticsQueryOperator
const (
	AnalyticsQueryOperator_IN AnalyticsQueryOperator = "IN"
	AnalyticsQueryOperator_EQ AnalyticsQueryOperator = "EQ"
	AnalyticsQueryOperator_NE AnalyticsQueryOperator = "NE"
	AnalyticsQueryOperator_LT AnalyticsQueryOperator = "LT"
	AnalyticsQueryOperator_LTE AnalyticsQueryOperator = "LTE"
	AnalyticsQueryOperator_GT AnalyticsQueryOperator = "GT"
	AnalyticsQueryOperator_GTE AnalyticsQueryOperator = "GTE"
	AnalyticsQueryOperator_CONTAINS AnalyticsQueryOperator = "CONTAINS"
	AnalyticsQueryOperator_NOTCONTAINS AnalyticsQueryOperator = "NOTCONTAINS"
)