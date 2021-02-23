package mediaconvert

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	mc "github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/cbsinteractive/transcode-orchestrator/db"
)

const (
	defaultGopUnitVP8 = "frames"
)

func vp8CodecSettingsFrom(preset job.Preset) (*mc.VideoCodecSettings, error) {
	if gu := preset.Video.GopUnit; len(gu) > 0 && gu != defaultGopUnitVP8 {
		return nil, fmt.Errorf("can't configure gop unit: %v with vp8. Must use frames", preset.Video.GopUnit)
	}

	gopSize := preset.Video.GopSize
	bitrate := int64(preset.Video.Bitrate)

	return &mc.VideoCodecSettings{
		Codec: mc.VideoCodecVp8,
		Vp8Settings: &mc.Vp8Settings{
			Bitrate:          aws.Int64(bitrate),
			GopSize:          aws.Float64(gopSize),
			RateControlMode:  mc.Vp8RateControlModeVbr,
			FramerateControl: mc.Vp8FramerateControlInitializeFromSource,
			ParControl:       mc.Vp8ParControlSpecified,
			ParNumerator:     aws.Int64(1),
			ParDenominator:   aws.Int64(1),
		},
	}, nil
}
