package codec

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bitmovin/bitmovin-api-sdk-go"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/video-transcoding-api/db"
	"github.com/pkg/errors"
)

var h264Levels = []model.LevelH264{
	model.LevelH264_L1, model.LevelH264_L1b, model.LevelH264_L1_1, model.LevelH264_L1_2, model.LevelH264_L1_3,
	model.LevelH264_L2, model.LevelH264_L2_1, model.LevelH264_L2_2, model.LevelH264_L3, model.LevelH264_L3_1,
	model.LevelH264_L3_2, model.LevelH264_L4, model.LevelH264_L4_1, model.LevelH264_L4_2, model.LevelH264_L5,
	model.LevelH264_L5_1, model.LevelH264_L5_2,
}

// NewH264 creates a H.264 codec configuration and returns its ID
func NewH264(api *bitmovin.BitmovinApi, preset db.Preset) (string, error) {
	newVidCfg, err := h264ConfigFrom(preset)
	if err != nil {
		return "", errors.Wrap(err, "creating h264 config object")
	}

	vidCfg, err := api.Encoding.Configurations.Video.H264.Create(newVidCfg)
	if err != nil {
		return "", errors.Wrap(err, "creating h264 config with the API")
	}

	return vidCfg.Id, nil
}

func h264ConfigFrom(preset db.Preset) (model.H264VideoConfiguration, error) {
	cfg := model.H264VideoConfiguration{}

	cfg.Name = strings.ToLower(preset.Name)

	profile, err := profileFrom(preset.Video.Profile)
	if err != nil {
		return model.H264VideoConfiguration{}, err
	}
	cfg.Profile = profile

	level, err := levelFrom(preset.Video.ProfileLevel)
	if err != nil {
		return model.H264VideoConfiguration{}, err
	}
	cfg.Level = level

	presetWidth := preset.Video.Width
	if presetWidth != "" {
		width, err := dimensionFrom(presetWidth)
		if err != nil {
			return model.H264VideoConfiguration{}, err
		}
		cfg.Width = width
	}

	presetHeight := preset.Video.Height
	if presetHeight != "" {
		height, err := dimensionFrom(presetHeight)
		if err != nil {
			return model.H264VideoConfiguration{}, err
		}
		cfg.Height = height
	}

	bitrate, err := bitrateFrom(preset.Video.Bitrate)
	if err != nil {
		return model.H264VideoConfiguration{}, err
	}
	cfg.Bitrate = bitrate

	presetGOPSize := preset.Video.GopSize
	if presetGOPSize != "" {
		switch strings.ToLower(preset.Video.GopUnit) {
		case db.GopUnitFrames, "":
			gopSize, err := gopSizeFrom(presetGOPSize)
			if err != nil {
				return model.H264VideoConfiguration{}, err
			}
			cfg.MinGop = gopSize
			cfg.MaxGop = gopSize
		case db.GopUnitSeconds:
			gopSize, err := keyIntervalFrom(presetGOPSize)
			if err != nil {
				return model.H264VideoConfiguration{}, err
			}
			cfg.MinKeyframeInterval = gopSize
			cfg.MaxKeyframeInterval = gopSize
		default:
			return model.H264VideoConfiguration{}, fmt.Errorf("GopUnit %v not recognized", preset.Video.GopUnit)
		}

		cfg.SceneCutThreshold = int32ToPtr(int32(0))
	}

	return cfg, nil
}

func profileFrom(presetProfile string) (model.ProfileH264, error) {
	presetProfile = strings.ToLower(presetProfile)
	switch presetProfile {
	case "high", "":
		return model.ProfileH264_HIGH, nil
	case "main":
		return model.ProfileH264_MAIN, nil
	case "baseline":
		return model.ProfileH264_BASELINE, nil
	default:
		return "", fmt.Errorf("unrecognized h264 profile: %v", presetProfile)
	}
}

func levelFrom(presetLevel string) (model.LevelH264, error) {
	if presetLevel == "" {
		return "", fmt.Errorf("h264 codec level is missing")
	}

	for _, l := range h264Levels {
		if string(l) == presetLevel {
			return l, nil
		}
	}

	return "", fmt.Errorf("level %q cannot be mapped to a bitmovin level", presetLevel)
}

func gopSizeFrom(presetGOPSize string) (*int32, error) {
	dim, err := strconv.ParseInt(presetGOPSize, 10, 32)
	if err != nil {
		return nil, err
	}

	return int32ToPtr(int32(dim)), nil
}

func keyIntervalFrom(presetGOPSize string) (*float64, error) {
	dim, err := strconv.ParseFloat(presetGOPSize, 64)
	if err != nil {
		return nil, err
	}

	return floatToPtr(dim), nil
}
