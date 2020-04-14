package model

type AnalyticsInsightsOrganizationSettingsRequest struct {
	// Whether the organization's data is being contributed to industry insights
	IncludeInInsights *bool `json:"includeInInsights,omitempty"`
}

