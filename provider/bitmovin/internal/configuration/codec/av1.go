package codec

import (
	"strings"

	"errors"
	"github.com/bitmovin/bitmovin-api-sdk-go"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/transcode-orchestrator/db"
)

var ErrGopFramesOnly = errors.New("gop unit must be frames")

type CodecAV1 struct {
	Codec
	cfg *model.Av1VideoConfiguration
}

func (c CodecAV1) New(dst db.Preset) CodecAV1 {
	c.set(dst)
	return c
}

func (c *CodecAV1) Create(api *bitmovin.BitmovinApi) (ok bool) {
	create := api.Encoding.Configurations.Video.Av1.Create
	if c.ok() {
		c.cfg, c.err = create(*c.cfg)
	}
	if c.ok() {
		c.id = c.cfg.Id
	}
	return c.ok()
}

func (c *CodecAV1) set(preset db.Preset) (ok bool) {
	if !c.setCommon(ConfigPTR{
		Name:    &c.cfg.Name,
		Width:   &c.cfg.Width,
		Height:  &c.cfg.Height,
		Bitrate: &c.cfg.Bitrate,
		// in av1, this is a group of fields?
		MinGop: &c.cfg.MinGfInterval,
		MaxGop: &c.cfg.MinGfInterval,
	}, preset) {
		return false
	}

	if strings.ToLower(preset.Video.GopUnit) == db.GopUnitSeconds {
		return c.error(ErrGopFramesOnly)
	}
	// Single-pass encoding throws an error
	c.cfg.EncodingMode = model.EncodingMode_TWO_PASS
	return c.ok()
}
