package analytics
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type AnalyticsInsightsOrganizationsApi struct {
    apiClient *common.ApiClient
    Settings *AnalyticsInsightsOrganizationsSettingsApi
}

func NewAnalyticsInsightsOrganizationsApi(configs ...func(*common.ApiClient)) (*AnalyticsInsightsOrganizationsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AnalyticsInsightsOrganizationsApi{apiClient: apiClient}

    settingsApi, err := NewAnalyticsInsightsOrganizationsSettingsApi(configs...)
    api.Settings = settingsApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

