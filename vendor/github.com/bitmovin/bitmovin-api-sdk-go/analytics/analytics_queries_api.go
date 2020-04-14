package analytics
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type AnalyticsQueriesApi struct {
    apiClient *common.ApiClient
    Count *AnalyticsQueriesCountApi
    Sum *AnalyticsQueriesSumApi
    Avg *AnalyticsQueriesAvgApi
    Min *AnalyticsQueriesMinApi
    Max *AnalyticsQueriesMaxApi
    Stddev *AnalyticsQueriesStddevApi
    Percentile *AnalyticsQueriesPercentileApi
    Variance *AnalyticsQueriesVarianceApi
    Median *AnalyticsQueriesMedianApi
}

func NewAnalyticsQueriesApi(configs ...func(*common.ApiClient)) (*AnalyticsQueriesApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AnalyticsQueriesApi{apiClient: apiClient}

    countApi, err := NewAnalyticsQueriesCountApi(configs...)
    api.Count = countApi
    sumApi, err := NewAnalyticsQueriesSumApi(configs...)
    api.Sum = sumApi
    avgApi, err := NewAnalyticsQueriesAvgApi(configs...)
    api.Avg = avgApi
    minApi, err := NewAnalyticsQueriesMinApi(configs...)
    api.Min = minApi
    maxApi, err := NewAnalyticsQueriesMaxApi(configs...)
    api.Max = maxApi
    stddevApi, err := NewAnalyticsQueriesStddevApi(configs...)
    api.Stddev = stddevApi
    percentileApi, err := NewAnalyticsQueriesPercentileApi(configs...)
    api.Percentile = percentileApi
    varianceApi, err := NewAnalyticsQueriesVarianceApi(configs...)
    api.Variance = varianceApi
    medianApi, err := NewAnalyticsQueriesMedianApi(configs...)
    api.Median = medianApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

