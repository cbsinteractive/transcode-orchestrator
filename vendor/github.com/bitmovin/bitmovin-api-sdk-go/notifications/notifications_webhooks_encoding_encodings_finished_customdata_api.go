package notifications
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type NotificationsWebhooksEncodingEncodingsFinishedCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewNotificationsWebhooksEncodingEncodingsFinishedCustomdataApi(configs ...func(*common.ApiClient)) (*NotificationsWebhooksEncodingEncodingsFinishedCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &NotificationsWebhooksEncodingEncodingsFinishedCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *NotificationsWebhooksEncodingEncodingsFinishedCustomdataApi) GetCustomDataByEncodingIdAndWebhookId(encodingId string, webhookId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["webhook_id"] = webhookId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/{encoding_id}/finished/{webhook_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *NotificationsWebhooksEncodingEncodingsFinishedCustomdataApi) GetCustomDataByWebhookId(webhookId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["webhook_id"] = webhookId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/finished/{webhook_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

