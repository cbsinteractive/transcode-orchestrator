package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingConfigurationsAudioAc3CustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingConfigurationsAudioAc3CustomdataApi(configs ...func(*common.ApiClient)) (*EncodingConfigurationsAudioAc3CustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsAudioAc3CustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsAudioAc3CustomdataApi) Get(configurationId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/configurations/audio/ac3/{configuration_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

