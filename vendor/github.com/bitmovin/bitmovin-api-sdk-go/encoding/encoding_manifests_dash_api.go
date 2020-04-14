package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingManifestsDashApi struct {
    apiClient *common.ApiClient
    Default *EncodingManifestsDashDefaultApi
    Customdata *EncodingManifestsDashCustomdataApi
    Periods *EncodingManifestsDashPeriodsApi
}

func NewEncodingManifestsDashApi(configs ...func(*common.ApiClient)) (*EncodingManifestsDashApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsDashApi{apiClient: apiClient}

    defaultApi, err := NewEncodingManifestsDashDefaultApi(configs...)
    api.Default = defaultApi
    customdataApi, err := NewEncodingManifestsDashCustomdataApi(configs...)
    api.Customdata = customdataApi
    periodsApi, err := NewEncodingManifestsDashPeriodsApi(configs...)
    api.Periods = periodsApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsDashApi) Create(dashManifest model.DashManifest) (*model.DashManifest, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.DashManifest
    err := api.apiClient.Post("/encoding/manifests/dash", &dashManifest, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsDashApi) Delete(manifestId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/manifests/dash/{manifest_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsDashApi) Get(manifestId string) (*model.DashManifest, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }

    var responseModel *model.DashManifest
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsDashApi) List(queryParams ...func(*query.DashManifestListQueryParams)) (*pagination.DashManifestsListPagination, error) {
    queryParameters := &query.DashManifestListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.DashManifestsListPagination
    err := api.apiClient.Get("/encoding/manifests/dash", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsDashApi) Start(manifestId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/start", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsDashApi) Status(manifestId string) (*model.ModelTask, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }

    var responseModel *model.ModelTask
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/status", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsDashApi) Stop(manifestId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/stop", nil, &responseModel, reqParams)
    return responseModel, err
}

