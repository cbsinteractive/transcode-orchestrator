package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingInputsRedundantRtmpApi struct {
    apiClient *common.ApiClient
}

func NewEncodingInputsRedundantRtmpApi(configs ...func(*common.ApiClient)) (*EncodingInputsRedundantRtmpApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInputsRedundantRtmpApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInputsRedundantRtmpApi) Create(redundantRtmpInput model.RedundantRtmpInput) (*model.RedundantRtmpInput, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.RedundantRtmpInput
    err := api.apiClient.Post("/encoding/inputs/redundant-rtmp", &redundantRtmpInput, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsRedundantRtmpApi) Delete(inputId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/inputs/redundant-rtmp/{input_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsRedundantRtmpApi) Get(inputId string) (*model.RedundantRtmpInput, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel *model.RedundantRtmpInput
    err := api.apiClient.Get("/encoding/inputs/redundant-rtmp/{input_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInputsRedundantRtmpApi) List(queryParams ...func(*query.RedundantRtmpInputListQueryParams)) (*pagination.RedundantRtmpInputsListPagination, error) {
    queryParameters := &query.RedundantRtmpInputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.RedundantRtmpInputsListPagination
    err := api.apiClient.Get("/encoding/inputs/redundant-rtmp", nil, &responseModel, reqParams)
    return responseModel, err
}

