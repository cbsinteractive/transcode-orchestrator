package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsCaptionsSccApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsCaptionsSccCustomdataApi
}

func NewEncodingEncodingsCaptionsSccApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsCaptionsSccApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsCaptionsSccApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsCaptionsSccCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsCaptionsSccApi) Create(encodingId string, convertSccCaption model.ConvertSccCaption) (*model.ConvertSccCaption, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.ConvertSccCaption
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/captions/scc", &convertSccCaption, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsCaptionsSccApi) Delete(encodingId string, captionsId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["captions_id"] = captionsId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/captions/scc/{captions_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsCaptionsSccApi) Get(encodingId string, captionsId string) (*model.ConvertSccCaption, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["captions_id"] = captionsId
    }

    var responseModel *model.ConvertSccCaption
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/captions/scc/{captions_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsCaptionsSccApi) List(encodingId string, queryParams ...func(*query.ConvertSccCaptionListQueryParams)) (*pagination.ConvertSccCaptionsListPagination, error) {
    queryParameters := &query.ConvertSccCaptionListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.ConvertSccCaptionsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/captions/scc", nil, &responseModel, reqParams)
    return responseModel, err
}

