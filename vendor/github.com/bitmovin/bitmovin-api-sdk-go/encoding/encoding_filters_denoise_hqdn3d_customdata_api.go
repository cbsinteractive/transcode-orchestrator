package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingFiltersDenoiseHqdn3dCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingFiltersDenoiseHqdn3dCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingFiltersDenoiseHqdn3dCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingFiltersDenoiseHqdn3dCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingFiltersDenoiseHqdn3dCustomdataApi) Get(filterId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/filters/denoise-hqdn3d/{filter_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

