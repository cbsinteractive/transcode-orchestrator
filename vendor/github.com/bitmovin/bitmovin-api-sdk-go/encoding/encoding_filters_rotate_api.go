package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingFiltersRotateApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingFiltersRotateCustomdataApi
}

func NewEncodingFiltersRotateApi(configs ...func(*common.ApiClient)) (*EncodingFiltersRotateApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingFiltersRotateApi{apiClient: apiClient}

    customdataApi, err := NewEncodingFiltersRotateCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingFiltersRotateApi) Create(rotateFilter model.RotateFilter) (*model.RotateFilter, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.RotateFilter
    err := api.apiClient.Post("/encoding/filters/rotate", &rotateFilter, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersRotateApi) Delete(filterId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/filters/rotate/{filter_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersRotateApi) Get(filterId string) (*model.RotateFilter, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel *model.RotateFilter
    err := api.apiClient.Get("/encoding/filters/rotate/{filter_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersRotateApi) List(queryParams ...func(*query.RotateFilterListQueryParams)) (*pagination.RotateFiltersListPagination, error) {
    queryParameters := &query.RotateFilterListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.RotateFiltersListPagination
    err := api.apiClient.Get("/encoding/filters/rotate", nil, &responseModel, reqParams)
    return responseModel, err
}

