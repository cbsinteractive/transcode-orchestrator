package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingStatisticsEncodingsLiveStatisticsApi struct {
    apiClient *common.ApiClient
    Events *EncodingStatisticsEncodingsLiveStatisticsEventsApi
    Streams *EncodingStatisticsEncodingsLiveStatisticsStreamsApi
    Srt *EncodingStatisticsEncodingsLiveStatisticsSrtApi
}

func NewEncodingStatisticsEncodingsLiveStatisticsApi(configs ...func(*common.ApiClient)) (*EncodingStatisticsEncodingsLiveStatisticsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingStatisticsEncodingsLiveStatisticsApi{apiClient: apiClient}

    eventsApi, err := NewEncodingStatisticsEncodingsLiveStatisticsEventsApi(configs...)
    api.Events = eventsApi
    streamsApi, err := NewEncodingStatisticsEncodingsLiveStatisticsStreamsApi(configs...)
    api.Streams = streamsApi
    srtApi, err := NewEncodingStatisticsEncodingsLiveStatisticsSrtApi(configs...)
    api.Srt = srtApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingStatisticsEncodingsLiveStatisticsApi) Get(encodingId string) (*model.LiveEncodingStats, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.LiveEncodingStats
    err := api.apiClient.Get("/encoding/statistics/encodings/{encoding_id}/live-statistics", nil, &responseModel, reqParams)
    return responseModel, err
}

