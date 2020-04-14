package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingConfigurationsVideoAv1Api struct {
    apiClient *common.ApiClient
    Customdata *EncodingConfigurationsVideoAv1CustomdataApi
}

func NewEncodingConfigurationsVideoAv1Api(configs ...func(*common.ApiClient)) (*EncodingConfigurationsVideoAv1Api, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsVideoAv1Api{apiClient: apiClient}

    customdataApi, err := NewEncodingConfigurationsVideoAv1CustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsVideoAv1Api) Create(av1VideoConfiguration model.Av1VideoConfiguration) (*model.Av1VideoConfiguration, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.Av1VideoConfiguration
    err := api.apiClient.Post("/encoding/configurations/video/av1", &av1VideoConfiguration, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsVideoAv1Api) Delete(configurationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/configurations/video/av1/{configuration_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsVideoAv1Api) Get(configurationId string) (*model.Av1VideoConfiguration, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.Av1VideoConfiguration
    err := api.apiClient.Get("/encoding/configurations/video/av1/{configuration_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsVideoAv1Api) List(queryParams ...func(*query.Av1VideoConfigurationListQueryParams)) (*pagination.Av1VideoConfigurationsListPagination, error) {
    queryParameters := &query.Av1VideoConfigurationListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.Av1VideoConfigurationsListPagination
    err := api.apiClient.Get("/encoding/configurations/video/av1", nil, &responseModel, reqParams)
    return responseModel, err
}

