package notifications
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type NotificationsWebhooksEncodingManifestErrorApi struct {
    apiClient *common.ApiClient
}

func NewNotificationsWebhooksEncodingManifestErrorApi(configs ...func(*common.ApiClient)) (*NotificationsWebhooksEncodingManifestErrorApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &NotificationsWebhooksEncodingManifestErrorApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *NotificationsWebhooksEncodingManifestErrorApi) Create(webhook model.Webhook) (*pagination.WebhooksCreatePagination, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *pagination.WebhooksCreatePagination
    err := api.apiClient.Post("/notifications/webhooks/encoding/manifest/error", &webhook, &responseModel, reqParams)
    return responseModel, err
}

func (api *NotificationsWebhooksEncodingManifestErrorApi) CreateByManifestId(manifestId string, webhook model.Webhook) (*model.Webhook, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }

    var responseModel *model.Webhook
    err := api.apiClient.Post("/notifications/webhooks/encoding/manifest/{manifest_id}/error", &webhook, &responseModel, reqParams)
    return responseModel, err
}

func (api *NotificationsWebhooksEncodingManifestErrorApi) Update(notificationId string, webhook model.Webhook) (*model.Webhook, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["notification_id"] = notificationId
    }

    var responseModel *model.Webhook
    err := api.apiClient.Put("/notifications/webhooks/encoding/manifest/error/{notification_id}", &webhook, &responseModel, reqParams)
    return responseModel, err
}

