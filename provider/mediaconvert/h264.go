package mediaconvert

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/NYTimes/video-transcoding-api/db"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/pkg/errors"
)

func h264CodecSettingsFrom(preset db.Preset) (*mediaconvert.VideoCodecSettings, error) {
	bitrate, err := strconv.ParseInt(preset.Video.Bitrate, 10, 64)
	if err != nil {
		return nil, errors.Wrapf(err, "parsing video bitrate %q to int64", preset.Video.Bitrate)
	}

	gopSize, err := strconv.ParseFloat(preset.Video.GopSize, 64)
	if err != nil {
		return nil, errors.Wrapf(err, "parsing gop size %q to int64", preset.Video.GopSize)
	}

	rateControl, err := h264RateControlModeFrom(preset.RateControl)
	if err != nil {
		return nil, err
	}

	profile, err := h264CodecProfileFrom(preset.Video.Profile)
	if err != nil {
		return nil, err
	}

	level, err := h264CodecLevelFrom(preset.Video.ProfileLevel)
	if err != nil {
		return nil, err
	}

	interlaceMode, err := h264InterlaceModeFrom(preset.Video.InterlaceMode)
	if err != nil {
		return nil, err
	}

	tuning := mediaconvert.H264QualityTuningLevelSinglePassHq
	if preset.TwoPass {
		tuning = mediaconvert.H264QualityTuningLevelMultiPassHq
	}

	return &mediaconvert.VideoCodecSettings{
		Codec: mediaconvert.VideoCodecH264,
		H264Settings: &mediaconvert.H264Settings{
			Bitrate:            aws.Int64(bitrate),
			GopSize:            aws.Float64(gopSize),
			RateControlMode:    rateControl,
			CodecProfile:       profile,
			CodecLevel:         level,
			InterlaceMode:      interlaceMode,
			QualityTuningLevel: tuning,
		},
	}, nil
}

func h264RateControlModeFrom(rateControl string) (mediaconvert.H264RateControlMode, error) {
	rateControl = strings.ToLower(rateControl)
	switch rateControl {
	case "vbr":
		return mediaconvert.H264RateControlModeVbr, nil
	case "", "cbr":
		return mediaconvert.H264RateControlModeCbr, nil
	case "qvbr":
		return mediaconvert.H264RateControlModeQvbr, nil
	default:
		return "", fmt.Errorf("rate control mode %q is not supported with mediaconvert", rateControl)
	}
}

func h264CodecProfileFrom(profile string) (mediaconvert.H264CodecProfile, error) {
	profile = strings.ToLower(profile)
	switch profile {
	case "baseline":
		return mediaconvert.H264CodecProfileBaseline, nil
	case "main":
		return mediaconvert.H264CodecProfileMain, nil
	case "", "high":
		return mediaconvert.H264CodecProfileHigh, nil
	default:
		return "", fmt.Errorf("h264 profile %q is not supported with mediaconvert", profile)
	}
}

func h264CodecLevelFrom(level string) (mediaconvert.H264CodecLevel, error) {
	switch level {
	case "":
		return mediaconvert.H264CodecLevelAuto, nil
	case "1", "1.0":
		return mediaconvert.H264CodecLevelLevel1, nil
	case "1.1":
		return mediaconvert.H264CodecLevelLevel11, nil
	case "1.2":
		return mediaconvert.H264CodecLevelLevel12, nil
	case "1.3":
		return mediaconvert.H264CodecLevelLevel13, nil
	case "2", "2.0":
		return mediaconvert.H264CodecLevelLevel2, nil
	case "2.1":
		return mediaconvert.H264CodecLevelLevel21, nil
	case "2.2":
		return mediaconvert.H264CodecLevelLevel22, nil
	case "3", "3.0":
		return mediaconvert.H264CodecLevelLevel3, nil
	case "3.1":
		return mediaconvert.H264CodecLevelLevel31, nil
	case "3.2":
		return mediaconvert.H264CodecLevelLevel32, nil
	case "4", "4.0":
		return mediaconvert.H264CodecLevelLevel4, nil
	case "4.1":
		return mediaconvert.H264CodecLevelLevel41, nil
	case "4.2":
		return mediaconvert.H264CodecLevelLevel42, nil
	case "5", "5.0":
		return mediaconvert.H264CodecLevelLevel5, nil
	case "5.1":
		return mediaconvert.H264CodecLevelLevel51, nil
	case "5.2":
		return mediaconvert.H264CodecLevelLevel52, nil
	default:
		return "", fmt.Errorf("h264 level %q is not supported with mediaconvert", level)
	}
}

func h264InterlaceModeFrom(mode string) (mediaconvert.H264InterlaceMode, error) {
	mode = strings.ToLower(mode)
	switch mode {
	case "", "progressive":
		return mediaconvert.H264InterlaceModeProgressive, nil
	default:
		return "", fmt.Errorf("h264 interlace mode %q is not supported with mediaconvert", mode)
	}
}
