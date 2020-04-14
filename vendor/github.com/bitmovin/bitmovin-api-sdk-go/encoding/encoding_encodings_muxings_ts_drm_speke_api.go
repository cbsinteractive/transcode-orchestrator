package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsTsDrmSpekeApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsTsDrmSpekeCustomdataApi
}

func NewEncodingEncodingsMuxingsTsDrmSpekeApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsTsDrmSpekeApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsTsDrmSpekeApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsTsDrmSpekeCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsTsDrmSpekeApi) Create(encodingId string, muxingId string, spekeDrm model.SpekeDrm) (*model.SpekeDrm, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.SpekeDrm
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/drm/speke", &spekeDrm, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsTsDrmSpekeApi) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/drm/speke/{drm_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsTsDrmSpekeApi) Get(encodingId string, muxingId string, drmId string) (*model.SpekeDrm, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel *model.SpekeDrm
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/drm/speke/{drm_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsTsDrmSpekeApi) List(encodingId string, muxingId string, queryParams ...func(*query.SpekeDrmListQueryParams)) (*pagination.SpekeDrmsListPagination, error) {
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
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/ts/{muxing_id}/drm/speke", nil, &responseModel, reqParams)
    return responseModel, err
}

