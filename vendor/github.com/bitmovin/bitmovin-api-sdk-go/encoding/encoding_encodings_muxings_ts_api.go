package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsTsApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsTsCustomdataApi
    Drm *EncodingEncodingsMuxingsTsDrmApi
}

func NewEncodingEncodingsMuxingsTsApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsTsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsTsApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsTsCustomdataApi(configs...)
    api.Customdata = customdataApi
    drmApi, err := NewEncodingEncodingsMuxingsTsDrmApi(configs...)
    api.Drm = drmApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsTsApi) Create(encodingId string, tsMuxing model.TsMuxing) (*model.TsMuxing, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.TsMuxing
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/ts", &tsMuxing, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsTsApi) Delete(encodingId string, muxingId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsTsApi) Get(encodingId string, muxingId string) (*model.TsMuxing, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.TsMuxing
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsTsApi) List(encodingId string, queryParams ...func(*query.TsMuxingListQueryParams)) (*pagination.TsMuxingsListPagination, error) {
    queryParameters := &query.TsMuxingListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.TsMuxingsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/ts", nil, &responseModel, reqParams)
    return responseModel, err
}

