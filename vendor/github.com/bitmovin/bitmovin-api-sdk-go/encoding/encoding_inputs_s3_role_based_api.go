package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingInputsS3RoleBasedApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingInputsS3RoleBasedCustomdataApi
}

func NewEncodingInputsS3RoleBasedApi(configs ...func(*common.ApiClient)) (*EncodingInputsS3RoleBasedApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInputsS3RoleBasedApi{apiClient: apiClient}

    customdataApi, err := NewEncodingInputsS3RoleBasedCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInputsS3RoleBasedApi) Create(s3RoleBasedInput model.S3RoleBasedInput) (*model.S3RoleBasedInput, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.S3RoleBasedInput
    err := api.apiClient.Post("/encoding/inputs/s3-role-based", &s3RoleBasedInput, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsS3RoleBasedApi) Delete(inputId string) (*model.S3RoleBasedInput, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel *model.S3RoleBasedInput
    err := api.apiClient.Delete("/encoding/inputs/s3-role-based/{input_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsS3RoleBasedApi) Get(inputId string) (*model.S3RoleBasedInput, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel *model.S3RoleBasedInput
    err := api.apiClient.Get("/encoding/inputs/s3-role-based/{input_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsS3RoleBasedApi) List(queryParams ...func(*query.S3RoleBasedInputListQueryParams)) (*pagination.S3RoleBasedInputsListPagination, error) {
    queryParameters := &query.S3RoleBasedInputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.S3RoleBasedInputsListPagination
    err := api.apiClient.Get("/encoding/inputs/s3-role-based", nil, &responseModel, reqParams)
    return responseModel, err
}

