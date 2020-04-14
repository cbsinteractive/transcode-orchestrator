package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsStreamsInputApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsStreamsInputApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsStreamsInputApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsStreamsInputApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsStreamsInputApi) Get(encodingId string, streamId string) (*model.EncodingStreamInputDetails, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
    }

    var responseModel *model.EncodingStreamInputDetails
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/input", nil, &responseModel, reqParams)
    return responseModel, err
}

