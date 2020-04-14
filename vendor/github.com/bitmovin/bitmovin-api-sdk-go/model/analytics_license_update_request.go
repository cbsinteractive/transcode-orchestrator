package model

type AnalyticsLicenseUpdateRequest struct {
	Name string `json:"name,omitempty"`
	IgnoreDNT *bool `json:"ignoreDNT,omitempty"`
	TimeZone string `json:"timeZone,omitempty"`
	// Labels for CustomData fields
	CustomDataFieldLabels *AnalyticsLicenseCustomDataFieldLabels `json:"customDataFieldLabels,omitempty"`
}

