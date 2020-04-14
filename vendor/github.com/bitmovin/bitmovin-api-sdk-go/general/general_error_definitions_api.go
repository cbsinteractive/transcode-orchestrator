package general
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type GeneralErrorDefinitionsApi struct {
    apiClient *common.ApiClient
}

func NewGeneralErrorDefinitionsApi(configs ...func(*common.ApiClient)) (*GeneralErrorDefinitionsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &GeneralErrorDefinitionsApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *GeneralErrorDefinitionsApi) List(queryParams ...func(*query.ApiErrorDefinitionListQueryParams)) (*pagination.ApiErrorDefinitionsListPagination, error) {
    queryParameters := &query.ApiErrorDefinitionListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.ApiErrorDefinitionsListPagination
    err := api.apiClient.Get("/general/error-definitions", nil, &responseModel, reqParams)
    return responseModel, err
}

