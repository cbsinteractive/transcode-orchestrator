package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsMp4DrmMarlinApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsMp4DrmMarlinCustomdataApi
}

func NewEncodingEncodingsMuxingsMp4DrmMarlinApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsMp4DrmMarlinApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsMp4DrmMarlinApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsMp4DrmMarlinCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsMp4DrmMarlinApi) Create(encodingId string, muxingId string, marlinDrm model.MarlinDrm) (*model.MarlinDrm, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.MarlinDrm
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/marlin", &marlinDrm, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsMp4DrmMarlinApi) Delete(encodingId string, muxingId string, drmId string) (*model.MarlinDrm, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel *model.MarlinDrm
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/marlin/{drm_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsMp4DrmMarlinApi) Get(encodingId string, muxingId string, drmId string) (*model.MarlinDrm, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel *model.MarlinDrm
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/marlin/{drm_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsMp4DrmMarlinApi) List(encodingId string, muxingId string, queryParams ...func(*query.MarlinDrmListQueryParams)) (*pagination.MarlinDrmsListPagination, error) {
    queryParameters := &query.MarlinDrmListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.MarlinDrmsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/marlin", nil, &responseModel, reqParams)
    return responseModel, err
}

