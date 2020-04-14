package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingInfrastructureGceApi struct {
    apiClient *common.ApiClient
}

func NewEncodingInfrastructureGceApi(configs ...func(*common.ApiClient)) (*EncodingInfrastructureGceApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInfrastructureGceApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInfrastructureGceApi) Create(gceAccount model.GceAccount) (*model.GceAccount, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.GceAccount
    err := api.apiClient.Post("/encoding/infrastructure/gce", &gceAccount, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInfrastructureGceApi) Delete(infrastructureId string) (*model.GceAccount, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["infrastructure_id"] = infrastructureId
    }

    var responseModel *model.GceAccount
    err := api.apiClient.Delete("/encoding/infrastructure/gce/{infrastructure_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInfrastructureGceApi) Get(infrastructureId string) (*model.GceAccount, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["infrastructure_id"] = infrastructureId
    }

    var responseModel *model.GceAccount
    err := api.apiClient.Get("/encoding/infrastructure/gce/{infrastructure_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInfrastructureGceApi) List(queryParams ...func(*query.GceAccountListQueryParams)) (*pagination.GceAccountsListPagination, error) {
    queryParameters := &query.GceAccountListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.GceAccountsListPagination
    err := api.apiClient.Get("/encoding/infrastructure/gce", nil, &responseModel, reqParams)
    return responseModel, err
}

