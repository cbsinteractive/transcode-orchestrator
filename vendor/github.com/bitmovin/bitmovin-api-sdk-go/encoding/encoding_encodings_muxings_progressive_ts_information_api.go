package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsMuxingsProgressiveTsInformationApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsMuxingsProgressiveTsInformationApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsProgressiveTsInformationApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsProgressiveTsInformationApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsProgressiveTsInformationApi) Get(encodingId string, muxingId string) (*model.ProgressiveTsMuxingInformation, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.ProgressiveTsMuxingInformation
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/information", nil, &responseModel, reqParams)
    return responseModel, err
}

