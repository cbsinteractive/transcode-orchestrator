package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingFiltersTextApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingFiltersTextCustomdataApi
}

func NewEncodingFiltersTextApi(configs ...func(*common.ApiClient)) (*EncodingFiltersTextApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingFiltersTextApi{apiClient: apiClient}

    customdataApi, err := NewEncodingFiltersTextCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingFiltersTextApi) Create(textFilter model.TextFilter) (*model.TextFilter, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.TextFilter
    err := api.apiClient.Post("/encoding/filters/text", &textFilter, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersTextApi) Delete(filterId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/filters/text/{filter_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersTextApi) Get(filterId string) (*model.TextFilter, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel *model.TextFilter
    err := api.apiClient.Get("/encoding/filters/text/{filter_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersTextApi) List(queryParams ...func(*query.TextFilterListQueryParams)) (*pagination.TextFiltersListPagination, error) {
    queryParameters := &query.TextFilterListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.TextFiltersListPagination
    err := api.apiClient.Get("/encoding/filters/text", nil, &responseModel, reqParams)
    return responseModel, err
}

