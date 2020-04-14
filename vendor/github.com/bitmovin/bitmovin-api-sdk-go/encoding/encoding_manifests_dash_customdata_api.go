package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingManifestsDashCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingManifestsDashCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingManifestsDashCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsDashCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsDashCustomdataApi) Get(manifestId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

