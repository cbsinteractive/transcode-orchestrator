package player
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type PlayerChannelsVersionsApi struct {
    apiClient *common.ApiClient
    Latest *PlayerChannelsVersionsLatestApi
}

func NewPlayerChannelsVersionsApi(configs ...func(*common.ApiClient)) (*PlayerChannelsVersionsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &PlayerChannelsVersionsApi{apiClient: apiClient}

    latestApi, err := NewPlayerChannelsVersionsLatestApi(configs...)
    api.Latest = latestApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *PlayerChannelsVersionsApi) List(channelName string) (*pagination.PlayerVersionsListPagination, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["channel_name"] = channelName
    }

    var responseModel *pagination.PlayerVersionsListPagination
    err := api.apiClient.Get("/player/channels/{channel_name}/versions", nil, &responseModel, reqParams)
    return responseModel, err
}

