package mediaconvert

import (
	"fmt"
	"strings"

	mc "github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/cbsinteractive/transcode-orchestrator/db"
)

func av1CodecSettingsFrom(preset db.Preset) (*mc.VideoCodecSettings, error) {
	bitrate := int64(preset.Video.Bitrate)
	gopSize, err := av1GopSizeFrom(preset.Video.GopUnit, preset.Video.GopSize)
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

func av1GopSizeFrom(gopUnit string, gopSize float64) (float64, error) {
	switch strings.ToLower(gopUnit) {
	case "", db.GopUnitFrames:
		return gopSize, nil
	case db.GopUnitSeconds:
		return 0, fmt.Errorf("gop unit %q is not supported with mediaconvert and AV1", gopUnit)
	default:
		return 0, fmt.Errorf("gop unit %q is not supported with mediaconvert", gopUnit)
	}
}
