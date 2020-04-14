package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingInputsHttpsApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingInputsHttpsCustomdataApi
}

func NewEncodingInputsHttpsApi(configs ...func(*common.ApiClient)) (*EncodingInputsHttpsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInputsHttpsApi{apiClient: apiClient}

    customdataApi, err := NewEncodingInputsHttpsCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInputsHttpsApi) Create(httpsInput model.HttpsInput) (*model.HttpsInput, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.HttpsInput
    err := api.apiClient.Post("/encoding/inputs/https", &httpsInput, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsHttpsApi) Delete(inputId string) (*model.HttpsInput, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel *model.HttpsInput
    err := api.apiClient.Delete("/encoding/inputs/https/{input_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsHttpsApi) Get(inputId string) (*model.HttpsInput, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel *model.HttpsInput
    err := api.apiClient.Get("/encoding/inputs/https/{input_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsHttpsApi) List(queryParams ...func(*query.HttpsInputListQueryParams)) (*pagination.HttpsInputsListPagination, error) {
    queryParameters := &query.HttpsInputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.HttpsInputsListPagination
    err := api.apiClient.Get("/encoding/inputs/https", nil, &responseModel, reqParams)
    return responseModel, err
}

