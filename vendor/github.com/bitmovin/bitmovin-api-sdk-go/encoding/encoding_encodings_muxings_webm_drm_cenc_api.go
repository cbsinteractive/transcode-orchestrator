package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsWebmDrmCencApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsWebmDrmCencCustomdataApi
}

func NewEncodingEncodingsMuxingsWebmDrmCencApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsWebmDrmCencApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsWebmDrmCencApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsWebmDrmCencCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsWebmDrmCencApi) Create(encodingId string, muxingId string, cencDrm model.CencDrm) (*model.CencDrm, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.CencDrm
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/webm/{muxing_id}/drm/cenc", &cencDrm, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsWebmDrmCencApi) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/webm/{muxing_id}/drm/cenc/{drm_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsWebmDrmCencApi) Get(encodingId string, muxingId string, drmId string) (*model.CencDrm, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel *model.CencDrm
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/webm/{muxing_id}/drm/cenc/{drm_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsWebmDrmCencApi) List(encodingId string, muxingId string, queryParams ...func(*query.CencDrmListQueryParams)) (*pagination.CencDrmsListPagination, error) {
    queryParameters := &query.CencDrmListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.CencDrmsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/webm/{muxing_id}/drm/cenc", nil, &responseModel, reqParams)
    return responseModel, err
}

