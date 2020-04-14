package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4DrmApi struct {
    apiClient *common.ApiClient
}

func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4DrmApi(configs ...func(*common.ApiClient)) (*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4DrmApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4DrmApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4DrmApi) Create(manifestId string, periodId string, adaptationsetId string, dashMp4DrmRepresentation model.DashMp4DrmRepresentation) (*model.DashMp4DrmRepresentation, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
    }

    var responseModel *model.DashMp4DrmRepresentation
    err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/mp4/drm", &dashMp4DrmRepresentation, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4DrmApi) Delete(manifestId string, periodId string, adaptationsetId string, representationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
        params.PathParams["representation_id"] = representationId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/mp4/drm/{representation_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4DrmApi) Get(manifestId string, periodId string, adaptationsetId string, representationId string) (*model.DashMp4DrmRepresentation, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
        params.PathParams["representation_id"] = representationId
    }

    var responseModel *model.DashMp4DrmRepresentation
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/mp4/drm/{representation_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsMp4DrmApi) List(manifestId string, periodId string, adaptationsetId string, queryParams ...func(*query.DashMp4DrmRepresentationListQueryParams)) (*pagination.DashMp4DrmRepresentationsListPagination, error) {
    queryParameters := &query.DashMp4DrmRepresentationListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.DashMp4DrmRepresentationsListPagination
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/mp4/drm", nil, &responseModel, reqParams)
    return responseModel, err
}

