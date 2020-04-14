package analytics
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type AnalyticsMetricsMaxConcurrentviewersApi struct {
    apiClient *common.ApiClient
}

func NewAnalyticsMetricsMaxConcurrentviewersApi(configs ...func(*common.ApiClient)) (*AnalyticsMetricsMaxConcurrentviewersApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AnalyticsMetricsMaxConcurrentviewersApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AnalyticsMetricsMaxConcurrentviewersApi) Create(analyticsMetricsQueryRequest model.AnalyticsMetricsQueryRequest) (*model.AnalyticsMaxConcurrentViewersResponse, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.AnalyticsMaxConcurrentViewersResponse
    err := api.apiClient.Post("/analytics/metrics/max-concurrentviewers", &analyticsMetricsQueryRequest, &responseModel, reqParams)
    return responseModel, err
}

