package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingStatisticsEncodingsLiveStatisticsSrtApi struct {
    apiClient *common.ApiClient
}

func NewEncodingStatisticsEncodingsLiveStatisticsSrtApi(configs ...func(*common.ApiClient)) (*EncodingStatisticsEncodingsLiveStatisticsSrtApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingStatisticsEncodingsLiveStatisticsSrtApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingStatisticsEncodingsLiveStatisticsSrtApi) List(encodingId string, queryParams ...func(*query.SrtStatisticsListQueryParams)) (*pagination.SrtStatisticssListPagination, error) {
    queryParameters := &query.SrtStatisticsListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.SrtStatisticssListPagination
    err := api.apiClient.Get("/encoding/statistics/encodings/{encoding_id}/live-statistics/srt", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingStatisticsEncodingsLiveStatisticsSrtApi) ListBySrtInputId(encodingId string, srtInputId string, queryParams ...func(*query.SrtStatisticsListBySrtInputIdQueryParams)) (*pagination.SrtStatisticssListBySrtInputIdPagination, error) {
    queryParameters := &query.SrtStatisticsListBySrtInputIdQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["srt_input_id"] = srtInputId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.SrtStatisticssListBySrtInputIdPagination
    err := api.apiClient.Get("/encoding/statistics/encodings/{encoding_id}/live-statistics/srt/{srt_input_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

