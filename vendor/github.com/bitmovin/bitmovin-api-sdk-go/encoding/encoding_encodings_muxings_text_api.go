package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsTextApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsTextCustomdataApi
}

func NewEncodingEncodingsMuxingsTextApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsTextApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsTextApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsTextCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsTextApi) Create(encodingId string, textMuxing model.TextMuxing) (*model.TextMuxing, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.TextMuxing
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/text", &textMuxing, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsTextApi) Delete(encodingId string, muxingId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/text/{muxing_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsTextApi) Get(encodingId string, muxingId string) (*model.TextMuxing, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.TextMuxing
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/text/{muxing_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsTextApi) List(encodingId string, queryParams ...func(*query.TextMuxingListQueryParams)) (*pagination.TextMuxingsListPagination, error) {
    queryParameters := &query.TextMuxingListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.TextMuxingsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/text", nil, &responseModel, reqParams)
    return responseModel, err
}

