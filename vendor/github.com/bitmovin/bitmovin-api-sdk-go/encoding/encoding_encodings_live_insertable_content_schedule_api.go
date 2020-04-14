package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsLiveInsertableContentScheduleApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsLiveInsertableContentScheduleApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsLiveInsertableContentScheduleApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsLiveInsertableContentScheduleApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsLiveInsertableContentScheduleApi) Create(encodingId string, contentId string, scheduledInsertableContent model.ScheduledInsertableContent) (*model.ScheduledInsertableContent, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["content_id"] = contentId
    }

    var responseModel *model.ScheduledInsertableContent
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/live/insertable-content/{content_id}/schedule", &scheduledInsertableContent, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsLiveInsertableContentScheduleApi) Delete(encodingId string, contentId string, scheduledContentId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["content_id"] = contentId
        params.PathParams["scheduled_content_id"] = scheduledContentId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/live/insertable-content/{content_id}/schedule/{scheduled_content_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

