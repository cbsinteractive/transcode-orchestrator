package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsInputStreamsConcatenationApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsInputStreamsConcatenationApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsInputStreamsConcatenationApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsInputStreamsConcatenationApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsInputStreamsConcatenationApi) Create(encodingId string, concatenationInputStream model.ConcatenationInputStream) (*model.ConcatenationInputStream, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.ConcatenationInputStream
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/input-streams/concatenation", &concatenationInputStream, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsInputStreamsConcatenationApi) Delete(encodingId string, inputStreamId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/input-streams/concatenation/{input_stream_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsInputStreamsConcatenationApi) Get(encodingId string, inputStreamId string) (*model.ConcatenationInputStream, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
    }

    var responseModel *model.ConcatenationInputStream
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/concatenation/{input_stream_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsInputStreamsConcatenationApi) List(encodingId string, queryParams ...func(*query.ConcatenationInputStreamListQueryParams)) (*pagination.ConcatenationInputStreamsListPagination, error) {
    queryParameters := &query.ConcatenationInputStreamListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.ConcatenationInputStreamsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/concatenation", nil, &responseModel, reqParams)
    return responseModel, err
}

