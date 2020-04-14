package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsProgressiveMovApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsProgressiveMovCustomdataApi
    Information *EncodingEncodingsMuxingsProgressiveMovInformationApi
}

func NewEncodingEncodingsMuxingsProgressiveMovApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsProgressiveMovApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsProgressiveMovApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsProgressiveMovCustomdataApi(configs...)
    api.Customdata = customdataApi
    informationApi, err := NewEncodingEncodingsMuxingsProgressiveMovInformationApi(configs...)
    api.Information = informationApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsProgressiveMovApi) Create(encodingId string, progressiveMovMuxing model.ProgressiveMovMuxing) (*model.ProgressiveMovMuxing, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.ProgressiveMovMuxing
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/progressive-mov", &progressiveMovMuxing, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsProgressiveMovApi) Delete(encodingId string, muxingId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/progressive-mov/{muxing_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsProgressiveMovApi) Get(encodingId string, muxingId string) (*model.ProgressiveMovMuxing, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.ProgressiveMovMuxing
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-mov/{muxing_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsProgressiveMovApi) List(encodingId string, queryParams ...func(*query.ProgressiveMovMuxingListQueryParams)) (*pagination.ProgressiveMovMuxingsListPagination, error) {
    queryParameters := &query.ProgressiveMovMuxingListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.ProgressiveMovMuxingsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-mov", nil, &responseModel, reqParams)
    return responseModel, err
}

