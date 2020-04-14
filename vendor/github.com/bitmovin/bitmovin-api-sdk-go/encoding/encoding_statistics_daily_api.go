package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
    "time"
)

type EncodingStatisticsDailyApi struct {
    apiClient *common.ApiClient
}

func NewEncodingStatisticsDailyApi(configs ...func(*common.ApiClient)) (*EncodingStatisticsDailyApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingStatisticsDailyApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingStatisticsDailyApi) List(queryParams ...func(*query.DailyStatisticsListQueryParams)) (*pagination.DailyStatisticssListPagination, error) {
    queryParameters := &query.DailyStatisticsListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.DailyStatisticssListPagination
    err := api.apiClient.Get("/encoding/statistics/daily", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingStatisticsDailyApi) ListByDateRange(from time.Time, to time.Time, queryParams ...func(*query.DailyStatisticsListByDateRangeQueryParams)) (*pagination.DailyStatisticssListByDateRangePagination, error) {
    queryParameters := &query.DailyStatisticsListByDateRangeQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["from"] = from
        params.PathParams["to"] = to
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.DailyStatisticssListByDateRangePagination
    err := api.apiClient.Get("/encoding/statistics/daily/{from}/{to}", nil, &responseModel, reqParams)
    return responseModel, err
}

