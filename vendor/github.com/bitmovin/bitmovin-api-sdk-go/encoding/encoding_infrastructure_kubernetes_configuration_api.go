package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingInfrastructureKubernetesConfigurationApi struct {
    apiClient *common.ApiClient
}

func NewEncodingInfrastructureKubernetesConfigurationApi(configs ...func(*common.ApiClient)) (*EncodingInfrastructureKubernetesConfigurationApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInfrastructureKubernetesConfigurationApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInfrastructureKubernetesConfigurationApi) Get(infrastructureId string) (*model.KubernetesClusterConfiguration, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["infrastructure_id"] = infrastructureId
    }

    var responseModel *model.KubernetesClusterConfiguration
    err := api.apiClient.Get("/encoding/infrastructure/kubernetes/{infrastructure_id}/configuration", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInfrastructureKubernetesConfigurationApi) Update(infrastructureId string, kubernetesClusterConfiguration model.KubernetesClusterConfiguration) (*model.KubernetesClusterConfiguration, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["infrastructure_id"] = infrastructureId
    }

    var responseModel *model.KubernetesClusterConfiguration
    err := api.apiClient.Put("/encoding/infrastructure/kubernetes/{infrastructure_id}/configuration", &kubernetesClusterConfiguration, &responseModel, reqParams)
    return responseModel, err
}

