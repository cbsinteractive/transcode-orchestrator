package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingInputsSftpApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingInputsSftpCustomdataApi
}

func NewEncodingInputsSftpApi(configs ...func(*common.ApiClient)) (*EncodingInputsSftpApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInputsSftpApi{apiClient: apiClient}

    customdataApi, err := NewEncodingInputsSftpCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInputsSftpApi) Create(sftpInput model.SftpInput) (*model.SftpInput, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.SftpInput
    err := api.apiClient.Post("/encoding/inputs/sftp", &sftpInput, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsSftpApi) Delete(inputId string) (*model.SftpInput, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel *model.SftpInput
    err := api.apiClient.Delete("/encoding/inputs/sftp/{input_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsSftpApi) Get(inputId string) (*model.SftpInput, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel *model.SftpInput
    err := api.apiClient.Get("/encoding/inputs/sftp/{input_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsSftpApi) List(queryParams ...func(*query.SftpInputListQueryParams)) (*pagination.SftpInputsListPagination, error) {
    queryParameters := &query.SftpInputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.SftpInputsListPagination
    err := api.apiClient.Get("/encoding/inputs/sftp", nil, &responseModel, reqParams)
    return responseModel, err
}

