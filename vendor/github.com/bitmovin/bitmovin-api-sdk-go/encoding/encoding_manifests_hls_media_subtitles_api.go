package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingManifestsHlsMediaSubtitlesApi struct {
    apiClient *common.ApiClient
}

func NewEncodingManifestsHlsMediaSubtitlesApi(configs ...func(*common.ApiClient)) (*EncodingManifestsHlsMediaSubtitlesApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsHlsMediaSubtitlesApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsHlsMediaSubtitlesApi) Create(manifestId string, subtitlesMediaInfo model.SubtitlesMediaInfo) (*model.SubtitlesMediaInfo, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }

    var responseModel *model.SubtitlesMediaInfo
    err := api.apiClient.Post("/encoding/manifests/hls/{manifest_id}/media/subtitles", &subtitlesMediaInfo, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsHlsMediaSubtitlesApi) Delete(manifestId string, mediaId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["media_id"] = mediaId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/manifests/hls/{manifest_id}/media/subtitles/{media_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsHlsMediaSubtitlesApi) Get(manifestId string, mediaId string) (*model.SubtitlesMediaInfo, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["media_id"] = mediaId
    }

    var responseModel *model.SubtitlesMediaInfo
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/media/subtitles/{media_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsHlsMediaSubtitlesApi) List(manifestId string, queryParams ...func(*query.SubtitlesMediaInfoListQueryParams)) (*pagination.SubtitlesMediaInfosListPagination, error) {
    queryParameters := &query.SubtitlesMediaInfoListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.SubtitlesMediaInfosListPagination
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/media/subtitles", nil, &responseModel, reqParams)
    return responseModel, err
}

