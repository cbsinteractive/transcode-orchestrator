package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsInputStreamsFileApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsInputStreamsFileApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsInputStreamsFileApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsInputStreamsFileApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsInputStreamsFileApi) Create(encodingId string, fileInputStream model.FileInputStream) (*model.FileInputStream, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.FileInputStream
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/input-streams/file", &fileInputStream, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsInputStreamsFileApi) Delete(encodingId string, inputStreamId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/input-streams/file/{input_stream_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsInputStreamsFileApi) Get(encodingId string, inputStreamId string) (*model.FileInputStream, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
    }

    var responseModel *model.FileInputStream
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/file/{input_stream_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsInputStreamsFileApi) List(encodingId string, queryParams ...func(*query.FileInputStreamListQueryParams)) (*pagination.FileInputStreamsListPagination, error) {
    queryParameters := &query.FileInputStreamListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.FileInputStreamsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/file", nil, &responseModel, reqParams)
    return responseModel, err
}

