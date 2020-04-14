package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type EncodingEncodingsCaptionsApi struct {
    apiClient *common.ApiClient
    Scc *EncodingEncodingsCaptionsSccApi
}

func NewEncodingEncodingsCaptionsApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsCaptionsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsCaptionsApi{apiClient: apiClient}

    sccApi, err := NewEncodingEncodingsCaptionsSccApi(configs...)
    api.Scc = sccApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

