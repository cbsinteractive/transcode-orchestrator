package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingOutputsS3RoleBasedApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingOutputsS3RoleBasedCustomdataApi
}

func NewEncodingOutputsS3RoleBasedApi(configs ...func(*common.ApiClient)) (*EncodingOutputsS3RoleBasedApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingOutputsS3RoleBasedApi{apiClient: apiClient}

    customdataApi, err := NewEncodingOutputsS3RoleBasedCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingOutputsS3RoleBasedApi) Create(s3RoleBasedOutput model.S3RoleBasedOutput) (*model.S3RoleBasedOutput, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.S3RoleBasedOutput
    err := api.apiClient.Post("/encoding/outputs/s3-role-based", &s3RoleBasedOutput, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingOutputsS3RoleBasedApi) Delete(outputId string) (*model.S3RoleBasedOutput, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel *model.S3RoleBasedOutput
    err := api.apiClient.Delete("/encoding/outputs/s3-role-based/{output_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingOutputsS3RoleBasedApi) Get(outputId string) (*model.S3RoleBasedOutput, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel *model.S3RoleBasedOutput
    err := api.apiClient.Get("/encoding/outputs/s3-role-based/{output_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingOutputsS3RoleBasedApi) List(queryParams ...func(*query.S3RoleBasedOutputListQueryParams)) (*pagination.S3RoleBasedOutputsListPagination, error) {
    queryParameters := &query.S3RoleBasedOutputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.S3RoleBasedOutputsListPagination
    err := api.apiClient.Get("/encoding/outputs/s3-role-based", nil, &responseModel, reqParams)
    return responseModel, err
}

