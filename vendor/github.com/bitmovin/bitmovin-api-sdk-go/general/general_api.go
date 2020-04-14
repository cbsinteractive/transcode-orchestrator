package general
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type GeneralApi struct {
    apiClient *common.ApiClient
    ErrorDefinitions *GeneralErrorDefinitionsApi
}

func NewGeneralApi(configs ...func(*common.ApiClient)) (*GeneralApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &GeneralApi{apiClient: apiClient}

    errorDefinitionsApi, err := NewGeneralErrorDefinitionsApi(configs...)
    api.ErrorDefinitions = errorDefinitionsApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

