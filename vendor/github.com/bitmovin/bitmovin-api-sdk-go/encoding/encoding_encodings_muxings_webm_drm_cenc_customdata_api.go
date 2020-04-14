package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsMuxingsWebmDrmCencCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsMuxingsWebmDrmCencCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsWebmDrmCencCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsWebmDrmCencCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsWebmDrmCencCustomdataApi) Get(encodingId string, muxingId string, drmId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/webm/{muxing_id}/drm/cenc/{drm_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

