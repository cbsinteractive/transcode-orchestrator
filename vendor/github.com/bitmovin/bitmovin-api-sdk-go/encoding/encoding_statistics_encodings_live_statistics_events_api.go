package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingStatisticsEncodingsLiveStatisticsEventsApi struct {
    apiClient *common.ApiClient
}

func NewEncodingStatisticsEncodingsLiveStatisticsEventsApi(configs ...func(*common.ApiClient)) (*EncodingStatisticsEncodingsLiveStatisticsEventsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingStatisticsEncodingsLiveStatisticsEventsApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingStatisticsEncodingsLiveStatisticsEventsApi) List(encodingId string, queryParams ...func(*query.LiveEncodingStatsEventListQueryParams)) (*pagination.LiveEncodingStatsEventsListPagination, error) {
    queryParameters := &query.LiveEncodingStatsEventListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.LiveEncodingStatsEventsListPagination
    err := api.apiClient.Get("/encoding/statistics/encodings/{encoding_id}/live-statistics/events", nil, &responseModel, reqParams)
    return responseModel, err
}

