package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsTsDrmApi struct {
    apiClient *common.ApiClient
    Fairplay *EncodingEncodingsMuxingsTsDrmFairplayApi
    Aes *EncodingEncodingsMuxingsTsDrmAesApi
    Speke *EncodingEncodingsMuxingsTsDrmSpekeApi
}

func NewEncodingEncodingsMuxingsTsDrmApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsTsDrmApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsTsDrmApi{apiClient: apiClient}

    fairplayApi, err := NewEncodingEncodingsMuxingsTsDrmFairplayApi(configs...)
    api.Fairplay = fairplayApi
    aesApi, err := NewEncodingEncodingsMuxingsTsDrmAesApi(configs...)
    api.Aes = aesApi
    spekeApi, err := NewEncodingEncodingsMuxingsTsDrmSpekeApi(configs...)
    api.Speke = spekeApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsTsDrmApi) List(encodingId string, muxingId string) (*pagination.DrmsListPagination, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *pagination.DrmsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/drm", nil, &responseModel, reqParams)
    return responseModel, err
}

