package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingOutputsApi struct {
    apiClient *common.ApiClient
    Type *EncodingOutputsTypeApi
    S3 *EncodingOutputsS3Api
    S3RoleBased *EncodingOutputsS3RoleBasedApi
    GenericS3 *EncodingOutputsGenericS3Api
    Local *EncodingOutputsLocalApi
    Gcs *EncodingOutputsGcsApi
    Azure *EncodingOutputsAzureApi
    Ftp *EncodingOutputsFtpApi
    Sftp *EncodingOutputsSftpApi
    AkamaiMsl *EncodingOutputsAkamaiMslApi
    AkamaiNetstorage *EncodingOutputsAkamaiNetstorageApi
    LiveMediaIngest *EncodingOutputsLiveMediaIngestApi
}

func NewEncodingOutputsApi(configs ...func(*common.ApiClient)) (*EncodingOutputsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingOutputsApi{apiClient: apiClient}

    typeApi, err := NewEncodingOutputsTypeApi(configs...)
    api.Type = typeApi
    s3Api, err := NewEncodingOutputsS3Api(configs...)
    api.S3 = s3Api
    s3RoleBasedApi, err := NewEncodingOutputsS3RoleBasedApi(configs...)
    api.S3RoleBased = s3RoleBasedApi
    genericS3Api, err := NewEncodingOutputsGenericS3Api(configs...)
    api.GenericS3 = genericS3Api
    localApi, err := NewEncodingOutputsLocalApi(configs...)
    api.Local = localApi
    gcsApi, err := NewEncodingOutputsGcsApi(configs...)
    api.Gcs = gcsApi
    azureApi, err := NewEncodingOutputsAzureApi(configs...)
    api.Azure = azureApi
    ftpApi, err := NewEncodingOutputsFtpApi(configs...)
    api.Ftp = ftpApi
    sftpApi, err := NewEncodingOutputsSftpApi(configs...)
    api.Sftp = sftpApi
    akamaiMslApi, err := NewEncodingOutputsAkamaiMslApi(configs...)
    api.AkamaiMsl = akamaiMslApi
    akamaiNetstorageApi, err := NewEncodingOutputsAkamaiNetstorageApi(configs...)
    api.AkamaiNetstorage = akamaiNetstorageApi
    liveMediaIngestApi, err := NewEncodingOutputsLiveMediaIngestApi(configs...)
    api.LiveMediaIngest = liveMediaIngestApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingOutputsApi) List(queryParams ...func(*query.OutputListQueryParams)) (*pagination.OutputsListPagination, error) {
    queryParameters := &query.OutputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.OutputsListPagination
    err := api.apiClient.Get("/encoding/outputs", nil, &responseModel, reqParams)
    return responseModel, err
}

