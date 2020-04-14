package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingInputsTcpApi struct {
    apiClient *common.ApiClient
}

func NewEncodingInputsTcpApi(configs ...func(*common.ApiClient)) (*EncodingInputsTcpApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInputsTcpApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInputsTcpApi) Get(inputId string) (*model.TcpInput, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel *model.TcpInput
    err := api.apiClient.Get("/encoding/inputs/tcp/{input_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsTcpApi) List(queryParams ...func(*query.TcpInputListQueryParams)) (*pagination.TcpInputsListPagination, error) {
    queryParameters := &query.TcpInputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.TcpInputsListPagination
    err := api.apiClient.Get("/encoding/inputs/tcp", nil, &responseModel, reqParams)
    return responseModel, err
}

