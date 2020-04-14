package player
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type PlayerCustomBuildsWebStatusApi struct {
    apiClient *common.ApiClient
}

func NewPlayerCustomBuildsWebStatusApi(configs ...func(*common.ApiClient)) (*PlayerCustomBuildsWebStatusApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &PlayerCustomBuildsWebStatusApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *PlayerCustomBuildsWebStatusApi) Get(customBuildId string) (*model.CustomPlayerBuildStatus, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["custom_build_id"] = customBuildId
    }

    var responseModel *model.CustomPlayerBuildStatus
    err := api.apiClient.Get("/player/custom-builds/web/{custom_build_id}/status", nil, &responseModel, reqParams)
    return responseModel, err
}

