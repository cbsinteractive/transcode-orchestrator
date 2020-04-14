package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingConfigurationsAudioAc3Api struct {
    apiClient *common.ApiClient
    Customdata *EncodingConfigurationsAudioAc3CustomdataApi
}

func NewEncodingConfigurationsAudioAc3Api(configs ...func(*common.ApiClient)) (*EncodingConfigurationsAudioAc3Api, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsAudioAc3Api{apiClient: apiClient}

    customdataApi, err := NewEncodingConfigurationsAudioAc3CustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsAudioAc3Api) Create(ac3AudioConfiguration model.Ac3AudioConfiguration) (*model.Ac3AudioConfiguration, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.Ac3AudioConfiguration
    err := api.apiClient.Post("/encoding/configurations/audio/ac3", &ac3AudioConfiguration, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsAudioAc3Api) Delete(configurationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/configurations/audio/ac3/{configuration_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsAudioAc3Api) Get(configurationId string) (*model.Ac3AudioConfiguration, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.Ac3AudioConfiguration
    err := api.apiClient.Get("/encoding/configurations/audio/ac3/{configuration_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsAudioAc3Api) List(queryParams ...func(*query.Ac3AudioConfigurationListQueryParams)) (*pagination.Ac3AudioConfigurationsListPagination, error) {
    queryParameters := &query.Ac3AudioConfigurationListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.Ac3AudioConfigurationsListPagination
    err := api.apiClient.Get("/encoding/configurations/audio/ac3", nil, &responseModel, reqParams)
    return responseModel, err
}

