package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type EncodingEncodingsInputStreamsSidecarApi struct {
    apiClient *common.ApiClient
    DolbyVisionMetadataIngest *EncodingEncodingsInputStreamsSidecarDolbyVisionMetadataIngestApi
}

func NewEncodingEncodingsInputStreamsSidecarApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsInputStreamsSidecarApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsInputStreamsSidecarApi{apiClient: apiClient}

    dolbyVisionMetadataIngestApi, err := NewEncodingEncodingsInputStreamsSidecarDolbyVisionMetadataIngestApi(configs...)
    api.DolbyVisionMetadataIngest = dolbyVisionMetadataIngestApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

