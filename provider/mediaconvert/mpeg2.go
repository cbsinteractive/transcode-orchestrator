package mediaconvert

import (
	"encoding/json"

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

	/*
		FramerateDenominator int64
		FramerateNumerator int64

		HrdBufferInitialFillPercentage
		HrdBufferSize int64
		MaxBitrate int64
		ParDenominator int64
		ParNumerator int64
		Softness int64
	*/
}

func (m mpeg2) apply(p db.Preset) mpeg2 {
	return m //TODO
}

func (m mpeg2) generate() (*mediaconvert.VideoCodecSettings, error) {
	s := &mediaconvert.VideoCodecSettings{
		Codec:         mediaconvert.VideoCodecMpeg2,
		Mpeg2Settings: &mediaconvert.Mpeg2Settings{},
	}
	data, _ := json.Marshal(m)
	if err := json.Unmarshal(data, s.Mpeg2Settings); err != nil{
		return s, err
	}
	return s, s.Validate()
}

var mpeg2default = mpeg2{
	CodecProfile:    "PROFILE_422",
	Bitrate:         50000000,
	GopSize:         12,
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
