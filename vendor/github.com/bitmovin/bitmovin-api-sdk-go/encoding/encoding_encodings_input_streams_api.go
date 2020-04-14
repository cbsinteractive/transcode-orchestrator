package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsInputStreamsApi struct {
    apiClient *common.ApiClient
    Type *EncodingEncodingsInputStreamsTypeApi
    AudioMix *EncodingEncodingsInputStreamsAudioMixApi
    Ingest *EncodingEncodingsInputStreamsIngestApi
    Sidecar *EncodingEncodingsInputStreamsSidecarApi
    Concatenation *EncodingEncodingsInputStreamsConcatenationApi
    File *EncodingEncodingsInputStreamsFileApi
    Trimming *EncodingEncodingsInputStreamsTrimmingApi
    Subtitles *EncodingEncodingsInputStreamsSubtitlesApi
    Captions *EncodingEncodingsInputStreamsCaptionsApi
}

func NewEncodingEncodingsInputStreamsApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsInputStreamsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsInputStreamsApi{apiClient: apiClient}

    typeApi, err := NewEncodingEncodingsInputStreamsTypeApi(configs...)
    api.Type = typeApi
    audioMixApi, err := NewEncodingEncodingsInputStreamsAudioMixApi(configs...)
    api.AudioMix = audioMixApi
    ingestApi, err := NewEncodingEncodingsInputStreamsIngestApi(configs...)
    api.Ingest = ingestApi
    sidecarApi, err := NewEncodingEncodingsInputStreamsSidecarApi(configs...)
    api.Sidecar = sidecarApi
    concatenationApi, err := NewEncodingEncodingsInputStreamsConcatenationApi(configs...)
    api.Concatenation = concatenationApi
    fileApi, err := NewEncodingEncodingsInputStreamsFileApi(configs...)
    api.File = fileApi
    trimmingApi, err := NewEncodingEncodingsInputStreamsTrimmingApi(configs...)
    api.Trimming = trimmingApi
    subtitlesApi, err := NewEncodingEncodingsInputStreamsSubtitlesApi(configs...)
    api.Subtitles = subtitlesApi
    captionsApi, err := NewEncodingEncodingsInputStreamsCaptionsApi(configs...)
    api.Captions = captionsApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsInputStreamsApi) Get(encodingId string, inputStreamId string) (*model.InputStream, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
    }

    var responseModel *model.InputStream
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/{input_stream_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsInputStreamsApi) List(encodingId string, queryParams ...func(*query.InputStreamListQueryParams)) (*pagination.InputStreamsListPagination, error) {
    queryParameters := &query.InputStreamListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.InputStreamsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams", nil, &responseModel, reqParams)
    return responseModel, err
}

