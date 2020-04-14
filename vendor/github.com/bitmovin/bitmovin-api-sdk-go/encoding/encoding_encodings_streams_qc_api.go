package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type EncodingEncodingsStreamsQcApi struct {
    apiClient *common.ApiClient
    Psnr *EncodingEncodingsStreamsQcPsnrApi
}

func NewEncodingEncodingsStreamsQcApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsStreamsQcApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsStreamsQcApi{apiClient: apiClient}

    psnrApi, err := NewEncodingEncodingsStreamsQcPsnrApi(configs...)
    api.Psnr = psnrApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

