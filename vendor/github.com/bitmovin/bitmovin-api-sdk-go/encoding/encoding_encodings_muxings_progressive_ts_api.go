package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsProgressiveTsApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsProgressiveTsCustomdataApi
    Information *EncodingEncodingsMuxingsProgressiveTsInformationApi
    Id3 *EncodingEncodingsMuxingsProgressiveTsId3Api
    Drm *EncodingEncodingsMuxingsProgressiveTsDrmApi
}

func NewEncodingEncodingsMuxingsProgressiveTsApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsProgressiveTsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsProgressiveTsApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsProgressiveTsCustomdataApi(configs...)
    api.Customdata = customdataApi
    informationApi, err := NewEncodingEncodingsMuxingsProgressiveTsInformationApi(configs...)
    api.Information = informationApi
    id3Api, err := NewEncodingEncodingsMuxingsProgressiveTsId3Api(configs...)
    api.Id3 = id3Api
    drmApi, err := NewEncodingEncodingsMuxingsProgressiveTsDrmApi(configs...)
    api.Drm = drmApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsProgressiveTsApi) Create(encodingId string, progressiveTsMuxing model.ProgressiveTsMuxing) (*model.ProgressiveTsMuxing, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.ProgressiveTsMuxing
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/progressive-ts", &progressiveTsMuxing, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsProgressiveTsApi) Delete(encodingId string, muxingId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsProgressiveTsApi) Get(encodingId string, muxingId string) (*model.ProgressiveTsMuxing, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.ProgressiveTsMuxing
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsProgressiveTsApi) List(encodingId string, queryParams ...func(*query.ProgressiveTsMuxingListQueryParams)) (*pagination.ProgressiveTsMuxingsListPagination, error) {
    queryParameters := &query.ProgressiveTsMuxingListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.ProgressiveTsMuxingsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-ts", nil, &responseModel, reqParams)
    return responseModel, err
}

