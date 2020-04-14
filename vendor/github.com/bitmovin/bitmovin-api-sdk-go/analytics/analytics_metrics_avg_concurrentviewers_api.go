package analytics
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type AnalyticsMetricsAvgConcurrentviewersApi struct {
    apiClient *common.ApiClient
}

func NewAnalyticsMetricsAvgConcurrentviewersApi(configs ...func(*common.ApiClient)) (*AnalyticsMetricsAvgConcurrentviewersApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AnalyticsMetricsAvgConcurrentviewersApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AnalyticsMetricsAvgConcurrentviewersApi) Create(analyticsMetricsQueryRequest model.AnalyticsMetricsQueryRequest) (*model.AnalyticsAvgConcurrentViewersResponse, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.AnalyticsAvgConcurrentViewersResponse
    err := api.apiClient.Post("/analytics/metrics/avg-concurrentviewers", &analyticsMetricsQueryRequest, &responseModel, reqParams)
    return responseModel, err
}

