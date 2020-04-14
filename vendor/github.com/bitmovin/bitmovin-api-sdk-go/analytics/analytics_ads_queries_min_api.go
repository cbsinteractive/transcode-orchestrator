package analytics
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type AnalyticsAdsQueriesMinApi struct {
    apiClient *common.ApiClient
}

func NewAnalyticsAdsQueriesMinApi(configs ...func(*common.ApiClient)) (*AnalyticsAdsQueriesMinApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AnalyticsAdsQueriesMinApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AnalyticsAdsQueriesMinApi) Create(adAnalyticsMinQueryRequest model.AdAnalyticsMinQueryRequest) (*model.AnalyticsResponse, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.AnalyticsResponse
    err := api.apiClient.Post("/analytics/ads/queries/min", &adAnalyticsMinQueryRequest, &responseModel, reqParams)
    return responseModel, err
}

