package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingConfigurationsAudioOpusCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingConfigurationsAudioOpusCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingConfigurationsAudioOpusCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsAudioOpusCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsAudioOpusCustomdataApi) Get(configurationId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/configurations/audio/opus/{configuration_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

