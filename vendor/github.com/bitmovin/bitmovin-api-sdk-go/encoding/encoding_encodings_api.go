package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsApi struct {
    apiClient *common.ApiClient
    Live *EncodingEncodingsLiveApi
    MachineLearning *EncodingEncodingsMachineLearningApi
    Customdata *EncodingEncodingsCustomdataApi
    Streams *EncodingEncodingsStreamsApi
    InputStreams *EncodingEncodingsInputStreamsApi
    Muxings *EncodingEncodingsMuxingsApi
    Captions *EncodingEncodingsCaptionsApi
    Sidecars *EncodingEncodingsSidecarsApi
    Keyframes *EncodingEncodingsKeyframesApi
}

func NewEncodingEncodingsApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsApi{apiClient: apiClient}

    liveApi, err := NewEncodingEncodingsLiveApi(configs...)
    api.Live = liveApi
    machineLearningApi, err := NewEncodingEncodingsMachineLearningApi(configs...)
    api.MachineLearning = machineLearningApi
    customdataApi, err := NewEncodingEncodingsCustomdataApi(configs...)
    api.Customdata = customdataApi
    streamsApi, err := NewEncodingEncodingsStreamsApi(configs...)
    api.Streams = streamsApi
    inputStreamsApi, err := NewEncodingEncodingsInputStreamsApi(configs...)
    api.InputStreams = inputStreamsApi
    muxingsApi, err := NewEncodingEncodingsMuxingsApi(configs...)
    api.Muxings = muxingsApi
    captionsApi, err := NewEncodingEncodingsCaptionsApi(configs...)
    api.Captions = captionsApi
    sidecarsApi, err := NewEncodingEncodingsSidecarsApi(configs...)
    api.Sidecars = sidecarsApi
    keyframesApi, err := NewEncodingEncodingsKeyframesApi(configs...)
    api.Keyframes = keyframesApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsApi) Create(encoding model.Encoding) (*model.Encoding, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.Encoding
    err := api.apiClient.Post("/encoding/encodings", &encoding, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsApi) Delete(encodingId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsApi) Get(encodingId string) (*model.Encoding, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.Encoding
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsApi) GetStartRequest(encodingId string) (*model.StartEncodingRequest, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.StartEncodingRequest
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/start", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsApi) List(queryParams ...func(*query.EncodingListQueryParams)) (*pagination.EncodingsListPagination, error) {
    queryParameters := &query.EncodingListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.EncodingsListPagination
    err := api.apiClient.Get("/encoding/encodings", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsApi) Reprioritize(encodingId string, reprioritizeEncodingRequest model.ReprioritizeEncodingRequest) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/reprioritize", &reprioritizeEncodingRequest, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsApi) Reschedule(encodingId string, rescheduleEncodingRequest model.RescheduleEncodingRequest) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/reschedule", &rescheduleEncodingRequest, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsApi) Start(encodingId string, startEncodingRequest model.StartEncodingRequest) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/start", &startEncodingRequest, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsApi) Status(encodingId string) (*model.ModelTask, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.ModelTask
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/status", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsApi) Stop(encodingId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/stop", nil, &responseModel, reqParams)
    return responseModel, err
}

