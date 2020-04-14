package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsMuxingsTsCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsMuxingsTsCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsTsCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsTsCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsTsCustomdataApi) Get(encodingId string, muxingId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

