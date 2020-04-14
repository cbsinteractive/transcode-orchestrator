package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingInputsAkamaiNetstorageApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingInputsAkamaiNetstorageCustomdataApi
}

func NewEncodingInputsAkamaiNetstorageApi(configs ...func(*common.ApiClient)) (*EncodingInputsAkamaiNetstorageApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInputsAkamaiNetstorageApi{apiClient: apiClient}

    customdataApi, err := NewEncodingInputsAkamaiNetstorageCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInputsAkamaiNetstorageApi) Create(akamaiNetStorageInput model.AkamaiNetStorageInput) (*model.AkamaiNetStorageInput, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.AkamaiNetStorageInput
    err := api.apiClient.Post("/encoding/inputs/akamai-netstorage", &akamaiNetStorageInput, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsAkamaiNetstorageApi) Delete(inputId string) (*model.AkamaiNetStorageInput, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel *model.AkamaiNetStorageInput
    err := api.apiClient.Delete("/encoding/inputs/akamai-netstorage/{input_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsAkamaiNetstorageApi) Get(inputId string) (*model.AkamaiNetStorageInput, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel *model.AkamaiNetStorageInput
    err := api.apiClient.Get("/encoding/inputs/akamai-netstorage/{input_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsAkamaiNetstorageApi) List(queryParams ...func(*query.AkamaiNetStorageInputListQueryParams)) (*pagination.AkamaiNetStorageInputsListPagination, error) {
    queryParameters := &query.AkamaiNetStorageInputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.AkamaiNetStorageInputsListPagination
    err := api.apiClient.Get("/encoding/inputs/akamai-netstorage", nil, &responseModel, reqParams)
    return responseModel, err
}

