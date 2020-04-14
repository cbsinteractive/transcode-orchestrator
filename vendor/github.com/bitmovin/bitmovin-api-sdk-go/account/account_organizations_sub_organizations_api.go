package account
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type AccountOrganizationsSubOrganizationsApi struct {
    apiClient *common.ApiClient
}

func NewAccountOrganizationsSubOrganizationsApi(configs ...func(*common.ApiClient)) (*AccountOrganizationsSubOrganizationsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AccountOrganizationsSubOrganizationsApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AccountOrganizationsSubOrganizationsApi) List(organizationId string) (*pagination.OrganizationsListPagination, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["organization_id"] = organizationId
    }

    var responseModel *pagination.OrganizationsListPagination
    err := api.apiClient.Get("/account/organizations/{organization_id}/sub-organizations", nil, &responseModel, reqParams)
    return responseModel, err
}

