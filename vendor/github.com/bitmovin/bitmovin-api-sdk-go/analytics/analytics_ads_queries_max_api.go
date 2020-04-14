package analytics
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type AnalyticsAdsQueriesMaxApi struct {
    apiClient *common.ApiClient
}

func NewAnalyticsAdsQueriesMaxApi(configs ...func(*common.ApiClient)) (*AnalyticsAdsQueriesMaxApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AnalyticsAdsQueriesMaxApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AnalyticsAdsQueriesMaxApi) Create(adAnalyticsMaxQueryRequest model.AdAnalyticsMaxQueryRequest) (*model.AnalyticsResponse, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.AnalyticsResponse
    err := api.apiClient.Post("/analytics/ads/queries/max", &adAnalyticsMaxQueryRequest, &responseModel, reqParams)
    return responseModel, err
}

