package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingManifestsSmoothContentprotectionApi struct {
    apiClient *common.ApiClient
}

func NewEncodingManifestsSmoothContentprotectionApi(configs ...func(*common.ApiClient)) (*EncodingManifestsSmoothContentprotectionApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsSmoothContentprotectionApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsSmoothContentprotectionApi) Create(manifestId string, smoothManifestContentProtection model.SmoothManifestContentProtection) (*model.SmoothManifestContentProtection, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }

    var responseModel *model.SmoothManifestContentProtection
    err := api.apiClient.Post("/encoding/manifests/smooth/{manifest_id}/contentprotection", &smoothManifestContentProtection, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsSmoothContentprotectionApi) Delete(manifestId string, protectionId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["protection_id"] = protectionId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/manifests/smooth/{manifest_id}/contentprotection/{protection_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsSmoothContentprotectionApi) Get(manifestId string, protectionId string) (*model.SmoothManifestContentProtection, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["protection_id"] = protectionId
    }

    var responseModel *model.SmoothManifestContentProtection
    err := api.apiClient.Get("/encoding/manifests/smooth/{manifest_id}/contentprotection/{protection_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsSmoothContentprotectionApi) List(manifestId string, queryParams ...func(*query.SmoothManifestContentProtectionListQueryParams)) (*pagination.SmoothManifestContentProtectionsListPagination, error) {
    queryParameters := &query.SmoothManifestContentProtectionListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.SmoothManifestContentProtectionsListPagination
    err := api.apiClient.Get("/encoding/manifests/smooth/{manifest_id}/contentprotection", nil, &responseModel, reqParams)
    return responseModel, err
}

