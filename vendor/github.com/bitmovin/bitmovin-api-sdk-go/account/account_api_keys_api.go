package account
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type AccountApiKeysApi struct {
    apiClient *common.ApiClient
}

func NewAccountApiKeysApi(configs ...func(*common.ApiClient)) (*AccountApiKeysApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AccountApiKeysApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AccountApiKeysApi) Create() (*model.AccountApiKey, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.AccountApiKey
    err := api.apiClient.Post("/account/api-keys", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *AccountApiKeysApi) Delete(apiKeyId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["api_key_id"] = apiKeyId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/account/api-keys/{api_key_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *AccountApiKeysApi) Get(apiKeyId string) (*model.AccountApiKey, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["api_key_id"] = apiKeyId
    }

    var responseModel *model.AccountApiKey
    err := api.apiClient.Get("/account/api-keys/{api_key_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *AccountApiKeysApi) List() (*pagination.AccountApiKeysListPagination, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *pagination.AccountApiKeysListPagination
    err := api.apiClient.Get("/account/api-keys", nil, &responseModel, reqParams)
    return responseModel, err
}

