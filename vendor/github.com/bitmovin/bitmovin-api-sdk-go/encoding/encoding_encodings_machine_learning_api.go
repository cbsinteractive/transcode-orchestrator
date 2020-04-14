package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type EncodingEncodingsMachineLearningApi struct {
    apiClient *common.ApiClient
    ObjectDetection *EncodingEncodingsMachineLearningObjectDetectionApi
}

func NewEncodingEncodingsMachineLearningApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMachineLearningApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMachineLearningApi{apiClient: apiClient}

    objectDetectionApi, err := NewEncodingEncodingsMachineLearningObjectDetectionApi(configs...)
    api.ObjectDetection = objectDetectionApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

