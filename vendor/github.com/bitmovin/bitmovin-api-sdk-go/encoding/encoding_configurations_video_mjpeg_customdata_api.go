package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingConfigurationsVideoMjpegCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingConfigurationsVideoMjpegCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingConfigurationsVideoMjpegCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsVideoMjpegCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsVideoMjpegCustomdataApi) Get(configurationId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/configurations/video/mjpeg/{configuration_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

