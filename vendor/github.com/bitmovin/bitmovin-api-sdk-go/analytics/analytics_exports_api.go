package analytics
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type AnalyticsExportsApi struct {
    apiClient *common.ApiClient
}

func NewAnalyticsExportsApi(configs ...func(*common.ApiClient)) (*AnalyticsExportsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AnalyticsExportsApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AnalyticsExportsApi) Create(analyticsExportTask model.AnalyticsExportTask) (*model.AnalyticsExportTask, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.AnalyticsExportTask
    err := api.apiClient.Post("/analytics/exports", &analyticsExportTask, &responseModel, reqParams)
    return responseModel, err
}

func (api *AnalyticsExportsApi) Get(exportTaskId string) (*model.AnalyticsExportTask, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["export_task_id"] = exportTaskId
    }

    var responseModel *model.AnalyticsExportTask
    err := api.apiClient.Get("/analytics/exports/{export_task_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *AnalyticsExportsApi) List(queryParams ...func(*query.AnalyticsExportTaskListQueryParams)) (*pagination.AnalyticsExportTasksListPagination, error) {
    queryParameters := &query.AnalyticsExportTaskListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.AnalyticsExportTasksListPagination
    err := api.apiClient.Get("/analytics/exports", nil, &responseModel, reqParams)
    return responseModel, err
}

