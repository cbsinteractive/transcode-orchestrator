package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingStatisticsEncodingsLiveStatisticsStreamsApi struct {
    apiClient *common.ApiClient
}

func NewEncodingStatisticsEncodingsLiveStatisticsStreamsApi(configs ...func(*common.ApiClient)) (*EncodingStatisticsEncodingsLiveStatisticsStreamsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingStatisticsEncodingsLiveStatisticsStreamsApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingStatisticsEncodingsLiveStatisticsStreamsApi) List(encodingId string, queryParams ...func(*query.StreamInfosListQueryParams)) (*pagination.StreamInfossListPagination, error) {
    queryParameters := &query.StreamInfosListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.StreamInfossListPagination
    err := api.apiClient.Get("/encoding/statistics/encodings/{encoding_id}/live-statistics/streams", nil, &responseModel, reqParams)
    return responseModel, err
}

