package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsMp4DrmApi struct {
    apiClient *common.ApiClient
    Playready *EncodingEncodingsMuxingsMp4DrmPlayreadyApi
    Clearkey *EncodingEncodingsMuxingsMp4DrmClearkeyApi
    Widevine *EncodingEncodingsMuxingsMp4DrmWidevineApi
    Marlin *EncodingEncodingsMuxingsMp4DrmMarlinApi
    Cenc *EncodingEncodingsMuxingsMp4DrmCencApi
    Speke *EncodingEncodingsMuxingsMp4DrmSpekeApi
}

func NewEncodingEncodingsMuxingsMp4DrmApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsMp4DrmApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsMp4DrmApi{apiClient: apiClient}

    playreadyApi, err := NewEncodingEncodingsMuxingsMp4DrmPlayreadyApi(configs...)
    api.Playready = playreadyApi
    clearkeyApi, err := NewEncodingEncodingsMuxingsMp4DrmClearkeyApi(configs...)
    api.Clearkey = clearkeyApi
    widevineApi, err := NewEncodingEncodingsMuxingsMp4DrmWidevineApi(configs...)
    api.Widevine = widevineApi
    marlinApi, err := NewEncodingEncodingsMuxingsMp4DrmMarlinApi(configs...)
    api.Marlin = marlinApi
    cencApi, err := NewEncodingEncodingsMuxingsMp4DrmCencApi(configs...)
    api.Cenc = cencApi
    spekeApi, err := NewEncodingEncodingsMuxingsMp4DrmSpekeApi(configs...)
    api.Speke = spekeApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsMp4DrmApi) List(encodingId string, muxingId string) (*pagination.DrmsListPagination, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *pagination.DrmsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm", nil, &responseModel, reqParams)
    return responseModel, err
}

