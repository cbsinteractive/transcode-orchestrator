package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsStreamsBifsApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsStreamsBifsCustomdataApi
}

func NewEncodingEncodingsStreamsBifsApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsStreamsBifsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsStreamsBifsApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsStreamsBifsCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsStreamsBifsApi) Create(encodingId string, streamId string, bif model.Bif) (*model.Bif, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
    }

    var responseModel *model.Bif
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/streams/{stream_id}/bifs", &bif, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsStreamsBifsApi) Delete(encodingId string, streamId string, bifId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.PathParams["bif_id"] = bifId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/streams/{stream_id}/bifs/{bif_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsStreamsBifsApi) Get(encodingId string, streamId string, bifId string) (*model.Bif, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.PathParams["bif_id"] = bifId
    }

    var responseModel *model.Bif
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/bifs/{bif_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsStreamsBifsApi) List(encodingId string, streamId string, queryParams ...func(*query.BifListQueryParams)) (*pagination.BifsListPagination, error) {
    queryParameters := &query.BifListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.BifsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/bifs", nil, &responseModel, reqParams)
    return responseModel, err
}

