package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type EncodingEncodingsMuxingsCmafDrmApi struct {
    apiClient *common.ApiClient
    Speke *EncodingEncodingsMuxingsCmafDrmSpekeApi
}

func NewEncodingEncodingsMuxingsCmafDrmApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsCmafDrmApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsCmafDrmApi{apiClient: apiClient}

    spekeApi, err := NewEncodingEncodingsMuxingsCmafDrmSpekeApi(configs...)
    api.Speke = spekeApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

