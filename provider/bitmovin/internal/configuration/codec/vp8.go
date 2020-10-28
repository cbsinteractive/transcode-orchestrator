package codec

import (
	"github.com/bitmovin/bitmovin-api-sdk-go"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/transcode-orchestrator/db"
)

type CodecVP8 struct {
	Codec
	cfg *model.Vp8VideoConfiguration
}

func (c CodecVP8) New(dst db.Preset) CodecVP8 {
	c.set(dst)
	return c
}

func (c *CodecVP8) Create(api *bitmovin.BitmovinApi) (ok bool) {
	create := api.Encoding.Configurations.Video.Vp8.Create
	if c.ok() {
		c.cfg, c.err = create(*c.cfg)
	}
	if c.ok() {
		c.id = c.cfg.Id
	}
	return c.ok()
}

func (c *CodecVP8) set(preset db.Preset) (ok bool) {
	return c.setCommon(ConfigPTR{
		Name:         &c.cfg.Name,
		Width:        &c.cfg.Width,
		Height:       &c.cfg.Height,
		Bitrate:      &c.cfg.Bitrate,
		EncodingMode: &c.cfg.EncodingMode,
	}, preset)
}
