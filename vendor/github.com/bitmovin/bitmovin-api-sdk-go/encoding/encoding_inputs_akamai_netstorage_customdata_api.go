package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingInputsAkamaiNetstorageCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingInputsAkamaiNetstorageCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingInputsAkamaiNetstorageCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInputsAkamaiNetstorageCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInputsAkamaiNetstorageCustomdataApi) Get(inputId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/inputs/akamai-netstorage/{input_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

