package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingInputsLocalApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingInputsLocalCustomdataApi
}

func NewEncodingInputsLocalApi(configs ...func(*common.ApiClient)) (*EncodingInputsLocalApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInputsLocalApi{apiClient: apiClient}

    customdataApi, err := NewEncodingInputsLocalCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInputsLocalApi) Create(localInput model.LocalInput) (*model.LocalInput, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.LocalInput
    err := api.apiClient.Post("/encoding/inputs/local", &localInput, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsLocalApi) Delete(inputId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/inputs/local/{input_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsLocalApi) Get(inputId string) (*model.LocalInput, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel *model.LocalInput
    err := api.apiClient.Get("/encoding/inputs/local/{input_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsLocalApi) List(queryParams ...func(*query.LocalInputListQueryParams)) (*pagination.LocalInputsListPagination, error) {
    queryParameters := &query.LocalInputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.LocalInputsListPagination
    err := api.apiClient.Get("/encoding/inputs/local", nil, &responseModel, reqParams)
    return responseModel, err
}

