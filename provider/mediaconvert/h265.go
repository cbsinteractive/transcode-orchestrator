package mediaconvert

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/cbsinteractive/video-transcoding-api/db"
	"github.com/pkg/errors"
)

func h265CodecSettingsFrom(preset db.Preset) (*mediaconvert.VideoCodecSettings, error) {
	bitrate, err := strconv.ParseInt(preset.Video.Bitrate, 10, 64)
	if err != nil {
		return nil, errors.Wrapf(err, "parsing video bitrate %q to int64", preset.Video.Bitrate)
	}

	gopSize, err := strconv.ParseFloat(preset.Video.GopSize, 64)
	if err != nil {
		return nil, errors.Wrapf(err, "parsing gop size %q to float64", preset.Video.GopSize)
	}

	gopUnit, err := h265GopUnitFrom(preset.Video.GopUnit)
	if err != nil {
		return nil, err
	}

	rateControl, err := h265RateControlModeFrom(preset.RateControl)
	if err != nil {
		return nil, err
	}

	profile := mediaconvert.H265CodecProfileMainMain
	if preset.Video.HDR10Settings.Enabled {
		profile = mediaconvert.H265CodecProfileMain10Main
	}

	level, err := h265CodecLevelFrom(preset.Video.ProfileLevel)
	if err != nil {
		return nil, err
	}

	tuning := mediaconvert.H265QualityTuningLevelSinglePassHq
	if preset.TwoPass {
		tuning = mediaconvert.H265QualityTuningLevelMultiPassHq
	}

	return &mediaconvert.VideoCodecSettings{
		Codec: mediaconvert.VideoCodecH265,
		H265Settings: &mediaconvert.H265Settings{
			Bitrate:                        aws.Int64(bitrate),
			GopSize:                        aws.Float64(gopSize),
			GopSizeUnits:                   gopUnit,
			RateControlMode:                rateControl,
			CodecProfile:                   profile,
			CodecLevel:                     level,
			InterlaceMode:                  mediaconvert.H265InterlaceModeProgressive,
			QualityTuningLevel:             tuning,
			WriteMp4PackagingType:          mediaconvert.H265WriteMp4PackagingTypeHvc1,
			AlternateTransferFunctionSei:   mediaconvert.H265AlternateTransferFunctionSeiDisabled,
			SpatialAdaptiveQuantization:    mediaconvert.H265SpatialAdaptiveQuantizationEnabled,
			TemporalAdaptiveQuantization:   mediaconvert.H265TemporalAdaptiveQuantizationEnabled,
			FlickerAdaptiveQuantization:    mediaconvert.H265FlickerAdaptiveQuantizationEnabled,
			SceneChangeDetect:              mediaconvert.H265SceneChangeDetectEnabled,
			UnregisteredSeiTimecode:        mediaconvert.H265UnregisteredSeiTimecodeDisabled,
			SampleAdaptiveOffsetFilterMode: mediaconvert.H265SampleAdaptiveOffsetFilterModeAdaptive,
		},
	}, nil
}

func h265GopUnitFrom(gopUnit string) (mediaconvert.H265GopSizeUnits, error) {
	gopUnit = strings.ToLower(gopUnit)
	switch gopUnit {
	case "", db.GopUnitFrames:
		return mediaconvert.H265GopSizeUnitsFrames, nil
	case db.GopUnitSeconds:
		return mediaconvert.H265GopSizeUnitsSeconds, nil
	default:
		return "", fmt.Errorf("gop unit %q is not supported with mediaconvert", gopUnit)
	}
}

func h265RateControlModeFrom(rateControl string) (mediaconvert.H265RateControlMode, error) {
	rateControl = strings.ToLower(rateControl)
	switch rateControl {
	case "vbr":
		return mediaconvert.H265RateControlModeVbr, nil
	case "cbr":
		return mediaconvert.H265RateControlModeCbr, nil
	case "", "qvbr":
		return mediaconvert.H265RateControlModeQvbr, nil
	default:
		return "", fmt.Errorf("rate control mode %q is not supported with mediaconvert", rateControl)
	}
}

func h265CodecLevelFrom(level string) (mediaconvert.H265CodecLevel, error) {
	switch level {
	case "":
		return mediaconvert.H265CodecLevelAuto, nil
	case "1", "1.0":
		return mediaconvert.H265CodecLevelLevel1, nil
	case "2", "2.0":
		return mediaconvert.H265CodecLevelLevel2, nil
	case "2.1":
		return mediaconvert.H265CodecLevelLevel21, nil
	case "3", "3.0":
		return mediaconvert.H265CodecLevelLevel3, nil
	case "3.1":
		return mediaconvert.H265CodecLevelLevel31, nil
	case "4", "4.0":
		return mediaconvert.H265CodecLevelLevel4, nil
	case "4.1":
		return mediaconvert.H265CodecLevelLevel41, nil
	case "5", "5.0":
		return mediaconvert.H265CodecLevelLevel5, nil
	case "5.1":
		return mediaconvert.H265CodecLevelLevel51, nil
	case "5.2":
		return mediaconvert.H265CodecLevelLevel52, nil
	case "6", "6.0":
		return mediaconvert.H265CodecLevelLevel6, nil
	case "6.1":
		return mediaconvert.H265CodecLevelLevel61, nil
	case "6.2":
		return mediaconvert.H265CodecLevelLevel62, nil
	default:
		return "", fmt.Errorf("h265 level %q is not supported with mediaconvert", level)
	}
}
