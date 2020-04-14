package model
import (
	"time"
)

type AnalyticsImpressionsQuery struct {
	// Start of timeframe which is queried in UTC format.
	Start *time.Time `json:"start,omitempty"`
	// End of timeframe which is queried in UTC format.
	End *time.Time `json:"end,omitempty"`
	// Analytics license key
	LicenseKey string `json:"licenseKey,omitempty"`
	Filters []AnalyticsAbstractFilter `json:"filters,omitempty"`
}

