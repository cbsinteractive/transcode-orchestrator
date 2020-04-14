package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingInputsGenericS3Api struct {
    apiClient *common.ApiClient
    Customdata *EncodingInputsGenericS3CustomdataApi
}

func NewEncodingInputsGenericS3Api(configs ...func(*common.ApiClient)) (*EncodingInputsGenericS3Api, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInputsGenericS3Api{apiClient: apiClient}

    customdataApi, err := NewEncodingInputsGenericS3CustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInputsGenericS3Api) Create(genericS3Input model.GenericS3Input) (*model.GenericS3Input, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.GenericS3Input
    err := api.apiClient.Post("/encoding/inputs/generic-s3", &genericS3Input, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsGenericS3Api) Delete(inputId string) (*model.GenericS3Input, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel *model.GenericS3Input
    err := api.apiClient.Delete("/encoding/inputs/generic-s3/{input_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsGenericS3Api) Get(inputId string) (*model.GenericS3Input, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel *model.GenericS3Input
    err := api.apiClient.Get("/encoding/inputs/generic-s3/{input_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsGenericS3Api) List(queryParams ...func(*query.GenericS3InputListQueryParams)) (*pagination.GenericS3InputsListPagination, error) {
    queryParameters := &query.GenericS3InputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.GenericS3InputsListPagination
    err := api.apiClient.Get("/encoding/inputs/generic-s3", nil, &responseModel, reqParams)
    return responseModel, err
}

