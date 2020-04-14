package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingFiltersAudioMixApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingFiltersAudioMixCustomdataApi
}

func NewEncodingFiltersAudioMixApi(configs ...func(*common.ApiClient)) (*EncodingFiltersAudioMixApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingFiltersAudioMixApi{apiClient: apiClient}

    customdataApi, err := NewEncodingFiltersAudioMixCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingFiltersAudioMixApi) Create(audioMixFilter model.AudioMixFilter) (*model.AudioMixFilter, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.AudioMixFilter
    err := api.apiClient.Post("/encoding/filters/audio-mix", &audioMixFilter, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersAudioMixApi) Delete(filterId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/filters/audio-mix/{filter_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersAudioMixApi) Get(filterId string) (*model.AudioMixFilter, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel *model.AudioMixFilter
    err := api.apiClient.Get("/encoding/filters/audio-mix/{filter_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersAudioMixApi) List(queryParams ...func(*query.AudioMixFilterListQueryParams)) (*pagination.AudioMixFiltersListPagination, error) {
    queryParameters := &query.AudioMixFilterListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.AudioMixFiltersListPagination
    err := api.apiClient.Get("/encoding/filters/audio-mix", nil, &responseModel, reqParams)
    return responseModel, err
}

