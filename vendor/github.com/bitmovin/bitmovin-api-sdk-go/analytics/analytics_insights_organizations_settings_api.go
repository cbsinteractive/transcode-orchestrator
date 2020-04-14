package analytics
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type AnalyticsInsightsOrganizationsSettingsApi struct {
    apiClient *common.ApiClient
}

func NewAnalyticsInsightsOrganizationsSettingsApi(configs ...func(*common.ApiClient)) (*AnalyticsInsightsOrganizationsSettingsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AnalyticsInsightsOrganizationsSettingsApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AnalyticsInsightsOrganizationsSettingsApi) Get(organizationId string) (*model.AnalyticsInsightsOrganizationSettings, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["organization_id"] = organizationId
    }

    var responseModel *model.AnalyticsInsightsOrganizationSettings
    err := api.apiClient.Get("/analytics/insights/organizations/{organization_id}/settings", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *AnalyticsInsightsOrganizationsSettingsApi) Update(organizationId string, analyticsInsightsOrganizationSettingsRequest model.AnalyticsInsightsOrganizationSettingsRequest) (*model.AnalyticsInsightsOrganizationSettings, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["organization_id"] = organizationId
    }

    var responseModel *model.AnalyticsInsightsOrganizationSettings
    err := api.apiClient.Put("/analytics/insights/organizations/{organization_id}/settings", &analyticsInsightsOrganizationSettingsRequest, &responseModel, reqParams)
    return responseModel, err
}

