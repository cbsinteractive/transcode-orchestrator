package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingFiltersInterlaceCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingFiltersInterlaceCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingFiltersInterlaceCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingFiltersInterlaceCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingFiltersInterlaceCustomdataApi) Get(filterId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/filters/interlace/{filter_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

