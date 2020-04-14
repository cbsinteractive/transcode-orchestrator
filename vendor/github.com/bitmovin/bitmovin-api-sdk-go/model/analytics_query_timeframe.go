package model
import (
	"time"
)

type AnalyticsQueryTimeframe struct {
	// Start of timeframe which is queried in UTC format.
	Start *time.Time `json:"start,omitempty"`
	// End of timeframe which is queried in UTC format.
	End *time.Time `json:"end,omitempty"`
}

