package player
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type PlayerLicensesAnalyticsApi struct {
    apiClient *common.ApiClient
}

func NewPlayerLicensesAnalyticsApi(configs ...func(*common.ApiClient)) (*PlayerLicensesAnalyticsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &PlayerLicensesAnalyticsApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *PlayerLicensesAnalyticsApi) Create(licenseId string, playerLicenseAnalytics model.PlayerLicenseAnalytics) (*model.PlayerLicenseAnalytics, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["license_id"] = licenseId
    }

    var responseModel *model.PlayerLicenseAnalytics
    err := api.apiClient.Post("/player/licenses/{license_id}/analytics", &playerLicenseAnalytics, &responseModel, reqParams)
    return responseModel, err
}

func (api *PlayerLicensesAnalyticsApi) Delete(licenseId string) (*model.PlayerLicenseAnalytics, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["license_id"] = licenseId
    }

    var responseModel *model.PlayerLicenseAnalytics
    err := api.apiClient.Delete("/player/licenses/{license_id}/analytics", nil, &responseModel, reqParams)
    return responseModel, err
}

