package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingConfigurationsAudioEac3Api struct {
    apiClient *common.ApiClient
    Customdata *EncodingConfigurationsAudioEac3CustomdataApi
}

func NewEncodingConfigurationsAudioEac3Api(configs ...func(*common.ApiClient)) (*EncodingConfigurationsAudioEac3Api, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsAudioEac3Api{apiClient: apiClient}

    customdataApi, err := NewEncodingConfigurationsAudioEac3CustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsAudioEac3Api) Create(eac3AudioConfiguration model.Eac3AudioConfiguration) (*model.Eac3AudioConfiguration, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.Eac3AudioConfiguration
    err := api.apiClient.Post("/encoding/configurations/audio/eac3", &eac3AudioConfiguration, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsAudioEac3Api) Delete(configurationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/configurations/audio/eac3/{configuration_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsAudioEac3Api) Get(configurationId string) (*model.Eac3AudioConfiguration, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.Eac3AudioConfiguration
    err := api.apiClient.Get("/encoding/configurations/audio/eac3/{configuration_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsAudioEac3Api) List(queryParams ...func(*query.Eac3AudioConfigurationListQueryParams)) (*pagination.Eac3AudioConfigurationsListPagination, error) {
    queryParameters := &query.Eac3AudioConfigurationListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.Eac3AudioConfigurationsListPagination
    err := api.apiClient.Get("/encoding/configurations/audio/eac3", nil, &responseModel, reqParams)
    return responseModel, err
}

