package player
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type PlayerCustomBuildsApi struct {
    apiClient *common.ApiClient
    Web *PlayerCustomBuildsWebApi
}

func NewPlayerCustomBuildsApi(configs ...func(*common.ApiClient)) (*PlayerCustomBuildsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &PlayerCustomBuildsApi{apiClient: apiClient}

    webApi, err := NewPlayerCustomBuildsWebApi(configs...)
    api.Web = webApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

