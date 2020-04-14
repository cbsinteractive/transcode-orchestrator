package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingInputsHttpApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingInputsHttpCustomdataApi
}

func NewEncodingInputsHttpApi(configs ...func(*common.ApiClient)) (*EncodingInputsHttpApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInputsHttpApi{apiClient: apiClient}

    customdataApi, err := NewEncodingInputsHttpCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInputsHttpApi) Create(httpInput model.HttpInput) (*model.HttpInput, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.HttpInput
    err := api.apiClient.Post("/encoding/inputs/http", &httpInput, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsHttpApi) Delete(inputId string) (*model.HttpInput, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel *model.HttpInput
    err := api.apiClient.Delete("/encoding/inputs/http/{input_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsHttpApi) Get(inputId string) (*model.HttpInput, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel *model.HttpInput
    err := api.apiClient.Get("/encoding/inputs/http/{input_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsHttpApi) List(queryParams ...func(*query.HttpInputListQueryParams)) (*pagination.HttpInputsListPagination, error) {
    queryParameters := &query.HttpInputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.HttpInputsListPagination
    err := api.apiClient.Get("/encoding/inputs/http", nil, &responseModel, reqParams)
    return responseModel, err
}

