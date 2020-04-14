package model
type AnalyticsInterval string

// List of AnalyticsInterval
const (
	AnalyticsInterval_MINUTE AnalyticsInterval = "MINUTE"
	AnalyticsInterval_HOUR AnalyticsInterval = "HOUR"
	AnalyticsInterval_DAY AnalyticsInterval = "DAY"
	AnalyticsInterval_MONTH AnalyticsInterval = "MONTH"
)