package mediaconvert

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/cbsinteractive/pkg/video"
	"github.com/cbsinteractive/transcode-orchestrator/db"
)

func Test_vbrLevel(t *testing.T) {
	tests := []struct {
		name      string
		bitrate   int64
		wantLevel int64
	}{
		{
			name:      "NotSet",
			bitrate:   0,
			wantLevel: 4,
		},
		{
			name:      "45Kbps",
			bitrate:   45000,
			wantLevel: -1,
		},
		{
			name:      "128Kbps",
			bitrate:   128000,
			wantLevel: 4,
		},
		{
			name:      "196Kbps",
			bitrate:   196000,
			wantLevel: 6,
		},
		{
			name:      "256Kbps",
			bitrate:   256000,
			wantLevel: 8,
		},
		{
			name:      "500Kbps",
			bitrate:   500000,
			wantLevel: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := vbrLevel(tt.bitrate); got != tt.wantLevel {
				t.Errorf("vbrLevel() = %v, want %v", got, tt.wantLevel)
			}
		})
	}
}

func TestSetterScanType(t *testing.T) {
	dst := db.Preset{}
	src := db.File{}

	for _, tt := range []struct {
		name, src, dst string
		want           *mediaconvert.Deinterlacer
	}{
		{"i2p", "interlaced", "progressive", &deinterlacerStandard},
		{"u2p", "unknown", "progressive", &deinterlacerAdaptive},
		{"p2u", "progressive", "unknown", nil},
		{"i2i", "interlaced", "interlaced", nil},
		{"p2i", "progressive", "interlaced", nil},
		{"p2p", "progressive", "progressive", nil},
	} {
		t.Run(tt.name, func(t *testing.T) {
			src.ScanType = db.ScanType(tt.src)
			dst.Video.InterlaceMode = tt.dst
			v := setter{dst, src}.ScanType(nil)
			if have := v.VideoPreprocessors.Deinterlacer; have != tt.want {
				t.Logf("bad deinterlacer:\n\t\thave: %#v\n\t\twant: %#v", have, tt.want)
			}
		})
	}
}

func TestCrop(t *testing.T) {
	type (
		dims struct{ height, width uint }
		rect struct{ height, width, x, y int64 }
	)
	for _, tt := range []struct {
		name string
		src  dims
		crop video.Crop
		want rect
	}{
		{
			"Crop",
			dims{height: 150, width: 300},
			video.Crop{Left: 50, Top: 10, Right: 40, Bottom: 20},
			rect{height: 120, width: 210, x: 50, y: 10},
		},
		{
			"CropOdd2Even",
			dims{height: 150, width: 300},
			video.Crop{Left: 55, Top: 11, Right: 49, Bottom: 25},
			rect{height: 114, width: 196, x: 56, y: 12},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			have, err := videoPresetFrom(
				db.Preset{
					Video: db.VideoPreset{Bitrate: "12000", GopSize: "2", Width: "30", Height: "15", Codec: "h264", Crop: tt.crop},
				},
				db.File{Width: tt.src.width, Height: tt.src.height},
			)
			if err != nil {
				t.Fatal(err)
			}
			want := &mediaconvert.Rectangle{Height: &tt.want.height, Width: &tt.want.width, X: &tt.want.x, Y: &tt.want.y}
			if !reflect.DeepEqual(have.Crop, want) {
				t.Errorf("bad crop rect:\nhave: %+v\nwant: %+v", have.Crop, tt.want)
			}
		})
	}
}

func TestAudioSplit(t *testing.T) {
	input := mediaconvert.AudioDescription{
		CodecSettings: &mediaconvert.AudioCodecSettings{
			Codec: mediaconvert.AudioCodecWav,
			WavSettings: &mediaconvert.WavSettings{
				BitDepth:   aws.Int64(24),
				Channels:   aws.Int64(2),
				SampleRate: aws.Int64(48000),
				Format:     "RIFF",
			},
		},
	}

	want := []mediaconvert.AudioDescription{{
		RemixSettings: &mediaconvert.RemixSettings{
			ChannelMapping: &mediaconvert.ChannelMapping{
				OutputChannels: []mediaconvert.OutputChannelMapping{{
					InputChannels: []int64{0, -60},
				},
				}},
			ChannelsIn:  aws.Int64(2),
			ChannelsOut: aws.Int64(1),
		},
		CodecSettings: &mediaconvert.AudioCodecSettings{
			Codec: mediaconvert.AudioCodecWav,
			WavSettings: &mediaconvert.WavSettings{
				BitDepth:   aws.Int64(24),
				Channels:   aws.Int64(1),
				SampleRate: aws.Int64(48000),
				Format:     "RIFF",
			},
		},
	}, {
		RemixSettings: &mediaconvert.RemixSettings{
			ChannelMapping: &mediaconvert.ChannelMapping{
				OutputChannels: []mediaconvert.OutputChannelMapping{{
					InputChannels: []int64{-60, 0},
				},
				}},
			ChannelsIn:  aws.Int64(2),
			ChannelsOut: aws.Int64(1),
		},
		CodecSettings: &mediaconvert.AudioCodecSettings{
			Codec: mediaconvert.AudioCodecWav,
			WavSettings: &mediaconvert.WavSettings{
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
