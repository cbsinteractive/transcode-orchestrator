package bitmovin
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/account"
    "github.com/bitmovin/bitmovin-api-sdk-go/analytics"
    "github.com/bitmovin/bitmovin-api-sdk-go/encoding"
    "github.com/bitmovin/bitmovin-api-sdk-go/general"
    "github.com/bitmovin/bitmovin-api-sdk-go/notifications"
    "github.com/bitmovin/bitmovin-api-sdk-go/player"
)

type BitmovinApi struct {
    apiClient *common.ApiClient
    Account *account.AccountApi
    Analytics *analytics.AnalyticsApi
    Encoding *encoding.EncodingApi
    General *general.GeneralApi
    Notifications *notifications.NotificationsApi
    Player *player.PlayerApi
}

func NewBitmovinApi(configs ...func(*common.ApiClient)) (*BitmovinApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &BitmovinApi{apiClient: apiClient}

    accountApi, err := account.NewAccountApi(configs...)
    api.Account = accountApi
    analyticsApi, err := analytics.NewAnalyticsApi(configs...)
    api.Analytics = analyticsApi
    encodingApi, err := encoding.NewEncodingApi(configs...)
    api.Encoding = encodingApi
    generalApi, err := general.NewGeneralApi(configs...)
    api.General = generalApi
    notificationsApi, err := notifications.NewNotificationsApi(configs...)
    api.Notifications = notificationsApi
    playerApi, err := player.NewPlayerApi(configs...)
    api.Player = playerApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

