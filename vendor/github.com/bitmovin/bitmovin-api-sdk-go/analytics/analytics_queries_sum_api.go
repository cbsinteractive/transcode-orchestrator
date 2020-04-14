package analytics
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type AnalyticsQueriesSumApi struct {
    apiClient *common.ApiClient
}

func NewAnalyticsQueriesSumApi(configs ...func(*common.ApiClient)) (*AnalyticsQueriesSumApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AnalyticsQueriesSumApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AnalyticsQueriesSumApi) Create(analyticsSumQueryRequest model.AnalyticsSumQueryRequest) (*model.AnalyticsResponse, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.AnalyticsResponse
    err := api.apiClient.Post("/analytics/queries/sum", &analyticsSumQueryRequest, &responseModel, reqParams)
    return responseModel, err
}

