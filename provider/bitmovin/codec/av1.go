package codec

import (
	"strings"

	"errors"

	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/transcode-orchestrator/job"
)

var ErrGopFramesOnly = errors.New("gop unit must be frames")

type CodecAV1 struct {
	codec
	cfg model.Av1VideoConfiguration
}

func (c *CodecAV1) set(preset job.Preset) bool {
	if !c.setVideo(VideoPTR{
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

	if strings.ToLower(preset.Video.GopUnit) == job.GopUnitSeconds {
		return c.error(ErrGopFramesOnly)
	}
	// Single-pass encoding throws an error
	c.cfg.EncodingMode = model.EncodingMode_TWO_PASS
	return c.ok()
}
