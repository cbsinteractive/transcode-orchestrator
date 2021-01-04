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
	if gu := preset.Video.GopUnit; len(gu) > 0 && gu != defaultGopUnitVP8 {
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

	settings := &mediaconvert.VideoCodecSettings{
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
	}

	if fr := preset.Video.Framerate; !fr.Empty() {
		settings.Vp8Settings.FramerateControl = mediaconvert.Vp8FramerateControlSpecified
		settings.Vp8Settings.FramerateConversionAlgorithm = mediaconvert.Vp8FramerateConversionAlgorithmInterpolate
		settings.Vp8Settings.FramerateNumerator = aws.Int64(int64(fr.Numerator))
		settings.Vp8Settings.FramerateDenominator = aws.Int64(int64(fr.Denominator))
	}

	return settings, nil
}
