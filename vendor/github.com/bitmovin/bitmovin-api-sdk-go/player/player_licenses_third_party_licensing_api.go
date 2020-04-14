package player
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type PlayerLicensesThirdPartyLicensingApi struct {
    apiClient *common.ApiClient
}

func NewPlayerLicensesThirdPartyLicensingApi(configs ...func(*common.ApiClient)) (*PlayerLicensesThirdPartyLicensingApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &PlayerLicensesThirdPartyLicensingApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *PlayerLicensesThirdPartyLicensingApi) Create(licenseId string, playerThirdPartyLicensing model.PlayerThirdPartyLicensing) (*model.PlayerThirdPartyLicensing, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["license_id"] = licenseId
    }

    var responseModel *model.PlayerThirdPartyLicensing
    err := api.apiClient.Post("/player/licenses/{license_id}/third-party-licensing", &playerThirdPartyLicensing, &responseModel, reqParams)
    return responseModel, err
}

func (api *PlayerLicensesThirdPartyLicensingApi) Delete(licenseId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["license_id"] = licenseId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/player/licenses/{license_id}/third-party-licensing", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *PlayerLicensesThirdPartyLicensingApi) Get(licenseId string) (*model.PlayerThirdPartyLicensing, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["license_id"] = licenseId
    }

    var responseModel *model.PlayerThirdPartyLicensing
    err := api.apiClient.Get("/player/licenses/{license_id}/third-party-licensing", nil, &responseModel, reqParams)
    return responseModel, err
}

