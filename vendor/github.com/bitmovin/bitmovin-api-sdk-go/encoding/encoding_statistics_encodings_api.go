package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingStatisticsEncodingsApi struct {
    apiClient *common.ApiClient
    Live *EncodingStatisticsEncodingsLiveApi
    Vod *EncodingStatisticsEncodingsVodApi
    LiveStatistics *EncodingStatisticsEncodingsLiveStatisticsApi
}

func NewEncodingStatisticsEncodingsApi(configs ...func(*common.ApiClient)) (*EncodingStatisticsEncodingsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingStatisticsEncodingsApi{apiClient: apiClient}

    liveApi, err := NewEncodingStatisticsEncodingsLiveApi(configs...)
    api.Live = liveApi
    vodApi, err := NewEncodingStatisticsEncodingsVodApi(configs...)
    api.Vod = vodApi
    liveStatisticsApi, err := NewEncodingStatisticsEncodingsLiveStatisticsApi(configs...)
    api.LiveStatistics = liveStatisticsApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingStatisticsEncodingsApi) Get(encodingId string) (*model.EncodingStats, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.EncodingStats
    err := api.apiClient.Get("/encoding/statistics/encodings/{encoding_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

