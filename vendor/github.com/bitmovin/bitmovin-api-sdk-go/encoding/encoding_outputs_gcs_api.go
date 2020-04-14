package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingOutputsGcsApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingOutputsGcsCustomdataApi
}

func NewEncodingOutputsGcsApi(configs ...func(*common.ApiClient)) (*EncodingOutputsGcsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingOutputsGcsApi{apiClient: apiClient}

    customdataApi, err := NewEncodingOutputsGcsCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingOutputsGcsApi) Create(gcsOutput model.GcsOutput) (*model.GcsOutput, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.GcsOutput
    err := api.apiClient.Post("/encoding/outputs/gcs", &gcsOutput, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingOutputsGcsApi) Delete(outputId string) (*model.GcsOutput, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel *model.GcsOutput
    err := api.apiClient.Delete("/encoding/outputs/gcs/{output_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingOutputsGcsApi) Get(outputId string) (*model.GcsOutput, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel *model.GcsOutput
    err := api.apiClient.Get("/encoding/outputs/gcs/{output_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingOutputsGcsApi) List(queryParams ...func(*query.GcsOutputListQueryParams)) (*pagination.GcsOutputsListPagination, error) {
    queryParameters := &query.GcsOutputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.GcsOutputsListPagination
    err := api.apiClient.Get("/encoding/outputs/gcs", nil, &responseModel, reqParams)
    return responseModel, err
}

