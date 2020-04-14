package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingInfrastructureAwsApi struct {
    apiClient *common.ApiClient
    Regions *EncodingInfrastructureAwsRegionsApi
}

func NewEncodingInfrastructureAwsApi(configs ...func(*common.ApiClient)) (*EncodingInfrastructureAwsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInfrastructureAwsApi{apiClient: apiClient}

    regionsApi, err := NewEncodingInfrastructureAwsRegionsApi(configs...)
    api.Regions = regionsApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInfrastructureAwsApi) Create(awsAccount model.AwsAccount) (*model.AwsAccount, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.AwsAccount
    err := api.apiClient.Post("/encoding/infrastructure/aws", &awsAccount, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInfrastructureAwsApi) Delete(infrastructureId string) (*model.AwsAccount, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["infrastructure_id"] = infrastructureId
    }

    var responseModel *model.AwsAccount
    err := api.apiClient.Delete("/encoding/infrastructure/aws/{infrastructure_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInfrastructureAwsApi) Get(infrastructureId string) (*model.AwsAccount, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["infrastructure_id"] = infrastructureId
    }

    var responseModel *model.AwsAccount
    err := api.apiClient.Get("/encoding/infrastructure/aws/{infrastructure_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInfrastructureAwsApi) List(queryParams ...func(*query.AwsAccountListQueryParams)) (*pagination.AwsAccountsListPagination, error) {
    queryParameters := &query.AwsAccountListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.AwsAccountsListPagination
    err := api.apiClient.Get("/encoding/infrastructure/aws", nil, &responseModel, reqParams)
    return responseModel, err
}

