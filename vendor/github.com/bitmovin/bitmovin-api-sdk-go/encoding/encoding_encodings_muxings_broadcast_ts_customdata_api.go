package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsMuxingsBroadcastTsCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsMuxingsBroadcastTsCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsBroadcastTsCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsBroadcastTsCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsBroadcastTsCustomdataApi) Get(encodingId string, muxingId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/broadcast-ts/{muxing_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

