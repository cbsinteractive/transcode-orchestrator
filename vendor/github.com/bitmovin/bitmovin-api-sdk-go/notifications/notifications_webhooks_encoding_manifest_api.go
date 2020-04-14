package notifications
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type NotificationsWebhooksEncodingManifestApi struct {
    apiClient *common.ApiClient
    Error *NotificationsWebhooksEncodingManifestErrorApi
    Finished *NotificationsWebhooksEncodingManifestFinishedApi
}

func NewNotificationsWebhooksEncodingManifestApi(configs ...func(*common.ApiClient)) (*NotificationsWebhooksEncodingManifestApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &NotificationsWebhooksEncodingManifestApi{apiClient: apiClient}

    errorApi, err := NewNotificationsWebhooksEncodingManifestErrorApi(configs...)
    api.Error = errorApi
    finishedApi, err := NewNotificationsWebhooksEncodingManifestFinishedApi(configs...)
    api.Finished = finishedApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *NotificationsWebhooksEncodingManifestApi) List(manifestId string, queryParams ...func(*query.NotificationListQueryParams)) (*pagination.NotificationsListPagination, error) {
    queryParameters := &query.NotificationListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.NotificationsListPagination
    err := api.apiClient.Get("/notifications/webhooks/encoding/manifest/{manifest_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

