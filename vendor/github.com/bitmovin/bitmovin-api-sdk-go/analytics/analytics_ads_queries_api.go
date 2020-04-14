package analytics
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type AnalyticsAdsQueriesApi struct {
    apiClient *common.ApiClient
    Count *AnalyticsAdsQueriesCountApi
    Sum *AnalyticsAdsQueriesSumApi
    Avg *AnalyticsAdsQueriesAvgApi
    Min *AnalyticsAdsQueriesMinApi
    Max *AnalyticsAdsQueriesMaxApi
    Stddev *AnalyticsAdsQueriesStddevApi
    Percentile *AnalyticsAdsQueriesPercentileApi
    Variance *AnalyticsAdsQueriesVarianceApi
    Median *AnalyticsAdsQueriesMedianApi
}

func NewAnalyticsAdsQueriesApi(configs ...func(*common.ApiClient)) (*AnalyticsAdsQueriesApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AnalyticsAdsQueriesApi{apiClient: apiClient}

    countApi, err := NewAnalyticsAdsQueriesCountApi(configs...)
    api.Count = countApi
    sumApi, err := NewAnalyticsAdsQueriesSumApi(configs...)
    api.Sum = sumApi
    avgApi, err := NewAnalyticsAdsQueriesAvgApi(configs...)
    api.Avg = avgApi
    minApi, err := NewAnalyticsAdsQueriesMinApi(configs...)
    api.Min = minApi
    maxApi, err := NewAnalyticsAdsQueriesMaxApi(configs...)
    api.Max = maxApi
    stddevApi, err := NewAnalyticsAdsQueriesStddevApi(configs...)
    api.Stddev = stddevApi
    percentileApi, err := NewAnalyticsAdsQueriesPercentileApi(configs...)
    api.Percentile = percentileApi
    varianceApi, err := NewAnalyticsAdsQueriesVarianceApi(configs...)
    api.Variance = varianceApi
    medianApi, err := NewAnalyticsAdsQueriesMedianApi(configs...)
    api.Median = medianApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

