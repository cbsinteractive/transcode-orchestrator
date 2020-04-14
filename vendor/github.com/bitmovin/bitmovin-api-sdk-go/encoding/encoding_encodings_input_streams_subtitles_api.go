package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type EncodingEncodingsInputStreamsSubtitlesApi struct {
    apiClient *common.ApiClient
    DvbTeletext *EncodingEncodingsInputStreamsSubtitlesDvbTeletextApi
}

func NewEncodingEncodingsInputStreamsSubtitlesApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsInputStreamsSubtitlesApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsInputStreamsSubtitlesApi{apiClient: apiClient}

    dvbTeletextApi, err := NewEncodingEncodingsInputStreamsSubtitlesDvbTeletextApi(configs...)
    api.DvbTeletext = dvbTeletextApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

