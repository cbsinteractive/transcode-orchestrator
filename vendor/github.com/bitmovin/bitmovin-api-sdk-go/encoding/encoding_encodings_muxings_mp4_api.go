package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsMp4Api struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsMp4CustomdataApi
    Information *EncodingEncodingsMuxingsMp4InformationApi
    Drm *EncodingEncodingsMuxingsMp4DrmApi
}

func NewEncodingEncodingsMuxingsMp4Api(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsMp4Api, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsMp4Api{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsMp4CustomdataApi(configs...)
    api.Customdata = customdataApi
    informationApi, err := NewEncodingEncodingsMuxingsMp4InformationApi(configs...)
    api.Information = informationApi
    drmApi, err := NewEncodingEncodingsMuxingsMp4DrmApi(configs...)
    api.Drm = drmApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsMp4Api) Create(encodingId string, mp4Muxing model.Mp4Muxing) (*model.Mp4Muxing, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.Mp4Muxing
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/mp4", &mp4Muxing, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsMp4Api) Delete(encodingId string, muxingId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsMp4Api) Get(encodingId string, muxingId string) (*model.Mp4Muxing, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.Mp4Muxing
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp4/{muxing_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsMp4Api) List(encodingId string, queryParams ...func(*query.Mp4MuxingListQueryParams)) (*pagination.Mp4MuxingsListPagination, error) {
    queryParameters := &query.Mp4MuxingListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.Mp4MuxingsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp4", nil, &responseModel, reqParams)
    return responseModel, err
}

