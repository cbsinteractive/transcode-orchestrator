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

func TestVideoMapping(t *testing.T) {
	i64 := func(i int64) *int64 { return &i }

	for _, tt := range []struct {
		name   string
		video  db.VideoPreset
		assert func(t *testing.T, desc *mediaconvert.VideoDescription)
	}{
		{
			name: "Crop",
			video: db.VideoPreset{
				Bitrate: "12000", GopSize: "2", Width: "300", Height: "150", Codec: "h264",
				Crop: video.Crop{
					Top:    10,
					Bottom: 20,
					Right:  40,
					Left:   50,
				},
			},
			assert: func(t *testing.T, got *mediaconvert.VideoDescription) {
				want := &mediaconvert.Rectangle{
					Height: i64(120),
					Width:  i64(210),
					X:      i64(50),
					Y:      i64(10),
				}
				if !reflect.DeepEqual(got.Crop, want) {
					t.Errorf("bad crop rect:\nhave: %+v\nwant: %+v", got.Crop, want)
				}
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			have, err := videoPresetFrom(db.Preset{Video: tt.video}, db.File{})
			if err != nil {
				t.Fatal(err)
			}
			if tt.assert != nil {
				tt.assert(t, have)
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
