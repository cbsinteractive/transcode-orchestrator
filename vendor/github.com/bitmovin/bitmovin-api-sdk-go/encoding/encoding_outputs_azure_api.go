package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingOutputsAzureApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingOutputsAzureCustomdataApi
}

func NewEncodingOutputsAzureApi(configs ...func(*common.ApiClient)) (*EncodingOutputsAzureApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingOutputsAzureApi{apiClient: apiClient}

    customdataApi, err := NewEncodingOutputsAzureCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingOutputsAzureApi) Create(azureOutput model.AzureOutput) (*model.AzureOutput, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.AzureOutput
    err := api.apiClient.Post("/encoding/outputs/azure", &azureOutput, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingOutputsAzureApi) Delete(outputId string) (*model.AzureOutput, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel *model.AzureOutput
    err := api.apiClient.Delete("/encoding/outputs/azure/{output_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingOutputsAzureApi) Get(outputId string) (*model.AzureOutput, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel *model.AzureOutput
    err := api.apiClient.Get("/encoding/outputs/azure/{output_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingOutputsAzureApi) List(queryParams ...func(*query.AzureOutputListQueryParams)) (*pagination.AzureOutputsListPagination, error) {
    queryParameters := &query.AzureOutputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.AzureOutputsListPagination
    err := api.apiClient.Get("/encoding/outputs/azure", nil, &responseModel, reqParams)
    return responseModel, err
}

