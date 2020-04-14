package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsMp4DrmSpekeApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsMp4DrmSpekeCustomdataApi
}

func NewEncodingEncodingsMuxingsMp4DrmSpekeApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsMp4DrmSpekeApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsMp4DrmSpekeApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsMp4DrmSpekeCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsMp4DrmSpekeApi) Create(encodingId string, muxingId string, spekeDrm model.SpekeDrm) (*model.SpekeDrm, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.SpekeDrm
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/speke", &spekeDrm, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsMp4DrmSpekeApi) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/speke/{drm_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsMp4DrmSpekeApi) Get(encodingId string, muxingId string, drmId string) (*model.SpekeDrm, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel *model.SpekeDrm
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/speke/{drm_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsMp4DrmSpekeApi) List(encodingId string, muxingId string, queryParams ...func(*query.SpekeDrmListQueryParams)) (*pagination.SpekeDrmsListPagination, error) {
    queryParameters := &query.SpekeDrmListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.SpekeDrmsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/speke", nil, &responseModel, reqParams)
    return responseModel, err
}

