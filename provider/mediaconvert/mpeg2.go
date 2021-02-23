package mediaconvert

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	mc "github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/cbsinteractive/transcode-orchestrator/db"
)

var ErrProfileUnsupported = errors.New("unsupported profile")

var mpeg2profiles = map[string]mc.Mpeg2CodecProfile{
	"hd422": mc.Mpeg2CodecProfileProfile422,
}

func atoi(a string) int64 {
	i, _ := strconv.Atoi(a)
	return int64(i)
}

func (m mpeg2) validate(p job.Preset) error {
	if profile := p.Video.Profile; profile != "" {
		if _, ok := mpeg2profiles[profile]; !ok {
			return fmt.Errorf("%w: %q", ErrProfileUnsupported, profile)
		}
	}
	return nil
}

func (m mpeg2) apply(p job.Preset) mpeg2 {
	if p.Video.Profile != "" {
		m.CodecProfile = mpeg2profiles[p.Video.Profile]
	}
	if p.Video.Bitrate != 0 {
		m.Bitrate = aws.Int64(int64(p.Video.Bitrate))
	}
	if p.Video.GopSize != 0 {
		m.GopSize = &p.Video.GopSize
	}
	if p.RateControl != "" {
		m.RateControlMode = mc.Mpeg2RateControlMode(p.RateControl)
	}
	return m
}

func (m mpeg2) generate(p job.Preset) (*mc.VideoCodecSettings, error) {
	s := &mc.VideoCodecSettings{
		Codec:         mc.VideoCodecMpeg2,
		Mpeg2Settings: &mc.Mpeg2Settings{},
	}
	if err := m.validate(p); err != nil {
		return s, err
	}
	m = m.apply(p)
	data, _ := json.Marshal(m)
	if err := json.Unmarshal(data, s.Mpeg2Settings); err != nil {
		return s, err
	}
	return s, s.Validate()
}

type mpeg2 mc.Mpeg2Settings

var mpeg2default = mpeg2{
	Bitrate:         aws.Int64(50000000),
	GopSize:         aws.Float64(60),
	CodecProfile:    mc.Mpeg2CodecProfileProfile422,
	CodecLevel:      mc.Mpeg2CodecLevelHigh,
	InterlaceMode:   mc.Mpeg2InterlaceModeTopField,
	RateControlMode: mc.Mpeg2RateControlModeCbr,
	GopSizeUnits:    mc.Mpeg2GopSizeUnitsFrames,
}

var mpeg2XDCAM = mpeg2default
