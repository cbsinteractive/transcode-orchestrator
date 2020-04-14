package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsMp3Api struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsMp3CustomdataApi
    Information *EncodingEncodingsMuxingsMp3InformationApi
}

func NewEncodingEncodingsMuxingsMp3Api(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsMp3Api, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsMp3Api{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsMp3CustomdataApi(configs...)
    api.Customdata = customdataApi
    informationApi, err := NewEncodingEncodingsMuxingsMp3InformationApi(configs...)
    api.Information = informationApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsMp3Api) Create(encodingId string, mp3Muxing model.Mp3Muxing) (*model.Mp3Muxing, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.Mp3Muxing
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/mp3", &mp3Muxing, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsMp3Api) Delete(encodingId string, muxingId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/mp3/{muxing_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsMp3Api) Get(encodingId string, muxingId string) (*model.Mp3Muxing, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.Mp3Muxing
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp3/{muxing_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsMuxingsMp3Api) List(encodingId string, queryParams ...func(*query.Mp3MuxingListQueryParams)) (*pagination.Mp3MuxingsListPagination, error) {
    queryParameters := &query.Mp3MuxingListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.Mp3MuxingsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/mp3", nil, &responseModel, reqParams)
    return responseModel, err
}

