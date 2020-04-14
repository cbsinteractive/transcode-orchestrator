package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsProgressiveTsDrmSpekeApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsProgressiveTsDrmSpekeCustomdataApi
}

func NewEncodingEncodingsMuxingsProgressiveTsDrmSpekeApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsProgressiveTsDrmSpekeApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsProgressiveTsDrmSpekeApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsProgressiveTsDrmSpekeCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsProgressiveTsDrmSpekeApi) Create(encodingId string, muxingId string, spekeDrm model.SpekeDrm) (*model.SpekeDrm, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.SpekeDrm
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/drm/speke", &spekeDrm, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsProgressiveTsDrmSpekeApi) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/drm/speke/{drm_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsProgressiveTsDrmSpekeApi) Get(encodingId string, muxingId string, drmId string) (*model.SpekeDrm, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel *model.SpekeDrm
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/drm/speke/{drm_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsProgressiveTsDrmSpekeApi) List(encodingId string, muxingId string, queryParams ...func(*query.SpekeDrmListQueryParams)) (*pagination.SpekeDrmsListPagination, error) {
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
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/drm/speke", nil, &responseModel, reqParams)
    return responseModel, err
}

