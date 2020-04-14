package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingConfigurationsVideoVp9CustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingConfigurationsVideoVp9CustomdataApi(configs ...func(*common.ApiClient)) (*EncodingConfigurationsVideoVp9CustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsVideoVp9CustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsVideoVp9CustomdataApi) Get(configurationId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/configurations/video/vp9/{configuration_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

