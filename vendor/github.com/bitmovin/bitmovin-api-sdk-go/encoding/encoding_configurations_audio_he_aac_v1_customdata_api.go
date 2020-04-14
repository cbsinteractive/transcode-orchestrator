package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingConfigurationsAudioHeAacV1CustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingConfigurationsAudioHeAacV1CustomdataApi(configs ...func(*common.ApiClient)) (*EncodingConfigurationsAudioHeAacV1CustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsAudioHeAacV1CustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsAudioHeAacV1CustomdataApi) Get(configurationId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/configurations/audio/he-aac-v1/{configuration_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

