package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type EncodingEncodingsStreamsCaptionsApi struct {
    apiClient *common.ApiClient
    Cea *EncodingEncodingsStreamsCaptionsCeaApi
}

func NewEncodingEncodingsStreamsCaptionsApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsStreamsCaptionsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsStreamsCaptionsApi{apiClient: apiClient}

    ceaApi, err := NewEncodingEncodingsStreamsCaptionsCeaApi(configs...)
    api.Cea = ceaApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

