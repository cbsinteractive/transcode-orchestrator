package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingInputsGcsApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingInputsGcsCustomdataApi
}

func NewEncodingInputsGcsApi(configs ...func(*common.ApiClient)) (*EncodingInputsGcsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInputsGcsApi{apiClient: apiClient}

    customdataApi, err := NewEncodingInputsGcsCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInputsGcsApi) Create(gcsInput model.GcsInput) (*model.GcsInput, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.GcsInput
    err := api.apiClient.Post("/encoding/inputs/gcs", &gcsInput, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsGcsApi) Delete(inputId string) (*model.GcsInput, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel *model.GcsInput
    err := api.apiClient.Delete("/encoding/inputs/gcs/{input_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsGcsApi) Get(inputId string) (*model.GcsInput, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel *model.GcsInput
    err := api.apiClient.Get("/encoding/inputs/gcs/{input_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsGcsApi) List(queryParams ...func(*query.GcsInputListQueryParams)) (*pagination.GcsInputsListPagination, error) {
    queryParameters := &query.GcsInputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.GcsInputsListPagination
    err := api.apiClient.Get("/encoding/inputs/gcs", nil, &responseModel, reqParams)
    return responseModel, err
}

