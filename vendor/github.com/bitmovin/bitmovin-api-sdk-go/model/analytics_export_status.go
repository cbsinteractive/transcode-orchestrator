package model
type AnalyticsExportStatus string

// List of AnalyticsExportStatus
const (
	AnalyticsExportStatus_STARTED AnalyticsExportStatus = "STARTED"
	AnalyticsExportStatus_FINISHED AnalyticsExportStatus = "FINISHED"
	AnalyticsExportStatus_QUEUED AnalyticsExportStatus = "QUEUED"
	AnalyticsExportStatus_ERROR AnalyticsExportStatus = "ERROR"
)