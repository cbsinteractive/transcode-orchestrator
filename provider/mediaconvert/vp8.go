package mediaconvert

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/cbsinteractive/transcode-orchestrator/db"
	"github.com/pkg/errors"
)

const (
	defaultGopUnitVP8 = "frames"
)

func vp8CodecSettingsFrom(preset db.Preset) (*mediaconvert.VideoCodecSettings, error) {
	if preset.Video.GopUnit != defaultGopUnitVP8 {
		return nil, fmt.Errorf("can't configure gop unit: %v with vp8. Must use frames", preset.Video.GopUnit)
	}

	gopSize, err := strconv.ParseFloat(preset.Video.GopSize, 64)
	if err != nil {
		return nil, errors.Wrapf(err, "parsing gop size %q to float64", preset.Video.GopSize)
	}

	bitrate, err := strconv.ParseInt(preset.Video.Bitrate, 10, 64)
	if err != nil {
		return nil, errors.Wrapf(err, "parsing video bitrate %q to int64", preset.Video.Bitrate)
	}

	return &mediaconvert.VideoCodecSettings{
		Codec: mediaconvert.VideoCodecVp8,
		Vp8Settings: &mediaconvert.Vp8Settings{
			Bitrate:         aws.Int64(bitrate),
			GopSize:         aws.Float64(gopSize),
			RateControlMode: mediaconvert.Vp8RateControlModeVbr,
			//TODO: which of these are needed below
			//FramerateControl
			//FramerateConversionAlgorithm
			//FramerateDenominator
			//FramerateNumerator
			//HrdBufferSize
			//MaxBitrate
			//ParControl
			//ParDenominator
			//ParNumerator
			//QualitytuningLevel
		},
	}, nil
}
