package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingFiltersDeinterlaceApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingFiltersDeinterlaceCustomdataApi
}

func NewEncodingFiltersDeinterlaceApi(configs ...func(*common.ApiClient)) (*EncodingFiltersDeinterlaceApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingFiltersDeinterlaceApi{apiClient: apiClient}

    customdataApi, err := NewEncodingFiltersDeinterlaceCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingFiltersDeinterlaceApi) Create(deinterlaceFilter model.DeinterlaceFilter) (*model.DeinterlaceFilter, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.DeinterlaceFilter
    err := api.apiClient.Post("/encoding/filters/deinterlace", &deinterlaceFilter, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersDeinterlaceApi) Delete(filterId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/filters/deinterlace/{filter_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersDeinterlaceApi) Get(filterId string) (*model.DeinterlaceFilter, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel *model.DeinterlaceFilter
    err := api.apiClient.Get("/encoding/filters/deinterlace/{filter_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersDeinterlaceApi) List(queryParams ...func(*query.DeinterlaceFilterListQueryParams)) (*pagination.DeinterlaceFiltersListPagination, error) {
    queryParameters := &query.DeinterlaceFilterListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.DeinterlaceFiltersListPagination
    err := api.apiClient.Get("/encoding/filters/deinterlace", nil, &responseModel, reqParams)
    return responseModel, err
}

