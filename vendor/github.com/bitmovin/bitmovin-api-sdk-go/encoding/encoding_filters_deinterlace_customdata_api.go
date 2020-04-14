package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingFiltersDeinterlaceCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingFiltersDeinterlaceCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingFiltersDeinterlaceCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingFiltersDeinterlaceCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingFiltersDeinterlaceCustomdataApi) Get(filterId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/filters/deinterlace/{filter_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

