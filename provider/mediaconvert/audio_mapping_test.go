package mediaconvert

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/cbsinteractive/transcode-orchestrator/db"
)

func Test_mcProvider_audioSelectorsFrom(t *testing.T) {
	tests := []struct {
		name         string
		audioDownmix db.AudioDownmix
		want         map[string]mediaconvert.AudioSelector
	}{
		{
			name:         "when audio downmix is not set",
			audioDownmix: db.AudioDownmix{},
			want: map[string]mediaconvert.AudioSelector{
				"Audio Selector 1": mediaconvert.AudioSelector{
					DefaultSelection: mediaconvert.AudioDefaultSelectionDefault,
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			p := &mcProvider{}
			if got := p.audioSelectorsFrom(tc.audioDownmix); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got:\n%v\nwant:\n%v", got, tc.want)
			}
		})
	}
}

func Test_mcProvider_stereoAudioChannelMappingFrom(t *testing.T) {
	tests := []struct {
		name    string
		audioDM db.AudioDownmix
		want    *mediaconvert.ChannelMapping
	}{
		{
			name: "5.1",
			audioDM: db.AudioDownmix{
				SrcChannels: []db.AudioChannel{
					{TrackIdx: 1, ChannelIdx: 1, Layout: "L"},
					{TrackIdx: 1, ChannelIdx: 2, Layout: "R"},
					{TrackIdx: 1, ChannelIdx: 3, Layout: "C"},
					{TrackIdx: 1, ChannelIdx: 4, Layout: "LFE"},
					{TrackIdx: 1, ChannelIdx: 5, Layout: "Ls"},
					{TrackIdx: 1, ChannelIdx: 6, Layout: "Rs"},
				},
				DestChannels: []db.AudioChannel{
					{TrackIdx: 1, ChannelIdx: 1, Layout: "L"},
					{TrackIdx: 1, ChannelIdx: 2, Layout: "R"},
				},
			},
			want: &mediaconvert.ChannelMapping{
				OutputChannels: []mediaconvert.OutputChannelMapping{
					{InputChannels: []int64{-60, 0, -60, 0, -60, 0}},
					{InputChannels: []int64{0, -60, -60, 0, 0, -60}},
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			p := &mcProvider{}

			if got := p.audioChannelMappingFrom(tc.audioDM); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got\n%v\nwant\n %v", got, tc.want)
			}
		})
	}
}
