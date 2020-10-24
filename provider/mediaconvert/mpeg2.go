package mediaconvert

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/cbsinteractive/transcode-orchestrator/db"
)

type mpeg2 struct {
	InterlaceMode                       string
	Syntax                              string
	GopClosedCadence                    int64
	GopSize                             float64
	SlowPal                             string
	SpatialAdaptiveQuantization         string
	TemporalAdaptiveQuantization        string
	Bitrate                             int64
	IntraDcPrecision                    string
	FramerateControl                    string
	RateControlMode                     string
	CodecProfile                        string
	Telecine                            string
	MinIInterval                        int64
	AdaptiveQuantization                string
	CodecLevel                          string
	SceneChangeDetect                   string
	QualityTuningLevel                  string
	FramerateConversionAlgorithm        string
	GopSizeUnits                        string
	ParControl                          string
	NumberBFramesBetweenReferenceFrames int64
	DynamicSubGop                       string
}

var ErrProfileUnsupported = errors.New("unsupported profile")

var mpeg2profiles = map[string]string{
	"hd422": "PROFILE_422",
}

func atoi(a string) int64 {
	i, _ := strconv.Atoi(a)
	return int64(i)
}

func (m mpeg2) validate(p db.Preset) error {
	if c := p.Video.Codec; c != "" {
		if _, ok := mpeg2profiles[c]; !ok {
			return fmt.Errorf("%w: %q", ErrProfileUnsupported, c)
		}
	}
	return nil
}

func (m mpeg2) apply(p db.Preset) mpeg2 {
	if p.Video.Codec != "" {
		m.CodecProfile = mpeg2profiles[p.Video.Codec]
	}
	if p.Video.Bitrate != "" {
		m.Bitrate = atoi(p.Video.Bitrate)
	}
	if p.Video.GopSize != "" {
		m.GopSize = float64(atoi(p.Video.GopSize))
	}
	if p.RateControl != "" {
		m.RateControlMode = p.RateControl
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

var mpeg2default = mpeg2{
	CodecProfile:    "PROFILE_422",
	Bitrate:         50000000,
	GopSize:         60,
	InterlaceMode:   "TOP_FIELD",
	RateControlMode: "CBR",

	Syntax:                              "DEFAULT",
	GopClosedCadence:                    1,
	SlowPal:                             "DISABLED",
	SpatialAdaptiveQuantization:         "ENABLED",
	TemporalAdaptiveQuantization:        "ENABLED",
	IntraDcPrecision:                    "AUTO",
	FramerateControl:                    "INITIALIZE_FROM_SOURCE",
	Telecine:                            "NONE",
	MinIInterval:                        0,
	AdaptiveQuantization:                "HIGH",
	CodecLevel:                          "HIGH",
	SceneChangeDetect:                   "ENABLED",
	QualityTuningLevel:                  "SINGLE_PASS",
	FramerateConversionAlgorithm:        "DUPLICATE_DROP",
	GopSizeUnits:                        "FRAMES",
	ParControl:                          "INITIALIZE_FROM_SOURCE",
	NumberBFramesBetweenReferenceFrames: 2,
	DynamicSubGop:                       "STATIC",
}

var mpeg2XDCAM = mpeg2default
