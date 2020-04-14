package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsKeyframesApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsKeyframesApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsKeyframesApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsKeyframesApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsKeyframesApi) Create(encodingId string, keyframe model.Keyframe) (*model.Keyframe, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.Keyframe
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/keyframes", &keyframe, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsKeyframesApi) Delete(encodingId string, keyframeId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["keyframe_id"] = keyframeId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/keyframes/{keyframe_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsKeyframesApi) Get(encodingId string, keyframeId string) (*model.Keyframe, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["keyframe_id"] = keyframeId
    }

    var responseModel *model.Keyframe
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/keyframes/{keyframe_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsKeyframesApi) List(encodingId string, queryParams ...func(*query.KeyframeListQueryParams)) (*pagination.KeyframesListPagination, error) {
    queryParameters := &query.KeyframeListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.KeyframesListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/keyframes", nil, &responseModel, reqParams)
    return responseModel, err
}

