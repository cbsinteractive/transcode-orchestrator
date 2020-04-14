package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type EncodingManifestsSmoothRepresentationsApi struct {
    apiClient *common.ApiClient
    Mp4 *EncodingManifestsSmoothRepresentationsMp4Api
}

func NewEncodingManifestsSmoothRepresentationsApi(configs ...func(*common.ApiClient)) (*EncodingManifestsSmoothRepresentationsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsSmoothRepresentationsApi{apiClient: apiClient}

    mp4Api, err := NewEncodingManifestsSmoothRepresentationsMp4Api(configs...)
    api.Mp4 = mp4Api

	if err != nil {
		return nil, err
	}

	return api, nil
}

