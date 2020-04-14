package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsMuxingsMp4InformationApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsMuxingsMp4InformationApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsMp4InformationApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsMp4InformationApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsMp4InformationApi) Get(encodingId string, muxingId string) (*model.Mp4MuxingInformation, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.Mp4MuxingInformation
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/information", nil, &responseModel, reqParams)
    return responseModel, err
}

