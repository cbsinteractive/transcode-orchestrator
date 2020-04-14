package analytics
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type AnalyticsAdsQueriesSumApi struct {
    apiClient *common.ApiClient
}

func NewAnalyticsAdsQueriesSumApi(configs ...func(*common.ApiClient)) (*AnalyticsAdsQueriesSumApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AnalyticsAdsQueriesSumApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AnalyticsAdsQueriesSumApi) Create(adAnalyticsSumQueryRequest model.AdAnalyticsSumQueryRequest) (*model.AnalyticsResponse, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.AnalyticsResponse
    err := api.apiClient.Post("/analytics/ads/queries/sum", &adAnalyticsSumQueryRequest, &responseModel, reqParams)
    return responseModel, err
}

