package analytics
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type AnalyticsAdsApi struct {
    apiClient *common.ApiClient
    Queries *AnalyticsAdsQueriesApi
}

func NewAnalyticsAdsApi(configs ...func(*common.ApiClient)) (*AnalyticsAdsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AnalyticsAdsApi{apiClient: apiClient}

    queriesApi, err := NewAnalyticsAdsQueriesApi(configs...)
    api.Queries = queriesApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

