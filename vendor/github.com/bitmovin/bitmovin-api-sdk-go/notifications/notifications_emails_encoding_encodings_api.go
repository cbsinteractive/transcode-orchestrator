package notifications
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type NotificationsEmailsEncodingEncodingsApi struct {
    apiClient *common.ApiClient
    LiveInputStreamChanged *NotificationsEmailsEncodingEncodingsLiveInputStreamChangedApi
    Error *NotificationsEmailsEncodingEncodingsErrorApi
}

func NewNotificationsEmailsEncodingEncodingsApi(configs ...func(*common.ApiClient)) (*NotificationsEmailsEncodingEncodingsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &NotificationsEmailsEncodingEncodingsApi{apiClient: apiClient}

    liveInputStreamChangedApi, err := NewNotificationsEmailsEncodingEncodingsLiveInputStreamChangedApi(configs...)
    api.LiveInputStreamChanged = liveInputStreamChangedApi
    errorApi, err := NewNotificationsEmailsEncodingEncodingsErrorApi(configs...)
    api.Error = errorApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *NotificationsEmailsEncodingEncodingsApi) List(encodingId string, queryParams ...func(*query.EmailNotificationWithStreamConditionsListQueryParams)) (*pagination.EmailNotificationWithStreamConditionssListPagination, error) {
    queryParameters := &query.EmailNotificationWithStreamConditionsListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.EmailNotificationWithStreamConditionssListPagination
    err := api.apiClient.Get("/notifications/emails/encoding/encodings/{encoding_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

