package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsStreamsApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsStreamsCustomdataApi
    Input *EncodingEncodingsStreamsInputApi
    Inputs *EncodingEncodingsStreamsInputsApi
    Filters *EncodingEncodingsStreamsFiltersApi
    BurnInSubtitles *EncodingEncodingsStreamsBurnInSubtitlesApi
    Captions *EncodingEncodingsStreamsCaptionsApi
    Bifs *EncodingEncodingsStreamsBifsApi
    Hdr *EncodingEncodingsStreamsHdrApi
    Thumbnails *EncodingEncodingsStreamsThumbnailsApi
    Sprites *EncodingEncodingsStreamsSpritesApi
    Qc *EncodingEncodingsStreamsQcApi
}

func NewEncodingEncodingsStreamsApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsStreamsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsStreamsApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsStreamsCustomdataApi(configs...)
    api.Customdata = customdataApi
    inputApi, err := NewEncodingEncodingsStreamsInputApi(configs...)
    api.Input = inputApi
    inputsApi, err := NewEncodingEncodingsStreamsInputsApi(configs...)
    api.Inputs = inputsApi
    filtersApi, err := NewEncodingEncodingsStreamsFiltersApi(configs...)
    api.Filters = filtersApi
    burnInSubtitlesApi, err := NewEncodingEncodingsStreamsBurnInSubtitlesApi(configs...)
    api.BurnInSubtitles = burnInSubtitlesApi
    captionsApi, err := NewEncodingEncodingsStreamsCaptionsApi(configs...)
    api.Captions = captionsApi
    bifsApi, err := NewEncodingEncodingsStreamsBifsApi(configs...)
    api.Bifs = bifsApi
    hdrApi, err := NewEncodingEncodingsStreamsHdrApi(configs...)
    api.Hdr = hdrApi
    thumbnailsApi, err := NewEncodingEncodingsStreamsThumbnailsApi(configs...)
    api.Thumbnails = thumbnailsApi
    spritesApi, err := NewEncodingEncodingsStreamsSpritesApi(configs...)
    api.Sprites = spritesApi
    qcApi, err := NewEncodingEncodingsStreamsQcApi(configs...)
    api.Qc = qcApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsStreamsApi) Create(encodingId string, stream model.Stream) (*model.Stream, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.Stream
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/streams", &stream, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsStreamsApi) Delete(encodingId string, streamId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/streams/{stream_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsStreamsApi) Get(encodingId string, streamId string) (*model.Stream, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
    }

    var responseModel *model.Stream
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsStreamsApi) List(encodingId string, queryParams ...func(*query.StreamListQueryParams)) (*pagination.StreamsListPagination, error) {
    queryParameters := &query.StreamListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.StreamsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams", nil, &responseModel, reqParams)
    return responseModel, err
}

