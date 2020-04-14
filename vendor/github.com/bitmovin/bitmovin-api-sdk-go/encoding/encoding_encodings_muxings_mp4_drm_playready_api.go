package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsMp4DrmPlayreadyApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsMp4DrmPlayreadyCustomdataApi
}

func NewEncodingEncodingsMuxingsMp4DrmPlayreadyApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsMp4DrmPlayreadyApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsMp4DrmPlayreadyApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsMp4DrmPlayreadyCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsMp4DrmPlayreadyApi) Create(encodingId string, muxingId string, playReadyDrm model.PlayReadyDrm) (*model.PlayReadyDrm, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.PlayReadyDrm
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/playready", &playReadyDrm, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsMp4DrmPlayreadyApi) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/playready/{drm_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsMp4DrmPlayreadyApi) Get(encodingId string, muxingId string, drmId string) (*model.PlayReadyDrm, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel *model.PlayReadyDrm
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/playready/{drm_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsMp4DrmPlayreadyApi) List(encodingId string, muxingId string, queryParams ...func(*query.PlayReadyDrmListQueryParams)) (*pagination.PlayReadyDrmsListPagination, error) {
    queryParameters := &query.PlayReadyDrmListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.PlayReadyDrmsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}/drm/playready", nil, &responseModel, reqParams)
    return responseModel, err
}

