package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingOutputsAkamaiNetstorageApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingOutputsAkamaiNetstorageCustomdataApi
}

func NewEncodingOutputsAkamaiNetstorageApi(configs ...func(*common.ApiClient)) (*EncodingOutputsAkamaiNetstorageApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingOutputsAkamaiNetstorageApi{apiClient: apiClient}

    customdataApi, err := NewEncodingOutputsAkamaiNetstorageCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingOutputsAkamaiNetstorageApi) Create(akamaiNetStorageOutput model.AkamaiNetStorageOutput) (*model.AkamaiNetStorageOutput, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.AkamaiNetStorageOutput
    err := api.apiClient.Post("/encoding/outputs/akamai-netstorage", &akamaiNetStorageOutput, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingOutputsAkamaiNetstorageApi) Delete(outputId string) (*model.AkamaiNetStorageOutput, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel *model.AkamaiNetStorageOutput
    err := api.apiClient.Delete("/encoding/outputs/akamai-netstorage/{output_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingOutputsAkamaiNetstorageApi) Get(outputId string) (*model.AkamaiNetStorageOutput, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel *model.AkamaiNetStorageOutput
    err := api.apiClient.Get("/encoding/outputs/akamai-netstorage/{output_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingOutputsAkamaiNetstorageApi) List(queryParams ...func(*query.AkamaiNetStorageOutputListQueryParams)) (*pagination.AkamaiNetStorageOutputsListPagination, error) {
    queryParameters := &query.AkamaiNetStorageOutputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.AkamaiNetStorageOutputsListPagination
    err := api.apiClient.Get("/encoding/outputs/akamai-netstorage", nil, &responseModel, reqParams)
    return responseModel, err
}

