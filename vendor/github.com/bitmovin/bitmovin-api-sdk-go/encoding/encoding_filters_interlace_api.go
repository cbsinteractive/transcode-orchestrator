package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingFiltersInterlaceApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingFiltersInterlaceCustomdataApi
}

func NewEncodingFiltersInterlaceApi(configs ...func(*common.ApiClient)) (*EncodingFiltersInterlaceApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingFiltersInterlaceApi{apiClient: apiClient}

    customdataApi, err := NewEncodingFiltersInterlaceCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingFiltersInterlaceApi) Create(interlaceFilter model.InterlaceFilter) (*model.InterlaceFilter, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.InterlaceFilter
    err := api.apiClient.Post("/encoding/filters/interlace", &interlaceFilter, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersInterlaceApi) Delete(filterId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/filters/interlace/{filter_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersInterlaceApi) Get(filterId string) (*model.InterlaceFilter, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel *model.InterlaceFilter
    err := api.apiClient.Get("/encoding/filters/interlace/{filter_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersInterlaceApi) List(queryParams ...func(*query.InterlaceFilterListQueryParams)) (*pagination.InterlaceFiltersListPagination, error) {
    queryParameters := &query.InterlaceFilterListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.InterlaceFiltersListPagination
    err := api.apiClient.Get("/encoding/filters/interlace", nil, &responseModel, reqParams)
    return responseModel, err
}

