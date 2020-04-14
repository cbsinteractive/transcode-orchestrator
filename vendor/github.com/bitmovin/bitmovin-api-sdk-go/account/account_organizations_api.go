package account
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type AccountOrganizationsApi struct {
    apiClient *common.ApiClient
    SubOrganizations *AccountOrganizationsSubOrganizationsApi
    Groups *AccountOrganizationsGroupsApi
}

func NewAccountOrganizationsApi(configs ...func(*common.ApiClient)) (*AccountOrganizationsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AccountOrganizationsApi{apiClient: apiClient}

    subOrganizationsApi, err := NewAccountOrganizationsSubOrganizationsApi(configs...)
    api.SubOrganizations = subOrganizationsApi
    groupsApi, err := NewAccountOrganizationsGroupsApi(configs...)
    api.Groups = groupsApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AccountOrganizationsApi) Create(organization model.Organization) (*model.Organization, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.Organization
    err := api.apiClient.Post("/account/organizations", &organization, &responseModel, reqParams)
    return responseModel, err
}

func (api *AccountOrganizationsApi) Get(organizationId string) (*model.Organization, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["organization_id"] = organizationId
    }

    var responseModel *model.Organization
    err := api.apiClient.Get("/account/organizations/{organization_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *AccountOrganizationsApi) List() (*pagination.OrganizationsListPagination, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *pagination.OrganizationsListPagination
    err := api.apiClient.Get("/account/organizations", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *AccountOrganizationsApi) Update(organizationId string, updateOrganizationRequest model.UpdateOrganizationRequest) (*model.Organization, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["organization_id"] = organizationId
    }

    var responseModel *model.Organization
    err := api.apiClient.Put("/account/organizations/{organization_id}", &updateOrganizationRequest, &responseModel, reqParams)
    return responseModel, err
}

