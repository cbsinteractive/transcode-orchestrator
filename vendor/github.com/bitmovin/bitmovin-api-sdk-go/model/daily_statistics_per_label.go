package model
import (
	"time"
)

type DailyStatisticsPerLabel struct {
	// Date, format. yyyy-MM-dd (required)
	Date *time.Time `json:"date,omitempty"`
	// List of labels and their aggregated statistics (required)
	Labels []DailyStatistics `json:"labels,omitempty"`
}

