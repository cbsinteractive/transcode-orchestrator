package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsMuxingsProgressiveTsId3RawCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsMuxingsProgressiveTsId3RawCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsProgressiveTsId3RawCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsProgressiveTsId3RawCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsProgressiveTsId3RawCustomdataApi) Get(encodingId string, muxingId string, id3TagId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["id3_tag_id"] = id3TagId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3/raw/{id3_tag_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

