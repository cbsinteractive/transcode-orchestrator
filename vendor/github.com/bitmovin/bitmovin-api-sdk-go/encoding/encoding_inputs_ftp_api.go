package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingInputsFtpApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingInputsFtpCustomdataApi
}

func NewEncodingInputsFtpApi(configs ...func(*common.ApiClient)) (*EncodingInputsFtpApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInputsFtpApi{apiClient: apiClient}

    customdataApi, err := NewEncodingInputsFtpCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInputsFtpApi) Create(ftpInput model.FtpInput) (*model.FtpInput, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.FtpInput
    err := api.apiClient.Post("/encoding/inputs/ftp", &ftpInput, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsFtpApi) Delete(inputId string) (*model.FtpInput, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel *model.FtpInput
    err := api.apiClient.Delete("/encoding/inputs/ftp/{input_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsFtpApi) Get(inputId string) (*model.FtpInput, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel *model.FtpInput
    err := api.apiClient.Get("/encoding/inputs/ftp/{input_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsFtpApi) List(queryParams ...func(*query.FtpInputListQueryParams)) (*pagination.FtpInputsListPagination, error) {
    queryParameters := &query.FtpInputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.FtpInputsListPagination
    err := api.apiClient.Get("/encoding/inputs/ftp", nil, &responseModel, reqParams)
    return responseModel, err
}

