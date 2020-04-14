package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMachineLearningObjectDetectionResultsByTimestampApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsMachineLearningObjectDetectionResultsByTimestampApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMachineLearningObjectDetectionResultsByTimestampApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMachineLearningObjectDetectionResultsByTimestampApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMachineLearningObjectDetectionResultsByTimestampApi) List(encodingId string, objectDetectionId string, queryParams ...func(*query.ObjectDetectionTimestampResultListQueryParams)) (*pagination.ObjectDetectionTimestampResultsListPagination, error) {
    queryParameters := &query.ObjectDetectionTimestampResultListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["object_detection_id"] = objectDetectionId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.ObjectDetectionTimestampResultsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/machine-learning/object-detection/{object_detection_id}/results/by-timestamp", nil, &responseModel, reqParams)
    return responseModel, err
}

