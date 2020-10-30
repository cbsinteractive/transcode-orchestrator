package codec

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/transcode-orchestrator/db"
	"github.com/pkg/errors"
)

var h265Profiles = list(
	model.ProfileH265_MAIN,   // "main"
	model.ProfileH265_MAIN10, // "main10"
)
var h265Levels = list(
	model.LevelH265_L1,
	model.LevelH265_L2, model.LevelH265_L2_1,
	model.LevelH265_L3, model.LevelH265_L3_1,
	model.LevelH265_L4, model.LevelH265_L4_1,
	model.LevelH265_L5, model.LevelH265_L5_1, model.LevelH265_L5_2,
	model.LevelH265_L6, model.LevelH265_L6_1, model.LevelH265_L6_2,
)

type CodecH265 struct {
	codec
	cfg model.H265VideoConfiguration
}

func (c *CodecH265) set(preset db.Preset) (ok bool) {
	c.Profiles = h265Profiles
	c.Levels = h265Levels
	if !c.setVideo(VideoPTR{
		Name:                &c.cfg.Name,
		Width:               &c.cfg.Width,
		Height:              &c.cfg.Height,
		Bitrate:             &c.cfg.Bitrate,
		MinGop:              &c.cfg.MinGop,
		MaxGop:              &c.cfg.MaxGop,
		MinKeyframeInterval: &c.cfg.MinKeyframeInterval,
		MaxKeyframeInterval: &c.cfg.MaxKeyframeInterval,
		EncodingMode:        &c.cfg.EncodingMode,
		Profile:             &c.cfg.Profile,
		Level:               &c.cfg.Level,
	}, preset) {
		return false
	}

	if preset.Video.GopSize != 0 {
		c.cfg.SceneCutThreshold = &zero
	}

	if hdr10 := preset.Video.HDR10Settings; hdr10.Enabled {
		if !c.setHDR10(hdr10, preset.Video.Profile) {
			return false
		}
	}

	return c.ok()
}

func (c *CodecH265) setHDR10(hdr10 db.HDR10Settings, requestedProfile string) bool {
	c.cfg.ColorConfig = &model.ColorConfig{
		ColorTransfer:  model.ColorTransfer_SMPTE2084,
		ColorPrimaries: model.ColorPrimaries_BT2020,
		ColorSpace:     model.ColorSpace_BT2020_NCL,
	}

	if hdr10.MasterDisplay != "" {
		c.cfg.MasterDisplay = hdr10.MasterDisplay
	}

	if hdr10.MaxCLL != 0 {
		c.cfg.MaxContentLightLevel = ptr(hdr10.MaxCLL)
	}

	if hdr10.MaxFALL != 0 {
		c.cfg.MaxPictureAverageLightLevel = ptr(hdr10.MaxFALL)
	}

	c.cfg.PixelFormat = model.PixelFormat_YUV420P10LE

	if requestedProfile == "" {
		c.cfg.Profile = model.ProfileH265_MAIN10
	}
	if c.cfg.Profile != model.ProfileH265_MAIN10 {
		c.err = errors.New("for HDR10 jobs outputting HEVC, profile must be main10")
	}
	return c.ok()

}
