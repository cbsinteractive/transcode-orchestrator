package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsStreamsBurnInSubtitlesSrtApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsStreamsBurnInSubtitlesSrtApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsStreamsBurnInSubtitlesSrtApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsStreamsBurnInSubtitlesSrtApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsStreamsBurnInSubtitlesSrtApi) Create(encodingId string, streamId string, burnInSubtitleSrt model.BurnInSubtitleSrt) (*model.BurnInSubtitleSrt, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
    }

    var responseModel *model.BurnInSubtitleSrt
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/streams/{stream_id}/burn-in-subtitles/srt", &burnInSubtitleSrt, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsStreamsBurnInSubtitlesSrtApi) Delete(encodingId string, streamId string, subtitleId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.PathParams["subtitle_id"] = subtitleId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/streams/{stream_id}/burn-in-subtitles/srt/{subtitle_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsStreamsBurnInSubtitlesSrtApi) Get(encodingId string, streamId string, subtitleId string) (*model.BurnInSubtitleSrt, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.PathParams["subtitle_id"] = subtitleId
    }

    var responseModel *model.BurnInSubtitleSrt
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/burn-in-subtitles/srt/{subtitle_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsStreamsBurnInSubtitlesSrtApi) List(encodingId string, streamId string, queryParams ...func(*query.BurnInSubtitleSrtListQueryParams)) (*pagination.BurnInSubtitleSrtsListPagination, error) {
    queryParameters := &query.BurnInSubtitleSrtListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.BurnInSubtitleSrtsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/burn-in-subtitles/srt", nil, &responseModel, reqParams)
    return responseModel, err
}

