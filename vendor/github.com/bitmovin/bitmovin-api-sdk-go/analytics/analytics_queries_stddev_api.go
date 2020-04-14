package analytics
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type AnalyticsQueriesStddevApi struct {
    apiClient *common.ApiClient
}

func NewAnalyticsQueriesStddevApi(configs ...func(*common.ApiClient)) (*AnalyticsQueriesStddevApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AnalyticsQueriesStddevApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AnalyticsQueriesStddevApi) Create(analyticsStddevQueryRequest model.AnalyticsStddevQueryRequest) (*model.AnalyticsResponse, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.AnalyticsResponse
    err := api.apiClient.Post("/analytics/queries/stddev", &analyticsStddevQueryRequest, &responseModel, reqParams)
    return responseModel, err
}

