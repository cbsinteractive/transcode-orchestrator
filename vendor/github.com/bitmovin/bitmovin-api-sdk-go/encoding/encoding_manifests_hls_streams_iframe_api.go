package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingManifestsHlsStreamsIframeApi struct {
    apiClient *common.ApiClient
}

func NewEncodingManifestsHlsStreamsIframeApi(configs ...func(*common.ApiClient)) (*EncodingManifestsHlsStreamsIframeApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsHlsStreamsIframeApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsHlsStreamsIframeApi) Create(manifestId string, streamId string, iFramePlaylist model.IFramePlaylist) (*model.IFramePlaylist, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["stream_id"] = streamId
    }

    var responseModel *model.IFramePlaylist
    err := api.apiClient.Post("/encoding/manifests/hls/{manifest_id}/streams/{stream_id}/iframe", &iFramePlaylist, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsHlsStreamsIframeApi) Delete(manifestId string, streamId string, iframeId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["stream_id"] = streamId
        params.PathParams["iframe_id"] = iframeId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/manifests/hls/{manifest_id}/streams/{stream_id}/iframe/{iframe_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsHlsStreamsIframeApi) Get(manifestId string, streamId string, iframeId string) (*model.IFramePlaylist, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["stream_id"] = streamId
        params.PathParams["iframe_id"] = iframeId
    }

    var responseModel *model.IFramePlaylist
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/streams/{stream_id}/iframe/{iframe_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsHlsStreamsIframeApi) List(manifestId string, streamId string, queryParams ...func(*query.IFramePlaylistListQueryParams)) (*pagination.IFramePlaylistsListPagination, error) {
    queryParameters := &query.IFramePlaylistListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["stream_id"] = streamId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.IFramePlaylistsListPagination
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/streams/{stream_id}/iframe", nil, &responseModel, reqParams)
    return responseModel, err
}

