package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsBroadcastTsApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsBroadcastTsCustomdataApi
    Information *EncodingEncodingsMuxingsBroadcastTsInformationApi
}

func NewEncodingEncodingsMuxingsBroadcastTsApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsBroadcastTsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsBroadcastTsApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsBroadcastTsCustomdataApi(configs...)
    api.Customdata = customdataApi
    informationApi, err := NewEncodingEncodingsMuxingsBroadcastTsInformationApi(configs...)
    api.Information = informationApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsBroadcastTsApi) Create(encodingId string, broadcastTsMuxing model.BroadcastTsMuxing) (*model.BroadcastTsMuxing, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.BroadcastTsMuxing
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/broadcast-ts", &broadcastTsMuxing, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsBroadcastTsApi) Delete(encodingId string, muxingId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/broadcast-ts/{muxing_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsBroadcastTsApi) Get(encodingId string, muxingId string) (*model.BroadcastTsMuxing, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.BroadcastTsMuxing
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/broadcast-ts/{muxing_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsBroadcastTsApi) List(encodingId string, queryParams ...func(*query.BroadcastTsMuxingListQueryParams)) (*pagination.BroadcastTsMuxingsListPagination, error) {
    queryParameters := &query.BroadcastTsMuxingListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.BroadcastTsMuxingsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/broadcast-ts", nil, &responseModel, reqParams)
    return responseModel, err
}

