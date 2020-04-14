package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
    "time"
)

type EncodingStatisticsLabelsApi struct {
    apiClient *common.ApiClient
    Daily *EncodingStatisticsLabelsDailyApi
}

func NewEncodingStatisticsLabelsApi(configs ...func(*common.ApiClient)) (*EncodingStatisticsLabelsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingStatisticsLabelsApi{apiClient: apiClient}

    dailyApi, err := NewEncodingStatisticsLabelsDailyApi(configs...)
    api.Daily = dailyApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingStatisticsLabelsApi) List(queryParams ...func(*query.StatisticsPerLabelListQueryParams)) (*pagination.StatisticsPerLabelsListPagination, error) {
    queryParameters := &query.StatisticsPerLabelListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.StatisticsPerLabelsListPagination
    err := api.apiClient.Get("/encoding/statistics/labels/", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingStatisticsLabelsApi) ListByDateRange(from time.Time, to time.Time, queryParams ...func(*query.StatisticsPerLabelListByDateRangeQueryParams)) (*pagination.StatisticsPerLabelsListByDateRangePagination, error) {
    queryParameters := &query.StatisticsPerLabelListByDateRangeQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["from"] = from
        params.PathParams["to"] = to
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.StatisticsPerLabelsListByDateRangePagination
    err := api.apiClient.Get("/encoding/statistics/labels/{from}/{to}", nil, &responseModel, reqParams)
    return responseModel, err
}

