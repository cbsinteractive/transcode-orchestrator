package codec

import (
	"fmt"
	"strings"

	"github.com/bitmovin/bitmovin-api-sdk-go"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/transcode-orchestrator/db"
)

var h264Levels = []model.LevelH264{
	model.LevelH264_L1, model.LevelH264_L1b, model.LevelH264_L1_1, model.LevelH264_L1_2, model.LevelH264_L1_3,
	model.LevelH264_L2, model.LevelH264_L2_1, model.LevelH264_L2_2, model.LevelH264_L3, model.LevelH264_L3_1,
	model.LevelH264_L3_2, model.LevelH264_L4, model.LevelH264_L4_1, model.LevelH264_L4_2, model.LevelH264_L5,
	model.LevelH264_L5_1, model.LevelH264_L5_2,
}

type CodecH264 struct {
	Codec
	cfg *model.H264VideoConfiguration
}

func (c CodecH264) New(dst db.Preset) CodecH264 {
	c.set(dst)
	return c
}

func (c *CodecH264) Create(api *bitmovin.BitmovinApi) (ok bool) {
	create := api.Encoding.Configurations.Video.H264.Create
	if c.ok() {
		c.cfg, c.err = create(*c.cfg)
	}
	if c.ok() {
		c.id = c.cfg.Id
	}
	return c.ok()
}

func (c *CodecH264) set(preset db.Preset) (ok bool) {
	if !c.setCommon(ConfigPTR{
		Name:                &c.cfg.Name,
		Width:               &c.cfg.Width,
		Height:              &c.cfg.Height,
		Bitrate:             &c.cfg.Bitrate,
		MinGop:              &c.cfg.MinGop,
		MaxGop:              &c.cfg.MaxGop,
		MinKeyframeInterval: &c.cfg.MinKeyframeInterval,
		MaxKeyframeInterval: &c.cfg.MaxKeyframeInterval,
		EncodingMode:        &c.cfg.EncodingMode,
	}, preset) {
		return false
	}

	if preset.Video.GopSize != 0 {
		c.cfg.SceneCutThreshold = int32ToPtr(int32(0))
	}

	profile, err := profileFrom(preset.Video.Profile)
	if err != nil {
		return c.error(err)
	}
	c.cfg.Profile = profile

	level, err := levelFrom(preset.Video.ProfileLevel)
	if err != nil {
		return c.error(err)
	}
	c.cfg.Level = level
	return c.ok()
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
