package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingConfigurationsAudioMp3CustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingConfigurationsAudioMp3CustomdataApi(configs ...func(*common.ApiClient)) (*EncodingConfigurationsAudioMp3CustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsAudioMp3CustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsAudioMp3CustomdataApi) Get(configurationId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/configurations/audio/mp3/{configuration_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

