package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingOutputsLiveMediaIngestApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingOutputsLiveMediaIngestCustomdataApi
}

func NewEncodingOutputsLiveMediaIngestApi(configs ...func(*common.ApiClient)) (*EncodingOutputsLiveMediaIngestApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingOutputsLiveMediaIngestApi{apiClient: apiClient}

    customdataApi, err := NewEncodingOutputsLiveMediaIngestCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingOutputsLiveMediaIngestApi) Create(liveMediaIngestOutput model.LiveMediaIngestOutput) (*model.LiveMediaIngestOutput, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.LiveMediaIngestOutput
    err := api.apiClient.Post("/encoding/outputs/live-media-ingest", &liveMediaIngestOutput, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingOutputsLiveMediaIngestApi) Delete(outputId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/outputs/live-media-ingest/{output_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingOutputsLiveMediaIngestApi) Get(outputId string) (*model.LiveMediaIngestOutput, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel *model.LiveMediaIngestOutput
    err := api.apiClient.Get("/encoding/outputs/live-media-ingest/{output_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingOutputsLiveMediaIngestApi) List(queryParams ...func(*query.LiveMediaIngestOutputListQueryParams)) (*pagination.LiveMediaIngestOutputsListPagination, error) {
    queryParameters := &query.LiveMediaIngestOutputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.LiveMediaIngestOutputsListPagination
    err := api.apiClient.Get("/encoding/outputs/live-media-ingest", nil, &responseModel, reqParams)
    return responseModel, err
}

