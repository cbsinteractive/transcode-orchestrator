package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingFiltersCropApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingFiltersCropCustomdataApi
}

func NewEncodingFiltersCropApi(configs ...func(*common.ApiClient)) (*EncodingFiltersCropApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingFiltersCropApi{apiClient: apiClient}

    customdataApi, err := NewEncodingFiltersCropCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingFiltersCropApi) Create(cropFilter model.CropFilter) (*model.CropFilter, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.CropFilter
    err := api.apiClient.Post("/encoding/filters/crop", &cropFilter, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersCropApi) Delete(filterId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/filters/crop/{filter_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersCropApi) Get(filterId string) (*model.CropFilter, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel *model.CropFilter
    err := api.apiClient.Get("/encoding/filters/crop/{filter_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersCropApi) List(queryParams ...func(*query.CropFilterListQueryParams)) (*pagination.CropFiltersListPagination, error) {
    queryParameters := &query.CropFilterListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.CropFiltersListPagination
    err := api.apiClient.Get("/encoding/filters/crop", nil, &responseModel, reqParams)
    return responseModel, err
}

