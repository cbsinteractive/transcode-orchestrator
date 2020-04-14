package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsApi struct {
    apiClient *common.ApiClient
    Vtt *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsVttApi
    Fmp4 *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4Api
    Cmaf *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafApi
    Mp4 *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4Api
    Webm *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmApi
}

func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsApi(configs ...func(*common.ApiClient)) (*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsApi{apiClient: apiClient}

    vttApi, err := NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsVttApi(configs...)
    api.Vtt = vttApi
    fmp4Api, err := NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4Api(configs...)
    api.Fmp4 = fmp4Api
    cmafApi, err := NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafApi(configs...)
    api.Cmaf = cmafApi
    mp4Api, err := NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4Api(configs...)
    api.Mp4 = mp4Api
    webmApi, err := NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmApi(configs...)
    api.Webm = webmApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

