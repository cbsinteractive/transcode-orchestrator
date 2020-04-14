package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingInfrastructureKubernetesCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingInfrastructureKubernetesCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingInfrastructureKubernetesCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInfrastructureKubernetesCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInfrastructureKubernetesCustomdataApi) Get(infrastructureId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["infrastructure_id"] = infrastructureId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/infrastructure/kubernetes/{infrastructure_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

