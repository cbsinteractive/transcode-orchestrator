package notifications
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type NotificationsWebhooksEncodingEncodingsErrorApi struct {
    apiClient *common.ApiClient
    Customdata *NotificationsWebhooksEncodingEncodingsErrorCustomdataApi
}

func NewNotificationsWebhooksEncodingEncodingsErrorApi(configs ...func(*common.ApiClient)) (*NotificationsWebhooksEncodingEncodingsErrorApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &NotificationsWebhooksEncodingEncodingsErrorApi{apiClient: apiClient}

    customdataApi, err := NewNotificationsWebhooksEncodingEncodingsErrorCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *NotificationsWebhooksEncodingEncodingsErrorApi) Create(webhook model.Webhook) (*model.Webhook, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.Webhook
    err := api.apiClient.Post("/notifications/webhooks/encoding/encodings/error", &webhook, &responseModel, reqParams)
    return responseModel, err
}

func (api *NotificationsWebhooksEncodingEncodingsErrorApi) CreateByEncodingId(encodingId string, webhook model.Webhook) (*model.Webhook, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.Webhook
    err := api.apiClient.Post("/notifications/webhooks/encoding/encodings/{encoding_id}/error", &webhook, &responseModel, reqParams)
    return responseModel, err
}

func (api *NotificationsWebhooksEncodingEncodingsErrorApi) DeleteByEncodingIdAndWebhookId(encodingId string, webhookId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["webhook_id"] = webhookId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/notifications/webhooks/encoding/encodings/{encoding_id}/error/{webhook_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *NotificationsWebhooksEncodingEncodingsErrorApi) DeleteByWebhookId(webhookId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["webhook_id"] = webhookId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/notifications/webhooks/encoding/encodings/error/{webhook_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *NotificationsWebhooksEncodingEncodingsErrorApi) GetByEncodingIdAndWebhookId(encodingId string, webhookId string) (*model.Webhook, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["webhook_id"] = webhookId
    }

    var responseModel *model.Webhook
    err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/{encoding_id}/error/{webhook_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *NotificationsWebhooksEncodingEncodingsErrorApi) GetByWebhookId(webhookId string) (*model.Webhook, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["webhook_id"] = webhookId
    }

    var responseModel *model.Webhook
    err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/error/{webhook_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *NotificationsWebhooksEncodingEncodingsErrorApi) List(queryParams ...func(*query.WebhookListQueryParams)) (*pagination.WebhooksListPagination, error) {
    queryParameters := &query.WebhookListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.WebhooksListPagination
    err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/error", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *NotificationsWebhooksEncodingEncodingsErrorApi) ListByEncodingId(encodingId string, queryParams ...func(*query.WebhookListByEncodingIdQueryParams)) (*pagination.WebhooksListByEncodingIdPagination, error) {
    queryParameters := &query.WebhookListByEncodingIdQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.WebhooksListByEncodingIdPagination
    err := api.apiClient.Get("/notifications/webhooks/encoding/encodings/{encoding_id}/error", nil, &responseModel, reqParams)
    return responseModel, err
}

