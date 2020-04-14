package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsMuxingsProgressiveWebmInformationApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsMuxingsProgressiveWebmInformationApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsProgressiveWebmInformationApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsProgressiveWebmInformationApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsProgressiveWebmInformationApi) Get(encodingId string, muxingId string) (*model.ProgressiveWebmMuxingInformation, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.ProgressiveWebmMuxingInformation
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-webm/{muxing_id}/information", nil, &responseModel, reqParams)
    return responseModel, err
}

