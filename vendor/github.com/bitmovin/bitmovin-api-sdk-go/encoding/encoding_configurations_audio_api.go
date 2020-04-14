package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type EncodingConfigurationsAudioApi struct {
    apiClient *common.ApiClient
    Aac *EncodingConfigurationsAudioAacApi
    HeAacV1 *EncodingConfigurationsAudioHeAacV1Api
    HeAacV2 *EncodingConfigurationsAudioHeAacV2Api
    Vorbis *EncodingConfigurationsAudioVorbisApi
    Opus *EncodingConfigurationsAudioOpusApi
    Ac3 *EncodingConfigurationsAudioAc3Api
    Eac3 *EncodingConfigurationsAudioEac3Api
    Mp2 *EncodingConfigurationsAudioMp2Api
    Mp3 *EncodingConfigurationsAudioMp3Api
}

func NewEncodingConfigurationsAudioApi(configs ...func(*common.ApiClient)) (*EncodingConfigurationsAudioApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsAudioApi{apiClient: apiClient}

    aacApi, err := NewEncodingConfigurationsAudioAacApi(configs...)
    api.Aac = aacApi
    heAacV1Api, err := NewEncodingConfigurationsAudioHeAacV1Api(configs...)
    api.HeAacV1 = heAacV1Api
    heAacV2Api, err := NewEncodingConfigurationsAudioHeAacV2Api(configs...)
    api.HeAacV2 = heAacV2Api
    vorbisApi, err := NewEncodingConfigurationsAudioVorbisApi(configs...)
    api.Vorbis = vorbisApi
    opusApi, err := NewEncodingConfigurationsAudioOpusApi(configs...)
    api.Opus = opusApi
    ac3Api, err := NewEncodingConfigurationsAudioAc3Api(configs...)
    api.Ac3 = ac3Api
    eac3Api, err := NewEncodingConfigurationsAudioEac3Api(configs...)
    api.Eac3 = eac3Api
    mp2Api, err := NewEncodingConfigurationsAudioMp2Api(configs...)
    api.Mp2 = mp2Api
    mp3Api, err := NewEncodingConfigurationsAudioMp3Api(configs...)
    api.Mp3 = mp3Api

	if err != nil {
		return nil, err
	}

	return api, nil
}

