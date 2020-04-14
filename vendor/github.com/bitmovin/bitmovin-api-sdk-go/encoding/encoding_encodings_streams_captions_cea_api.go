package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type EncodingEncodingsStreamsCaptionsCeaApi struct {
    apiClient *common.ApiClient
    Scc *EncodingEncodingsStreamsCaptionsCeaSccApi
}

func NewEncodingEncodingsStreamsCaptionsCeaApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsStreamsCaptionsCeaApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsStreamsCaptionsCeaApi{apiClient: apiClient}

    sccApi, err := NewEncodingEncodingsStreamsCaptionsCeaSccApi(configs...)
    api.Scc = sccApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

