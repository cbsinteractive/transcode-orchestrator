package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsMp4DrmWidevineApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsMp4DrmWidevineCustomdataApi
}

func NewEncodingEncodingsMuxingsMp4DrmWidevineApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsMp4DrmWidevineApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsMp4DrmWidevineApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsMp4DrmWidevineCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsMp4DrmWidevineApi) Create(encodingId string, muxingId string, widevineDrm model.WidevineDrm) (*model.WidevineDrm, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.WidevineDrm
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/widevine", &widevineDrm, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsMp4DrmWidevineApi) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/widevine/{drm_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsMp4DrmWidevineApi) Get(encodingId string, muxingId string, drmId string) (*model.WidevineDrm, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel *model.WidevineDrm
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/widevine/{drm_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsMp4DrmWidevineApi) List(encodingId string, muxingId string, queryParams ...func(*query.WidevineDrmListQueryParams)) (*pagination.WidevineDrmsListPagination, error) {
    queryParameters := &query.WidevineDrmListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.WidevineDrmsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/widevine", nil, &responseModel, reqParams)
    return responseModel, err
}

