package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsProgressiveTsId3RawApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsProgressiveTsId3RawCustomdataApi
}

func NewEncodingEncodingsMuxingsProgressiveTsId3RawApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsProgressiveTsId3RawApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsProgressiveTsId3RawApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsProgressiveTsId3RawCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsProgressiveTsId3RawApi) Create(encodingId string, muxingId string, rawId3Tag model.RawId3Tag) (*model.RawId3Tag, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.RawId3Tag
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3/raw", &rawId3Tag, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsProgressiveTsId3RawApi) Delete(encodingId string, muxingId string, id3TagId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["id3_tag_id"] = id3TagId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3/raw/{id3_tag_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsProgressiveTsId3RawApi) Get(encodingId string, muxingId string, id3TagId string) (*model.RawId3Tag, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["id3_tag_id"] = id3TagId
    }

    var responseModel *model.RawId3Tag
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3/raw/{id3_tag_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsProgressiveTsId3RawApi) List(encodingId string, muxingId string, queryParams ...func(*query.RawId3TagListQueryParams)) (*pagination.RawId3TagsListPagination, error) {
    queryParameters := &query.RawId3TagListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.RawId3TagsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3/raw", nil, &responseModel, reqParams)
    return responseModel, err
}

