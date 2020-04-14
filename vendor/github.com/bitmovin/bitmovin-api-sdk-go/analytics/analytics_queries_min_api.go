package analytics
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type AnalyticsQueriesMinApi struct {
    apiClient *common.ApiClient
}

func NewAnalyticsQueriesMinApi(configs ...func(*common.ApiClient)) (*AnalyticsQueriesMinApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AnalyticsQueriesMinApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AnalyticsQueriesMinApi) Create(analyticsMinQueryRequest model.AnalyticsMinQueryRequest) (*model.AnalyticsResponse, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.AnalyticsResponse
    err := api.apiClient.Post("/analytics/queries/min", &analyticsMinQueryRequest, &responseModel, reqParams)
    return responseModel, err
}

