package account
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type AccountOrganizationsGroupsInvitationsApi struct {
    apiClient *common.ApiClient
}

func NewAccountOrganizationsGroupsInvitationsApi(configs ...func(*common.ApiClient)) (*AccountOrganizationsGroupsInvitationsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AccountOrganizationsGroupsInvitationsApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AccountOrganizationsGroupsInvitationsApi) Create(organizationId string, groupId string, invitation model.Invitation) (*model.Invitation, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["organization_id"] = organizationId
        params.PathParams["group_id"] = groupId
    }

    var responseModel *model.Invitation
    err := api.apiClient.Post("/account/organizations/{organization_id}/groups/{group_id}/invitations", &invitation, &responseModel, reqParams)
    return responseModel, err
}

func (api *AccountOrganizationsGroupsInvitationsApi) List(organizationId string, groupId string) (*pagination.InvitationsListPagination, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["organization_id"] = organizationId
        params.PathParams["group_id"] = groupId
    }

    var responseModel *pagination.InvitationsListPagination
    err := api.apiClient.Get("/account/organizations/{organization_id}/groups/{group_id}/invitations", nil, &responseModel, reqParams)
    return responseModel, err
}

