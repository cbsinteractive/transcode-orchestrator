package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingManifestsApi struct {
    apiClient *common.ApiClient
    Type *EncodingManifestsTypeApi
    Dash *EncodingManifestsDashApi
    Hls *EncodingManifestsHlsApi
    Smooth *EncodingManifestsSmoothApi
}

func NewEncodingManifestsApi(configs ...func(*common.ApiClient)) (*EncodingManifestsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsApi{apiClient: apiClient}

    typeApi, err := NewEncodingManifestsTypeApi(configs...)
    api.Type = typeApi
    dashApi, err := NewEncodingManifestsDashApi(configs...)
    api.Dash = dashApi
    hlsApi, err := NewEncodingManifestsHlsApi(configs...)
    api.Hls = hlsApi
    smoothApi, err := NewEncodingManifestsSmoothApi(configs...)
    api.Smooth = smoothApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsApi) List(queryParams ...func(*query.ManifestListQueryParams)) (*pagination.ManifestsListPagination, error) {
    queryParameters := &query.ManifestListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.ManifestsListPagination
    err := api.apiClient.Get("/encoding/manifests", nil, &responseModel, reqParams)
    return responseModel, err
}

