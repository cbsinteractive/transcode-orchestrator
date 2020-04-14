package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsMuxingsBroadcastTsInformationApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsMuxingsBroadcastTsInformationApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsBroadcastTsInformationApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsBroadcastTsInformationApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsBroadcastTsInformationApi) Get(encodingId string, muxingId string) (*model.BroadcastTsMuxingInformation, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.BroadcastTsMuxingInformation
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/broadcast-ts/{muxing_id}/information", nil, &responseModel, reqParams)
    return responseModel, err
}

