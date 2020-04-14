package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingManifestsDashPeriodsCustomXmlElementsApi struct {
    apiClient *common.ApiClient
}

func NewEncodingManifestsDashPeriodsCustomXmlElementsApi(configs ...func(*common.ApiClient)) (*EncodingManifestsDashPeriodsCustomXmlElementsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsDashPeriodsCustomXmlElementsApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsDashPeriodsCustomXmlElementsApi) Create(manifestId string, periodId string, customXmlElement model.CustomXmlElement) (*model.CustomXmlElement, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
    }

    var responseModel *model.CustomXmlElement
    err := api.apiClient.Post("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/custom-xml-elements", &customXmlElement, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsDashPeriodsCustomXmlElementsApi) Delete(manifestId string, periodId string, customXmlElementId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["custom_xml_element_id"] = customXmlElementId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/custom-xml-elements/{custom_xml_element_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsDashPeriodsCustomXmlElementsApi) Get(manifestId string, periodId string, customXmlElementId string) (*model.CustomXmlElement, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.PathParams["custom_xml_element_id"] = customXmlElementId
    }

    var responseModel *model.CustomXmlElement
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/custom-xml-elements/{custom_xml_element_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsDashPeriodsCustomXmlElementsApi) List(manifestId string, periodId string, queryParams ...func(*query.CustomXmlElementListQueryParams)) (*pagination.CustomXmlElementsListPagination, error) {
    queryParameters := &query.CustomXmlElementListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["period_id"] = periodId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.CustomXmlElementsListPagination
    err := api.apiClient.Get("/encoding/manifests/dash/{manifest_id}/periods/{period_id}/custom-xml-elements", nil, &responseModel, reqParams)
    return responseModel, err
}

