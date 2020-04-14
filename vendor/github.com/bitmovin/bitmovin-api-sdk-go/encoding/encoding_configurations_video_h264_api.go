package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingConfigurationsVideoH264Api struct {
    apiClient *common.ApiClient
    Customdata *EncodingConfigurationsVideoH264CustomdataApi
}

func NewEncodingConfigurationsVideoH264Api(configs ...func(*common.ApiClient)) (*EncodingConfigurationsVideoH264Api, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsVideoH264Api{apiClient: apiClient}

    customdataApi, err := NewEncodingConfigurationsVideoH264CustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsVideoH264Api) Create(h264VideoConfiguration model.H264VideoConfiguration) (*model.H264VideoConfiguration, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.H264VideoConfiguration
    err := api.apiClient.Post("/encoding/configurations/video/h264", &h264VideoConfiguration, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsVideoH264Api) Delete(configurationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/configurations/video/h264/{configuration_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsVideoH264Api) Get(configurationId string) (*model.H264VideoConfiguration, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.H264VideoConfiguration
    err := api.apiClient.Get("/encoding/configurations/video/h264/{configuration_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsVideoH264Api) List(queryParams ...func(*query.H264VideoConfigurationListQueryParams)) (*pagination.H264VideoConfigurationsListPagination, error) {
    queryParameters := &query.H264VideoConfigurationListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.H264VideoConfigurationsListPagination
    err := api.apiClient.Get("/encoding/configurations/video/h264", nil, &responseModel, reqParams)
    return responseModel, err
}

