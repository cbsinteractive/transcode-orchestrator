package codec

import (
	"fmt"

	"github.com/bitmovin/bitmovin-api-sdk-go"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/transcode-orchestrator/db"
)

type CodecOpus struct {
	Codec
	cfg *model.OpusAudioConfiguration
}

func (c CodecOpus) New(dst db.Preset) CodecOpus {
	c.set(dst)
	return c
}

func (c *CodecOpus) Create(api *bitmovin.BitmovinApi) (ok bool) {
	create := api.Encoding.Configurations.Audio.Opus.Create
	if c.ok() {
		c.cfg, c.err = create(*c.cfg)
	}
	if c.ok() {
		c.id = c.cfg.Id
	}
	return c.ok()
}

func (c *CodecOpus) set(p db.Preset) (ok bool) {
	abr := int64(p.Audio.Bitrate)
	c.cfg.Name = fmt.Sprintf("opus_%d_%d", abr, int(defaultSampleRate))
	c.cfg.Bitrate = &abr
	c.cfg.Rate = &defaultSampleRate
	return c.ok()
}
