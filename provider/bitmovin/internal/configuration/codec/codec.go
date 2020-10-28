package codec

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/transcode-orchestrator/db"
)

var ErrUnsupportedValue = errors.New("unsupported value")
var ErrEmptyList = errors.New("empty list")

var defaultSampleRate = 48000.

type enum []string

func (e enum) Set(src string, dst interface{}) error {
	if len(e) == 0 {
		return ErrEmptyList
	}
	if src == "" && len(e) > 0 {
		src = e[0]
	}
	for _, v := range e {
		if strings.EqualFold(v, src) {
			if err := json.Unmarshal([]byte(v), dst); err != nil {
				return err
			}
		}
	}
	return ErrUnsupportedValue
}

func list(v ...interface{}) (s enum) {
	for _, v := range v {
		s = append(s, v.(string))
	}
	return s
}

type Codec struct {
	Profiles enum
	Levels   enum
	id       string
	err      error
}

type ConfigPTR struct {
	Name                                     *string
	Width, Height                            **int32
	Bitrate                                  **int64
	MinGop, MaxGop                           **int32
	MinKeyframeInterval, MaxKeyframeInterval **float64
	EncodingMode                             *model.EncodingMode

	Profile, Level interface{}
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

	if cfg.Profile != nil {
		if err := c.Profiles.Set(p.Video.Profile, cfg.Profile); err != nil {
			return c.errorf("%s: profile: %w", *cfg.Name, err)
		}
	}
	if cfg.Level != nil {
		if err := c.Levels.Set(p.Video.ProfileLevel, cfg.Level); err != nil {
			return c.errorf("%s: level: %w", *cfg.Name, err)
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

var zero = int32(0)

func ptr(n uint) *int32 {
	i := int32(n)
	return &i
}
