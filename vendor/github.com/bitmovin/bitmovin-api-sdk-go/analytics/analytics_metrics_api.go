package analytics
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type AnalyticsMetricsApi struct {
    apiClient *common.ApiClient
    MaxConcurrentviewers *AnalyticsMetricsMaxConcurrentviewersApi
    AvgConcurrentviewers *AnalyticsMetricsAvgConcurrentviewersApi
    AvgDroppedFrames *AnalyticsMetricsAvgDroppedFramesApi
}

func NewAnalyticsMetricsApi(configs ...func(*common.ApiClient)) (*AnalyticsMetricsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AnalyticsMetricsApi{apiClient: apiClient}

    maxConcurrentviewersApi, err := NewAnalyticsMetricsMaxConcurrentviewersApi(configs...)
    api.MaxConcurrentviewers = maxConcurrentviewersApi
    avgConcurrentviewersApi, err := NewAnalyticsMetricsAvgConcurrentviewersApi(configs...)
    api.AvgConcurrentviewers = avgConcurrentviewersApi
    avgDroppedFramesApi, err := NewAnalyticsMetricsAvgDroppedFramesApi(configs...)
    api.AvgDroppedFrames = avgDroppedFramesApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

