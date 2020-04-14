package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsProgressiveTsId3FrameIdApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsProgressiveTsId3FrameIdCustomdataApi
}

func NewEncodingEncodingsMuxingsProgressiveTsId3FrameIdApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsProgressiveTsId3FrameIdApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsProgressiveTsId3FrameIdApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsProgressiveTsId3FrameIdCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsProgressiveTsId3FrameIdApi) Create(encodingId string, muxingId string, frameIdId3Tag model.FrameIdId3Tag) (*model.FrameIdId3Tag, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.FrameIdId3Tag
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3/frame-id", &frameIdId3Tag, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsProgressiveTsId3FrameIdApi) Delete(encodingId string, muxingId string, id3TagId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["id3_tag_id"] = id3TagId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3/frame-id/{id3_tag_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsProgressiveTsId3FrameIdApi) Get(encodingId string, muxingId string, id3TagId string) (*model.FrameIdId3Tag, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["id3_tag_id"] = id3TagId
    }

    var responseModel *model.FrameIdId3Tag
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3/frame-id/{id3_tag_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsProgressiveTsId3FrameIdApi) List(encodingId string, muxingId string, queryParams ...func(*query.FrameIdId3TagListQueryParams)) (*pagination.FrameIdId3TagsListPagination, error) {
    queryParameters := &query.FrameIdId3TagListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.FrameIdId3TagsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3/frame-id", nil, &responseModel, reqParams)
    return responseModel, err
}

