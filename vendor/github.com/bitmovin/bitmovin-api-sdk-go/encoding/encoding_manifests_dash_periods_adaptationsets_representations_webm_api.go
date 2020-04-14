package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmApi struct {
    apiClient *common.ApiClient
    Contentprotection *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmContentprotectionApi
}

func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmApi(configs ...func(*common.ApiClient)) (*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmApi{apiClient: apiClient}

    contentprotectionApi, err := NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmContentprotectionApi(configs...)
    api.Contentprotection = contentprotectionApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmApi) Create(manifestId string, periodId string, adaptationsetId string, dashWebmRepresentation model.DashWebmRepresentation) (*model.DashWebmRepresentation, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
    }

    var responseModel *model.DashWebmRepresentation
    err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/webm", &dashWebmRepresentation, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmApi) Delete(manifestId string, periodId string, adaptationsetId string, representationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
        params.PathParams["representation_id"] = representationId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/webm/{representation_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmApi) Get(manifestId string, periodId string, adaptationsetId string, representationId string) (*model.DashWebmRepresentation, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
        params.PathParams["representation_id"] = representationId
    }

    var responseModel *model.DashWebmRepresentation
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/webm/{representation_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsWebmApi) List(manifestId string, periodId string, adaptationsetId string, queryParams ...func(*query.DashWebmRepresentationListQueryParams)) (*pagination.DashWebmRepresentationsListPagination, error) {
    queryParameters := &query.DashWebmRepresentationListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.DashWebmRepresentationsListPagination
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/webm", nil, &responseModel, reqParams)
    return responseModel, err
}

