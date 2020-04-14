package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type EncodingEncodingsLiveInsertableContentStopApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsLiveInsertableContentStopApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsLiveInsertableContentStopApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsLiveInsertableContentStopApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsLiveInsertableContentStopApi) Create(encodingId string) (error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/live/insertable-content/stop", nil, nil, reqParams)
    return err
}

