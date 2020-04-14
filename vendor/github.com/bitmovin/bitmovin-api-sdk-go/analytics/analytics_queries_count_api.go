package analytics
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type AnalyticsQueriesCountApi struct {
    apiClient *common.ApiClient
}

func NewAnalyticsQueriesCountApi(configs ...func(*common.ApiClient)) (*AnalyticsQueriesCountApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AnalyticsQueriesCountApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AnalyticsQueriesCountApi) Create(analyticsCountQueryRequest model.AnalyticsCountQueryRequest) (*model.AnalyticsResponse, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.AnalyticsResponse
    err := api.apiClient.Post("/analytics/queries/count", &analyticsCountQueryRequest, &responseModel, reqParams)
    return responseModel, err
}

