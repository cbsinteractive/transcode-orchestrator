package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingManifestsHlsMediaTypeApi struct {
    apiClient *common.ApiClient
}

func NewEncodingManifestsHlsMediaTypeApi(configs ...func(*common.ApiClient)) (*EncodingManifestsHlsMediaTypeApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsHlsMediaTypeApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsHlsMediaTypeApi) Get(manifestId string, mediaId string) (*model.MediaInfoTypeResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["media_id"] = mediaId
    }

    var responseModel *model.MediaInfoTypeResponse
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/media/{media_id}/type", nil, &responseModel, reqParams)
    return responseModel, err
}

