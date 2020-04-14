package analytics
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type AnalyticsQueriesMaxApi struct {
    apiClient *common.ApiClient
}

func NewAnalyticsQueriesMaxApi(configs ...func(*common.ApiClient)) (*AnalyticsQueriesMaxApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AnalyticsQueriesMaxApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AnalyticsQueriesMaxApi) Create(analyticsMaxQueryRequest model.AnalyticsMaxQueryRequest) (*model.AnalyticsResponse, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.AnalyticsResponse
    err := api.apiClient.Post("/analytics/queries/max", &analyticsMaxQueryRequest, &responseModel, reqParams)
    return responseModel, err
}

