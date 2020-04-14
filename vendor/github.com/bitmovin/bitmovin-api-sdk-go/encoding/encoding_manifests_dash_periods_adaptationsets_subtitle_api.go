package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingManifestsDashPeriodsAdaptationsetsSubtitleApi struct {
    apiClient *common.ApiClient
}

func NewEncodingManifestsDashPeriodsAdaptationsetsSubtitleApi(configs ...func(*common.ApiClient)) (*EncodingManifestsDashPeriodsAdaptationsetsSubtitleApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsDashPeriodsAdaptationsetsSubtitleApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsDashPeriodsAdaptationsetsSubtitleApi) Create(manifestId string, periodId string, subtitleAdaptationSet model.SubtitleAdaptationSet) (*model.SubtitleAdaptationSet, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
    }

    var responseModel *model.SubtitleAdaptationSet
    err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/subtitle", &subtitleAdaptationSet, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsDashPeriodsAdaptationsetsSubtitleApi) Delete(manifestId string, periodId string, adaptationsetId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/subtitle/{adaptationset_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsDashPeriodsAdaptationsetsSubtitleApi) Get(manifestId string, periodId string, adaptationsetId string) (*model.SubtitleAdaptationSet, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["adaptationset_id"] = adaptationsetId
    }

    var responseModel *model.SubtitleAdaptationSet
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/subtitle/{adaptationset_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsDashPeriodsAdaptationsetsSubtitleApi) List(manifestId string, periodId string, queryParams ...func(*query.SubtitleAdaptationSetListQueryParams)) (*pagination.SubtitleAdaptationSetsListPagination, error) {
    queryParameters := &query.SubtitleAdaptationSetListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.SubtitleAdaptationSetsListPagination
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/adaptationsets/subtitle", nil, &responseModel, reqParams)
    return responseModel, err
}

