package mediaconvert

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/cbsinteractive/transcode-orchestrator/db"
)

const (
	defaultGopUnitVP8 = "frames"
)

func vp8CodecSettingsFrom(preset db.Preset) (*mediaconvert.VideoCodecSettings, error) {
	if gu := preset.Video.GopUnit; len(gu) > 0 && gu != defaultGopUnitVP8 {
		return nil, fmt.Errorf("can't configure gop unit: %v with vp8. Must use frames", preset.Video.GopUnit)
	}

	gopSize := preset.Video.GopSize
	bitrate := int64(preset.Video.Bitrate)

	return &mediaconvert.VideoCodecSettings{
		Codec: mediaconvert.VideoCodecVp8,
		Vp8Settings: &mediaconvert.Vp8Settings{
			Bitrate:          aws.Int64(bitrate),
			GopSize:          aws.Float64(gopSize),
			RateControlMode:  mediaconvert.Vp8RateControlModeVbr,
			FramerateControl: mediaconvert.Vp8FramerateControlInitializeFromSource,
			ParControl:       mediaconvert.Vp8ParControlSpecified,
			ParNumerator:     aws.Int64(1),
			ParDenominator:   aws.Int64(1),
		},
	}, nil
}
