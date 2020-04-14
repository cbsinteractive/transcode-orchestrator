package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingConfigurationsAudioMp3Api struct {
    apiClient *common.ApiClient
    Customdata *EncodingConfigurationsAudioMp3CustomdataApi
}

func NewEncodingConfigurationsAudioMp3Api(configs ...func(*common.ApiClient)) (*EncodingConfigurationsAudioMp3Api, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsAudioMp3Api{apiClient: apiClient}

    customdataApi, err := NewEncodingConfigurationsAudioMp3CustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsAudioMp3Api) Create(mp3AudioConfiguration model.Mp3AudioConfiguration) (*model.Mp3AudioConfiguration, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.Mp3AudioConfiguration
    err := api.apiClient.Post("/encoding/configurations/audio/mp3", &mp3AudioConfiguration, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsAudioMp3Api) Delete(configurationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/configurations/audio/mp3/{configuration_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsAudioMp3Api) Get(configurationId string) (*model.Mp3AudioConfiguration, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.Mp3AudioConfiguration
    err := api.apiClient.Get("/encoding/configurations/audio/mp3/{configuration_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsAudioMp3Api) List(queryParams ...func(*query.Mp3AudioConfigurationListQueryParams)) (*pagination.Mp3AudioConfigurationsListPagination, error) {
    queryParameters := &query.Mp3AudioConfigurationListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.Mp3AudioConfigurationsListPagination
    err := api.apiClient.Get("/encoding/configurations/audio/mp3", nil, &responseModel, reqParams)
    return responseModel, err
}

