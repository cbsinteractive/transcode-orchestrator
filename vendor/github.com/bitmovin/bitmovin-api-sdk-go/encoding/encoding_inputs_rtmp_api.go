package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingInputsRtmpApi struct {
    apiClient *common.ApiClient
}

func NewEncodingInputsRtmpApi(configs ...func(*common.ApiClient)) (*EncodingInputsRtmpApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInputsRtmpApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInputsRtmpApi) Get(inputId string) (*model.RtmpInput, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel *model.RtmpInput
    err := api.apiClient.Get("/encoding/inputs/rtmp/{input_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsRtmpApi) List(queryParams ...func(*query.RtmpInputListQueryParams)) (*pagination.RtmpInputsListPagination, error) {
    queryParameters := &query.RtmpInputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.RtmpInputsListPagination
    err := api.apiClient.Get("/encoding/inputs/rtmp", nil, &responseModel, reqParams)
    return responseModel, err
}

