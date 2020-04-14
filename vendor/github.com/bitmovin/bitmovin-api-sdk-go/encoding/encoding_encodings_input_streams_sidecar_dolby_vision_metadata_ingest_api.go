package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsInputStreamsSidecarDolbyVisionMetadataIngestApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsInputStreamsSidecarDolbyVisionMetadataIngestApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsInputStreamsSidecarDolbyVisionMetadataIngestApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsInputStreamsSidecarDolbyVisionMetadataIngestApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsInputStreamsSidecarDolbyVisionMetadataIngestApi) Create(encodingId string, dolbyVisionMetadataIngestInputStream model.DolbyVisionMetadataIngestInputStream) (*model.DolbyVisionMetadataIngestInputStream, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.DolbyVisionMetadataIngestInputStream
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/input-streams/sidecar/dolby-vision-metadata-ingest", &dolbyVisionMetadataIngestInputStream, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsInputStreamsSidecarDolbyVisionMetadataIngestApi) Delete(encodingId string, inputStreamId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/input-streams/sidecar/dolby-vision-metadata-ingest/{input_stream_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsInputStreamsSidecarDolbyVisionMetadataIngestApi) Get(encodingId string, inputStreamId string) (*model.DolbyVisionMetadataIngestInputStream, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
    }

    var responseModel *model.DolbyVisionMetadataIngestInputStream
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/sidecar/dolby-vision-metadata-ingest/{input_stream_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsInputStreamsSidecarDolbyVisionMetadataIngestApi) List(encodingId string, queryParams ...func(*query.DolbyVisionMetadataIngestInputStreamListQueryParams)) (*pagination.DolbyVisionMetadataIngestInputStreamsListPagination, error) {
    queryParameters := &query.DolbyVisionMetadataIngestInputStreamListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.DolbyVisionMetadataIngestInputStreamsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/sidecar/dolby-vision-metadata-ingest", nil, &responseModel, reqParams)
    return responseModel, err
}

