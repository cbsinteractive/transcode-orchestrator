package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingOutputsFtpApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingOutputsFtpCustomdataApi
}

func NewEncodingOutputsFtpApi(configs ...func(*common.ApiClient)) (*EncodingOutputsFtpApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingOutputsFtpApi{apiClient: apiClient}

    customdataApi, err := NewEncodingOutputsFtpCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingOutputsFtpApi) Create(ftpOutput model.FtpOutput) (*model.FtpOutput, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.FtpOutput
    err := api.apiClient.Post("/encoding/outputs/ftp", &ftpOutput, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingOutputsFtpApi) Delete(outputId string) (*model.FtpOutput, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel *model.FtpOutput
    err := api.apiClient.Delete("/encoding/outputs/ftp/{output_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingOutputsFtpApi) Get(outputId string) (*model.FtpOutput, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel *model.FtpOutput
    err := api.apiClient.Get("/encoding/outputs/ftp/{output_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingOutputsFtpApi) List(queryParams ...func(*query.FtpOutputListQueryParams)) (*pagination.FtpOutputsListPagination, error) {
    queryParameters := &query.FtpOutputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.FtpOutputsListPagination
    err := api.apiClient.Get("/encoding/outputs/ftp", nil, &responseModel, reqParams)
    return responseModel, err
}

