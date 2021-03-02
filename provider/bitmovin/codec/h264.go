package codec

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/transcode-orchestrator/job"
)

var h264Profiles = list(
	model.ProfileH264_HIGH,     // "HIGH"
	model.ProfileH264_MAIN,     // "MAIN"
	model.ProfileH264_BASELINE, // "BASELINE"
)
var h264Levels = list(
	model.LevelH264_L1, model.LevelH264_L1b, model.LevelH264_L1_1, model.LevelH264_L1_2, model.LevelH264_L1_3,
	model.LevelH264_L2, model.LevelH264_L2_1, model.LevelH264_L2_2,
	model.LevelH264_L3, model.LevelH264_L3_1, model.LevelH264_L3_2,
	model.LevelH264_L4, model.LevelH264_L4_1, model.LevelH264_L4_2,
	model.LevelH264_L5, model.LevelH264_L5_1, model.LevelH264_L5_2,
)

type CodecH264 struct {
	codec
	cfg model.H264VideoConfiguration
}

func (c *CodecH264) set(preset job.File) (ok bool) {
	c.Profiles = h264Profiles
	c.Levels = h264Levels
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
	if preset.Video.Gop.Size != 0 {
		c.cfg.SceneCutThreshold = &zero
	}
	return c.ok()
}
