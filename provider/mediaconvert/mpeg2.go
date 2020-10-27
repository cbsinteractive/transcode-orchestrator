package mediaconvert

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/cbsinteractive/transcode-orchestrator/db"
)

var ErrProfileUnsupported = errors.New("unsupported profile")

var mpeg2profiles = map[string]mediaconvert.Mpeg2CodecProfile{
	"hd422": mediaconvert.Mpeg2CodecProfileProfile422,
}

func atoi(a string) int64 {
	i, _ := strconv.Atoi(a)
	return int64(i)
}

func (m mpeg2) validate(p db.Preset) error {
	if profile := p.Video.Profile; profile != "" {
		if _, ok := mpeg2profiles[profile]; !ok {
			return fmt.Errorf("%w: %q", ErrProfileUnsupported, profile)
		}
	}
	return nil
}

func (m mpeg2) apply(p db.Preset) mpeg2 {
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
		m.RateControlMode = mediaconvert.Mpeg2RateControlMode(p.RateControl)
	}
	return m
}

func (m mpeg2) generate(p db.Preset) (*mediaconvert.VideoCodecSettings, error) {
	s := &mediaconvert.VideoCodecSettings{
		Codec:         mediaconvert.VideoCodecMpeg2,
		Mpeg2Settings: &mediaconvert.Mpeg2Settings{},
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

type mpeg2 mediaconvert.Mpeg2Settings

var mpeg2default = mpeg2{
	Bitrate:         aws.Int64(50000000),
	GopSize:         aws.Float64(60),
	CodecProfile:    mediaconvert.Mpeg2CodecProfileProfile422,
	CodecLevel:      mediaconvert.Mpeg2CodecLevelHigh,
	InterlaceMode:   mediaconvert.Mpeg2InterlaceModeTopField,
	RateControlMode: mediaconvert.Mpeg2RateControlModeCbr,
	GopSizeUnits:    mediaconvert.Mpeg2GopSizeUnitsFrames,
}

var mpeg2XDCAM = mpeg2default
