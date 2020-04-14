package notifications
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type NotificationsEmailsEncodingEncodingsLiveInputStreamChangedApi struct {
    apiClient *common.ApiClient
}

func NewNotificationsEmailsEncodingEncodingsLiveInputStreamChangedApi(configs ...func(*common.ApiClient)) (*NotificationsEmailsEncodingEncodingsLiveInputStreamChangedApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &NotificationsEmailsEncodingEncodingsLiveInputStreamChangedApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *NotificationsEmailsEncodingEncodingsLiveInputStreamChangedApi) Create(emailNotificationWithStreamConditionsRequest model.EmailNotificationWithStreamConditionsRequest) (*model.EmailNotificationWithStreamConditions, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.EmailNotificationWithStreamConditions
    err := api.apiClient.Post("/notifications/emails/encoding/encodings/live-input-stream-changed", &emailNotificationWithStreamConditionsRequest, &responseModel, reqParams)
    return responseModel, err
}

func (api *NotificationsEmailsEncodingEncodingsLiveInputStreamChangedApi) CreateByEncodingId(encodingId string, emailNotificationWithStreamConditionsRequest model.EmailNotificationWithStreamConditionsRequest) (*model.EmailNotificationWithStreamConditions, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.EmailNotificationWithStreamConditions
    err := api.apiClient.Post("/notifications/emails/encoding/encodings/{encoding_id}/live-input-stream-changed", &emailNotificationWithStreamConditionsRequest, &responseModel, reqParams)
    return responseModel, err
}

func (api *NotificationsEmailsEncodingEncodingsLiveInputStreamChangedApi) Update(notificationId string, emailNotificationWithStreamConditionsRequest model.EmailNotificationWithStreamConditionsRequest) (*model.EmailNotificationWithStreamConditions, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["notification_id"] = notificationId
    }

    var responseModel *model.EmailNotificationWithStreamConditions
    err := api.apiClient.Put("/notifications/emails/encoding/encodings/live-input-stream-changed/{notification_id}", &emailNotificationWithStreamConditionsRequest, &responseModel, reqParams)
    return responseModel, err
}

