package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsInputStreamsCaptionsCea708Api struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsInputStreamsCaptionsCea708Api(configs ...func(*common.ApiClient)) (*EncodingEncodingsInputStreamsCaptionsCea708Api, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsInputStreamsCaptionsCea708Api{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsInputStreamsCaptionsCea708Api) Create(encodingId string, cea708CaptionInputStream model.Cea708CaptionInputStream) (*model.Cea708CaptionInputStream, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.Cea708CaptionInputStream
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/input-streams/captions/cea708", &cea708CaptionInputStream, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsInputStreamsCaptionsCea708Api) Delete(encodingId string, inputStreamId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/input-streams/captions/cea708/{input_stream_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsInputStreamsCaptionsCea708Api) Get(encodingId string, inputStreamId string) (*model.Cea708CaptionInputStream, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
    }

    var responseModel *model.Cea708CaptionInputStream
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/captions/cea708/{input_stream_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsInputStreamsCaptionsCea708Api) List(encodingId string, queryParams ...func(*query.Cea708CaptionInputStreamListQueryParams)) (*pagination.Cea708CaptionInputStreamsListPagination, error) {
    queryParameters := &query.Cea708CaptionInputStreamListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.Cea708CaptionInputStreamsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/captions/cea708", nil, &responseModel, reqParams)
    return responseModel, err
}

