package player
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type PlayerCustomBuildsWebDomainsApi struct {
    apiClient *common.ApiClient
}

func NewPlayerCustomBuildsWebDomainsApi(configs ...func(*common.ApiClient)) (*PlayerCustomBuildsWebDomainsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &PlayerCustomBuildsWebDomainsApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *PlayerCustomBuildsWebDomainsApi) Create(customWebPlayerBuildDomain model.CustomWebPlayerBuildDomain) (*model.CustomWebPlayerBuildDomain, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.CustomWebPlayerBuildDomain
    err := api.apiClient.Post("/player/custom-builds/web/domains", &customWebPlayerBuildDomain, &responseModel, reqParams)
    return responseModel, err
}

func (api *PlayerCustomBuildsWebDomainsApi) Delete(domainId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["domain_id"] = domainId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/player/custom-builds/web/domains/{domain_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *PlayerCustomBuildsWebDomainsApi) Get(domainId string) (*model.CustomWebPlayerBuildDomain, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["domain_id"] = domainId
    }

    var responseModel *model.CustomWebPlayerBuildDomain
    err := api.apiClient.Get("/player/custom-builds/web/domains/{domain_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *PlayerCustomBuildsWebDomainsApi) List() (*pagination.CustomWebPlayerBuildDomainsListPagination, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *pagination.CustomWebPlayerBuildDomainsListPagination
    err := api.apiClient.Get("/player/custom-builds/web/domains", nil, &responseModel, reqParams)
    return responseModel, err
}

