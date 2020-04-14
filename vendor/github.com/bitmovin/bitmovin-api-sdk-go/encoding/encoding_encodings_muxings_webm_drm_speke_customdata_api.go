package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsMuxingsWebmDrmSpekeCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsMuxingsWebmDrmSpekeCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsWebmDrmSpekeCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsWebmDrmSpekeCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsWebmDrmSpekeCustomdataApi) Get(encodingId string, muxingId string, drmId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/webm/{muxing_id}/drm/speke/{drm_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

