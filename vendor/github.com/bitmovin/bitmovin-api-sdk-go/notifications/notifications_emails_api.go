package notifications
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type NotificationsEmailsApi struct {
    apiClient *common.ApiClient
    UsageReports *NotificationsEmailsUsageReportsApi
    Encoding *NotificationsEmailsEncodingApi
}

func NewNotificationsEmailsApi(configs ...func(*common.ApiClient)) (*NotificationsEmailsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &NotificationsEmailsApi{apiClient: apiClient}

    usageReportsApi, err := NewNotificationsEmailsUsageReportsApi(configs...)
    api.UsageReports = usageReportsApi
    encodingApi, err := NewNotificationsEmailsEncodingApi(configs...)
    api.Encoding = encodingApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *NotificationsEmailsApi) List(queryParams ...func(*query.NotificationListQueryParams)) (*pagination.NotificationsListPagination, error) {
    queryParameters := &query.NotificationListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.NotificationsListPagination
    err := api.apiClient.Get("/notifications/emails", nil, &responseModel, reqParams)
    return responseModel, err
}

