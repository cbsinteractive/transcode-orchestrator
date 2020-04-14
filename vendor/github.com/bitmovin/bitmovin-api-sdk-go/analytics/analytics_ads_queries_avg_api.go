package analytics
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type AnalyticsAdsQueriesAvgApi struct {
    apiClient *common.ApiClient
}

func NewAnalyticsAdsQueriesAvgApi(configs ...func(*common.ApiClient)) (*AnalyticsAdsQueriesAvgApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AnalyticsAdsQueriesAvgApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AnalyticsAdsQueriesAvgApi) Create(adAnalyticsAvgQueryRequest model.AdAnalyticsAvgQueryRequest) (*model.AnalyticsResponse, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.AnalyticsResponse
    err := api.apiClient.Post("/analytics/ads/queries/avg", &adAnalyticsAvgQueryRequest, &responseModel, reqParams)
    return responseModel, err
}

