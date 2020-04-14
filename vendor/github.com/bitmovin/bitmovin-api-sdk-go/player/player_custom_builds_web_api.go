package player
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type PlayerCustomBuildsWebApi struct {
    apiClient *common.ApiClient
    Domains *PlayerCustomBuildsWebDomainsApi
    Status *PlayerCustomBuildsWebStatusApi
    Download *PlayerCustomBuildsWebDownloadApi
}

func NewPlayerCustomBuildsWebApi(configs ...func(*common.ApiClient)) (*PlayerCustomBuildsWebApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &PlayerCustomBuildsWebApi{apiClient: apiClient}

    domainsApi, err := NewPlayerCustomBuildsWebDomainsApi(configs...)
    api.Domains = domainsApi
    statusApi, err := NewPlayerCustomBuildsWebStatusApi(configs...)
    api.Status = statusApi
    downloadApi, err := NewPlayerCustomBuildsWebDownloadApi(configs...)
    api.Download = downloadApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *PlayerCustomBuildsWebApi) Create(customPlayerBuildDetails model.CustomPlayerBuildDetails) (*model.CustomPlayerBuildDetails, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.CustomPlayerBuildDetails
    err := api.apiClient.Post("/player/custom-builds/web", &customPlayerBuildDetails, &responseModel, reqParams)
    return responseModel, err
}

func (api *PlayerCustomBuildsWebApi) Get(customBuildId string) (*model.CustomPlayerBuildStatus, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["custom_build_id"] = customBuildId
    }

    var responseModel *model.CustomPlayerBuildStatus
    err := api.apiClient.Get("/player/custom-builds/web/{custom_build_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *PlayerCustomBuildsWebApi) List() (*pagination.CustomPlayerBuildDetailssListPagination, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *pagination.CustomPlayerBuildDetailssListPagination
    err := api.apiClient.Get("/player/custom-builds/web", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *PlayerCustomBuildsWebApi) Start(customBuildId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["custom_build_id"] = customBuildId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Post("/player/custom-builds/web/{custom_build_id}/start", nil, &responseModel, reqParams)
    return responseModel, err
}

