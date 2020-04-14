package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsStreamsFiltersApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsStreamsFiltersApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsStreamsFiltersApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsStreamsFiltersApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsStreamsFiltersApi) Create(encodingId string, streamId string, streamFilter []model.StreamFilter) (*model.StreamFilterList, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
    }

    var responseModel *model.StreamFilterList
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/streams/{stream_id}/filters", &streamFilter, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsStreamsFiltersApi) Delete(encodingId string, streamId string, filterId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.PathParams["filter_id"] = filterId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/streams/{stream_id}/filters/{filter_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsStreamsFiltersApi) DeleteAll(encodingId string, streamId string) (*model.BitmovinResponseList, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
    }

    var responseModel *model.BitmovinResponseList
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/streams/{stream_id}/filters", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsStreamsFiltersApi) List(encodingId string, streamId string, queryParams ...func(*query.StreamFilterListListQueryParams)) (*model.StreamFilterList, error) {
    queryParameters := &query.StreamFilterListListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.QueryParams = queryParameters
    }

    var responseModel *model.StreamFilterList
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/filters", nil, &responseModel, reqParams)
    return responseModel, err
}

