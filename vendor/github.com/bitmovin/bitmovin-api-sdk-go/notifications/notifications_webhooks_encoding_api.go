package notifications
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type NotificationsWebhooksEncodingApi struct {
    apiClient *common.ApiClient
    Encodings *NotificationsWebhooksEncodingEncodingsApi
    Manifest *NotificationsWebhooksEncodingManifestApi
}

func NewNotificationsWebhooksEncodingApi(configs ...func(*common.ApiClient)) (*NotificationsWebhooksEncodingApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &NotificationsWebhooksEncodingApi{apiClient: apiClient}

    encodingsApi, err := NewNotificationsWebhooksEncodingEncodingsApi(configs...)
    api.Encodings = encodingsApi
    manifestApi, err := NewNotificationsWebhooksEncodingManifestApi(configs...)
    api.Manifest = manifestApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

