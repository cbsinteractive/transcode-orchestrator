package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsStreamsSpritesApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsStreamsSpritesCustomdataApi
}

func NewEncodingEncodingsStreamsSpritesApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsStreamsSpritesApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsStreamsSpritesApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsStreamsSpritesCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsStreamsSpritesApi) Create(encodingId string, streamId string, sprite model.Sprite) (*model.Sprite, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
    }

    var responseModel *model.Sprite
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/streams/{stream_id}/sprites", &sprite, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsStreamsSpritesApi) Delete(encodingId string, streamId string, spriteId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.PathParams["sprite_id"] = spriteId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/streams/{stream_id}/sprites/{sprite_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsStreamsSpritesApi) Get(encodingId string, streamId string, spriteId string) (*model.Sprite, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.PathParams["sprite_id"] = spriteId
    }

    var responseModel *model.Sprite
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/sprites/{sprite_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsStreamsSpritesApi) List(encodingId string, streamId string, queryParams ...func(*query.SpriteListQueryParams)) (*pagination.SpritesListPagination, error) {
    queryParameters := &query.SpriteListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.SpritesListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/sprites", nil, &responseModel, reqParams)
    return responseModel, err
}

