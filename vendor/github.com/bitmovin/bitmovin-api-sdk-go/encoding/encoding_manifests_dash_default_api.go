package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingManifestsDashDefaultApi struct {
    apiClient *common.ApiClient
}

func NewEncodingManifestsDashDefaultApi(configs ...func(*common.ApiClient)) (*EncodingManifestsDashDefaultApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsDashDefaultApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsDashDefaultApi) Create(dashManifestDefault model.DashManifestDefault) (*model.DashManifestDefault, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.DashManifestDefault
    err := api.apiClient.Post("/encoding/manifests/dash/default", &dashManifestDefault, &responseModel, reqParams)
    return responseModel, err
}

