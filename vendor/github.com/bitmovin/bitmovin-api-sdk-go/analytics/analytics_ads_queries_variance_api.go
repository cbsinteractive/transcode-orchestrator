package analytics
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type AnalyticsAdsQueriesVarianceApi struct {
    apiClient *common.ApiClient
}

func NewAnalyticsAdsQueriesVarianceApi(configs ...func(*common.ApiClient)) (*AnalyticsAdsQueriesVarianceApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AnalyticsAdsQueriesVarianceApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AnalyticsAdsQueriesVarianceApi) Create(adAnalyticsVarianceQueryRequest model.AdAnalyticsVarianceQueryRequest) (*model.AnalyticsResponse, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.AnalyticsResponse
    err := api.apiClient.Post("/analytics/ads/queries/variance", &adAnalyticsVarianceQueryRequest, &responseModel, reqParams)
    return responseModel, err
}

