package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingConfigurationsVideoVp9Api struct {
    apiClient *common.ApiClient
    Customdata *EncodingConfigurationsVideoVp9CustomdataApi
}

func NewEncodingConfigurationsVideoVp9Api(configs ...func(*common.ApiClient)) (*EncodingConfigurationsVideoVp9Api, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsVideoVp9Api{apiClient: apiClient}

    customdataApi, err := NewEncodingConfigurationsVideoVp9CustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsVideoVp9Api) Create(vp9VideoConfiguration model.Vp9VideoConfiguration) (*model.Vp9VideoConfiguration, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.Vp9VideoConfiguration
    err := api.apiClient.Post("/encoding/configurations/video/vp9", &vp9VideoConfiguration, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsVideoVp9Api) Delete(configurationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/configurations/video/vp9/{configuration_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsVideoVp9Api) Get(configurationId string) (*model.Vp9VideoConfiguration, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.Vp9VideoConfiguration
    err := api.apiClient.Get("/encoding/configurations/video/vp9/{configuration_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsVideoVp9Api) List(queryParams ...func(*query.Vp9VideoConfigurationListQueryParams)) (*pagination.Vp9VideoConfigurationsListPagination, error) {
    queryParameters := &query.Vp9VideoConfigurationListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.Vp9VideoConfigurationsListPagination
    err := api.apiClient.Get("/encoding/configurations/video/vp9", nil, &responseModel, reqParams)
    return responseModel, err
}

