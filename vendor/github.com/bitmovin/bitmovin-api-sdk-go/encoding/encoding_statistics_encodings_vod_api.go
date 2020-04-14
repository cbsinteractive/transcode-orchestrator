package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
    "time"
)

type EncodingStatisticsEncodingsVodApi struct {
    apiClient *common.ApiClient
}

func NewEncodingStatisticsEncodingsVodApi(configs ...func(*common.ApiClient)) (*EncodingStatisticsEncodingsVodApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingStatisticsEncodingsVodApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingStatisticsEncodingsVodApi) List(queryParams ...func(*query.EncodingStatisticsVodListQueryParams)) (*pagination.EncodingStatisticsVodsListPagination, error) {
    queryParameters := &query.EncodingStatisticsVodListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.EncodingStatisticsVodsListPagination
    err := api.apiClient.Get("/encoding/statistics/encodings/vod", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingStatisticsEncodingsVodApi) ListByDateRange(from time.Time, to time.Time, queryParams ...func(*query.EncodingStatisticsVodListByDateRangeQueryParams)) (*pagination.EncodingStatisticsVodsListByDateRangePagination, error) {
    queryParameters := &query.EncodingStatisticsVodListByDateRangeQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["from"] = from
        params.PathParams["to"] = to
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.EncodingStatisticsVodsListByDateRangePagination
    err := api.apiClient.Get("/encoding/statistics/encodings/vod/{from}/{to}", nil, &responseModel, reqParams)
    return responseModel, err
}

