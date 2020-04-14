package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingInputsAzureApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingInputsAzureCustomdataApi
}

func NewEncodingInputsAzureApi(configs ...func(*common.ApiClient)) (*EncodingInputsAzureApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInputsAzureApi{apiClient: apiClient}

    customdataApi, err := NewEncodingInputsAzureCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInputsAzureApi) Create(azureInput model.AzureInput) (*model.AzureInput, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.AzureInput
    err := api.apiClient.Post("/encoding/inputs/azure", &azureInput, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsAzureApi) Delete(inputId string) (*model.AzureInput, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel *model.AzureInput
    err := api.apiClient.Delete("/encoding/inputs/azure/{input_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsAzureApi) Get(inputId string) (*model.AzureInput, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel *model.AzureInput
    err := api.apiClient.Get("/encoding/inputs/azure/{input_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsAzureApi) List(queryParams ...func(*query.AzureInputListQueryParams)) (*pagination.AzureInputsListPagination, error) {
    queryParameters := &query.AzureInputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.AzureInputsListPagination
    err := api.apiClient.Get("/encoding/inputs/azure", nil, &responseModel, reqParams)
    return responseModel, err
}

