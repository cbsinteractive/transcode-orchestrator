package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingManifestsSmoothCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingManifestsSmoothCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingManifestsSmoothCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsSmoothCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsSmoothCustomdataApi) Get(manifestId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/manifests/smooth/{manifest_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

