package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type EncodingManifestsDashPeriodsAdaptationsetsApi struct {
    apiClient *common.ApiClient
    Audio *EncodingManifestsDashPeriodsAdaptationsetsAudioApi
    Video *EncodingManifestsDashPeriodsAdaptationsetsVideoApi
    Subtitle *EncodingManifestsDashPeriodsAdaptationsetsSubtitleApi
    Representations *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsApi
    Contentprotection *EncodingManifestsDashPeriodsAdaptationsetsContentprotectionApi
}

func NewEncodingManifestsDashPeriodsAdaptationsetsApi(configs ...func(*common.ApiClient)) (*EncodingManifestsDashPeriodsAdaptationsetsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsDashPeriodsAdaptationsetsApi{apiClient: apiClient}

    audioApi, err := NewEncodingManifestsDashPeriodsAdaptationsetsAudioApi(configs...)
    api.Audio = audioApi
    videoApi, err := NewEncodingManifestsDashPeriodsAdaptationsetsVideoApi(configs...)
    api.Video = videoApi
    subtitleApi, err := NewEncodingManifestsDashPeriodsAdaptationsetsSubtitleApi(configs...)
    api.Subtitle = subtitleApi
    representationsApi, err := NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsApi(configs...)
    api.Representations = representationsApi
    contentprotectionApi, err := NewEncodingManifestsDashPeriodsAdaptationsetsContentprotectionApi(configs...)
    api.Contentprotection = contentprotectionApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

