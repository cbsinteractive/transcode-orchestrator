package codec

import (
	"fmt"
	"strings"

	"github.com/bitmovin/bitmovin-api-sdk-go"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/transcode-orchestrator/db"
	"github.com/pkg/errors"
)

var h265Levels = []model.LevelH265{
	model.LevelH265_L1, model.LevelH265_L2, model.LevelH265_L2_1, model.LevelH265_L3, model.LevelH265_L3_1,
	model.LevelH265_L4, model.LevelH265_L4_1, model.LevelH265_L5, model.LevelH265_L5_1, model.LevelH265_L5_2,
	model.LevelH265_L6, model.LevelH265_L6_1, model.LevelH265_L6_2,
}

type Codec struct {
	id  string
	err error
}

func (c Codec) ok() bool   { return c.err == nil }
func (c Codec) Err() error { return c.err }
func (c Codec) ID() string { return c.id }

type CodecH265 struct {
	Codec
	cfg *model.H265VideoConfiguration
}

func (c CodecH265) New(dst db.Preset) CodecH265 {
	c.configFrom(dst)
	return c
}

func (c *CodecH265) Create(api *bitmovin.BitmovinApi) (ok bool) {
	create := api.Encoding.Configurations.Video.H265.Create
	if c.ok() {
		c.cfg, c.err = create(*c.cfg)
	}
	if c.ok() {
		c.id = c.cfg.Id
	}
	return c.ok()
}

func (c *CodecH265) configFrom(preset db.Preset) (ok bool) {
	c.cfg.Name = strings.ToLower(preset.Name)
	c.cfg.Profile, c.err = h265ProfileFrom(preset.Video.Profile)
	if !c.ok() {
		return false
	}
	c.cfg.Level, c.err = h265LevelFrom(preset.Video.ProfileLevel)
	if !c.ok() {
		return false
	}

	if n := int32(preset.Video.Width); n != 0 {
		c.cfg.Width = &n
	}
	if n := int32(preset.Video.Height); n != 0 {
		c.cfg.Height = &n
	}
	bitrate := int64(preset.Video.Bitrate)
	c.cfg.Bitrate = &bitrate

	gopSize := int32(preset.Video.GopSize)
	if gopSize != 0 {
		switch strings.ToLower(preset.Video.GopUnit) {
		case db.GopUnitFrames, "":
			c.cfg.MinGop = &gopSize
			c.cfg.MaxGop = &gopSize
		case db.GopUnitSeconds:
			c.cfg.MinKeyframeInterval = &preset.Video.GopSize
			c.cfg.MaxKeyframeInterval = &preset.Video.GopSize
		default:
			c.err = fmt.Errorf("GopUnit %v not recognized", preset.Video.GopUnit)
			return false
		}
		c.cfg.SceneCutThreshold = int32ToPtr(int32(0))
	}

	if hdr10 := preset.Video.HDR10Settings; hdr10.Enabled {
		if !c.setHDR10(hdr10, preset.Video.Profile) {
			return false
		}
	}

	c.cfg.EncodingMode = model.EncodingMode_SINGLE_PASS
	if preset.TwoPass {
		c.cfg.EncodingMode = model.EncodingMode_TWO_PASS
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
		c.cfg.MaxContentLightLevel = int32ToPtr(int32(hdr10.MaxCLL))
	}

	if hdr10.MaxFALL != 0 {
		c.cfg.MaxPictureAverageLightLevel = int32ToPtr(int32(hdr10.MaxFALL))
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
