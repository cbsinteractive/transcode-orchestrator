package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type EncodingEncodingsStreamsBurnInSubtitlesApi struct {
    apiClient *common.ApiClient
    Dvbsub *EncodingEncodingsStreamsBurnInSubtitlesDvbsubApi
    Srt *EncodingEncodingsStreamsBurnInSubtitlesSrtApi
}

func NewEncodingEncodingsStreamsBurnInSubtitlesApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsStreamsBurnInSubtitlesApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsStreamsBurnInSubtitlesApi{apiClient: apiClient}

    dvbsubApi, err := NewEncodingEncodingsStreamsBurnInSubtitlesDvbsubApi(configs...)
    api.Dvbsub = dvbsubApi
    srtApi, err := NewEncodingEncodingsStreamsBurnInSubtitlesSrtApi(configs...)
    api.Srt = srtApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

