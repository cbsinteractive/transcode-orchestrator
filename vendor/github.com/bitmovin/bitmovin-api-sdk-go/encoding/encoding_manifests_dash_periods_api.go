package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingManifestsDashPeriodsApi struct {
    apiClient *common.ApiClient
    CustomXmlElements *EncodingManifestsDashPeriodsCustomXmlElementsApi
    Adaptationsets *EncodingManifestsDashPeriodsAdaptationsetsApi
}

func NewEncodingManifestsDashPeriodsApi(configs ...func(*common.ApiClient)) (*EncodingManifestsDashPeriodsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsDashPeriodsApi{apiClient: apiClient}

    customXmlElementsApi, err := NewEncodingManifestsDashPeriodsCustomXmlElementsApi(configs...)
    api.CustomXmlElements = customXmlElementsApi
    adaptationsetsApi, err := NewEncodingManifestsDashPeriodsAdaptationsetsApi(configs...)
    api.Adaptationsets = adaptationsetsApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsDashPeriodsApi) Create(manifestId string, period model.Period) (*model.Period, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }

    var responseModel *model.Period
    err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/periods", &period, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsDashPeriodsApi) Delete(manifestId string, periodId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/manifests/dash/{manifest_id}/periods/{period_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsDashPeriodsApi) Get(manifestId string, periodId string) (*model.Period, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
    }

    var responseModel *model.Period
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsDashPeriodsApi) List(manifestId string, queryParams ...func(*query.PeriodListQueryParams)) (*pagination.PeriodsListPagination, error) {
    queryParameters := &query.PeriodListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.PeriodsListPagination
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods", nil, &responseModel, reqParams)
    return responseModel, err
}

