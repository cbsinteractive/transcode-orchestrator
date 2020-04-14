package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsStreamsInputsApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsStreamsInputsApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsStreamsInputsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsStreamsInputsApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsStreamsInputsApi) List(encodingId string, streamId string) (*pagination.StreamDetailssListPagination, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
    }

    var responseModel *pagination.StreamDetailssListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/inputs", nil, &responseModel, reqParams)
    return responseModel, err
}

