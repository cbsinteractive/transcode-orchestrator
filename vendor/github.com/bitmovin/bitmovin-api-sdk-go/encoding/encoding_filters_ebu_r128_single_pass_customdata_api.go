package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingFiltersEbuR128SinglePassCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingFiltersEbuR128SinglePassCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingFiltersEbuR128SinglePassCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingFiltersEbuR128SinglePassCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingFiltersEbuR128SinglePassCustomdataApi) Get(filterId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/filters/ebu-r128-single-pass/{filter_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

