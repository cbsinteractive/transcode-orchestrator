package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafApi struct {
    apiClient *common.ApiClient
    Drm *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafDrmApi
    Contentprotection *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafContentprotectionApi
}

func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafApi(configs ...func(*common.ApiClient)) (*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafApi{apiClient: apiClient}

    drmApi, err := NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafDrmApi(configs...)
    api.Drm = drmApi
    contentprotectionApi, err := NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafContentprotectionApi(configs...)
    api.Contentprotection = contentprotectionApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafApi) Create(manifestId string, periodId string, adaptationsetId string, dashCmafRepresentation model.DashCmafRepresentation) (*model.DashCmafRepresentation, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
    }

    var responseModel *model.DashCmafRepresentation
    err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/cmaf", &dashCmafRepresentation, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafApi) Delete(manifestId string, periodId string, adaptationsetId string, representationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
        params.PathParams["representation_id"] = representationId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/cmaf/{representation_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafApi) Get(manifestId string, periodId string, adaptationsetId string, representationId string) (*model.DashCmafRepresentation, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
        params.PathParams["representation_id"] = representationId
    }

    var responseModel *model.DashCmafRepresentation
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/cmaf/{representation_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsCmafApi) List(manifestId string, periodId string, adaptationsetId string, queryParams ...func(*query.DashCmafRepresentationListQueryParams)) (*pagination.DashCmafRepresentationsListPagination, error) {
    queryParameters := &query.DashCmafRepresentationListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.DashCmafRepresentationsListPagination
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/cmaf", nil, &responseModel, reqParams)
    return responseModel, err
}

