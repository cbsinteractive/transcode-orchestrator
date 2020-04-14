package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsProgressiveTsId3Api struct {
    apiClient *common.ApiClient
    Raw *EncodingEncodingsMuxingsProgressiveTsId3RawApi
    FrameId *EncodingEncodingsMuxingsProgressiveTsId3FrameIdApi
    PlainText *EncodingEncodingsMuxingsProgressiveTsId3PlainTextApi
}

func NewEncodingEncodingsMuxingsProgressiveTsId3Api(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsProgressiveTsId3Api, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsProgressiveTsId3Api{apiClient: apiClient}

    rawApi, err := NewEncodingEncodingsMuxingsProgressiveTsId3RawApi(configs...)
    api.Raw = rawApi
    frameIdApi, err := NewEncodingEncodingsMuxingsProgressiveTsId3FrameIdApi(configs...)
    api.FrameId = frameIdApi
    plainTextApi, err := NewEncodingEncodingsMuxingsProgressiveTsId3PlainTextApi(configs...)
    api.PlainText = plainTextApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsProgressiveTsId3Api) List(encodingId string, muxingId string, queryParams ...func(*query.Id3TagListQueryParams)) (*pagination.Id3TagsListPagination, error) {
    queryParameters := &query.Id3TagListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.Id3TagsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3", nil, &responseModel, reqParams)
    return responseModel, err
}

