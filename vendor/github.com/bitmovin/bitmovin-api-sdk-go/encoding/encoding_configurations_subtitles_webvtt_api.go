package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingConfigurationsSubtitlesWebvttApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingConfigurationsSubtitlesWebvttCustomdataApi
}

func NewEncodingConfigurationsSubtitlesWebvttApi(configs ...func(*common.ApiClient)) (*EncodingConfigurationsSubtitlesWebvttApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsSubtitlesWebvttApi{apiClient: apiClient}

    customdataApi, err := NewEncodingConfigurationsSubtitlesWebvttCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsSubtitlesWebvttApi) Create(webVttConfiguration model.WebVttConfiguration) (*model.WebVttConfiguration, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.WebVttConfiguration
    err := api.apiClient.Post("/encoding/configurations/subtitles/webvtt/", &webVttConfiguration, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsSubtitlesWebvttApi) Delete(configurationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/configurations/subtitles/webvtt/{configuration_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsSubtitlesWebvttApi) Get(configurationId string) (*model.WebVttConfiguration, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.WebVttConfiguration
    err := api.apiClient.Get("/encoding/configurations/subtitles/webvtt/{configuration_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsSubtitlesWebvttApi) List(queryParams ...func(*query.WebVttConfigurationListQueryParams)) (*pagination.WebVttConfigurationsListPagination, error) {
    queryParameters := &query.WebVttConfigurationListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.WebVttConfigurationsListPagination
    err := api.apiClient.Get("/encoding/configurations/subtitles/webvtt/", nil, &responseModel, reqParams)
    return responseModel, err
}

