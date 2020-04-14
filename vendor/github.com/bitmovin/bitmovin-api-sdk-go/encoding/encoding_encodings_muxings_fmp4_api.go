package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsFmp4Api struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsFmp4CustomdataApi
    Drm *EncodingEncodingsMuxingsFmp4DrmApi
}

func NewEncodingEncodingsMuxingsFmp4Api(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsFmp4Api, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsFmp4Api{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsFmp4CustomdataApi(configs...)
    api.Customdata = customdataApi
    drmApi, err := NewEncodingEncodingsMuxingsFmp4DrmApi(configs...)
    api.Drm = drmApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsFmp4Api) Create(encodingId string, fmp4Muxing model.Fmp4Muxing) (*model.Fmp4Muxing, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.Fmp4Muxing
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/fmp4", &fmp4Muxing, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsFmp4Api) Delete(encodingId string, muxingId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsFmp4Api) Get(encodingId string, muxingId string) (*model.Fmp4Muxing, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.Fmp4Muxing
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsFmp4Api) List(encodingId string, queryParams ...func(*query.Fmp4MuxingListQueryParams)) (*pagination.Fmp4MuxingsListPagination, error) {
    queryParameters := &query.Fmp4MuxingListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.Fmp4MuxingsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4", nil, &responseModel, reqParams)
    return responseModel, err
}

