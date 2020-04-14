package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type EncodingConfigurationsVideoApi struct {
    apiClient *common.ApiClient
    H264 *EncodingConfigurationsVideoH264Api
    H265 *EncodingConfigurationsVideoH265Api
    Vp8 *EncodingConfigurationsVideoVp8Api
    Vp9 *EncodingConfigurationsVideoVp9Api
    Av1 *EncodingConfigurationsVideoAv1Api
    Mjpeg *EncodingConfigurationsVideoMjpegApi
}

func NewEncodingConfigurationsVideoApi(configs ...func(*common.ApiClient)) (*EncodingConfigurationsVideoApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsVideoApi{apiClient: apiClient}

    h264Api, err := NewEncodingConfigurationsVideoH264Api(configs...)
    api.H264 = h264Api
    h265Api, err := NewEncodingConfigurationsVideoH265Api(configs...)
    api.H265 = h265Api
    vp8Api, err := NewEncodingConfigurationsVideoVp8Api(configs...)
    api.Vp8 = vp8Api
    vp9Api, err := NewEncodingConfigurationsVideoVp9Api(configs...)
    api.Vp9 = vp9Api
    av1Api, err := NewEncodingConfigurationsVideoAv1Api(configs...)
    api.Av1 = av1Api
    mjpegApi, err := NewEncodingConfigurationsVideoMjpegApi(configs...)
    api.Mjpeg = mjpegApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

