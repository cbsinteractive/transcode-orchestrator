package player
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type PlayerApi struct {
    apiClient *common.ApiClient
    Channels *PlayerChannelsApi
    Licenses *PlayerLicensesApi
    CustomBuilds *PlayerCustomBuildsApi
}

func NewPlayerApi(configs ...func(*common.ApiClient)) (*PlayerApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &PlayerApi{apiClient: apiClient}

    channelsApi, err := NewPlayerChannelsApi(configs...)
    api.Channels = channelsApi
    licensesApi, err := NewPlayerLicensesApi(configs...)
    api.Licenses = licensesApi
    customBuildsApi, err := NewPlayerCustomBuildsApi(configs...)
    api.CustomBuilds = customBuildsApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

