package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingManifestsHlsStreamsApi struct {
    apiClient *common.ApiClient
    CustomTags *EncodingManifestsHlsStreamsCustomTagsApi
    Iframe *EncodingManifestsHlsStreamsIframeApi
}

func NewEncodingManifestsHlsStreamsApi(configs ...func(*common.ApiClient)) (*EncodingManifestsHlsStreamsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsHlsStreamsApi{apiClient: apiClient}

    customTagsApi, err := NewEncodingManifestsHlsStreamsCustomTagsApi(configs...)
    api.CustomTags = customTagsApi
    iframeApi, err := NewEncodingManifestsHlsStreamsIframeApi(configs...)
    api.Iframe = iframeApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsHlsStreamsApi) Create(manifestId string, streamInfo model.StreamInfo) (*model.StreamInfo, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }

    var responseModel *model.StreamInfo
    err := api.apiClient.Post("/encoding/manifests/hls/{manifest_id}/streams", &streamInfo, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsHlsStreamsApi) Delete(manifestId string, streamId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["stream_id"] = streamId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/manifests/hls/{manifest_id}/streams/{stream_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsHlsStreamsApi) Get(manifestId string, streamId string) (*model.StreamInfo, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["stream_id"] = streamId
    }

    var responseModel *model.StreamInfo
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/streams/{stream_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsHlsStreamsApi) List(manifestId string, queryParams ...func(*query.StreamInfoListQueryParams)) (*pagination.StreamInfosListPagination, error) {
    queryParameters := &query.StreamInfoListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.StreamInfosListPagination
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/streams", nil, &responseModel, reqParams)
    return responseModel, err
}

