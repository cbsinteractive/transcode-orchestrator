package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsStreamsBurnInSubtitlesDvbsubApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsStreamsBurnInSubtitlesDvbsubApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsStreamsBurnInSubtitlesDvbsubApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsStreamsBurnInSubtitlesDvbsubApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsStreamsBurnInSubtitlesDvbsubApi) Create(encodingId string, streamId string, streamDvbSubSubtitle model.StreamDvbSubSubtitle) (*model.StreamDvbSubSubtitle, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
    }

    var responseModel *model.StreamDvbSubSubtitle
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/streams/{stream_id}/burn-in-subtitles/dvbsub", &streamDvbSubSubtitle, &responseModel, reqParams)
    return responseModel, err
}

