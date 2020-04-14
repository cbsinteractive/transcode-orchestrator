package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingFiltersWatermarkCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingFiltersWatermarkCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingFiltersWatermarkCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingFiltersWatermarkCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingFiltersWatermarkCustomdataApi) Get(filterId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/filters/watermark/{filter_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

