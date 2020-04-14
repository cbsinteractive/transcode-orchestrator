package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingOutputsLiveMediaIngestCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingOutputsLiveMediaIngestCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingOutputsLiveMediaIngestCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingOutputsLiveMediaIngestCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingOutputsLiveMediaIngestCustomdataApi) Get(outputId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/outputs/live-media-ingest/{output_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

