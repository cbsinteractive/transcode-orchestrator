package codec

import (
	"fmt"

	"github.com/bitmovin/bitmovin-api-sdk-go"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/transcode-orchestrator/db"
)

type CodecAAC struct {
	Codec
	cfg *model.AacAudioConfiguration
}

func (c CodecAAC) New(dst db.Preset) CodecAAC {
	c.set(dst)
	return c
}

func (c *CodecAAC) Create(api *bitmovin.BitmovinApi) (ok bool) {
	create := api.Encoding.Configurations.Audio.Aac.Create
	if c.ok() {
		c.cfg, c.err = create(*c.cfg)
	}
	if c.ok() {
		c.id = c.cfg.Id
	}
	return c.ok()
}

func (c *CodecAAC) set(p db.Preset) (ok bool) {
	abr := int64(p.Audio.Bitrate)
	c.cfg.Name = fmt.Sprintf("aac_%d_%d", abr, int(defaultSampleRate))
	c.cfg.Bitrate = &abr
	c.cfg.Rate = &defaultSampleRate
	return c.ok()
}
