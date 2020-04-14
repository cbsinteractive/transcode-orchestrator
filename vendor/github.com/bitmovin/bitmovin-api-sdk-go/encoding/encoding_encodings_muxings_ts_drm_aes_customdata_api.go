package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsMuxingsTsDrmAesCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsMuxingsTsDrmAesCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsTsDrmAesCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsTsDrmAesCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsTsDrmAesCustomdataApi) Get(encodingId string, muxingId string, drmId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/drm/aes/{drm_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

