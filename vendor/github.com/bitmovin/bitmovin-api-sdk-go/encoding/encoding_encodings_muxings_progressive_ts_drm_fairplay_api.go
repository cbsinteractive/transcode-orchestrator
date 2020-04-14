package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsProgressiveTsDrmFairplayApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsProgressiveTsDrmFairplayCustomdataApi
}

func NewEncodingEncodingsMuxingsProgressiveTsDrmFairplayApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsProgressiveTsDrmFairplayApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsProgressiveTsDrmFairplayApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsProgressiveTsDrmFairplayCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsProgressiveTsDrmFairplayApi) Create(encodingId string, muxingId string, fairPlayDrm model.FairPlayDrm) (*model.FairPlayDrm, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.FairPlayDrm
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/drm/fairplay", &fairPlayDrm, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsProgressiveTsDrmFairplayApi) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/drm/fairplay/{drm_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsProgressiveTsDrmFairplayApi) Get(encodingId string, muxingId string, drmId string) (*model.FairPlayDrm, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel *model.FairPlayDrm
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/drm/fairplay/{drm_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsProgressiveTsDrmFairplayApi) List(encodingId string, muxingId string, queryParams ...func(*query.FairPlayDrmListQueryParams)) (*pagination.FairPlayDrmsListPagination, error) {
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
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/drm/fairplay", nil, &responseModel, reqParams)
    return responseModel, err
}

