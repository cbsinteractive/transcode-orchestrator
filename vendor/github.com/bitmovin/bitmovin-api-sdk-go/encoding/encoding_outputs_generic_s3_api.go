package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingOutputsGenericS3Api struct {
    apiClient *common.ApiClient
    Customdata *EncodingOutputsGenericS3CustomdataApi
}

func NewEncodingOutputsGenericS3Api(configs ...func(*common.ApiClient)) (*EncodingOutputsGenericS3Api, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingOutputsGenericS3Api{apiClient: apiClient}

    customdataApi, err := NewEncodingOutputsGenericS3CustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingOutputsGenericS3Api) Create(genericS3Output model.GenericS3Output) (*model.GenericS3Output, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.GenericS3Output
    err := api.apiClient.Post("/encoding/outputs/generic-s3", &genericS3Output, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingOutputsGenericS3Api) Delete(outputId string) (*model.GenericS3Output, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel *model.GenericS3Output
    err := api.apiClient.Delete("/encoding/outputs/generic-s3/{output_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingOutputsGenericS3Api) Get(outputId string) (*model.GenericS3Output, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel *model.GenericS3Output
    err := api.apiClient.Get("/encoding/outputs/generic-s3/{output_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingOutputsGenericS3Api) List(queryParams ...func(*query.GenericS3OutputListQueryParams)) (*pagination.GenericS3OutputsListPagination, error) {
    queryParameters := &query.GenericS3OutputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.GenericS3OutputsListPagination
    err := api.apiClient.Get("/encoding/outputs/generic-s3", nil, &responseModel, reqParams)
    return responseModel, err
}

