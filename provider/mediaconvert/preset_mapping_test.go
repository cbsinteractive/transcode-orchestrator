package mediaconvert

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/aws/aws-sdk-go/aws"
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
