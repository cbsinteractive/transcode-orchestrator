package account
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type AccountOrganizationsGroupsApi struct {
    apiClient *common.ApiClient
    Tenants *AccountOrganizationsGroupsTenantsApi
    Invitations *AccountOrganizationsGroupsInvitationsApi
    Permissions *AccountOrganizationsGroupsPermissionsApi
}

func NewAccountOrganizationsGroupsApi(configs ...func(*common.ApiClient)) (*AccountOrganizationsGroupsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AccountOrganizationsGroupsApi{apiClient: apiClient}

    tenantsApi, err := NewAccountOrganizationsGroupsTenantsApi(configs...)
    api.Tenants = tenantsApi
    invitationsApi, err := NewAccountOrganizationsGroupsInvitationsApi(configs...)
    api.Invitations = invitationsApi
    permissionsApi, err := NewAccountOrganizationsGroupsPermissionsApi(configs...)
    api.Permissions = permissionsApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AccountOrganizationsGroupsApi) Create(organizationId string, group model.Group) (*model.Group, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["organization_id"] = organizationId
    }

    var responseModel *model.Group
    err := api.apiClient.Post("/account/organizations/{organization_id}/groups", &group, &responseModel, reqParams)
    return responseModel, err
}

func (api *AccountOrganizationsGroupsApi) Delete(organizationId string, groupId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["organization_id"] = organizationId
        params.PathParams["group_id"] = groupId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/account/organizations/{organization_id}/groups/{group_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *AccountOrganizationsGroupsApi) Get(organizationId string, groupId string) (*model.Group, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["organization_id"] = organizationId
        params.PathParams["group_id"] = groupId
    }

    var responseModel *model.Group
    err := api.apiClient.Get("/account/organizations/{organization_id}/groups/{group_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *AccountOrganizationsGroupsApi) List(organizationId string) (*pagination.GroupsListPagination, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["organization_id"] = organizationId
    }

    var responseModel *pagination.GroupsListPagination
    err := api.apiClient.Get("/account/organizations/{organization_id}/groups", nil, &responseModel, reqParams)
    return responseModel, err
}

