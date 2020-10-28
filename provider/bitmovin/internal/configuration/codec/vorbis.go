package codec

import (
	"fmt"

	"github.com/bitmovin/bitmovin-api-sdk-go"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/transcode-orchestrator/db"
)

var defaultSampleRate = 48000.

type CodecVorbis struct {
	Codec
	cfg *model.VorbisAudioConfiguration
}

func (c CodecVorbis) New(dst db.Preset) CodecVorbis {
	c.set(dst)
	return c
}

func (c *CodecVorbis) Create(api *bitmovin.BitmovinApi) (ok bool) {
	create := api.Encoding.Configurations.Audio.Vorbis.Create
	if c.ok() {
		c.cfg, c.err = create(*c.cfg)
	}
	if c.ok() {
		c.id = c.cfg.Id
	}
	return c.ok()
}

func (c *CodecVorbis) set(p db.Preset) (ok bool) {
	abr := int64(p.Audio.Bitrate)
	c.cfg.Name = fmt.Sprintf("vorbis_%d_%d", abr, int(defaultSampleRate))
	c.cfg.Bitrate = &abr
	c.cfg.Rate = &defaultSampleRate
	return c.ok()
}
