package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type EncodingEncodingsMuxingsCmafCaptionsApi struct {
    apiClient *common.ApiClient
    Ttml *EncodingEncodingsMuxingsCmafCaptionsTtmlApi
}

func NewEncodingEncodingsMuxingsCmafCaptionsApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsCmafCaptionsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsCmafCaptionsApi{apiClient: apiClient}

    ttmlApi, err := NewEncodingEncodingsMuxingsCmafCaptionsTtmlApi(configs...)
    api.Ttml = ttmlApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

