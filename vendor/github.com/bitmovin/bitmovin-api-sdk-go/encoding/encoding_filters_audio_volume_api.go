package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingFiltersAudioVolumeApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingFiltersAudioVolumeCustomdataApi
}

func NewEncodingFiltersAudioVolumeApi(configs ...func(*common.ApiClient)) (*EncodingFiltersAudioVolumeApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingFiltersAudioVolumeApi{apiClient: apiClient}

    customdataApi, err := NewEncodingFiltersAudioVolumeCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingFiltersAudioVolumeApi) Create(audioVolumeFilter model.AudioVolumeFilter) (*model.AudioVolumeFilter, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.AudioVolumeFilter
    err := api.apiClient.Post("/encoding/filters/audio-volume", &audioVolumeFilter, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersAudioVolumeApi) Delete(filterId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/filters/audio-volume/{filter_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersAudioVolumeApi) Get(filterId string) (*model.AudioVolumeFilter, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel *model.AudioVolumeFilter
    err := api.apiClient.Get("/encoding/filters/audio-volume/{filter_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingFiltersAudioVolumeApi) List(queryParams ...func(*query.AudioVolumeFilterListQueryParams)) (*pagination.AudioVolumeFiltersListPagination, error) {
    queryParameters := &query.AudioVolumeFilterListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.AudioVolumeFiltersListPagination
    err := api.apiClient.Get("/encoding/filters/audio-volume", nil, &responseModel, reqParams)
    return responseModel, err
}

