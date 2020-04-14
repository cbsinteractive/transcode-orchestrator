package analytics
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type AnalyticsMetricsAvgDroppedFramesApi struct {
    apiClient *common.ApiClient
}

func NewAnalyticsMetricsAvgDroppedFramesApi(configs ...func(*common.ApiClient)) (*AnalyticsMetricsAvgDroppedFramesApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AnalyticsMetricsAvgDroppedFramesApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AnalyticsMetricsAvgDroppedFramesApi) Create(analyticsMetricsQueryRequest model.AnalyticsMetricsQueryRequest) (*model.AnalyticsAvgDroppedFramesResponse, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.AnalyticsAvgDroppedFramesResponse
    err := api.apiClient.Post("/analytics/metrics/avg-dropped-frames", &analyticsMetricsQueryRequest, &responseModel, reqParams)
    return responseModel, err
}

