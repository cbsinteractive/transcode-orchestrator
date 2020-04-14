package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingManifestsHlsMediaClosedCaptionsApi struct {
    apiClient *common.ApiClient
}

func NewEncodingManifestsHlsMediaClosedCaptionsApi(configs ...func(*common.ApiClient)) (*EncodingManifestsHlsMediaClosedCaptionsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsHlsMediaClosedCaptionsApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsHlsMediaClosedCaptionsApi) Create(manifestId string, closedCaptionsMediaInfo model.ClosedCaptionsMediaInfo) (*model.ClosedCaptionsMediaInfo, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }

    var responseModel *model.ClosedCaptionsMediaInfo
    err := api.apiClient.Post("/encoding/manifests/hls/{manifest_id}/media/closed-captions", &closedCaptionsMediaInfo, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsHlsMediaClosedCaptionsApi) Delete(manifestId string, mediaId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["media_id"] = mediaId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/manifests/hls/{manifest_id}/media/closed-captions/{media_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsHlsMediaClosedCaptionsApi) Get(manifestId string, mediaId string) (*model.ClosedCaptionsMediaInfo, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["media_id"] = mediaId
    }

    var responseModel *model.ClosedCaptionsMediaInfo
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/media/closed-captions/{media_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsHlsMediaClosedCaptionsApi) List(manifestId string, queryParams ...func(*query.ClosedCaptionsMediaInfoListQueryParams)) (*pagination.ClosedCaptionsMediaInfosListPagination, error) {
    queryParameters := &query.ClosedCaptionsMediaInfoListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.ClosedCaptionsMediaInfosListPagination
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/media/closed-captions", nil, &responseModel, reqParams)
    return responseModel, err
}

