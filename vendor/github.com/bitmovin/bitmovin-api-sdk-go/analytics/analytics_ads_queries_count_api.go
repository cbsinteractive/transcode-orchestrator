package analytics
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type AnalyticsAdsQueriesCountApi struct {
    apiClient *common.ApiClient
}

func NewAnalyticsAdsQueriesCountApi(configs ...func(*common.ApiClient)) (*AnalyticsAdsQueriesCountApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AnalyticsAdsQueriesCountApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AnalyticsAdsQueriesCountApi) Create(adAnalyticsCountQueryRequest model.AdAnalyticsCountQueryRequest) (*model.AnalyticsResponse, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.AnalyticsResponse
    err := api.apiClient.Post("/analytics/ads/queries/count", &adAnalyticsCountQueryRequest, &responseModel, reqParams)
    return responseModel, err
}

