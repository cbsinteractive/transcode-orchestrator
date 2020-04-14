package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingOutputsLocalApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingOutputsLocalCustomdataApi
}

func NewEncodingOutputsLocalApi(configs ...func(*common.ApiClient)) (*EncodingOutputsLocalApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingOutputsLocalApi{apiClient: apiClient}

    customdataApi, err := NewEncodingOutputsLocalCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingOutputsLocalApi) Create(localOutput model.LocalOutput) (*model.LocalOutput, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.LocalOutput
    err := api.apiClient.Post("/encoding/outputs/local", &localOutput, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingOutputsLocalApi) Delete(outputId string) (*model.LocalOutput, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel *model.LocalOutput
    err := api.apiClient.Delete("/encoding/outputs/local/{output_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingOutputsLocalApi) Get(outputId string) (*model.LocalOutput, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel *model.LocalOutput
    err := api.apiClient.Get("/encoding/outputs/local/{output_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingOutputsLocalApi) List(queryParams ...func(*query.LocalOutputListQueryParams)) (*pagination.LocalOutputsListPagination, error) {
    queryParameters := &query.LocalOutputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.LocalOutputsListPagination
    err := api.apiClient.Get("/encoding/outputs/local", nil, &responseModel, reqParams)
    return responseModel, err
}

