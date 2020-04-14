package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingInfrastructureKubernetesPrewarmedDeploymentApi struct {
    apiClient *common.ApiClient
}

func NewEncodingInfrastructureKubernetesPrewarmedDeploymentApi(configs ...func(*common.ApiClient)) (*EncodingInfrastructureKubernetesPrewarmedDeploymentApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInfrastructureKubernetesPrewarmedDeploymentApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInfrastructureKubernetesPrewarmedDeploymentApi) Create(infrastructureId string, prewarmEncoderSettings model.PrewarmEncoderSettings) (*model.PrewarmEncoderSettings, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["infrastructure_id"] = infrastructureId
    }

    var responseModel *model.PrewarmEncoderSettings
    err := api.apiClient.Post("/encoding/infrastructure/kubernetes/{infrastructure_id}/prewarmed-deployment", &prewarmEncoderSettings, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInfrastructureKubernetesPrewarmedDeploymentApi) Delete(infrastructureId string, deploymentId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["infrastructure_id"] = infrastructureId
        params.PathParams["deployment_id"] = deploymentId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/infrastructure/kubernetes/{infrastructure_id}/prewarmed-deployment/{deployment_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInfrastructureKubernetesPrewarmedDeploymentApi) Get(infrastructureId string, deploymentId string) (*model.PrewarmEncoderSettings, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["infrastructure_id"] = infrastructureId
        params.PathParams["deployment_id"] = deploymentId
    }

    var responseModel *model.PrewarmEncoderSettings
    err := api.apiClient.Get("/encoding/infrastructure/kubernetes/{infrastructure_id}/prewarmed-deployment/{deployment_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInfrastructureKubernetesPrewarmedDeploymentApi) List(infrastructureId string, queryParams ...func(*query.PrewarmEncoderSettingsListQueryParams)) (*pagination.PrewarmEncoderSettingssListPagination, error) {
    queryParameters := &query.PrewarmEncoderSettingsListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["infrastructure_id"] = infrastructureId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.PrewarmEncoderSettingssListPagination
    err := api.apiClient.Get("/encoding/infrastructure/kubernetes/{infrastructure_id}/prewarmed-deployment", nil, &responseModel, reqParams)
    return responseModel, err
}

