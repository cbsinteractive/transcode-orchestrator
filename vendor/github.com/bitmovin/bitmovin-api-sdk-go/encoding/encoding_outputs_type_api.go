package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingOutputsTypeApi struct {
    apiClient *common.ApiClient
}

func NewEncodingOutputsTypeApi(configs ...func(*common.ApiClient)) (*EncodingOutputsTypeApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingOutputsTypeApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingOutputsTypeApi) Get(outputId string) (*model.OutputTypeResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel *model.OutputTypeResponse
    err := api.apiClient.Get("/encoding/outputs/{output_id}/type", nil, &responseModel, reqParams)
    return responseModel, err
}

