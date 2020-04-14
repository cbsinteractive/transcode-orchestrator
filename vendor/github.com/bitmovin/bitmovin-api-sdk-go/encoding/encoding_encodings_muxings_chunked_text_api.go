package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsChunkedTextApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsChunkedTextCustomdataApi
}

func NewEncodingEncodingsMuxingsChunkedTextApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsChunkedTextApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsChunkedTextApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsChunkedTextCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsChunkedTextApi) Create(encodingId string, chunkedTextMuxing model.ChunkedTextMuxing) (*model.ChunkedTextMuxing, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.ChunkedTextMuxing
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/chunked-text", &chunkedTextMuxing, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsChunkedTextApi) Delete(encodingId string, muxingId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/chunked-text/{muxing_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsChunkedTextApi) Get(encodingId string, muxingId string) (*model.ChunkedTextMuxing, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.ChunkedTextMuxing
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/chunked-text/{muxing_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsChunkedTextApi) List(encodingId string, queryParams ...func(*query.ChunkedTextMuxingListQueryParams)) (*pagination.ChunkedTextMuxingsListPagination, error) {
    queryParameters := &query.ChunkedTextMuxingListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.ChunkedTextMuxingsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/chunked-text", nil, &responseModel, reqParams)
    return responseModel, err
}

