package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsInputStreamsTrimmingH264PictureTimingApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsInputStreamsTrimmingH264PictureTimingApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsInputStreamsTrimmingH264PictureTimingApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsInputStreamsTrimmingH264PictureTimingApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsInputStreamsTrimmingH264PictureTimingApi) Create(encodingId string, h264PictureTimingTrimmingInputStream model.H264PictureTimingTrimmingInputStream) (*model.H264PictureTimingTrimmingInputStream, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.H264PictureTimingTrimmingInputStream
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/input-streams/trimming/h264-picture-timing", &h264PictureTimingTrimmingInputStream, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsInputStreamsTrimmingH264PictureTimingApi) Delete(encodingId string, inputStreamId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/input-streams/trimming/h264-picture-timing/{input_stream_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsInputStreamsTrimmingH264PictureTimingApi) Get(encodingId string, inputStreamId string) (*model.H264PictureTimingTrimmingInputStream, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
    }

    var responseModel *model.H264PictureTimingTrimmingInputStream
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/trimming/h264-picture-timing/{input_stream_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsInputStreamsTrimmingH264PictureTimingApi) List(encodingId string, queryParams ...func(*query.H264PictureTimingTrimmingInputStreamListQueryParams)) (*pagination.H264PictureTimingTrimmingInputStreamsListPagination, error) {
    queryParameters := &query.H264PictureTimingTrimmingInputStreamListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.H264PictureTimingTrimmingInputStreamsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/trimming/h264-picture-timing", nil, &responseModel, reqParams)
    return responseModel, err
}

