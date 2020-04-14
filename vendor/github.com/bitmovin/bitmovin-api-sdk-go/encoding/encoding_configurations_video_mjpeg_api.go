package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingConfigurationsVideoMjpegApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingConfigurationsVideoMjpegCustomdataApi
}

func NewEncodingConfigurationsVideoMjpegApi(configs ...func(*common.ApiClient)) (*EncodingConfigurationsVideoMjpegApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsVideoMjpegApi{apiClient: apiClient}

    customdataApi, err := NewEncodingConfigurationsVideoMjpegCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsVideoMjpegApi) Create(mjpegVideoConfiguration model.MjpegVideoConfiguration) (*model.MjpegVideoConfiguration, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.MjpegVideoConfiguration
    err := api.apiClient.Post("/encoding/configurations/video/mjpeg", &mjpegVideoConfiguration, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsVideoMjpegApi) Delete(configurationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/configurations/video/mjpeg/{configuration_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsVideoMjpegApi) Get(configurationId string) (*model.MjpegVideoConfiguration, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.MjpegVideoConfiguration
    err := api.apiClient.Get("/encoding/configurations/video/mjpeg/{configuration_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsVideoMjpegApi) List(queryParams ...func(*query.MjpegVideoConfigurationListQueryParams)) (*pagination.MjpegVideoConfigurationsListPagination, error) {
    queryParameters := &query.MjpegVideoConfigurationListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.MjpegVideoConfigurationsListPagination
    err := api.apiClient.Get("/encoding/configurations/video/mjpeg", nil, &responseModel, reqParams)
    return responseModel, err
}

