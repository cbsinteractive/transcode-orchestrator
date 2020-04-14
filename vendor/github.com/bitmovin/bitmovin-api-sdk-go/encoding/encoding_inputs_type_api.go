package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingInputsTypeApi struct {
    apiClient *common.ApiClient
}

func NewEncodingInputsTypeApi(configs ...func(*common.ApiClient)) (*EncodingInputsTypeApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInputsTypeApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInputsTypeApi) Get(inputId string) (*model.InputTypeResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel *model.InputTypeResponse
    err := api.apiClient.Get("/encoding/inputs/{input_id}/type", nil, &responseModel, reqParams)
    return responseModel, err
}

