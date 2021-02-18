package mediaconvert

import (
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	mc "github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/cbsinteractive/transcode-orchestrator/db"
)

var (
	ErrUnsupported = errors.New("unsupported")
	ErrInvalid     = errors.New("invalid")
)

func h265CodecSettingsFrom(preset db.Preset) (*mc.VideoCodecSettings, error) {
	bitrate := preset.Video.Bitrate
	gopSize := preset.Video.GopSize

	gopUnit, err := h265GopUnitFrom(preset.Video.GopUnit)
	if err != nil {
		return nil, err
	}

	rateControl, err := h265RateControlModeFrom(preset.RateControl)
	if err != nil {
		return nil, err
	}

	profile := mc.H265CodecProfileMainMain
	if preset.Video.HDR10Settings.Enabled {
		profile = mc.H265CodecProfileMain10Main
	}

	level, err := h265CodecLevelFrom(preset.Video.ProfileLevel)
	if err != nil {
		return nil, err
	}

	tuning := mc.H265QualityTuningLevelSinglePassHq
	if preset.TwoPass {
		tuning = mc.H265QualityTuningLevelMultiPassHq
	}

	return &mc.VideoCodecSettings{
		Codec: mc.VideoCodecH265,
		H265Settings: &mc.H265Settings{
			Bitrate:                        aws.Int64(int64(bitrate)),
			GopSize:                        aws.Float64(gopSize),
			GopSizeUnits:                   gopUnit,
			RateControlMode:                rateControl,
			CodecProfile:                   profile,
			CodecLevel:                     level,
			InterlaceMode:                  mc.H265InterlaceModeProgressive,
			ParControl:                     mc.H265ParControlSpecified,
			ParNumerator:                   aws.Int64(1),
			ParDenominator:                 aws.Int64(1),
			QualityTuningLevel:             tuning,
			WriteMp4PackagingType:          mc.H265WriteMp4PackagingTypeHvc1,
			AlternateTransferFunctionSei:   mc.H265AlternateTransferFunctionSeiDisabled,
			SpatialAdaptiveQuantization:    mc.H265SpatialAdaptiveQuantizationEnabled,
			TemporalAdaptiveQuantization:   mc.H265TemporalAdaptiveQuantizationEnabled,
			FlickerAdaptiveQuantization:    mc.H265FlickerAdaptiveQuantizationEnabled,
			SceneChangeDetect:              mc.H265SceneChangeDetectEnabled,
			UnregisteredSeiTimecode:        mc.H265UnregisteredSeiTimecodeDisabled,
			SampleAdaptiveOffsetFilterMode: mc.H265SampleAdaptiveOffsetFilterModeAdaptive,
		},
	}, nil
}

func h265GopUnitFrom(v string) (mc.H265GopSizeUnits, error) {
	switch strings.ToLower(v) {
	case "", db.GopUnitFrames:
		return mc.H265GopSizeUnitsFrames, nil
	case db.GopUnitSeconds:
		return mc.H265GopSizeUnitsSeconds, nil
	default:
		return "", fmt.Errorf("h265: %w: gop unit %q", ErrUnsupported, v)
	}
}

func h265RateControlModeFrom(v string) (mc.H265RateControlMode, error) {
	switch strings.ToLower(v) {
	case "vbr":
		return mc.H265RateControlModeVbr, nil
	case "", "cbr":
		return mc.H265RateControlModeCbr, nil
	case "qvbr":
		return mc.H265RateControlModeQvbr, nil
	default:
		return "", fmt.Errorf("h265: %w: rate control mode: %q", ErrUnsupported, v)
	}
}

func h265CodecLevelFrom(v string) (mc.H265CodecLevel, error) {
	switch v {
	case "":
		return mc.H265CodecLevelAuto, nil
	case "1", "1.0":
		return mc.H265CodecLevelLevel1, nil
	case "2", "2.0":
		return mc.H265CodecLevelLevel2, nil
	case "2.1":
		return mc.H265CodecLevelLevel21, nil
	case "3", "3.0":
		return mc.H265CodecLevelLevel3, nil
	case "3.1":
		return mc.H265CodecLevelLevel31, nil
	case "4", "4.0":
		return mc.H265CodecLevelLevel4, nil
	case "4.1":
		return mc.H265CodecLevelLevel41, nil
	case "5", "5.0":
		return mc.H265CodecLevelLevel5, nil
	case "5.1":
		return mc.H265CodecLevelLevel51, nil
	case "5.2":
		return mc.H265CodecLevelLevel52, nil
	case "6", "6.0":
		return mc.H265CodecLevelLevel6, nil
	case "6.1":
		return mc.H265CodecLevelLevel61, nil
	case "6.2":
		return mc.H265CodecLevelLevel62, nil
	default:
		return "", fmt.Errorf("h265: %w: level: %q", ErrUnsupported, v)
	}
}
