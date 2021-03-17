package codec

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/transcode-orchestrator/client/transcoding/job"
)

var ErrUnsupportedValue = errors.New("unsupported value")
var ErrEmptyList = errors.New("empty list")

var AudioSampleRate = 48000.

type enum []string

// Set sets dst to a matching enum value, keyed by src. The
// comparison is done in a case-insensitive way, so it is possible
// that dst.(string) != src on a successful call to Set.
//
// Set returns ErrUnsupportedValue if there is no match
func (e enum) Set(src string, dst interface{}) error {
	if len(e) == 0 {
		return ErrEmptyList
	}
	if src == "" && len(e) > 0 {
		src = e[0]
	}
	for _, v := range e {
		if strings.EqualFold(v, src) {
			_, err := fmt.Sscan(v, dst)
			return err
		}
	}
	return ErrUnsupportedValue
}

func list(v ...interface{}) (s enum) {
	for _, v := range v {
		s = append(s, fmt.Sprint(v))
	}
	return s
}

// VideoPTR: holds pointers to all these very similar codec structs,
// and is used to set them all in a "generic" way. For example, h264
// and h265 are seperate codec objects, but they both have a name,
// so the pointer to it is assigned to this object and then that's passed
// into codec.setVideo.
type VideoPTR struct {
	Name                                     *string
	Width, Height                            **int32
	Bitrate                                  **int64
	MinGop, MaxGop                           **int32
	MinKeyframeInterval, MaxKeyframeInterval **float64
	EncodingMode                             *model.EncodingMode

	Profile, Level interface{}
}

type codec struct {
	Profiles enum
	Levels   enum
	id       string
	err      error
}

// setVideo generalizes setting the very common fields across
// the supported set of codecs. The caller packs pointers to the
// the target struct fields in cfg.
//
// specific codecs, like h264.go:/CodecH264/ will call codec.setVideo,
// and then follow up with their own codec-specific checks. For example,
// the h264 codec originally set the "scene cut threshhold" to zero if the
// GOP size was not zero.
func (c *codec) setVideo(cfg VideoPTR, p job.File) bool {
	*cfg.Name = strings.ToLower(p.Name)
	if n := int32(p.Video.Width); n != 0 && cfg.Width != nil {
		*cfg.Width = &n
	}
	if n := int32(p.Video.Height); n != 0 && cfg.Height != nil {
		*cfg.Height = &n
	}
	if n := int64(p.Video.Bitrate.BPS); n != 0 && cfg.Bitrate != nil {
		*cfg.Bitrate = &n
	}

	size := int32(p.Video.Gop.Size)
	if size != 0 {
		if p.Video.Gop.Seconds() {
			if cfg.MinKeyframeInterval != nil && cfg.MaxKeyframeInterval != nil {
				*cfg.MinKeyframeInterval = &p.Video.Gop.Size
				*cfg.MaxKeyframeInterval = &p.Video.Gop.Size
			}
		} else {
			if cfg.MinGop != nil && cfg.MaxGop != nil {
				*cfg.MinGop = &size
				*cfg.MaxGop = &size
			}
		}
	}

	if cfg.EncodingMode != nil {
		// some of these can fail on single pass mode
		// check that in the specific codecs
		*cfg.EncodingMode = model.EncodingMode_SINGLE_PASS
		if p.Video.Bitrate.TwoPass {
			*cfg.EncodingMode = model.EncodingMode_TWO_PASS
		}
	}

	if cfg.Profile != nil {
		if err := c.Profiles.Set(p.Video.Profile, cfg.Profile); err != nil {
			return c.errorf("%s: profile: %w", *cfg.Name, err)
		}
	}
	if cfg.Level != nil {
		if err := c.Levels.Set(p.Video.Level, cfg.Level); err != nil {
			return c.errorf("%s: level: %w", *cfg.Name, err)
		}
	}

	return c.ok()
}

func (c codec) ok() bool   { return c.err == nil }
func (c codec) Err() error { return c.err }
func (c codec) ID() string { return c.id }
func (c *codec) error(err error) bool {
	c.err = err
	return c.ok()
}
func (c *codec) errorf(fm string, a ...interface{}) bool {
	return c.error(fmt.Errorf(fm, a...))
}

var zero = int32(0)

func ptr(n int) *int32 {
	i := int32(n)
	return &i
}
