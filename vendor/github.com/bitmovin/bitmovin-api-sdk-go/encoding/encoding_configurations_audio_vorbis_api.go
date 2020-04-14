package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingConfigurationsAudioVorbisApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingConfigurationsAudioVorbisCustomdataApi
}

func NewEncodingConfigurationsAudioVorbisApi(configs ...func(*common.ApiClient)) (*EncodingConfigurationsAudioVorbisApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsAudioVorbisApi{apiClient: apiClient}

    customdataApi, err := NewEncodingConfigurationsAudioVorbisCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsAudioVorbisApi) Create(vorbisAudioConfiguration model.VorbisAudioConfiguration) (*model.VorbisAudioConfiguration, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.VorbisAudioConfiguration
    err := api.apiClient.Post("/encoding/configurations/audio/vorbis", &vorbisAudioConfiguration, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsAudioVorbisApi) Delete(configurationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/configurations/audio/vorbis/{configuration_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsAudioVorbisApi) Get(configurationId string) (*model.VorbisAudioConfiguration, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.VorbisAudioConfiguration
    err := api.apiClient.Get("/encoding/configurations/audio/vorbis/{configuration_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsAudioVorbisApi) List(queryParams ...func(*query.VorbisAudioConfigurationListQueryParams)) (*pagination.VorbisAudioConfigurationsListPagination, error) {
    queryParameters := &query.VorbisAudioConfigurationListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.VorbisAudioConfigurationsListPagination
    err := api.apiClient.Get("/encoding/configurations/audio/vorbis", nil, &responseModel, reqParams)
    return responseModel, err
}

