package notifications
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type NotificationsEmailsUsageReportsApi struct {
    apiClient *common.ApiClient
}

func NewNotificationsEmailsUsageReportsApi(configs ...func(*common.ApiClient)) (*NotificationsEmailsUsageReportsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &NotificationsEmailsUsageReportsApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *NotificationsEmailsUsageReportsApi) List(queryParams ...func(*query.EmailNotificationListQueryParams)) (*pagination.EmailNotificationsListPagination, error) {
    queryParameters := &query.EmailNotificationListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.EmailNotificationsListPagination
    err := api.apiClient.Get("/notifications/emails/usage-reports", nil, &responseModel, reqParams)
    return responseModel, err
}

