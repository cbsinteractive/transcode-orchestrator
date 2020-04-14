package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingInputsGcsCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingInputsGcsCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingInputsGcsCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInputsGcsCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInputsGcsCustomdataApi) Get(inputId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/inputs/gcs/{input_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

