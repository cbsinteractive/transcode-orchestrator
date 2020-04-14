package notifications
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type NotificationsWebhooksApi struct {
    apiClient *common.ApiClient
    Encoding *NotificationsWebhooksEncodingApi
}

func NewNotificationsWebhooksApi(configs ...func(*common.ApiClient)) (*NotificationsWebhooksApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &NotificationsWebhooksApi{apiClient: apiClient}

    encodingApi, err := NewNotificationsWebhooksEncodingApi(configs...)
    api.Encoding = encodingApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

