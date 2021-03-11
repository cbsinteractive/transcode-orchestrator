package codec

import (
	"fmt"

	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/transcode-orchestrator/client/transcoding/job"
)

type CodecOpus struct {
	codec
	cfg model.OpusAudioConfiguration
}

func (c *CodecOpus) set(p job.File) (ok bool) {
	abr := int64(p.Audio.Bitrate)
	c.cfg.Name = fmt.Sprintf("opus_%d_%d", abr, int(AudioSampleRate))
	c.cfg.Bitrate = &abr
	c.cfg.Rate = &AudioSampleRate
	return c.ok()
}
