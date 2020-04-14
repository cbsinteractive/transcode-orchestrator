package analytics
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type AnalyticsInsightsApi struct {
    apiClient *common.ApiClient
    Organizations *AnalyticsInsightsOrganizationsApi
}

func NewAnalyticsInsightsApi(configs ...func(*common.ApiClient)) (*AnalyticsInsightsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AnalyticsInsightsApi{apiClient: apiClient}

    organizationsApi, err := NewAnalyticsInsightsOrganizationsApi(configs...)
    api.Organizations = organizationsApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

