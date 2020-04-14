package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingManifestsSmoothRepresentationsMp4Api struct {
    apiClient *common.ApiClient
}

func NewEncodingManifestsSmoothRepresentationsMp4Api(configs ...func(*common.ApiClient)) (*EncodingManifestsSmoothRepresentationsMp4Api, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsSmoothRepresentationsMp4Api{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsSmoothRepresentationsMp4Api) Create(manifestId string, smoothStreamingRepresentation model.SmoothStreamingRepresentation) (*model.SmoothStreamingRepresentation, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }

    var responseModel *model.SmoothStreamingRepresentation
    err := api.apiClient.Post("/encoding/manifests/smooth/{manifest_id}/representations/mp4", &smoothStreamingRepresentation, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsSmoothRepresentationsMp4Api) Delete(manifestId string, representationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["representation_id"] = representationId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/manifests/smooth/{manifest_id}/representations/mp4/{representation_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsSmoothRepresentationsMp4Api) Get(manifestId string, representationId string) (*model.SmoothStreamingRepresentation, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["representation_id"] = representationId
    }

    var responseModel *model.SmoothStreamingRepresentation
    err := api.apiClient.Get("/encoding/manifests/smooth/{manifest_id}/representations/mp4/{representation_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsSmoothRepresentationsMp4Api) List(manifestId string, queryParams ...func(*query.SmoothStreamingRepresentationListQueryParams)) (*pagination.SmoothStreamingRepresentationsListPagination, error) {
    queryParameters := &query.SmoothStreamingRepresentationListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.SmoothStreamingRepresentationsListPagination
    err := api.apiClient.Get("/encoding/manifests/smooth/{manifest_id}/representations/mp4", nil, &responseModel, reqParams)
    return responseModel, err
}

