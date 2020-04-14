package model
import (
	"time"
)

type AnalyticsLicense struct {
	// Id of the resource (required)
	Id string `json:"id,omitempty"`
	// Name of the Analytics License
	Name string `json:"name,omitempty"`
	// License Key
	LicenseKey string `json:"licenseKey,omitempty"`
	// Creation date of the Analytics License in UTC format
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Maximum number of impressions
	MaxImpressions *int64 `json:"maxImpressions,omitempty"`
	// Number of impressions recorded
	Impressions *int64 `json:"impressions,omitempty"`
	// Whitelisted domains
	Domains []AnalyticsLicenseDomain `json:"domains,omitempty"`
	// Whether the Do Not Track request from the browser should be ignored
	IgnoreDNT *bool `json:"ignoreDNT,omitempty"`
	// The timezone of the Analytics License
	TimeZone string `json:"timeZone,omitempty"`
	// Labels for CustomData fields
	CustomDataFieldLabels *AnalyticsLicenseCustomDataFieldLabels `json:"customDataFieldLabels,omitempty"`
}

