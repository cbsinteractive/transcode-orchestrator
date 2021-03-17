package codec

import (
	"errors"
	"strings"

	"github.com/bitmovin/bitmovin-api-sdk-go"
	"github.com/cbsinteractive/transcode-orchestrator/client/transcoding/job"
)

var ErrUnsupported = errors.New("codec unsupported")

// Preset is a bitmovin Preset. This is a summarized form of an output
// for a bitmovin job. The Fields ending with "ID" store the identifiers
// bitmovin provided to us as we created all these codec configurations.
//
// There may be room to cache very common configurations for job runs
// to save bandwidth/cpu cycles by re-using common presets. That can
// be done with a sync.Map quickly.
type Preset struct {
	Name          string
	Container     string
	VideoCodec    string
	VideoConfigID string
	VideoFilters  []string
	AudioCodec    string
	AudioConfigID string
}

func (j Preset) HasVideo() bool {
	return j.VideoConfigID != ""
}

func New(codec string, preset job.File) (Codec, error) {
	c := enabled[strings.ToUpper(codec)]
	if c == nil {
		return nil, ErrUnsupported
	}
	c = c.New(preset)
	return c, c.Err()
}

func Summary(c Codec, src job.File, dst Preset) Preset {
	if c.Err() != nil {
		return dst
	}
	dst.Name = src.Name
	dst.Container = src.Container

	switch c.(type) {
	case interface{ video() }:
		dst.VideoCodec = c.Kind()
		dst.VideoConfigID = c.ID()
	case interface{}:
		dst.AudioCodec = c.Kind()
		dst.AudioConfigID = c.ID()
	}

	return dst
}

var enabled = map[string]Codec{
	"AAC":    &CodecAAC{},
	"AV1":    &CodecAV1{},
	"H264":   &CodecH264{},
	"H265":   &CodecH265{},
	"OPUS":   &CodecOpus{},
	"VORBIS": &CodecVorbis{},
	"VP8":    &CodecVP8{},
}

type Codec interface {
	New(p job.File) Codec
	Create(*bitmovin.BitmovinApi) bool
	Err() error

	Kind() string
	Name() string
	ID() string
}

func (c CodecAAC) Kind() string    { return "AAC" }
func (c CodecAV1) Kind() string    { return "AV1" }
func (c CodecH264) Kind() string   { return "H264" }
func (c CodecH265) Kind() string   { return "H265" }
func (c CodecOpus) Kind() string   { return "OPUS" }
func (c CodecVorbis) Kind() string { return "VORBIS" }
func (c CodecVP8) Kind() string    { return "VP8" }

func (c CodecAAC) Name() string    { return c.cfg.Name }
func (c CodecAV1) Name() string    { return c.cfg.Name }
func (c CodecH264) Name() string   { return c.cfg.Name }
func (c CodecH265) Name() string   { return c.cfg.Name }
func (c CodecOpus) Name() string   { return c.cfg.Name }
func (c CodecVorbis) Name() string { return c.cfg.Name }
func (c CodecVP8) Name() string    { return c.cfg.Name }

func (c CodecAAC) ID() string    { return c.cfg.Id }
func (c CodecAV1) ID() string    { return c.cfg.Id }
func (c CodecH264) ID() string   { return c.cfg.Id }
func (c CodecH265) ID() string   { return c.cfg.Id }
func (c CodecOpus) ID() string   { return c.cfg.Id }
func (c CodecVorbis) ID() string { return c.cfg.Id }
func (c CodecVP8) ID() string    { return c.cfg.Id }

func (c CodecAV1) video()  {}
func (c CodecH264) video() {}
func (c CodecH265) video() {}
func (c CodecVP8) video()  {}

func (c CodecAAC) New(p job.File) Codec    { c.set(p); return &c }
func (c CodecAV1) New(p job.File) Codec    { c.set(p); return &c }
func (c CodecH264) New(p job.File) Codec   { c.set(p); return &c }
func (c CodecH265) New(p job.File) Codec   { c.set(p); return &c }
func (c CodecOpus) New(p job.File) Codec   { c.set(p); return &c }
func (c CodecVorbis) New(p job.File) Codec { c.set(p); return &c }
func (c CodecVP8) New(p job.File) Codec    { c.set(p); return &c }

func (c *CodecAAC) Create(api *bitmovin.BitmovinApi) (ok bool) {
	create := api.Encoding.Configurations.Audio.Aac.Create

	dst := &c.cfg
	dst, c.err = create(c.cfg)
	if c.ok() {
		c.cfg = *dst
	}
	return c.ok()
}
func (c *CodecAV1) Create(api *bitmovin.BitmovinApi) (ok bool) {
	create := api.Encoding.Configurations.Video.Av1.Create

	dst := &c.cfg
	dst, c.err = create(c.cfg)
	if c.ok() {
		c.cfg = *dst
	}
	return c.ok()
}
func (c *CodecH264) Create(api *bitmovin.BitmovinApi) (ok bool) {
	create := api.Encoding.Configurations.Video.H264.Create

	dst := &c.cfg
	dst, c.err = create(c.cfg)
	if c.ok() {
		c.cfg = *dst
	}
	return c.ok()
}
func (c *CodecH265) Create(api *bitmovin.BitmovinApi) (ok bool) {
	create := api.Encoding.Configurations.Video.H265.Create

	dst := &c.cfg
	dst, c.err = create(c.cfg)
	if c.ok() {
		c.cfg = *dst
	}
	return c.ok()
}
func (c *CodecOpus) Create(api *bitmovin.BitmovinApi) (ok bool) {
	create := api.Encoding.Configurations.Audio.Opus.Create

	dst := &c.cfg
	dst, c.err = create(c.cfg)
	if c.ok() {
		c.cfg = *dst
	}
	return c.ok()
}
func (c *CodecVorbis) Create(api *bitmovin.BitmovinApi) (ok bool) {
	create := api.Encoding.Configurations.Audio.Vorbis.Create

	dst := &c.cfg
	dst, c.err = create(c.cfg)
	if c.ok() {
		c.cfg = *dst
	}
	return c.ok()
}
func (c *CodecVP8) Create(api *bitmovin.BitmovinApi) (ok bool) {
	create := api.Encoding.Configurations.Video.Vp8.Create

	dst := &c.cfg
	dst, c.err = create(c.cfg)
	if c.ok() {
		c.cfg = *dst
	}
	return c.ok()
}
