package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type EncodingConfigurationsSubtitlesApi struct {
    apiClient *common.ApiClient
    Webvtt *EncodingConfigurationsSubtitlesWebvttApi
}

func NewEncodingConfigurationsSubtitlesApi(configs ...func(*common.ApiClient)) (*EncodingConfigurationsSubtitlesApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsSubtitlesApi{apiClient: apiClient}

    webvttApi, err := NewEncodingConfigurationsSubtitlesWebvttApi(configs...)
    api.Webvtt = webvttApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

