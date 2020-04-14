package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingConfigurationsAudioVorbisCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingConfigurationsAudioVorbisCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingConfigurationsAudioVorbisCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsAudioVorbisCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsAudioVorbisCustomdataApi) Get(configurationId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/configurations/audio/vorbis/{configuration_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

