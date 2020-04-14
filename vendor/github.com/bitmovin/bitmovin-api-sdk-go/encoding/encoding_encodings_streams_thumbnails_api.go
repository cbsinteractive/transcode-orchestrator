package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsStreamsThumbnailsApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsStreamsThumbnailsCustomdataApi
}

func NewEncodingEncodingsStreamsThumbnailsApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsStreamsThumbnailsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsStreamsThumbnailsApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsStreamsThumbnailsCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsStreamsThumbnailsApi) Create(encodingId string, streamId string, thumbnail model.Thumbnail) (*model.Thumbnail, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
    }

    var responseModel *model.Thumbnail
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/streams/{stream_id}/thumbnails", &thumbnail, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsStreamsThumbnailsApi) Delete(encodingId string, streamId string, thumbnailId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.PathParams["thumbnail_id"] = thumbnailId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/streams/{stream_id}/thumbnails/{thumbnail_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsStreamsThumbnailsApi) Get(encodingId string, streamId string, thumbnailId string) (*model.Thumbnail, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.PathParams["thumbnail_id"] = thumbnailId
    }

    var responseModel *model.Thumbnail
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/thumbnails/{thumbnail_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsStreamsThumbnailsApi) List(encodingId string, streamId string, queryParams ...func(*query.ThumbnailListQueryParams)) (*pagination.ThumbnailsListPagination, error) {
    queryParameters := &query.ThumbnailListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.ThumbnailsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/thumbnails", nil, &responseModel, reqParams)
    return responseModel, err
}

