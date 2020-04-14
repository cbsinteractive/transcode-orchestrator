package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type EncodingInfrastructureKubernetesStatusApi struct {
    apiClient *common.ApiClient
}

func NewEncodingInfrastructureKubernetesStatusApi(configs ...func(*common.ApiClient)) (*EncodingInfrastructureKubernetesStatusApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInfrastructureKubernetesStatusApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInfrastructureKubernetesStatusApi) Get(infrastructureId string) (error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["infrastructure_id"] = infrastructureId
    }

    
    err := api.apiClient.Get("/encoding/infrastructure/kubernetes/{infrastructure_id}/status", nil, nil, reqParams)
    return err
}

