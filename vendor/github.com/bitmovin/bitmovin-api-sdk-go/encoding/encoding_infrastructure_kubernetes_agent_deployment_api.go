package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type EncodingInfrastructureKubernetesAgentDeploymentApi struct {
    apiClient *common.ApiClient
}

func NewEncodingInfrastructureKubernetesAgentDeploymentApi(configs ...func(*common.ApiClient)) (*EncodingInfrastructureKubernetesAgentDeploymentApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInfrastructureKubernetesAgentDeploymentApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInfrastructureKubernetesAgentDeploymentApi) Get(infrastructureId string) (error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["infrastructure_id"] = infrastructureId
    }

    
    err := api.apiClient.Get("/encoding/infrastructure/kubernetes/{infrastructure_id}/agent-deployment", nil, nil, reqParams)
    return err
}

