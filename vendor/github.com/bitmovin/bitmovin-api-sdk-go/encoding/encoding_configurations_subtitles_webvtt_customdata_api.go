package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingConfigurationsSubtitlesWebvttCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingConfigurationsSubtitlesWebvttCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingConfigurationsSubtitlesWebvttCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsSubtitlesWebvttCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsSubtitlesWebvttCustomdataApi) Get(configurationId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/configurations/subtitles/webvtt/{configuration_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

