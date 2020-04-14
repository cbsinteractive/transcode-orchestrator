package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingFiltersEnhancedWatermarkApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingFiltersEnhancedWatermarkCustomdataApi
}

func NewEncodingFiltersEnhancedWatermarkApi(configs ...func(*common.ApiClient)) (*EncodingFiltersEnhancedWatermarkApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingFiltersEnhancedWatermarkApi{apiClient: apiClient}

    customdataApi, err := NewEncodingFiltersEnhancedWatermarkCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingFiltersEnhancedWatermarkApi) Create(enhancedWatermarkFilter model.EnhancedWatermarkFilter) (*model.EnhancedWatermarkFilter, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.EnhancedWatermarkFilter
    err := api.apiClient.Post("/encoding/filters/enhanced-watermark", &enhancedWatermarkFilter, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersEnhancedWatermarkApi) Delete(filterId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/filters/enhanced-watermark/{filter_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersEnhancedWatermarkApi) Get(filterId string) (*model.EnhancedWatermarkFilter, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel *model.EnhancedWatermarkFilter
    err := api.apiClient.Get("/encoding/filters/enhanced-watermark/{filter_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersEnhancedWatermarkApi) List(queryParams ...func(*query.EnhancedWatermarkFilterListQueryParams)) (*pagination.EnhancedWatermarkFiltersListPagination, error) {
    queryParameters := &query.EnhancedWatermarkFilterListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.EnhancedWatermarkFiltersListPagination
    err := api.apiClient.Get("/encoding/filters/enhanced-watermark", nil, &responseModel, reqParams)
    return responseModel, err
}

