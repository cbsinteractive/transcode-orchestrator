package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
    "time"
)

type EncodingStatisticsApi struct {
    apiClient *common.ApiClient
    Daily *EncodingStatisticsDailyApi
    Encodings *EncodingStatisticsEncodingsApi
    Labels *EncodingStatisticsLabelsApi
}

func NewEncodingStatisticsApi(configs ...func(*common.ApiClient)) (*EncodingStatisticsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingStatisticsApi{apiClient: apiClient}

    dailyApi, err := NewEncodingStatisticsDailyApi(configs...)
    api.Daily = dailyApi
    encodingsApi, err := NewEncodingStatisticsEncodingsApi(configs...)
    api.Encodings = encodingsApi
    labelsApi, err := NewEncodingStatisticsLabelsApi(configs...)
    api.Labels = labelsApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingStatisticsApi) Get() (*model.Statistics, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.Statistics
    err := api.apiClient.Get("/encoding/statistics", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingStatisticsApi) List(from time.Time, to time.Time, queryParams ...func(*query.StatisticsListQueryParams)) (*pagination.StatisticssListPagination, error) {
    queryParameters := &query.StatisticsListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["from"] = from
        params.PathParams["to"] = to
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.StatisticssListPagination
    err := api.apiClient.Get("/encoding/statistics/{from}/{to}", nil, &responseModel, reqParams)
    return responseModel, err
}

