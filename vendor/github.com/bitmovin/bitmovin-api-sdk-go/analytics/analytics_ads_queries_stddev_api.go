package analytics
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type AnalyticsAdsQueriesStddevApi struct {
    apiClient *common.ApiClient
}

func NewAnalyticsAdsQueriesStddevApi(configs ...func(*common.ApiClient)) (*AnalyticsAdsQueriesStddevApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AnalyticsAdsQueriesStddevApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AnalyticsAdsQueriesStddevApi) Create(adAnalyticsStddevQueryRequest model.AdAnalyticsStddevQueryRequest) (*model.AnalyticsResponse, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.AnalyticsResponse
    err := api.apiClient.Post("/analytics/ads/queries/stddev", &adAnalyticsStddevQueryRequest, &responseModel, reqParams)
    return responseModel, err
}

