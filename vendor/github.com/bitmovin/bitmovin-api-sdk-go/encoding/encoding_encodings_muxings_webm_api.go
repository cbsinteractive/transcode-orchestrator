package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsWebmApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsWebmCustomdataApi
    Drm *EncodingEncodingsMuxingsWebmDrmApi
}

func NewEncodingEncodingsMuxingsWebmApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsWebmApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsWebmApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsWebmCustomdataApi(configs...)
    api.Customdata = customdataApi
    drmApi, err := NewEncodingEncodingsMuxingsWebmDrmApi(configs...)
    api.Drm = drmApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsWebmApi) Create(encodingId string, webmMuxing model.WebmMuxing) (*model.WebmMuxing, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.WebmMuxing
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/webm", &webmMuxing, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsWebmApi) Delete(encodingId string, muxingId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/webm/{muxing_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsWebmApi) Get(encodingId string, muxingId string) (*model.WebmMuxing, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.WebmMuxing
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/webm/{muxing_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsWebmApi) List(encodingId string, queryParams ...func(*query.WebmMuxingListQueryParams)) (*pagination.WebmMuxingsListPagination, error) {
    queryParameters := &query.WebmMuxingListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.WebmMuxingsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/webm", nil, &responseModel, reqParams)
    return responseModel, err
}

