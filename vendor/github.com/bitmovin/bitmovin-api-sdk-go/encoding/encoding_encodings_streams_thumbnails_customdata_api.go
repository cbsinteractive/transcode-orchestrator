package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsStreamsThumbnailsCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsStreamsThumbnailsCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsStreamsThumbnailsCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsStreamsThumbnailsCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsStreamsThumbnailsCustomdataApi) Get(encodingId string, streamId string, thumbnailId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.PathParams["thumbnail_id"] = thumbnailId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/thumbnails/{thumbnail_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

