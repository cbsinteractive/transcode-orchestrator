package model
import (
	"time"
)

type AnalyticsVarianceQueryRequest struct {
	// Analytics license key (required)
	LicenseKey string `json:"licenseKey,omitempty"`
	Filters []AnalyticsAbstractFilter `json:"filters,omitempty"`
	OrderBy []AnalyticsOrderByEntry `json:"orderBy,omitempty"`
	Dimension AnalyticsAttribute `json:"dimension,omitempty"`
	Interval AnalyticsInterval `json:"interval,omitempty"`
	GroupBy []AnalyticsAttribute `json:"groupBy,omitempty"`
	// Maximum number of rows returned (max. 200)
	Limit *int64 `json:"limit,omitempty"`
	// Offset of data
	Offset *int64 `json:"offset,omitempty"`
	// Start of timeframe which is queried in UTC format.
	Start *time.Time `json:"start,omitempty"`
	// End of timeframe which is queried in UTC format.
	End *time.Time `json:"end,omitempty"`
}

