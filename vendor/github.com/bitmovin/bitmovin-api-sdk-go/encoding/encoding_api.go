package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type EncodingApi struct {
    apiClient *common.ApiClient
    Inputs *EncodingInputsApi
    Outputs *EncodingOutputsApi
    Configurations *EncodingConfigurationsApi
    Filters *EncodingFiltersApi
    Encodings *EncodingEncodingsApi
    Manifests *EncodingManifestsApi
    Infrastructure *EncodingInfrastructureApi
    Statistics *EncodingStatisticsApi
    ErrorDefinitions *EncodingErrorDefinitionsApi
}

func NewEncodingApi(configs ...func(*common.ApiClient)) (*EncodingApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingApi{apiClient: apiClient}

    inputsApi, err := NewEncodingInputsApi(configs...)
    api.Inputs = inputsApi
    outputsApi, err := NewEncodingOutputsApi(configs...)
    api.Outputs = outputsApi
    configurationsApi, err := NewEncodingConfigurationsApi(configs...)
    api.Configurations = configurationsApi
    filtersApi, err := NewEncodingFiltersApi(configs...)
    api.Filters = filtersApi
    encodingsApi, err := NewEncodingEncodingsApi(configs...)
    api.Encodings = encodingsApi
    manifestsApi, err := NewEncodingManifestsApi(configs...)
    api.Manifests = manifestsApi
    infrastructureApi, err := NewEncodingInfrastructureApi(configs...)
    api.Infrastructure = infrastructureApi
    statisticsApi, err := NewEncodingStatisticsApi(configs...)
    api.Statistics = statisticsApi
    errorDefinitionsApi, err := NewEncodingErrorDefinitionsApi(configs...)
    api.ErrorDefinitions = errorDefinitionsApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

