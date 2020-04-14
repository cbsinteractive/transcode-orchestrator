package notifications
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type NotificationsStatesApi struct {
    apiClient *common.ApiClient
}

func NewNotificationsStatesApi(configs ...func(*common.ApiClient)) (*NotificationsStatesApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &NotificationsStatesApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *NotificationsStatesApi) List(notificationId string, resourceId string, queryParams ...func(*query.NotificationStateEntryListQueryParams)) (*pagination.NotificationStateEntrysListPagination, error) {
    queryParameters := &query.NotificationStateEntryListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["notification_id"] = notificationId
        params.PathParams["resource_id"] = resourceId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.NotificationStateEntrysListPagination
    err := api.apiClient.Get("/notifications/{notification_id}/states/{resource_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

