package model

type AnalyticsInsightsOrganizationSettings struct {
	// Organization ID
	OrgId string `json:"orgId,omitempty"`
	// Whether the organization's data is included in the industry insights
	IncludeInInsights *bool `json:"includeInInsights,omitempty"`
	// Industry of the organization
	Industry string `json:"industry,omitempty"`
	// Subindustry of the organization
	SubIndustry string `json:"subIndustry,omitempty"`
	// Whether the organization is on an analytics trial plan
	IsTrial *bool `json:"isTrial,omitempty"`
}

