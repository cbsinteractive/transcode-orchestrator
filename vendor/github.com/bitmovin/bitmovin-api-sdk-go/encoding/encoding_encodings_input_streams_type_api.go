package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsInputStreamsTypeApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsInputStreamsTypeApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsInputStreamsTypeApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsInputStreamsTypeApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsInputStreamsTypeApi) Get(encodingId string, inputStreamId string) (*model.InputStreamTypeResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
    }

    var responseModel *model.InputStreamTypeResponse
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/{input_stream_id}/type", nil, &responseModel, reqParams)
    return responseModel, err
}

