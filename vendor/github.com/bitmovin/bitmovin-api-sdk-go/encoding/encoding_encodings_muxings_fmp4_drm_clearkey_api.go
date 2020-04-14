package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsFmp4DrmClearkeyApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsFmp4DrmClearkeyCustomdataApi
}

func NewEncodingEncodingsMuxingsFmp4DrmClearkeyApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsFmp4DrmClearkeyApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsFmp4DrmClearkeyApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsFmp4DrmClearkeyCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsFmp4DrmClearkeyApi) Create(encodingId string, muxingId string, clearKeyDrm model.ClearKeyDrm) (*model.ClearKeyDrm, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.ClearKeyDrm
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/clearkey", &clearKeyDrm, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsFmp4DrmClearkeyApi) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/clearkey/{drm_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsFmp4DrmClearkeyApi) Get(encodingId string, muxingId string, drmId string) (*model.ClearKeyDrm, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel *model.ClearKeyDrm
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/clearkey/{drm_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsFmp4DrmClearkeyApi) List(encodingId string, muxingId string, queryParams ...func(*query.ClearKeyDrmListQueryParams)) (*pagination.ClearKeyDrmsListPagination, error) {
    queryParameters := &query.ClearKeyDrmListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.ClearKeyDrmsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/clearkey", nil, &responseModel, reqParams)
    return responseModel, err
}

