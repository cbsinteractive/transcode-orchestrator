package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingFiltersScaleCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingFiltersScaleCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingFiltersScaleCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingFiltersScaleCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingFiltersScaleCustomdataApi) Get(filterId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/filters/scale/{filter_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

