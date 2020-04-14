package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingOutputsLocalCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingOutputsLocalCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingOutputsLocalCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingOutputsLocalCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingOutputsLocalCustomdataApi) Get(outputId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/outputs/local/{output_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

