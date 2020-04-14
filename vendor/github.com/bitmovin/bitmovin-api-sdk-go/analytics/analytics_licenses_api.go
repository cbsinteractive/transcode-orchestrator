package analytics
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type AnalyticsLicensesApi struct {
    apiClient *common.ApiClient
    Domains *AnalyticsLicensesDomainsApi
}

func NewAnalyticsLicensesApi(configs ...func(*common.ApiClient)) (*AnalyticsLicensesApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AnalyticsLicensesApi{apiClient: apiClient}

    domainsApi, err := NewAnalyticsLicensesDomainsApi(configs...)
    api.Domains = domainsApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AnalyticsLicensesApi) Create(analyticsLicense model.AnalyticsLicense) (*model.AnalyticsLicense, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.AnalyticsLicense
    err := api.apiClient.Post("/analytics/licenses", &analyticsLicense, &responseModel, reqParams)
    return responseModel, err
}

func (api *AnalyticsLicensesApi) Get(licenseId string) (*model.AnalyticsLicense, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["license_id"] = licenseId
    }

    var responseModel *model.AnalyticsLicense
    err := api.apiClient.Get("/analytics/licenses/{license_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *AnalyticsLicensesApi) List() (*pagination.AnalyticsLicensesListPagination, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *pagination.AnalyticsLicensesListPagination
    err := api.apiClient.Get("/analytics/licenses", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *AnalyticsLicensesApi) Update(licenseId string, analyticsLicenseUpdateRequest model.AnalyticsLicenseUpdateRequest) (*model.AnalyticsLicense, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["license_id"] = licenseId
    }

    var responseModel *model.AnalyticsLicense
    err := api.apiClient.Put("/analytics/licenses/{license_id}", &analyticsLicenseUpdateRequest, &responseModel, reqParams)
    return responseModel, err
}

