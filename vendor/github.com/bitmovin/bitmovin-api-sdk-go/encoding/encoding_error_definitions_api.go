package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingErrorDefinitionsApi struct {
    apiClient *common.ApiClient
}

func NewEncodingErrorDefinitionsApi(configs ...func(*common.ApiClient)) (*EncodingErrorDefinitionsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingErrorDefinitionsApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingErrorDefinitionsApi) List(queryParams ...func(*query.EncodingErrorDefinitionListQueryParams)) (*pagination.EncodingErrorDefinitionsListPagination, error) {
    queryParameters := &query.EncodingErrorDefinitionListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.EncodingErrorDefinitionsListPagination
    err := api.apiClient.Get("/encoding/error-definitions", nil, &responseModel, reqParams)
    return responseModel, err
}

