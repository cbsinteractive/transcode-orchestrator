package notifications
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type NotificationsApi struct {
    apiClient *common.ApiClient
    Webhooks *NotificationsWebhooksApi
    States *NotificationsStatesApi
    Emails *NotificationsEmailsApi
}

func NewNotificationsApi(configs ...func(*common.ApiClient)) (*NotificationsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &NotificationsApi{apiClient: apiClient}

    webhooksApi, err := NewNotificationsWebhooksApi(configs...)
    api.Webhooks = webhooksApi
    statesApi, err := NewNotificationsStatesApi(configs...)
    api.States = statesApi
    emailsApi, err := NewNotificationsEmailsApi(configs...)
    api.Emails = emailsApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *NotificationsApi) Delete(notificationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["notification_id"] = notificationId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/notifications/{notification_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *NotificationsApi) Get(notificationId string) (*model.Notification, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["notification_id"] = notificationId
    }

    var responseModel *model.Notification
    err := api.apiClient.Get("/notifications/{notification_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *NotificationsApi) List(queryParams ...func(*query.NotificationListQueryParams)) (*pagination.NotificationsListPagination, error) {
    queryParameters := &query.NotificationListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.NotificationsListPagination
    err := api.apiClient.Get("/notifications", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *NotificationsApi) ListByNotificationId(notificationId string, queryParams ...func(*query.NotificationStateEntryListByNotificationIdQueryParams)) (*pagination.NotificationStateEntrysListByNotificationIdPagination, error) {
    queryParameters := &query.NotificationStateEntryListByNotificationIdQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["notification_id"] = notificationId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.NotificationStateEntrysListByNotificationIdPagination
    err := api.apiClient.Get("/notifications/{notification_id}/states", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *NotificationsApi) Mute(notificationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["notification_id"] = notificationId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Post("/notifications/{notification_id}/mute", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *NotificationsApi) Unmute(notificationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["notification_id"] = notificationId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Post("/notifications/{notification_id}/unmute", nil, &responseModel, reqParams)
    return responseModel, err
}

