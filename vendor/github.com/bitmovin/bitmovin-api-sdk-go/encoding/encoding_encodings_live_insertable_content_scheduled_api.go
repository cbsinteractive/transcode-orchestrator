package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsLiveInsertableContentScheduledApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsLiveInsertableContentScheduledApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsLiveInsertableContentScheduledApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsLiveInsertableContentScheduledApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsLiveInsertableContentScheduledApi) List(encodingId string, queryParams ...func(*query.ScheduledInsertableContentListQueryParams)) (*pagination.ScheduledInsertableContentsListPagination, error) {
    queryParameters := &query.ScheduledInsertableContentListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.ScheduledInsertableContentsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/live/insertable-content/scheduled", nil, &responseModel, reqParams)
    return responseModel, err
}

