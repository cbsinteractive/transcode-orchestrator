package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsSidecarsApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsSidecarsCustomdataApi
    Webvtt *EncodingEncodingsSidecarsWebvttApi
}

func NewEncodingEncodingsSidecarsApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsSidecarsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsSidecarsApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsSidecarsCustomdataApi(configs...)
    api.Customdata = customdataApi
    webvttApi, err := NewEncodingEncodingsSidecarsWebvttApi(configs...)
    api.Webvtt = webvttApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsSidecarsApi) Create(encodingId string, sidecarFile model.SidecarFile) (*model.SidecarFile, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.SidecarFile
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/sidecars", &sidecarFile, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsSidecarsApi) Delete(encodingId string, sidecarId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["sidecar_id"] = sidecarId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/sidecars/{sidecar_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsSidecarsApi) Get(encodingId string, sidecarId string) (*model.SidecarFile, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["sidecar_id"] = sidecarId
    }

    var responseModel *model.SidecarFile
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/sidecars/{sidecar_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsSidecarsApi) List(encodingId string, queryParams ...func(*query.SidecarFileListQueryParams)) (*pagination.SidecarFilesListPagination, error) {
    queryParameters := &query.SidecarFileListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.SidecarFilesListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/sidecars", nil, &responseModel, reqParams)
    return responseModel, err
}

