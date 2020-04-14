package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsWebmDrmApi struct {
    apiClient *common.ApiClient
    Cenc *EncodingEncodingsMuxingsWebmDrmCencApi
    Speke *EncodingEncodingsMuxingsWebmDrmSpekeApi
}

func NewEncodingEncodingsMuxingsWebmDrmApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsWebmDrmApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsWebmDrmApi{apiClient: apiClient}

    cencApi, err := NewEncodingEncodingsMuxingsWebmDrmCencApi(configs...)
    api.Cenc = cencApi
    spekeApi, err := NewEncodingEncodingsMuxingsWebmDrmSpekeApi(configs...)
    api.Speke = spekeApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsWebmDrmApi) List(encodingId string, muxingId string) (*pagination.DrmsListPagination, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *pagination.DrmsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/webm/{muxing_id}/drm", nil, &responseModel, reqParams)
    return responseModel, err
}

