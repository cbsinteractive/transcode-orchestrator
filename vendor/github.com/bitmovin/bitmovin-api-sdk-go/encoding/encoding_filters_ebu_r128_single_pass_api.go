package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingFiltersEbuR128SinglePassApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingFiltersEbuR128SinglePassCustomdataApi
}

func NewEncodingFiltersEbuR128SinglePassApi(configs ...func(*common.ApiClient)) (*EncodingFiltersEbuR128SinglePassApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingFiltersEbuR128SinglePassApi{apiClient: apiClient}

    customdataApi, err := NewEncodingFiltersEbuR128SinglePassCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingFiltersEbuR128SinglePassApi) Create(ebuR128SinglePassFilter model.EbuR128SinglePassFilter) (*model.EbuR128SinglePassFilter, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.EbuR128SinglePassFilter
    err := api.apiClient.Post("/encoding/filters/ebu-r128-single-pass", &ebuR128SinglePassFilter, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersEbuR128SinglePassApi) Delete(filterId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/filters/ebu-r128-single-pass/{filter_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersEbuR128SinglePassApi) Get(filterId string) (*model.EbuR128SinglePassFilter, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel *model.EbuR128SinglePassFilter
    err := api.apiClient.Get("/encoding/filters/ebu-r128-single-pass/{filter_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersEbuR128SinglePassApi) List(queryParams ...func(*query.EbuR128SinglePassFilterListQueryParams)) (*pagination.EbuR128SinglePassFiltersListPagination, error) {
    queryParameters := &query.EbuR128SinglePassFilterListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.EbuR128SinglePassFiltersListPagination
    err := api.apiClient.Get("/encoding/filters/ebu-r128-single-pass", nil, &responseModel, reqParams)
    return responseModel, err
}

