package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsMuxingsProgressiveMovInformationApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsMuxingsProgressiveMovInformationApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsProgressiveMovInformationApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsProgressiveMovInformationApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsProgressiveMovInformationApi) Get(encodingId string, muxingId string) (*model.ProgressiveMovMuxingInformation, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.ProgressiveMovMuxingInformation
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-mov/{muxing_id}/information", nil, &responseModel, reqParams)
    return responseModel, err
}

