package analytics
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type AnalyticsOutputsS3RoleBasedCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewAnalyticsOutputsS3RoleBasedCustomdataApi(configs ...func(*common.ApiClient)) (*AnalyticsOutputsS3RoleBasedCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AnalyticsOutputsS3RoleBasedCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AnalyticsOutputsS3RoleBasedCustomdataApi) Get(outputId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/analytics/outputs/s3-role-based/{output_id}/customData", nil, &responseModel, reqParams)
    return responseModel, err
}

