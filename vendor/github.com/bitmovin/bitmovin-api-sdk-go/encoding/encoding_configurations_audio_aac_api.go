package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingConfigurationsAudioAacApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingConfigurationsAudioAacCustomdataApi
}

func NewEncodingConfigurationsAudioAacApi(configs ...func(*common.ApiClient)) (*EncodingConfigurationsAudioAacApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsAudioAacApi{apiClient: apiClient}

    customdataApi, err := NewEncodingConfigurationsAudioAacCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsAudioAacApi) Create(aacAudioConfiguration model.AacAudioConfiguration) (*model.AacAudioConfiguration, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.AacAudioConfiguration
    err := api.apiClient.Post("/encoding/configurations/audio/aac", &aacAudioConfiguration, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsAudioAacApi) Delete(configurationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/configurations/audio/aac/{configuration_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsAudioAacApi) Get(configurationId string) (*model.AacAudioConfiguration, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.AacAudioConfiguration
    err := api.apiClient.Get("/encoding/configurations/audio/aac/{configuration_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsAudioAacApi) List(queryParams ...func(*query.AacAudioConfigurationListQueryParams)) (*pagination.AacAudioConfigurationsListPagination, error) {
    queryParameters := &query.AacAudioConfigurationListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.AacAudioConfigurationsListPagination
    err := api.apiClient.Get("/encoding/configurations/audio/aac", nil, &responseModel, reqParams)
    return responseModel, err
}

