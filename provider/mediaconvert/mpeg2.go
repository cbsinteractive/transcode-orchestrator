package mediaconvert

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	mc "github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/cbsinteractive/transcode-orchestrator/client/transcoding/job"
)

var ErrProfileUnsupported = errors.New("unsupported profile")

var mpeg2profiles = map[string]mc.Mpeg2CodecProfile{
	"hd422": mc.Mpeg2CodecProfileProfile422,
}

func atoi(a string) int64 {
	i, _ := strconv.Atoi(a)
	return int64(i)
}

func (m mpeg2) validate(p job.File) error {
	if profile := p.Video.Profile; profile != "" {
		if _, ok := mpeg2profiles[profile]; !ok {
			return fmt.Errorf("%w: %q", ErrProfileUnsupported, profile)
		}
	}
	return nil
}

func (m mpeg2) apply(p job.File) mpeg2 {
	if v := p.Video.Profile; v != "" {
		m.CodecProfile = mpeg2profiles[v]
	}
	if v := p.Video.Bitrate.BPS; v != 0 {
		m.Bitrate = aws.Int64(int64(v))
	}
	if v := p.Video.Gop.Size; v != 0 {
		m.GopSize = &v
	}
	if v := p.Video.Bitrate.Control; v != "" {
		m.RateControlMode = mc.Mpeg2RateControlMode(v)
	}
	return m
}

func (m mpeg2) generate(p job.File) (*mc.VideoCodecSettings, error) {
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
