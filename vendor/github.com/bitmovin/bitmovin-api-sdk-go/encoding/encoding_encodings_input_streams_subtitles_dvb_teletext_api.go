package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsInputStreamsSubtitlesDvbTeletextApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsInputStreamsSubtitlesDvbTeletextApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsInputStreamsSubtitlesDvbTeletextApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsInputStreamsSubtitlesDvbTeletextApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsInputStreamsSubtitlesDvbTeletextApi) Create(encodingId string, dvbTeletextInputStream model.DvbTeletextInputStream) (*model.DvbTeletextInputStream, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.DvbTeletextInputStream
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/input-streams/subtitles/dvb-teletext", &dvbTeletextInputStream, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsInputStreamsSubtitlesDvbTeletextApi) Delete(encodingId string, inputStreamId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/input-streams/subtitles/dvb-teletext/{input_stream_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsInputStreamsSubtitlesDvbTeletextApi) Get(encodingId string, inputStreamId string) (*model.DvbTeletextInputStream, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
    }

    var responseModel *model.DvbTeletextInputStream
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/subtitles/dvb-teletext/{input_stream_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsInputStreamsSubtitlesDvbTeletextApi) List(encodingId string, queryParams ...func(*query.DvbTeletextInputStreamListQueryParams)) (*pagination.DvbTeletextInputStreamsListPagination, error) {
    queryParameters := &query.DvbTeletextInputStreamListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.DvbTeletextInputStreamsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/subtitles/dvb-teletext", nil, &responseModel, reqParams)
    return responseModel, err
}

