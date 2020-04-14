package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsMuxingsMp4DrmWidevineCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsMuxingsMp4DrmWidevineCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsMp4DrmWidevineCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsMp4DrmWidevineCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsMp4DrmWidevineCustomdataApi) Get(encodingId string, muxingId string, drmId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/widevine/{drm_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

