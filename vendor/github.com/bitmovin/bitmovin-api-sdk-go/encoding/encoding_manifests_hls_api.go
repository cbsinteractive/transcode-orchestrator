package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingManifestsHlsApi struct {
    apiClient *common.ApiClient
    Default *EncodingManifestsHlsDefaultApi
    Customdata *EncodingManifestsHlsCustomdataApi
    Streams *EncodingManifestsHlsStreamsApi
    Media *EncodingManifestsHlsMediaApi
}

func NewEncodingManifestsHlsApi(configs ...func(*common.ApiClient)) (*EncodingManifestsHlsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsHlsApi{apiClient: apiClient}

    defaultApi, err := NewEncodingManifestsHlsDefaultApi(configs...)
    api.Default = defaultApi
    customdataApi, err := NewEncodingManifestsHlsCustomdataApi(configs...)
    api.Customdata = customdataApi
    streamsApi, err := NewEncodingManifestsHlsStreamsApi(configs...)
    api.Streams = streamsApi
    mediaApi, err := NewEncodingManifestsHlsMediaApi(configs...)
    api.Media = mediaApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsHlsApi) Create(hlsManifest model.HlsManifest) (*model.HlsManifest, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.HlsManifest
    err := api.apiClient.Post("/encoding/manifests/hls", &hlsManifest, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsHlsApi) Delete(manifestId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/manifests/hls/{manifest_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsHlsApi) Get(manifestId string) (*model.HlsManifest, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }

    var responseModel *model.HlsManifest
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsHlsApi) List(queryParams ...func(*query.HlsManifestListQueryParams)) (*pagination.HlsManifestsListPagination, error) {
    queryParameters := &query.HlsManifestListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.HlsManifestsListPagination
    err := api.apiClient.Get("/encoding/manifests/hls", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsHlsApi) Start(manifestId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Post("/encoding/manifests/hls/{manifest_id}/start", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsHlsApi) Status(manifestId string) (*model.ModelTask, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }

    var responseModel *model.ModelTask
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/status", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsHlsApi) Stop(manifestId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Post("/encoding/manifests/hls/{manifest_id}/stop", nil, &responseModel, reqParams)
    return responseModel, err
}

