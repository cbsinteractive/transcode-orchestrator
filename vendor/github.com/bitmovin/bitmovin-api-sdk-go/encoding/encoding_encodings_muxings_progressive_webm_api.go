package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsProgressiveWebmApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsProgressiveWebmCustomdataApi
    Information *EncodingEncodingsMuxingsProgressiveWebmInformationApi
}

func NewEncodingEncodingsMuxingsProgressiveWebmApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsProgressiveWebmApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsProgressiveWebmApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsProgressiveWebmCustomdataApi(configs...)
    api.Customdata = customdataApi
    informationApi, err := NewEncodingEncodingsMuxingsProgressiveWebmInformationApi(configs...)
    api.Information = informationApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsProgressiveWebmApi) Create(encodingId string, progressiveWebmMuxing model.ProgressiveWebmMuxing) (*model.ProgressiveWebmMuxing, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.ProgressiveWebmMuxing
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/progressive-webm", &progressiveWebmMuxing, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsProgressiveWebmApi) Delete(encodingId string, muxingId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/progressive-webm/{muxing_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsProgressiveWebmApi) Get(encodingId string, muxingId string) (*model.ProgressiveWebmMuxing, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.ProgressiveWebmMuxing
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-webm/{muxing_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsProgressiveWebmApi) List(encodingId string, queryParams ...func(*query.ProgressiveWebmMuxingListQueryParams)) (*pagination.ProgressiveWebmMuxingsListPagination, error) {
    queryParameters := &query.ProgressiveWebmMuxingListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.ProgressiveWebmMuxingsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-webm", nil, &responseModel, reqParams)
    return responseModel, err
}

