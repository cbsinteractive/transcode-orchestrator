package mediaconvert

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/cbsinteractive/transcode-orchestrator/db"
)

func av1CodecSettingsFrom(preset db.Preset) (*mediaconvert.VideoCodecSettings, error) {
	bitrate := int64(preset.Video.Bitrate)
	gopSize, err := av1GopSizeFrom(preset.Video.GopUnit, preset.Video.GopSize)
	if err != nil {
		return nil, err
	}

	return &mediaconvert.VideoCodecSettings{
		Codec: mediaconvert.VideoCodecAv1,
		Av1Settings: &mediaconvert.Av1Settings{
			MaxBitrate: aws.Int64(bitrate),
			GopSize:    aws.Float64(gopSize),
			QvbrSettings: &mediaconvert.Av1QvbrSettings{
				QvbrQualityLevel:         aws.Int64(7),
				QvbrQualityLevelFineTune: aws.Float64(0),
			},
			RateControlMode: mediaconvert.Av1RateControlModeQvbr,
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
