package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmApi struct {
    apiClient *common.ApiClient
    Contentprotection *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmContentprotectionApi
}

func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmApi(configs ...func(*common.ApiClient)) (*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmApi{apiClient: apiClient}

    contentprotectionApi, err := NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmContentprotectionApi(configs...)
    api.Contentprotection = contentprotectionApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmApi) Create(manifestId string, periodId string, adaptationsetId string, dashFmp4DrmRepresentation model.DashFmp4DrmRepresentation) (*model.DashFmp4DrmRepresentation, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
    }

    var responseModel *model.DashFmp4DrmRepresentation
    err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/fmp4/drm", &dashFmp4DrmRepresentation, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmApi) Delete(manifestId string, periodId string, adaptationsetId string, representationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
        params.PathParams["representation_id"] = representationId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/fmp4/drm/{representation_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmApi) Get(manifestId string, periodId string, adaptationsetId string, representationId string) (*model.DashFmp4DrmRepresentation, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
        params.PathParams["representation_id"] = representationId
    }

    var responseModel *model.DashFmp4DrmRepresentation
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/fmp4/drm/{representation_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4DrmApi) List(manifestId string, periodId string, adaptationsetId string, queryParams ...func(*query.DashFmp4DrmRepresentationListQueryParams)) (*pagination.DashFmp4DrmRepresentationsListPagination, error) {
    queryParameters := &query.DashFmp4DrmRepresentationListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.DashFmp4DrmRepresentationsListPagination
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/fmp4/drm", nil, &responseModel, reqParams)
    return responseModel, err
}

