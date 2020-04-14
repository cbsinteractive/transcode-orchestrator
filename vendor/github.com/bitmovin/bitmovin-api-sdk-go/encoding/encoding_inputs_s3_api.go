package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingInputsS3Api struct {
    apiClient *common.ApiClient
    Customdata *EncodingInputsS3CustomdataApi
}

func NewEncodingInputsS3Api(configs ...func(*common.ApiClient)) (*EncodingInputsS3Api, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInputsS3Api{apiClient: apiClient}

    customdataApi, err := NewEncodingInputsS3CustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInputsS3Api) Create(s3Input model.S3Input) (*model.S3Input, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.S3Input
    err := api.apiClient.Post("/encoding/inputs/s3", &s3Input, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsS3Api) Delete(inputId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/inputs/s3/{input_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsS3Api) Get(inputId string) (*model.S3Input, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel *model.S3Input
    err := api.apiClient.Get("/encoding/inputs/s3/{input_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsS3Api) List(queryParams ...func(*query.S3InputListQueryParams)) (*pagination.S3InputsListPagination, error) {
    queryParameters := &query.S3InputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.S3InputsListPagination
    err := api.apiClient.Get("/encoding/inputs/s3", nil, &responseModel, reqParams)
    return responseModel, err
}

