package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsStreamsHdrDolbyVisionApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsStreamsHdrDolbyVisionApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsStreamsHdrDolbyVisionApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsStreamsHdrDolbyVisionApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsStreamsHdrDolbyVisionApi) Create(encodingId string, streamId string, dolbyVisionMetadata model.DolbyVisionMetadata) (*model.DolbyVisionMetadata, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
    }

    var responseModel *model.DolbyVisionMetadata
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/streams/{stream_id}/hdr/dolby-vision", &dolbyVisionMetadata, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsStreamsHdrDolbyVisionApi) Delete(encodingId string, streamId string, hdrId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.PathParams["hdr_id"] = hdrId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/streams/{stream_id}/hdr/dolby-vision/{hdr_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsStreamsHdrDolbyVisionApi) Get(encodingId string, streamId string, hdrId string) (*model.DolbyVisionMetadata, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.PathParams["hdr_id"] = hdrId
    }

    var responseModel *model.DolbyVisionMetadata
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/hdr/dolby-vision/{hdr_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsStreamsHdrDolbyVisionApi) List(encodingId string, streamId string, queryParams ...func(*query.DolbyVisionMetadataListQueryParams)) (*pagination.DolbyVisionMetadatasListPagination, error) {
    queryParameters := &query.DolbyVisionMetadataListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.DolbyVisionMetadatasListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/hdr/dolby-vision", nil, &responseModel, reqParams)
    return responseModel, err
}

