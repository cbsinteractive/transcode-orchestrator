package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingConfigurationsVideoH264CustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingConfigurationsVideoH264CustomdataApi(configs ...func(*common.ApiClient)) (*EncodingConfigurationsVideoH264CustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsVideoH264CustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsVideoH264CustomdataApi) Get(configurationId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/configurations/video/h264/{configuration_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

