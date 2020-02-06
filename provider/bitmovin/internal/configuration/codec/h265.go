package codec

import (
	"fmt"
	"strings"

	"github.com/bitmovin/bitmovin-api-sdk-go"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/video-transcoding-api/db"
	"github.com/pkg/errors"
)

var h265Levels = []model.LevelH265{
	model.LevelH265_L1, model.LevelH265_L2, model.LevelH265_L2_1, model.LevelH265_L3, model.LevelH265_L3_1,
	model.LevelH265_L4, model.LevelH265_L4_1, model.LevelH265_L5, model.LevelH265_L5_1, model.LevelH265_L5_2,
	model.LevelH265_L6, model.LevelH265_L6_1, model.LevelH265_L6_2,
}

// NewH265 creates a H.265 codec configuration and returns its ID
func NewH265(api *bitmovin.BitmovinApi, preset db.Preset) (string, error) {
	newVidCfg, err := h265ConfigFrom(preset)
	if err != nil {
		return "", errors.Wrap(err, "creating h265 config object")
	}

	vidCfg, err := api.Encoding.Configurations.Video.H265.Create(newVidCfg)
	if err != nil {
		return "", errors.Wrap(err, "creating h265 config with the API")
	}

	return vidCfg.Id, nil
}

func h265ConfigFrom(preset db.Preset) (model.H265VideoConfiguration, error) {
	cfg := model.H265VideoConfiguration{}

	cfg.Name = strings.ToLower(preset.Name)

	profile, err := h265ProfileFrom(preset.Video.Profile)
	if err != nil {
		return model.H265VideoConfiguration{}, err
	}
	cfg.Profile = profile

	level, err := h265LevelFrom(preset.Video.ProfileLevel)
	if err != nil {
		return model.H265VideoConfiguration{}, err
	}
	cfg.Level = level

	presetWidth := preset.Video.Width
	if presetWidth != "" {
		width, err := dimensionFrom(presetWidth)
		if err != nil {
			return model.H265VideoConfiguration{}, err
		}
		cfg.Width = width
	}

	presetHeight := preset.Video.Height
	if presetHeight != "" {
		height, err := dimensionFrom(presetHeight)
		if err != nil {
			return model.H265VideoConfiguration{}, err
		}
		cfg.Height = height
	}

	bitrate, err := bitrateFrom(preset.Video.Bitrate)
	if err != nil {
		return model.H265VideoConfiguration{}, err
	}
	cfg.Bitrate = bitrate

	presetGOPSize := preset.Video.GopSize
	if presetGOPSize != "" {
		switch strings.ToLower(preset.Video.GopUnit) {
		case db.GopUnitFrames, "":
			gopSize, err := gopSizeFrom(presetGOPSize)
			if err != nil {
				return model.H265VideoConfiguration{}, err
			}
			cfg.MinGop = gopSize
			cfg.MaxGop = gopSize
		case db.GopUnitSeconds:
			gopSize, err := keyIntervalFrom(presetGOPSize)
			if err != nil {
				return model.H265VideoConfiguration{}, err
			}
			cfg.MinKeyframeInterval = gopSize
			cfg.MaxKeyframeInterval = gopSize
		default:
			return model.H265VideoConfiguration{}, fmt.Errorf("GopUnit %v not recognized", preset.Video.GopUnit)
		}

		cfg.SceneCutThreshold = int32ToPtr(int32(0))
	}

	if hdr10 := preset.Video.HDR10Settings; hdr10.Enabled {
		cfg, err = enrichH265CfgWithHDR10Settings(cfg, hdr10, preset.Video.Profile)
		if err != nil {
			return model.H265VideoConfiguration{}, errors.Wrap(err, "setting HDR10 configs to HEVC codec")
		}
	}

	cfg.EncodingMode = model.EncodingMode_SINGLE_PASS
	if preset.TwoPass {
		cfg.EncodingMode = model.EncodingMode_TWO_PASS
	}

	return cfg, nil
}

func enrichH265CfgWithHDR10Settings(cfg model.H265VideoConfiguration, hdr10 db.HDR10Settings,
	requestedProfile string) (model.H265VideoConfiguration, error) {
	cfg.ColorConfig = &model.ColorConfig{
		ColorTransfer:  model.ColorTransfer_SMPTE2084,
		ColorPrimaries: model.ColorPrimaries_BT2020,
		ColorSpace:     model.ColorSpace_BT2020_NCL,
	}

	if hdr10.MasterDisplay != "" {
		cfg.MasterDisplay = hdr10.MasterDisplay
	}

	if hdr10.MaxCLL != 0 {
		cfg.MaxContentLightLevel = int32ToPtr(int32(hdr10.MaxCLL))
	}

	if hdr10.MaxFALL != 0 {
		cfg.MaxPictureAverageLightLevel = int32ToPtr(int32(hdr10.MaxFALL))
	}

	cfg.PixelFormat = model.PixelFormat_YUV420P10LE

	if requestedProfile == "" {
		cfg.Profile = model.ProfileH265_MAIN10
	}

	if cfg.Profile != model.ProfileH265_MAIN10 {
		return model.H265VideoConfiguration{}, errors.New("for HDR10 jobs outputting HEVC, " +
			"profile must be main10")
	}

	return cfg, nil
}

func h265ProfileFrom(presetProfile string) (model.ProfileH265, error) {
	presetProfile = strings.ToLower(presetProfile)
	switch presetProfile {
	case "main", "":
		return model.ProfileH265_MAIN, nil
	case "main10":
		return model.ProfileH265_MAIN10, nil
	default:
		return "", fmt.Errorf("unrecognized h265 profile: %v", presetProfile)
	}
}

func h265LevelFrom(presetLevel string) (model.LevelH265, error) {
	if presetLevel == "" {
		return "", fmt.Errorf("h265 codec level is missing")
	}

	for _, l := range h265Levels {
		if string(l) == presetLevel {
			return l, nil
		}
	}

	return "", fmt.Errorf("level %q cannot be mapped to a bitmovin level", presetLevel)
}
