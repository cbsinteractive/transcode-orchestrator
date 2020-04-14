package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingFiltersTextCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingFiltersTextCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingFiltersTextCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingFiltersTextCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingFiltersTextCustomdataApi) Get(filterId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/filters/text/{filter_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

