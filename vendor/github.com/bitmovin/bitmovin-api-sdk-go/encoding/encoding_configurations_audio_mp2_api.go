package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingConfigurationsAudioMp2Api struct {
    apiClient *common.ApiClient
    Customdata *EncodingConfigurationsAudioMp2CustomdataApi
}

func NewEncodingConfigurationsAudioMp2Api(configs ...func(*common.ApiClient)) (*EncodingConfigurationsAudioMp2Api, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsAudioMp2Api{apiClient: apiClient}

    customdataApi, err := NewEncodingConfigurationsAudioMp2CustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsAudioMp2Api) Create(mp2AudioConfiguration model.Mp2AudioConfiguration) (*model.Mp2AudioConfiguration, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.Mp2AudioConfiguration
    err := api.apiClient.Post("/encoding/configurations/audio/mp2", &mp2AudioConfiguration, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsAudioMp2Api) Delete(configurationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/configurations/audio/mp2/{configuration_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsAudioMp2Api) Get(configurationId string) (*model.Mp2AudioConfiguration, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.Mp2AudioConfiguration
    err := api.apiClient.Get("/encoding/configurations/audio/mp2/{configuration_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

