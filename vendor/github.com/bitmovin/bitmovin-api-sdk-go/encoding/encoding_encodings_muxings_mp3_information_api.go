package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsMuxingsMp3InformationApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsMuxingsMp3InformationApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsMp3InformationApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsMp3InformationApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsMp3InformationApi) Get(encodingId string, muxingId string) (*model.Mp3MuxingInformation, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.Mp3MuxingInformation
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp3/{muxing_id}/information", nil, &responseModel, reqParams)
    return responseModel, err
}

