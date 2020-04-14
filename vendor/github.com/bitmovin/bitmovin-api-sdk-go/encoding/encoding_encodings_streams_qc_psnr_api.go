package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsStreamsQcPsnrApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsStreamsQcPsnrApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsStreamsQcPsnrApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsStreamsQcPsnrApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsStreamsQcPsnrApi) Create(encodingId string, streamId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/streams/{stream_id}/qc/psnr", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsStreamsQcPsnrApi) List(encodingId string, streamId string, queryParams ...func(*query.PsnrQualityMetricListQueryParams)) (*pagination.PsnrQualityMetricsListPagination, error) {
    queryParameters := &query.PsnrQualityMetricListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.PsnrQualityMetricsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/qc/psnr", nil, &responseModel, reqParams)
    return responseModel, err
}

