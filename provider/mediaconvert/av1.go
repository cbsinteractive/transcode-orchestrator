package mediaconvert

import (
	"fmt"

	mc "github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/cbsinteractive/transcode-orchestrator/av"
)

func av1CodecSettingsFrom(f av.File) (*mc.VideoCodecSettings, error) {
	bitrate := int64(f.Video.Bitrate.BPS)
	gopSize, err := av1GopSizeFrom(f.Video.Gop)
	if err != nil {
		return nil, err
	}

	return &mc.VideoCodecSettings{
		Codec: mc.VideoCodecAv1,
		Av1Settings: &mc.Av1Settings{
			MaxBitrate: aws.Int64(bitrate),
			GopSize:    aws.Float64(gopSize),
			QvbrSettings: &mc.Av1QvbrSettings{
				QvbrQualityLevel:         aws.Int64(7),
				QvbrQualityLevelFineTune: aws.Float64(0),
			},
			RateControlMode: mc.Av1RateControlModeQvbr,
		},
	}, nil
}

func av1GopSizeFrom(g av.Gop) (float64, error) {
	if g.Seconds() {
		return 0, fmt.Errorf(`gop unit "seconds" is not supported with mediaconvert and AV1`)
	}
	return g.Size, nil
}
