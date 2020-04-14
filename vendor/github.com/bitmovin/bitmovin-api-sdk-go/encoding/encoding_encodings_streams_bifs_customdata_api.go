package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsStreamsBifsCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsStreamsBifsCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsStreamsBifsCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsStreamsBifsCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsStreamsBifsCustomdataApi) Get(encodingId string, streamId string, bifId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.PathParams["bif_id"] = bifId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/bifs/{bif_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

