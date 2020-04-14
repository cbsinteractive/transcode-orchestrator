package analytics
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type AnalyticsQueriesAvgApi struct {
    apiClient *common.ApiClient
}

func NewAnalyticsQueriesAvgApi(configs ...func(*common.ApiClient)) (*AnalyticsQueriesAvgApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AnalyticsQueriesAvgApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AnalyticsQueriesAvgApi) Create(analyticsAvgQueryRequest model.AnalyticsAvgQueryRequest) (*model.AnalyticsResponse, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.AnalyticsResponse
    err := api.apiClient.Post("/analytics/queries/avg", &analyticsAvgQueryRequest, &responseModel, reqParams)
    return responseModel, err
}

