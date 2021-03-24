package mediaconvert

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	mc "github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/cbsinteractive/transcode-orchestrator/av"
)

// H265RateControlMode
// H264RateControlMode
var RateControl = map[string]string{
	"":     "CBR",
	"cbr":  "CBR",
	"vbr":  "VBR",
	"qvbr": "QVBR",
}

func h264CodecSettingsFrom(f av.File) (*mc.VideoCodecSettings, error) {
	rateControl, err := h264RateControl(f.Video.Bitrate.Control)
	if err != nil {
		return nil, err
	}

	profile, err := h264CodecProfileFrom(f.Video.Profile)
	if err != nil {
		return nil, err
	}

	passes := mc.H264QualityTuningLevelSinglePassHq
	if f.Video.TwoPass {
		passes = mc.H264QualityTuningLevelMultiPassHq
	}

	return &mc.VideoCodecSettings{
		Codec: mc.VideoCodecH264,
		H264Settings: &mc.H264Settings{
			Bitrate:            aws.Int64(int64(f.Video.Bitrate.BPS)),
			GopSize:            aws.Float64(f.Video.Gop.Size),
			GopSizeUnits:       h264GopUnit(f.Video.Gop),
			RateControlMode:    rateControl,
			CodecProfile:       profile,
			CodecLevel:         mc.H264CodecLevelAuto,
			InterlaceMode:      mc.H264InterlaceModeProgressive,
			ParControl:         mc.H264ParControlSpecified,
			ParNumerator:       aws.Int64(1),
			ParDenominator:     aws.Int64(1),
			QualityTuningLevel: passes,
		},
	}, nil
}

func h264GopUnit(g av.Gop) mc.H264GopSizeUnits {
	if g.Seconds() {
		return mc.H264GopSizeUnitsSeconds
	}
	return mc.H264GopSizeUnitsFrames
}

func h264RateControl(v string) (mc.H264RateControlMode, error) {
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
