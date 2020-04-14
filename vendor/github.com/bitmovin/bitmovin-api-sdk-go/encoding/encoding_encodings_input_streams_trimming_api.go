package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type EncodingEncodingsInputStreamsTrimmingApi struct {
    apiClient *common.ApiClient
    TimeBased *EncodingEncodingsInputStreamsTrimmingTimeBasedApi
    TimecodeTrack *EncodingEncodingsInputStreamsTrimmingTimecodeTrackApi
    H264PictureTiming *EncodingEncodingsInputStreamsTrimmingH264PictureTimingApi
}

func NewEncodingEncodingsInputStreamsTrimmingApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsInputStreamsTrimmingApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsInputStreamsTrimmingApi{apiClient: apiClient}

    timeBasedApi, err := NewEncodingEncodingsInputStreamsTrimmingTimeBasedApi(configs...)
    api.TimeBased = timeBasedApi
    timecodeTrackApi, err := NewEncodingEncodingsInputStreamsTrimmingTimecodeTrackApi(configs...)
    api.TimecodeTrack = timecodeTrackApi
    h264PictureTimingApi, err := NewEncodingEncodingsInputStreamsTrimmingH264PictureTimingApi(configs...)
    api.H264PictureTiming = h264PictureTimingApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

