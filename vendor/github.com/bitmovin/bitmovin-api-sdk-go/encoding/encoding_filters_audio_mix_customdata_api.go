package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingFiltersAudioMixCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingFiltersAudioMixCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingFiltersAudioMixCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingFiltersAudioMixCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingFiltersAudioMixCustomdataApi) Get(filterId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/filters/audio-mix/{filter_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

