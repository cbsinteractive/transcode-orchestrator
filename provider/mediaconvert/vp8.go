package mediaconvert

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	mc "github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/cbsinteractive/transcode-orchestrator/client/transcoding/job"
)

const (
	defaultGopUnitVP8 = "frames"
)

func vp8CodecSettingsFrom(f job.File) (*mc.VideoCodecSettings, error) {
	if f.Video.Gop.Seconds() {
		return nil, fmt.Errorf("can't configure gop unit: seconds with vp8. Must use frames")
	}
	return &mc.VideoCodecSettings{
		Codec: mc.VideoCodecVp8,
		Vp8Settings: &mc.Vp8Settings{
			Bitrate:          aws.Int64(int64(f.Video.Bitrate.BPS)),
			GopSize:          aws.Float64(f.Video.Gop.Size),
			RateControlMode:  mc.Vp8RateControlModeVbr,
			FramerateControl: mc.Vp8FramerateControlInitializeFromSource,
			ParControl:       mc.Vp8ParControlSpecified,
			ParNumerator:     aws.Int64(1),
			ParDenominator:   aws.Int64(1),
		},
	}, nil
}
