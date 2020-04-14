package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingFiltersConformApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingFiltersConformCustomdataApi
}

func NewEncodingFiltersConformApi(configs ...func(*common.ApiClient)) (*EncodingFiltersConformApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingFiltersConformApi{apiClient: apiClient}

    customdataApi, err := NewEncodingFiltersConformCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingFiltersConformApi) Create(conformFilter model.ConformFilter) (*model.ConformFilter, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.ConformFilter
    err := api.apiClient.Post("/encoding/filters/conform", &conformFilter, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersConformApi) Delete(filterId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/filters/conform/{filter_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersConformApi) Get(filterId string) (*model.ConformFilter, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel *model.ConformFilter
    err := api.apiClient.Get("/encoding/filters/conform/{filter_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersConformApi) List(queryParams ...func(*query.ConformFilterListQueryParams)) (*pagination.ConformFiltersListPagination, error) {
    queryParameters := &query.ConformFilterListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.ConformFiltersListPagination
    err := api.apiClient.Get("/encoding/filters/conform", nil, &responseModel, reqParams)
    return responseModel, err
}

