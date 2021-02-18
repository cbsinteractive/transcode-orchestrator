package mediaconvert

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	mc "github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/cbsinteractive/transcode-orchestrator/db"
)

// H265RateControlMode
// H264RateControlMode
var RateControl = map[string]string{
	"":     "CBR",
	"cbr":  "CBR",
	"vbr":  "VBR",
	"qvbr": "QVBR",
}
var GopUnits = map[string]string{
	"":                "FRAMES",
	db.GopUnitFrames:  "FRAMES",
	db.GopUnitSeconds: "SECONDS",
}

func h264CodecSettingsFrom(preset db.Preset) (*mc.VideoCodecSettings, error) {
	bitrate := preset.Video.Bitrate
	gopSize := preset.Video.GopSize
	gopUnit, err := h264GopUnitFrom(preset.Video.GopUnit)
	if err != nil {
		return nil, err
	}

	rateControl, err := h264RateControlModeFrom(preset.RateControl)
	if err != nil {
		return nil, err
	}

	profile, err := h264CodecProfileFrom(preset.Video.Profile)
	if err != nil {
		return nil, err
	}

	tuning := mc.H264QualityTuningLevelSinglePassHq
	if preset.TwoPass {
		tuning = mc.H264QualityTuningLevelMultiPassHq
	}

	return &mc.VideoCodecSettings{
		Codec: mc.VideoCodecH264,
		H264Settings: &mc.H264Settings{
			Bitrate:            aws.Int64(int64(bitrate)),
			GopSize:            aws.Float64(gopSize),
			GopSizeUnits:       gopUnit,
			RateControlMode:    rateControl,
			CodecProfile:       profile,
			CodecLevel:         mc.H264CodecLevelAuto,
			InterlaceMode:      mc.H264InterlaceModeProgressive,
			ParControl:         mc.H264ParControlSpecified,
			ParNumerator:       aws.Int64(1),
			ParDenominator:     aws.Int64(1),
			QualityTuningLevel: tuning,
		},
	}, nil
}

func h264GopUnitFrom(v string) (mc.H264GopSizeUnits, error) {
	switch strings.ToLower(v) {
	case "", db.GopUnitFrames:
		return mc.H264GopSizeUnitsFrames, nil
	case db.GopUnitSeconds:
		return mc.H264GopSizeUnitsSeconds, nil
	default:
		return "", fmt.Errorf("h264: %w: gop unit: %q", ErrUnsupported, v)
	}
}

func h264RateControlModeFrom(v string) (mc.H264RateControlMode, error) {
	switch strings.ToLower(v) {
	case "vbr":
		return mc.H264RateControlModeVbr, nil
	case "", "cbr":
		return mc.H264RateControlModeCbr, nil
	case "qvbr":
		return mc.H264RateControlModeQvbr, nil
	default:
		return "", fmt.Errorf("h264: %w: rate control mode: %q", ErrUnsupported, v)
	}
}

func h264CodecProfileFrom(v string) (mc.H264CodecProfile, error) {
	switch strings.ToLower(v) {
	case "baseline":
		return mc.H264CodecProfileBaseline, nil
	case "main":
		return mc.H264CodecProfileMain, nil
	case "", "high":
		return mc.H264CodecProfileHigh, nil
	default:
		return "", fmt.Errorf("h264: %w: profile: %q", ErrUnsupported, v)
	}
}
