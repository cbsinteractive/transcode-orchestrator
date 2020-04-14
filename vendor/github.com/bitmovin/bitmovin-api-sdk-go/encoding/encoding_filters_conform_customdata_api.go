package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingFiltersConformCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingFiltersConformCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingFiltersConformCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingFiltersConformCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingFiltersConformCustomdataApi) Get(filterId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/filters/conform/{filter_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

