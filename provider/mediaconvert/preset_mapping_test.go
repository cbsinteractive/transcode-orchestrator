package mediaconvert

import (
	"reflect"
	"testing"

	mc "github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/cbsinteractive/transcode-orchestrator/job"
)

func TestBitrateLevel(t *testing.T) {
	for _, tt := range []struct {
		n       string
		bitrate int64
		level   int64
	}{
		{"0Kbps", 0, +4},
		{"45Kbps", 45000, -1},
		{"128Kbps", 128000, +4},
		{"196Kbps", 196000, +6},
		{"256Kbps", 256000, +8},
		{"500Kbps", 500000, +10},
	} {
		t.Run(tt.n, func(t *testing.T) {
			if h := vbrLevel(tt.bitrate); h != tt.level {
				t.Fatalf("have %v, want %v", h, tt.level)
			}
		})
	}
}

func TestScanType(t *testing.T) {
	dst := job.Preset{}
	src := job.File{}

	for _, tt := range []struct {
		name, src, dst string
		want           *mc.Deinterlacer
	}{
		{"i2p", "interlaced", "progressive", &deinterlacerStandard},
		{"u2p", "unknown", "progressive", &deinterlacerAdaptive},
		{"p2u", "progressive", "unknown", nil},
		{"i2i", "interlaced", "interlaced", nil},
		{"p2i", "progressive", "interlaced", nil},
		{"p2p", "progressive", "progressive", nil},
	} {
		t.Run(tt.name, func(t *testing.T) {
			src.ScanType = job.ScanType(tt.src)
			dst.Video.InterlaceMode = tt.dst
			v := setter{dst, src}.ScanType(nil)
			if have := v.VideoPreprocessors.Deinterlacer; have != tt.want {
				t.Logf("bad deinterlacer:\n\t\thave: %#v\n\t\twant: %#v", have, tt.want)
			}
		})
	}
}

func TestAudioSplit(t *testing.T) {
	input := mc.AudioDescription{
		CodecSettings: &mc.AudioCodecSettings{
			Codec: mc.AudioCodecWav,
			WavSettings: &mc.WavSettings{
				BitDepth:   aws.Int64(24),
				Channels:   aws.Int64(2),
				SampleRate: aws.Int64(48000),
				Format:     "RIFF",
			},
		},
	}

	want := []mc.AudioDescription{{
		RemixSettings: &mc.RemixSettings{
			ChannelMapping: &mc.ChannelMapping{
				OutputChannels: []mc.OutputChannelMapping{{
					InputChannels: []int64{0, -60},
				},
				}},
			ChannelsIn:  aws.Int64(2),
			ChannelsOut: aws.Int64(1),
		},
		CodecSettings: &mc.AudioCodecSettings{
			Codec: mc.AudioCodecWav,
			WavSettings: &mc.WavSettings{
				BitDepth:   aws.Int64(24),
				Channels:   aws.Int64(1),
				SampleRate: aws.Int64(48000),
				Format:     "RIFF",
			},
		},
	}, {
		RemixSettings: &mc.RemixSettings{
			ChannelMapping: &mc.ChannelMapping{
				OutputChannels: []mc.OutputChannelMapping{{
					InputChannels: []int64{-60, 0},
				},
				}},
			ChannelsIn:  aws.Int64(2),
			ChannelsOut: aws.Int64(1),
		},
		CodecSettings: &mc.AudioCodecSettings{
			Codec: mc.AudioCodecWav,
			WavSettings: &mc.WavSettings{
				BitDepth:   aws.Int64(24),
				Channels:   aws.Int64(1),
				SampleRate: aws.Int64(48000),
				Format:     "RIFF",
			},
		},
	}}

	have := audioSplit(input)
	if !reflect.DeepEqual(have, want) {
		t.Fatalf("bad split:\nhave:\t\t%v\nwant:\t\t%v", have, want)
	}
}
