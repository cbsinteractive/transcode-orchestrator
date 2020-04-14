package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type EncodingEncodingsInputStreamsCaptionsApi struct {
    apiClient *common.ApiClient
    Cea608 *EncodingEncodingsInputStreamsCaptionsCea608Api
    Cea708 *EncodingEncodingsInputStreamsCaptionsCea708Api
}

func NewEncodingEncodingsInputStreamsCaptionsApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsInputStreamsCaptionsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsInputStreamsCaptionsApi{apiClient: apiClient}

    cea608Api, err := NewEncodingEncodingsInputStreamsCaptionsCea608Api(configs...)
    api.Cea608 = cea608Api
    cea708Api, err := NewEncodingEncodingsInputStreamsCaptionsCea708Api(configs...)
    api.Cea708 = cea708Api

	if err != nil {
		return nil, err
	}

	return api, nil
}

