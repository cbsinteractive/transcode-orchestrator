package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingConfigurationsAudioOpusApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingConfigurationsAudioOpusCustomdataApi
}

func NewEncodingConfigurationsAudioOpusApi(configs ...func(*common.ApiClient)) (*EncodingConfigurationsAudioOpusApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsAudioOpusApi{apiClient: apiClient}

    customdataApi, err := NewEncodingConfigurationsAudioOpusCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsAudioOpusApi) Create(opusAudioConfiguration model.OpusAudioConfiguration) (*model.OpusAudioConfiguration, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.OpusAudioConfiguration
    err := api.apiClient.Post("/encoding/configurations/audio/opus", &opusAudioConfiguration, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsAudioOpusApi) Delete(configurationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/configurations/audio/opus/{configuration_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsAudioOpusApi) Get(configurationId string) (*model.OpusAudioConfiguration, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.OpusAudioConfiguration
    err := api.apiClient.Get("/encoding/configurations/audio/opus/{configuration_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsAudioOpusApi) List(queryParams ...func(*query.OpusAudioConfigurationListQueryParams)) (*pagination.OpusAudioConfigurationsListPagination, error) {
    queryParameters := &query.OpusAudioConfigurationListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.OpusAudioConfigurationsListPagination
    err := api.apiClient.Get("/encoding/configurations/audio/opus", nil, &responseModel, reqParams)
    return responseModel, err
}

