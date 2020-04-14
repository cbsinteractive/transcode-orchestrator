package model
import (
	"time"
)

type AnalyticsExportTask struct {
	// Id of the resource (required)
	Id string `json:"id,omitempty"`
	// Start of timeframe which is exported in UTC format (required)
	StartTime *time.Time `json:"startTime,omitempty"`
	// End of timeframe which is exported in UTC format (required)
	EndTime *time.Time `json:"endTime,omitempty"`
	// Name of the export task (required)
	Name string `json:"name,omitempty"`
	// Export task description
	Description string `json:"description,omitempty"`
	// License key (required)
	LicenseKey string `json:"licenseKey,omitempty"`
	Output *AnalyticsExportTaskOutputTarget `json:"output,omitempty"`
	// Progress of the export task
	Progress *int32 `json:"progress,omitempty"`
	Status AnalyticsExportStatus `json:"status,omitempty"`
	// UTC timestamp when the export task started
	StartedAt *time.Time `json:"startedAt,omitempty"`
	// UTC timestamp when the export task finished
	FinishedAt *time.Time `json:"finishedAt,omitempty"`
}

