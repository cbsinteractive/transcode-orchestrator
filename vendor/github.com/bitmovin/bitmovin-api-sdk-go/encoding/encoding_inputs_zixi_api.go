package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingInputsZixiApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingInputsZixiCustomdataApi
}

func NewEncodingInputsZixiApi(configs ...func(*common.ApiClient)) (*EncodingInputsZixiApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInputsZixiApi{apiClient: apiClient}

    customdataApi, err := NewEncodingInputsZixiCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInputsZixiApi) Create(zixiInput model.ZixiInput) (*model.ZixiInput, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.ZixiInput
    err := api.apiClient.Post("/encoding/inputs/zixi", &zixiInput, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsZixiApi) Delete(inputId string) (*model.ZixiInput, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel *model.ZixiInput
    err := api.apiClient.Delete("/encoding/inputs/zixi/{input_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsZixiApi) Get(inputId string) (*model.ZixiInput, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel *model.ZixiInput
    err := api.apiClient.Get("/encoding/inputs/zixi/{input_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsZixiApi) List(queryParams ...func(*query.ZixiInputListQueryParams)) (*pagination.ZixiInputsListPagination, error) {
    queryParameters := &query.ZixiInputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.ZixiInputsListPagination
    err := api.apiClient.Get("/encoding/inputs/zixi", nil, &responseModel, reqParams)
    return responseModel, err
}

