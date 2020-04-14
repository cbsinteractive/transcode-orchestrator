package analytics
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type AnalyticsOutputsApi struct {
    apiClient *common.ApiClient
    S3RoleBased *AnalyticsOutputsS3RoleBasedApi
}

func NewAnalyticsOutputsApi(configs ...func(*common.ApiClient)) (*AnalyticsOutputsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AnalyticsOutputsApi{apiClient: apiClient}

    s3RoleBasedApi, err := NewAnalyticsOutputsS3RoleBasedApi(configs...)
    api.S3RoleBased = s3RoleBasedApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

