package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingFiltersDenoiseHqdn3dApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingFiltersDenoiseHqdn3dCustomdataApi
}

func NewEncodingFiltersDenoiseHqdn3dApi(configs ...func(*common.ApiClient)) (*EncodingFiltersDenoiseHqdn3dApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingFiltersDenoiseHqdn3dApi{apiClient: apiClient}

    customdataApi, err := NewEncodingFiltersDenoiseHqdn3dCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingFiltersDenoiseHqdn3dApi) Create(denoiseHqdn3dFilter model.DenoiseHqdn3dFilter) (*model.DenoiseHqdn3dFilter, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.DenoiseHqdn3dFilter
    err := api.apiClient.Post("/encoding/filters/denoise-hqdn3d", &denoiseHqdn3dFilter, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersDenoiseHqdn3dApi) Delete(filterId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/filters/denoise-hqdn3d/{filter_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersDenoiseHqdn3dApi) Get(filterId string) (*model.DenoiseHqdn3dFilter, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel *model.DenoiseHqdn3dFilter
    err := api.apiClient.Get("/encoding/filters/denoise-hqdn3d/{filter_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersDenoiseHqdn3dApi) List(queryParams ...func(*query.DenoiseHqdn3dFilterListQueryParams)) (*pagination.DenoiseHqdn3dFiltersListPagination, error) {
    queryParameters := &query.DenoiseHqdn3dFilterListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.DenoiseHqdn3dFiltersListPagination
    err := api.apiClient.Get("/encoding/filters/denoise-hqdn3d", nil, &responseModel, reqParams)
    return responseModel, err
}

