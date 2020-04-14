package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingInputsS3RoleBasedCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingInputsS3RoleBasedCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingInputsS3RoleBasedCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInputsS3RoleBasedCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInputsS3RoleBasedCustomdataApi) Get(inputId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/inputs/s3-role-based/{input_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

