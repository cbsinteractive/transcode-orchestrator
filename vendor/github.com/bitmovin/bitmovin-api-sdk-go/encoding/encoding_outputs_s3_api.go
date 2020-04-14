package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingOutputsS3Api struct {
    apiClient *common.ApiClient
    Customdata *EncodingOutputsS3CustomdataApi
}

func NewEncodingOutputsS3Api(configs ...func(*common.ApiClient)) (*EncodingOutputsS3Api, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingOutputsS3Api{apiClient: apiClient}

    customdataApi, err := NewEncodingOutputsS3CustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingOutputsS3Api) Create(s3Output model.S3Output) (*model.S3Output, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.S3Output
    err := api.apiClient.Post("/encoding/outputs/s3", &s3Output, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingOutputsS3Api) Delete(outputId string) (*model.S3Output, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel *model.S3Output
    err := api.apiClient.Delete("/encoding/outputs/s3/{output_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingOutputsS3Api) Get(outputId string) (*model.S3Output, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel *model.S3Output
    err := api.apiClient.Get("/encoding/outputs/s3/{output_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingOutputsS3Api) List(queryParams ...func(*query.S3OutputListQueryParams)) (*pagination.S3OutputsListPagination, error) {
    queryParameters := &query.S3OutputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.S3OutputsListPagination
    err := api.apiClient.Get("/encoding/outputs/s3", nil, &responseModel, reqParams)
    return responseModel, err
}

