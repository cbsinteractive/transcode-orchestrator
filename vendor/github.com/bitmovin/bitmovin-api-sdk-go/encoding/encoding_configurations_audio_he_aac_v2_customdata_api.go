package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingConfigurationsAudioHeAacV2CustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingConfigurationsAudioHeAacV2CustomdataApi(configs ...func(*common.ApiClient)) (*EncodingConfigurationsAudioHeAacV2CustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsAudioHeAacV2CustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsAudioHeAacV2CustomdataApi) Get(configurationId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/configurations/audio/he-aac-v2/{configuration_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

