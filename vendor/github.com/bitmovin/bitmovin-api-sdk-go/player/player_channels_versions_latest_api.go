package player
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type PlayerChannelsVersionsLatestApi struct {
    apiClient *common.ApiClient
}

func NewPlayerChannelsVersionsLatestApi(configs ...func(*common.ApiClient)) (*PlayerChannelsVersionsLatestApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &PlayerChannelsVersionsLatestApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *PlayerChannelsVersionsLatestApi) Get(channelName string) (*model.PlayerVersion, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["channel_name"] = channelName
    }

    var responseModel *model.PlayerVersion
    err := api.apiClient.Get("/player/channels/{channel_name}/versions/latest", nil, &responseModel, reqParams)
    return responseModel, err
}

