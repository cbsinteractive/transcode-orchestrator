package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingFiltersAudioVolumeCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingFiltersAudioVolumeCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingFiltersAudioVolumeCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingFiltersAudioVolumeCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingFiltersAudioVolumeCustomdataApi) Get(filterId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/filters/audio-volume/{filter_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

