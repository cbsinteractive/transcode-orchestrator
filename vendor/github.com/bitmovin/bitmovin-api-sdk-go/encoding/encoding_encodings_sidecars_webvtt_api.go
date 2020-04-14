package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsSidecarsWebvttApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsSidecarsWebvttApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsSidecarsWebvttApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsSidecarsWebvttApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsSidecarsWebvttApi) Create(encodingId string, webVttSidecarFile model.WebVttSidecarFile) (*model.WebVttSidecarFile, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.WebVttSidecarFile
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/sidecars/webvtt", &webVttSidecarFile, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsSidecarsWebvttApi) Delete(encodingId string, sidecarId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["sidecar_id"] = sidecarId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/sidecars/webvtt/{sidecar_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsSidecarsWebvttApi) Get(encodingId string, sidecarId string) (*model.WebVttSidecarFile, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["sidecar_id"] = sidecarId
    }

    var responseModel *model.WebVttSidecarFile
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/sidecars/webvtt/{sidecar_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

