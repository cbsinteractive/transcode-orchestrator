package analytics
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type AnalyticsQueriesPercentileApi struct {
    apiClient *common.ApiClient
}

func NewAnalyticsQueriesPercentileApi(configs ...func(*common.ApiClient)) (*AnalyticsQueriesPercentileApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AnalyticsQueriesPercentileApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AnalyticsQueriesPercentileApi) Create(analyticsPercentileQueryRequest model.AnalyticsPercentileQueryRequest) (*model.AnalyticsResponse, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.AnalyticsResponse
    err := api.apiClient.Post("/analytics/queries/percentile", &analyticsPercentileQueryRequest, &responseModel, reqParams)
    return responseModel, err
}

