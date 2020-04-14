package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type EncodingEncodingsStreamsHdrApi struct {
    apiClient *common.ApiClient
    DolbyVision *EncodingEncodingsStreamsHdrDolbyVisionApi
}

func NewEncodingEncodingsStreamsHdrApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsStreamsHdrApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsStreamsHdrApi{apiClient: apiClient}

    dolbyVisionApi, err := NewEncodingEncodingsStreamsHdrDolbyVisionApi(configs...)
    api.DolbyVision = dolbyVisionApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

