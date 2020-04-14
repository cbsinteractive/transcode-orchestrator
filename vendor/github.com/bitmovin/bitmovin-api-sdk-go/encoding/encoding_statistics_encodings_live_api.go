package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
    "time"
)

type EncodingStatisticsEncodingsLiveApi struct {
    apiClient *common.ApiClient
}

func NewEncodingStatisticsEncodingsLiveApi(configs ...func(*common.ApiClient)) (*EncodingStatisticsEncodingsLiveApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingStatisticsEncodingsLiveApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingStatisticsEncodingsLiveApi) List(queryParams ...func(*query.EncodingStatisticsLiveListQueryParams)) (*pagination.EncodingStatisticsLivesListPagination, error) {
    queryParameters := &query.EncodingStatisticsLiveListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.EncodingStatisticsLivesListPagination
    err := api.apiClient.Get("/encoding/statistics/encodings/live", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingStatisticsEncodingsLiveApi) ListByDateRange(from time.Time, to time.Time, queryParams ...func(*query.EncodingStatisticsLiveListByDateRangeQueryParams)) (*pagination.EncodingStatisticsLivesListByDateRangePagination, error) {
    queryParameters := &query.EncodingStatisticsLiveListByDateRangeQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["from"] = from
        params.PathParams["to"] = to
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.EncodingStatisticsLivesListByDateRangePagination
    err := api.apiClient.Get("/encoding/statistics/encodings/live/{from}/{to}", nil, &responseModel, reqParams)
    return responseModel, err
}

