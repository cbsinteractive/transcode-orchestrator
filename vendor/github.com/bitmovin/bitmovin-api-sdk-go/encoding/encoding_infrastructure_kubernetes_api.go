package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingInfrastructureKubernetesApi struct {
    apiClient *common.ApiClient
    Status *EncodingInfrastructureKubernetesStatusApi
    Customdata *EncodingInfrastructureKubernetesCustomdataApi
    Configuration *EncodingInfrastructureKubernetesConfigurationApi
    AgentDeployment *EncodingInfrastructureKubernetesAgentDeploymentApi
    PrewarmedDeployment *EncodingInfrastructureKubernetesPrewarmedDeploymentApi
}

func NewEncodingInfrastructureKubernetesApi(configs ...func(*common.ApiClient)) (*EncodingInfrastructureKubernetesApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInfrastructureKubernetesApi{apiClient: apiClient}

    statusApi, err := NewEncodingInfrastructureKubernetesStatusApi(configs...)
    api.Status = statusApi
    customdataApi, err := NewEncodingInfrastructureKubernetesCustomdataApi(configs...)
    api.Customdata = customdataApi
    configurationApi, err := NewEncodingInfrastructureKubernetesConfigurationApi(configs...)
    api.Configuration = configurationApi
    agentDeploymentApi, err := NewEncodingInfrastructureKubernetesAgentDeploymentApi(configs...)
    api.AgentDeployment = agentDeploymentApi
    prewarmedDeploymentApi, err := NewEncodingInfrastructureKubernetesPrewarmedDeploymentApi(configs...)
    api.PrewarmedDeployment = prewarmedDeploymentApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInfrastructureKubernetesApi) Create(kubernetesCluster model.KubernetesCluster) (*model.KubernetesCluster, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.KubernetesCluster
    err := api.apiClient.Post("/encoding/infrastructure/kubernetes", &kubernetesCluster, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInfrastructureKubernetesApi) Delete(infrastructureId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["infrastructure_id"] = infrastructureId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/infrastructure/kubernetes/{infrastructure_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInfrastructureKubernetesApi) Get(infrastructureId string) (*model.KubernetesCluster, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["infrastructure_id"] = infrastructureId
    }

    var responseModel *model.KubernetesCluster
    err := api.apiClient.Get("/encoding/infrastructure/kubernetes/{infrastructure_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInfrastructureKubernetesApi) List(queryParams ...func(*query.KubernetesClusterListQueryParams)) (*pagination.KubernetesClustersListPagination, error) {
    queryParameters := &query.KubernetesClusterListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.KubernetesClustersListPagination
    err := api.apiClient.Get("/encoding/infrastructure/kubernetes", nil, &responseModel, reqParams)
    return responseModel, err
}

