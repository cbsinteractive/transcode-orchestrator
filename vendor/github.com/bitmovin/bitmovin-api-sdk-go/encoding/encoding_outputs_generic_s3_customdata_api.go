package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingOutputsGenericS3CustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingOutputsGenericS3CustomdataApi(configs ...func(*common.ApiClient)) (*EncodingOutputsGenericS3CustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingOutputsGenericS3CustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingOutputsGenericS3CustomdataApi) Get(outputId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/outputs/generic-s3/{output_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

