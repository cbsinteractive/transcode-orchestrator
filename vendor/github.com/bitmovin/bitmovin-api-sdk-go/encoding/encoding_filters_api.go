package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingFiltersApi struct {
    apiClient *common.ApiClient
    Conform *EncodingFiltersConformApi
    Watermark *EncodingFiltersWatermarkApi
    AudioVolume *EncodingFiltersAudioVolumeApi
    EnhancedWatermark *EncodingFiltersEnhancedWatermarkApi
    Crop *EncodingFiltersCropApi
    Rotate *EncodingFiltersRotateApi
    Deinterlace *EncodingFiltersDeinterlaceApi
    AudioMix *EncodingFiltersAudioMixApi
    DenoiseHqdn3d *EncodingFiltersDenoiseHqdn3dApi
    EbuR128SinglePass *EncodingFiltersEbuR128SinglePassApi
    Text *EncodingFiltersTextApi
    Interlace *EncodingFiltersInterlaceApi
    Unsharp *EncodingFiltersUnsharpApi
    Scale *EncodingFiltersScaleApi
    Type *EncodingFiltersTypeApi
}

func NewEncodingFiltersApi(configs ...func(*common.ApiClient)) (*EncodingFiltersApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingFiltersApi{apiClient: apiClient}

    conformApi, err := NewEncodingFiltersConformApi(configs...)
    api.Conform = conformApi
    watermarkApi, err := NewEncodingFiltersWatermarkApi(configs...)
    api.Watermark = watermarkApi
    audioVolumeApi, err := NewEncodingFiltersAudioVolumeApi(configs...)
    api.AudioVolume = audioVolumeApi
    enhancedWatermarkApi, err := NewEncodingFiltersEnhancedWatermarkApi(configs...)
    api.EnhancedWatermark = enhancedWatermarkApi
    cropApi, err := NewEncodingFiltersCropApi(configs...)
    api.Crop = cropApi
    rotateApi, err := NewEncodingFiltersRotateApi(configs...)
    api.Rotate = rotateApi
    deinterlaceApi, err := NewEncodingFiltersDeinterlaceApi(configs...)
    api.Deinterlace = deinterlaceApi
    audioMixApi, err := NewEncodingFiltersAudioMixApi(configs...)
    api.AudioMix = audioMixApi
    denoiseHqdn3dApi, err := NewEncodingFiltersDenoiseHqdn3dApi(configs...)
    api.DenoiseHqdn3d = denoiseHqdn3dApi
    ebuR128SinglePassApi, err := NewEncodingFiltersEbuR128SinglePassApi(configs...)
    api.EbuR128SinglePass = ebuR128SinglePassApi
    textApi, err := NewEncodingFiltersTextApi(configs...)
    api.Text = textApi
    interlaceApi, err := NewEncodingFiltersInterlaceApi(configs...)
    api.Interlace = interlaceApi
    unsharpApi, err := NewEncodingFiltersUnsharpApi(configs...)
    api.Unsharp = unsharpApi
    scaleApi, err := NewEncodingFiltersScaleApi(configs...)
    api.Scale = scaleApi
    typeApi, err := NewEncodingFiltersTypeApi(configs...)
    api.Type = typeApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingFiltersApi) List(queryParams ...func(*query.FilterListQueryParams)) (*pagination.FiltersListPagination, error) {
    queryParameters := &query.FilterListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.FiltersListPagination
    err := api.apiClient.Get("/encoding/filters", nil, &responseModel, reqParams)
    return responseModel, err
}

