package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingFiltersUnsharpCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingFiltersUnsharpCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingFiltersUnsharpCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingFiltersUnsharpCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingFiltersUnsharpCustomdataApi) Get(filterId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/filters/unsharp/{filter_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

