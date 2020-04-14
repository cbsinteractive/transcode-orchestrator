package notifications
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type NotificationsWebhooksEncodingEncodingsErrorCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewNotificationsWebhooksEncodingEncodingsErrorCustomdataApi(configs ...func(*common.ApiClient)) (*NotificationsWebhooksEncodingEncodingsErrorCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &NotificationsWebhooksEncodingEncodingsErrorCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *NotificationsWebhooksEncodingEncodingsErrorCustomdataApi) GetCustomDataByEncodingIdAndWebhookId(encodingId string, webhookId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["webhook_id"] = webhookId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/{encoding_id}/error/{webhook_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *NotificationsWebhooksEncodingEncodingsErrorCustomdataApi) GetCustomDataByWebhookId(webhookId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["webhook_id"] = webhookId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/error/{webhook_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

