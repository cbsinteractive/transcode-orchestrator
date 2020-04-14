package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingFiltersWatermarkApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingFiltersWatermarkCustomdataApi
}

func NewEncodingFiltersWatermarkApi(configs ...func(*common.ApiClient)) (*EncodingFiltersWatermarkApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingFiltersWatermarkApi{apiClient: apiClient}

    customdataApi, err := NewEncodingFiltersWatermarkCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingFiltersWatermarkApi) Create(watermarkFilter model.WatermarkFilter) (*model.WatermarkFilter, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.WatermarkFilter
    err := api.apiClient.Post("/encoding/filters/watermark", &watermarkFilter, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersWatermarkApi) Delete(filterId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/filters/watermark/{filter_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersWatermarkApi) Get(filterId string) (*model.WatermarkFilter, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel *model.WatermarkFilter
    err := api.apiClient.Get("/encoding/filters/watermark/{filter_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersWatermarkApi) List(queryParams ...func(*query.WatermarkFilterListQueryParams)) (*pagination.WatermarkFiltersListPagination, error) {
    queryParameters := &query.WatermarkFilterListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.WatermarkFiltersListPagination
    err := api.apiClient.Get("/encoding/filters/watermark", nil, &responseModel, reqParams)
    return responseModel, err
}

