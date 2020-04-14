package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingManifestsDashPeriodsAdaptationsetsVideoApi struct {
    apiClient *common.ApiClient
}

func NewEncodingManifestsDashPeriodsAdaptationsetsVideoApi(configs ...func(*common.ApiClient)) (*EncodingManifestsDashPeriodsAdaptationsetsVideoApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsDashPeriodsAdaptationsetsVideoApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsDashPeriodsAdaptationsetsVideoApi) Create(manifestId string, periodId string, videoAdaptationSet model.VideoAdaptationSet) (*model.VideoAdaptationSet, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
    }

    var responseModel *model.VideoAdaptationSet
    err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/video", &videoAdaptationSet, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsDashPeriodsAdaptationsetsVideoApi) Delete(manifestId string, periodId string, adaptationsetId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/video/{adaptationset_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsDashPeriodsAdaptationsetsVideoApi) Get(manifestId string, periodId string, adaptationsetId string) (*model.VideoAdaptationSet, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
    }

    var responseModel *model.VideoAdaptationSet
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/video/{adaptationset_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsDashPeriodsAdaptationsetsVideoApi) List(manifestId string, periodId string, queryParams ...func(*query.VideoAdaptationSetListQueryParams)) (*pagination.VideoAdaptationSetsListPagination, error) {
    queryParameters := &query.VideoAdaptationSetListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.VideoAdaptationSetsListPagination
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/video", nil, &responseModel, reqParams)
    return responseModel, err
}

