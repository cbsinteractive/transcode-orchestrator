package codec

import (
	"fmt"

	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/transcode-orchestrator/db"
)

type CodecAAC struct {
	codec
	cfg model.AacAudioConfiguration
}

func (c *CodecAAC) set(p db.Preset) (ok bool) {
	abr := int64(p.Audio.Bitrate)
	c.cfg.Name = fmt.Sprintf("aac_%d_%d", abr, int(AudioSampleRate))
	c.cfg.Bitrate = &abr
	c.cfg.Rate = &AudioSampleRate
	return
}
