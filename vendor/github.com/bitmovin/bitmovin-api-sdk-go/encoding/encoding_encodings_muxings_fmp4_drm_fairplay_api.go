package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsFmp4DrmFairplayApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsFmp4DrmFairplayCustomdataApi
}

func NewEncodingEncodingsMuxingsFmp4DrmFairplayApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsFmp4DrmFairplayApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsFmp4DrmFairplayApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsFmp4DrmFairplayCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsFmp4DrmFairplayApi) Create(encodingId string, muxingId string, fairPlayDrm model.FairPlayDrm) (*model.FairPlayDrm, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.FairPlayDrm
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/fairplay", &fairPlayDrm, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsFmp4DrmFairplayApi) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/fairplay/{drm_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsFmp4DrmFairplayApi) Get(encodingId string, muxingId string, drmId string) (*model.FairPlayDrm, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel *model.FairPlayDrm
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/fairplay/{drm_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsFmp4DrmFairplayApi) List(encodingId string, muxingId string, queryParams ...func(*query.FairPlayDrmListQueryParams)) (*pagination.FairPlayDrmsListPagination, error) {
    queryParameters := &query.FairPlayDrmListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.FairPlayDrmsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/fairplay", nil, &responseModel, reqParams)
    return responseModel, err
}

