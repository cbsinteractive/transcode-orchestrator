package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingConfigurationsVideoH265Api struct {
    apiClient *common.ApiClient
    Customdata *EncodingConfigurationsVideoH265CustomdataApi
}

func NewEncodingConfigurationsVideoH265Api(configs ...func(*common.ApiClient)) (*EncodingConfigurationsVideoH265Api, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsVideoH265Api{apiClient: apiClient}

    customdataApi, err := NewEncodingConfigurationsVideoH265CustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsVideoH265Api) Create(h265VideoConfiguration model.H265VideoConfiguration) (*model.H265VideoConfiguration, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.H265VideoConfiguration
    err := api.apiClient.Post("/encoding/configurations/video/h265", &h265VideoConfiguration, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsVideoH265Api) Delete(configurationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/configurations/video/h265/{configuration_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsVideoH265Api) Get(configurationId string) (*model.H265VideoConfiguration, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.H265VideoConfiguration
    err := api.apiClient.Get("/encoding/configurations/video/h265/{configuration_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsVideoH265Api) List(queryParams ...func(*query.H265VideoConfigurationListQueryParams)) (*pagination.H265VideoConfigurationsListPagination, error) {
    queryParameters := &query.H265VideoConfigurationListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.H265VideoConfigurationsListPagination
    err := api.apiClient.Get("/encoding/configurations/video/h265", nil, &responseModel, reqParams)
    return responseModel, err
}

