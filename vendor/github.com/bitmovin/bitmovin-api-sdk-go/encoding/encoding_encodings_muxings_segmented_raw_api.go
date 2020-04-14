package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsSegmentedRawApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsSegmentedRawCustomdataApi
}

func NewEncodingEncodingsMuxingsSegmentedRawApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsSegmentedRawApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsSegmentedRawApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsSegmentedRawCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsSegmentedRawApi) Create(encodingId string, segmentedRawMuxing model.SegmentedRawMuxing) (*model.SegmentedRawMuxing, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.SegmentedRawMuxing
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/segmented-raw", &segmentedRawMuxing, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsSegmentedRawApi) Delete(encodingId string, muxingId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/segmented-raw/{muxing_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsSegmentedRawApi) Get(encodingId string, muxingId string) (*model.SegmentedRawMuxing, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.SegmentedRawMuxing
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/segmented-raw/{muxing_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsSegmentedRawApi) List(encodingId string, queryParams ...func(*query.SegmentedRawMuxingListQueryParams)) (*pagination.SegmentedRawMuxingsListPagination, error) {
    queryParameters := &query.SegmentedRawMuxingListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.SegmentedRawMuxingsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/segmented-raw", nil, &responseModel, reqParams)
    return responseModel, err
}

