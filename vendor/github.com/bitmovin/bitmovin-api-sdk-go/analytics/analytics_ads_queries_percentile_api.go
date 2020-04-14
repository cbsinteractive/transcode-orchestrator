package analytics
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type AnalyticsAdsQueriesPercentileApi struct {
    apiClient *common.ApiClient
}

func NewAnalyticsAdsQueriesPercentileApi(configs ...func(*common.ApiClient)) (*AnalyticsAdsQueriesPercentileApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AnalyticsAdsQueriesPercentileApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AnalyticsAdsQueriesPercentileApi) Create(adAnalyticsPercentileQueryRequest model.AdAnalyticsPercentileQueryRequest) (*model.AnalyticsResponse, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.AnalyticsResponse
    err := api.apiClient.Post("/analytics/ads/queries/percentile", &adAnalyticsPercentileQueryRequest, &responseModel, reqParams)
    return responseModel, err
}

