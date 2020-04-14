package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingFiltersScaleApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingFiltersScaleCustomdataApi
}

func NewEncodingFiltersScaleApi(configs ...func(*common.ApiClient)) (*EncodingFiltersScaleApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingFiltersScaleApi{apiClient: apiClient}

    customdataApi, err := NewEncodingFiltersScaleCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingFiltersScaleApi) Create(scaleFilter model.ScaleFilter) (*model.ScaleFilter, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.ScaleFilter
    err := api.apiClient.Post("/encoding/filters/scale", &scaleFilter, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersScaleApi) Delete(filterId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/filters/scale/{filter_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersScaleApi) Get(filterId string) (*model.ScaleFilter, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel *model.ScaleFilter
    err := api.apiClient.Get("/encoding/filters/scale/{filter_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersScaleApi) List(queryParams ...func(*query.ScaleFilterListQueryParams)) (*pagination.ScaleFiltersListPagination, error) {
    queryParameters := &query.ScaleFilterListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.ScaleFiltersListPagination
    err := api.apiClient.Get("/encoding/filters/scale", nil, &responseModel, reqParams)
    return responseModel, err
}

