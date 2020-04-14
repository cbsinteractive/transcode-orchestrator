package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingFiltersUnsharpApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingFiltersUnsharpCustomdataApi
}

func NewEncodingFiltersUnsharpApi(configs ...func(*common.ApiClient)) (*EncodingFiltersUnsharpApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingFiltersUnsharpApi{apiClient: apiClient}

    customdataApi, err := NewEncodingFiltersUnsharpCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingFiltersUnsharpApi) Create(unsharpFilter model.UnsharpFilter) (*model.UnsharpFilter, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.UnsharpFilter
    err := api.apiClient.Post("/encoding/filters/unsharp", &unsharpFilter, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersUnsharpApi) Delete(filterId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/filters/unsharp/{filter_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersUnsharpApi) Get(filterId string) (*model.UnsharpFilter, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel *model.UnsharpFilter
    err := api.apiClient.Get("/encoding/filters/unsharp/{filter_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersUnsharpApi) List(queryParams ...func(*query.UnsharpFilterListQueryParams)) (*pagination.UnsharpFiltersListPagination, error) {
    queryParameters := &query.UnsharpFilterListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.UnsharpFiltersListPagination
    err := api.apiClient.Get("/encoding/filters/unsharp", nil, &responseModel, reqParams)
    return responseModel, err
}

