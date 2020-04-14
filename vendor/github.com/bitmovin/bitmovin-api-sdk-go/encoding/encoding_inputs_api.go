package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingInputsApi struct {
    apiClient *common.ApiClient
    Type *EncodingInputsTypeApi
    Rtmp *EncodingInputsRtmpApi
    RedundantRtmp *EncodingInputsRedundantRtmpApi
    S3 *EncodingInputsS3Api
    S3RoleBased *EncodingInputsS3RoleBasedApi
    GenericS3 *EncodingInputsGenericS3Api
    Local *EncodingInputsLocalApi
    Gcs *EncodingInputsGcsApi
    Azure *EncodingInputsAzureApi
    Ftp *EncodingInputsFtpApi
    Sftp *EncodingInputsSftpApi
    Http *EncodingInputsHttpApi
    Https *EncodingInputsHttpsApi
    Aspera *EncodingInputsAsperaApi
    AkamaiNetstorage *EncodingInputsAkamaiNetstorageApi
    Srt *EncodingInputsSrtApi
    Tcp *EncodingInputsTcpApi
    Udp *EncodingInputsUdpApi
    UdpMulticast *EncodingInputsUdpMulticastApi
    Zixi *EncodingInputsZixiApi
}

func NewEncodingInputsApi(configs ...func(*common.ApiClient)) (*EncodingInputsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInputsApi{apiClient: apiClient}

    typeApi, err := NewEncodingInputsTypeApi(configs...)
    api.Type = typeApi
    rtmpApi, err := NewEncodingInputsRtmpApi(configs...)
    api.Rtmp = rtmpApi
    redundantRtmpApi, err := NewEncodingInputsRedundantRtmpApi(configs...)
    api.RedundantRtmp = redundantRtmpApi
    s3Api, err := NewEncodingInputsS3Api(configs...)
    api.S3 = s3Api
    s3RoleBasedApi, err := NewEncodingInputsS3RoleBasedApi(configs...)
    api.S3RoleBased = s3RoleBasedApi
    genericS3Api, err := NewEncodingInputsGenericS3Api(configs...)
    api.GenericS3 = genericS3Api
    localApi, err := NewEncodingInputsLocalApi(configs...)
    api.Local = localApi
    gcsApi, err := NewEncodingInputsGcsApi(configs...)
    api.Gcs = gcsApi
    azureApi, err := NewEncodingInputsAzureApi(configs...)
    api.Azure = azureApi
    ftpApi, err := NewEncodingInputsFtpApi(configs...)
    api.Ftp = ftpApi
    sftpApi, err := NewEncodingInputsSftpApi(configs...)
    api.Sftp = sftpApi
    httpApi, err := NewEncodingInputsHttpApi(configs...)
    api.Http = httpApi
    httpsApi, err := NewEncodingInputsHttpsApi(configs...)
    api.Https = httpsApi
    asperaApi, err := NewEncodingInputsAsperaApi(configs...)
    api.Aspera = asperaApi
    akamaiNetstorageApi, err := NewEncodingInputsAkamaiNetstorageApi(configs...)
    api.AkamaiNetstorage = akamaiNetstorageApi
    srtApi, err := NewEncodingInputsSrtApi(configs...)
    api.Srt = srtApi
    tcpApi, err := NewEncodingInputsTcpApi(configs...)
    api.Tcp = tcpApi
    udpApi, err := NewEncodingInputsUdpApi(configs...)
    api.Udp = udpApi
    udpMulticastApi, err := NewEncodingInputsUdpMulticastApi(configs...)
    api.UdpMulticast = udpMulticastApi
    zixiApi, err := NewEncodingInputsZixiApi(configs...)
    api.Zixi = zixiApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInputsApi) List(queryParams ...func(*query.InputListQueryParams)) (*pagination.InputsListPagination, error) {
    queryParameters := &query.InputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.InputsListPagination
    err := api.apiClient.Get("/encoding/inputs", nil, &responseModel, reqParams)
    return responseModel, err
}

