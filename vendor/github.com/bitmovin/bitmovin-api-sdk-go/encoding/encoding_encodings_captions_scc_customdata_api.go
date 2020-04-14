package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsCaptionsSccCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsCaptionsSccCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsCaptionsSccCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsCaptionsSccCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsCaptionsSccCustomdataApi) Get(encodingId string, captionsId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["captions_id"] = captionsId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/captions/scc/{captions_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

