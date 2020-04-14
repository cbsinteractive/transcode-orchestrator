package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingInfrastructureAwsRegionsApi struct {
    apiClient *common.ApiClient
}

func NewEncodingInfrastructureAwsRegionsApi(configs ...func(*common.ApiClient)) (*EncodingInfrastructureAwsRegionsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInfrastructureAwsRegionsApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInfrastructureAwsRegionsApi) Create(infrastructureId string, region model.AwsCloudRegion, awsAccountRegionSettings model.AwsAccountRegionSettings) (*model.AwsAccountRegionSettings, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["infrastructure_id"] = infrastructureId
        params.PathParams["region"] = string(region)
    }

    var responseModel *model.AwsAccountRegionSettings
    err := api.apiClient.Post("/encoding/infrastructure/aws/{infrastructure_id}/regions/{region}", &awsAccountRegionSettings, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInfrastructureAwsRegionsApi) Delete(infrastructureId string, region model.AwsCloudRegion) (*model.AwsAccountRegionSettings, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["infrastructure_id"] = infrastructureId
        params.PathParams["region"] = string(region)
    }

    var responseModel *model.AwsAccountRegionSettings
    err := api.apiClient.Delete("/encoding/infrastructure/aws/{infrastructure_id}/regions/{region}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInfrastructureAwsRegionsApi) Get(infrastructureId string, region model.AwsCloudRegion) (*model.AwsAccountRegionSettings, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["infrastructure_id"] = infrastructureId
        params.PathParams["region"] = string(region)
    }

    var responseModel *model.AwsAccountRegionSettings
    err := api.apiClient.Get("/encoding/infrastructure/aws/{infrastructure_id}/regions/{region}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInfrastructureAwsRegionsApi) List(infrastructureId string, queryParams ...func(*query.AwsAccountRegionSettingsListQueryParams)) (*pagination.AwsAccountRegionSettingssListPagination, error) {
    queryParameters := &query.AwsAccountRegionSettingsListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["infrastructure_id"] = infrastructureId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.AwsAccountRegionSettingssListPagination
    err := api.apiClient.Get("/encoding/infrastructure/aws/{infrastructure_id}/regions", nil, &responseModel, reqParams)
    return responseModel, err
}

