package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingManifestsHlsMediaVideoApi struct {
    apiClient *common.ApiClient
}

func NewEncodingManifestsHlsMediaVideoApi(configs ...func(*common.ApiClient)) (*EncodingManifestsHlsMediaVideoApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsHlsMediaVideoApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsHlsMediaVideoApi) Create(manifestId string, videoMediaInfo model.VideoMediaInfo) (*model.VideoMediaInfo, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }

    var responseModel *model.VideoMediaInfo
    err := api.apiClient.Post("/encoding/manifests/hls/{manifest_id}/media/video", &videoMediaInfo, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsHlsMediaVideoApi) Delete(manifestId string, mediaId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["media_id"] = mediaId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/manifests/hls/{manifest_id}/media/video/{media_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsHlsMediaVideoApi) Get(manifestId string, mediaId string) (*model.VideoMediaInfo, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["media_id"] = mediaId
    }

    var responseModel *model.VideoMediaInfo
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/media/video/{media_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsHlsMediaVideoApi) List(manifestId string, queryParams ...func(*query.VideoMediaInfoListQueryParams)) (*pagination.VideoMediaInfosListPagination, error) {
    queryParameters := &query.VideoMediaInfoListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.VideoMediaInfosListPagination
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/media/video", nil, &responseModel, reqParams)
    return responseModel, err
}

