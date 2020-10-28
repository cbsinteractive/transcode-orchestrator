package codec

import (
	"fmt"
	"strings"

	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/transcode-orchestrator/db"
)

type Codec struct {
	id  string
	err error
}

type ConfigPTR struct {
	Name                                     *string
	Width, Height                            **int32
	Bitrate                                  **int64
	MinGop, MaxGop                           **int32
	MinKeyframeInterval, MaxKeyframeInterval **float64
	EncodingMode                             *model.EncodingMode
}

func (c Codec) setCommon(cfg ConfigPTR, p db.Preset) bool {
	*cfg.Name = strings.ToLower(p.Name)
	if n := int32(p.Video.Width); n != 0 && cfg.Width != nil {
		*cfg.Width = &n
	}
	if n := int32(p.Video.Height); n != 0 && cfg.Height != nil {
		*cfg.Height = &n
	}
	if n := int64(p.Video.Bitrate); n != 0 && cfg.Bitrate != nil {
		*cfg.Bitrate = &n
	}

	gopSize := int32(p.Video.GopSize)
	if gopSize != 0 {
		switch strings.ToLower(p.Video.GopUnit) {
		case db.GopUnitFrames, "":
			if cfg.MinGop != nil && cfg.MaxGop != nil {
				*cfg.MinGop = &gopSize
				*cfg.MaxGop = &gopSize
			}
		case db.GopUnitSeconds:
			if cfg.MinKeyframeInterval != nil && cfg.MaxKeyframeInterval != nil {
				*cfg.MinKeyframeInterval = &p.Video.GopSize
				*cfg.MaxKeyframeInterval = &p.Video.GopSize
			}
		default:
			return c.errorf("GopUnit %v not recognized", p.Video.GopUnit)
		}
	}

	if cfg.EncodingMode != nil {
		// some of these can fail on single pass mode
		// check that in the specific codecs
		*cfg.EncodingMode = model.EncodingMode_SINGLE_PASS
		if p.TwoPass {
			*cfg.EncodingMode = model.EncodingMode_TWO_PASS
		}
	}
	return c.ok()
}

func (c Codec) ok() bool   { return c.err == nil }
func (c Codec) Err() error { return c.err }
func (c Codec) ID() string { return c.id }
func (c *Codec) error(err error) bool {
	c.err = err
	return c.ok()
}
func (c *Codec) errorf(fm string, a ...interface{}) bool {
	return c.error(fmt.Errorf(fm, a...))
}
