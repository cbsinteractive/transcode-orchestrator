package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingManifestsHlsMediaAudioApi struct {
    apiClient *common.ApiClient
}

func NewEncodingManifestsHlsMediaAudioApi(configs ...func(*common.ApiClient)) (*EncodingManifestsHlsMediaAudioApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsHlsMediaAudioApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsHlsMediaAudioApi) Create(manifestId string, audioMediaInfo model.AudioMediaInfo) (*model.AudioMediaInfo, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }

    var responseModel *model.AudioMediaInfo
    err := api.apiClient.Post("/encoding/manifests/hls/{manifest_id}/media/audio", &audioMediaInfo, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsHlsMediaAudioApi) Delete(manifestId string, mediaId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["media_id"] = mediaId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/manifests/hls/{manifest_id}/media/audio/{media_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsHlsMediaAudioApi) Get(manifestId string, mediaId string) (*model.AudioMediaInfo, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["media_id"] = mediaId
    }

    var responseModel *model.AudioMediaInfo
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/media/audio/{media_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsHlsMediaAudioApi) List(manifestId string, queryParams ...func(*query.AudioMediaInfoListQueryParams)) (*pagination.AudioMediaInfosListPagination, error) {
    queryParameters := &query.AudioMediaInfoListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.AudioMediaInfosListPagination
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/media/audio", nil, &responseModel, reqParams)
    return responseModel, err
}

