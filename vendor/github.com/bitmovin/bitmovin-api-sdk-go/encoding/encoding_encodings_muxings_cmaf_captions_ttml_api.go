package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsMuxingsCmafCaptionsTtmlApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsMuxingsCmafCaptionsTtmlApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsCmafCaptionsTtmlApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsCmafCaptionsTtmlApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsCmafCaptionsTtmlApi) Delete(encodingId string, muxingId string, captionsId string) (*model.TtmlEmbed, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["captions_id"] = captionsId
    }

    var responseModel *model.TtmlEmbed
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}/captions/ttml/{captions_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsCmafCaptionsTtmlApi) Get(encodingId string, muxingId string, captionsId string) (*model.TtmlEmbed, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["captions_id"] = captionsId
    }

    var responseModel *model.TtmlEmbed
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}/captions/ttml/{captions_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

