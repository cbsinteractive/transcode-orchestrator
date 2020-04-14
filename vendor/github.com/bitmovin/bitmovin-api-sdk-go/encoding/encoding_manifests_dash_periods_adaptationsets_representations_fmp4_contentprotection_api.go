package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4ContentprotectionApi struct {
    apiClient *common.ApiClient
}

func NewEncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4ContentprotectionApi(configs ...func(*common.ApiClient)) (*EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4ContentprotectionApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4ContentprotectionApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4ContentprotectionApi) Create(manifestId string, periodId string, adaptationsetId string, representationId string, contentProtection model.ContentProtection) (*model.ContentProtection, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
        params.PathParams["representation_id"] = representationId
    }

    var responseModel *model.ContentProtection
    err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/fmp4/{representation_id}/contentprotection", &contentProtection, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4ContentprotectionApi) Delete(manifestId string, periodId string, adaptationsetId string, representationId string, contentprotectionId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
        params.PathParams["representation_id"] = representationId
        params.PathParams["contentprotection_id"] = contentprotectionId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/fmp4/{representation_id}/contentprotection/{contentprotection_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4ContentprotectionApi) Get(manifestId string, periodId string, adaptationsetId string, representationId string, contentprotectionId string) (*model.ContentProtection, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
        params.PathParams["representation_id"] = representationId
        params.PathParams["contentprotection_id"] = contentprotectionId
    }

    var responseModel *model.ContentProtection
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/fmp4/{representation_id}/contentprotection/{contentprotection_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsDashPeriodsAdaptationsetsRepresentationsFmp4ContentprotectionApi) List(manifestId string, periodId string, adaptationsetId string, representationId string, queryParams ...func(*query.ContentProtectionListQueryParams)) (*pagination.ContentProtectionsListPagination, error) {
    queryParameters := &query.ContentProtectionListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
        params.PathParams["representation_id"] = representationId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.ContentProtectionsListPagination
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/{adaptationset_id}/representations/fmp4/{representation_id}/contentprotection", nil, &responseModel, reqParams)
    return responseModel, err
}

