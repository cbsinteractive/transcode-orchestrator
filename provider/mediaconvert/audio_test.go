package mediaconvert

import (
	"reflect"
	"testing"

	mc "github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/cbsinteractive/transcode-orchestrator/client/transcoding/job"
	"github.com/google/go-cmp/cmp"
)

func TestAudio(t *testing.T) {
	audioSelectorKey := "Audio Selector 1"
	defaultSelectorMap := map[string]mc.AudioSelector{
		audioSelectorKey: {
			DefaultSelection: mc.AudioDefaultSelectionDefault,
		},
	}

	tests := []struct {
		name         string
		audioDownmix job.Downmix
		want         mc.AudioSelector
		wantErr      bool
	}{
		{
			name: "SingleTrack5.1",
			audioDownmix: job.Downmix{
				Src: []job.AudioChannel{
					{TrackIdx: 1, ChannelIdx: 1, Layout: "L"},
					{TrackIdx: 1, ChannelIdx: 2, Layout: "R"},
					{TrackIdx: 1, ChannelIdx: 3, Layout: "C"},
					{TrackIdx: 1, ChannelIdx: 4, Layout: "LFE"},
					{TrackIdx: 1, ChannelIdx: 5, Layout: "Ls"},
					{TrackIdx: 1, ChannelIdx: 6, Layout: "Rs"},
				},
				Dst: []job.AudioChannel{
					{TrackIdx: 1, ChannelIdx: 1, Layout: "L"},
					{TrackIdx: 1, ChannelIdx: 2, Layout: "R"},
				},
			},
			want: getAudioSelector(6, 2, []int64{1}, []mc.OutputChannelMapping{
				{InputChannels: []int64{0, -60, 0, -60, 0, -60}},
				{InputChannels: []int64{-60, 0, 0, -60, -60, 0}},
			}),
		},
		{
			name: "DiscreteTrack5.1",
			audioDownmix: job.Downmix{
				Src: []job.AudioChannel{
					{TrackIdx: 1, ChannelIdx: 1, Layout: "L"},
					{TrackIdx: 2, ChannelIdx: 1, Layout: "R"},
					{TrackIdx: 3, ChannelIdx: 1, Layout: "C"},
					{TrackIdx: 4, ChannelIdx: 1, Layout: "LFE"},
					{TrackIdx: 5, ChannelIdx: 1, Layout: "Ls"},
					{TrackIdx: 6, ChannelIdx: 1, Layout: "Rs"},
				},
				Dst: []job.AudioChannel{
					{TrackIdx: 1, ChannelIdx: 1, Layout: "L"},
					{TrackIdx: 1, ChannelIdx: 2, Layout: "R"},
				},
			},
			want: getAudioSelector(6, 2, []int64{1, 2, 3, 4, 5, 6}, []mc.OutputChannelMapping{
				{InputChannels: []int64{0, -60, 0, -60, 0, -60}},
				{InputChannels: []int64{-60, 0, 0, -60, -60, 0}},
			}),
		},
		{
			name: "DiscreteTrack7.1",
			audioDownmix: job.Downmix{
				Src: []job.AudioChannel{
					{TrackIdx: 1, ChannelIdx: 1, Layout: "L"},
					{TrackIdx: 2, ChannelIdx: 1, Layout: "R"},
					{TrackIdx: 3, ChannelIdx: 1, Layout: "C"},
					{TrackIdx: 4, ChannelIdx: 1, Layout: "LFE"},
					{TrackIdx: 5, ChannelIdx: 1, Layout: "Ls"},
					{TrackIdx: 6, ChannelIdx: 1, Layout: "Rs"},
					{TrackIdx: 7, ChannelIdx: 1, Layout: "Lb"},
					{TrackIdx: 8, ChannelIdx: 1, Layout: "Rb"},
				},
				Dst: []job.AudioChannel{
					{TrackIdx: 1, ChannelIdx: 1, Layout: "L"},
					{TrackIdx: 1, ChannelIdx: 2, Layout: "R"},
				},
			},
			want: getAudioSelector(8, 2, []int64{1, 2, 3, 4, 5, 6, 7, 8}, []mc.OutputChannelMapping{
				{InputChannels: []int64{0, -60, 0, -60, 0, -60, 0, -60}},
				{InputChannels: []int64{-60, 0, 0, -60, -60, 0, -60, 0}},
			}),
		},
		{
			name: "NotStereo",
			audioDownmix: job.Downmix{
				Src: []job.AudioChannel{
					{TrackIdx: 1, ChannelIdx: 1, Layout: "L"},
					{TrackIdx: 2, ChannelIdx: 1, Layout: "R"},
					{TrackIdx: 3, ChannelIdx: 1, Layout: "C"},
					{TrackIdx: 4, ChannelIdx: 1, Layout: "LFE"},
					{TrackIdx: 5, ChannelIdx: 1, Layout: "Ls"},
					{TrackIdx: 6, ChannelIdx: 1, Layout: "Rs"},
					{TrackIdx: 7, ChannelIdx: 1, Layout: "Lb"},
					{TrackIdx: 8, ChannelIdx: 1, Layout: "Rb"},
				},
				Dst: []job.AudioChannel{
					{TrackIdx: 1, ChannelIdx: 1, Layout: "L"},
					{TrackIdx: 1, ChannelIdx: 2, Layout: "R"},
					{TrackIdx: 1, ChannelIdx: 3, Layout: "C"},
				},
			},
			wantErr: true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			audioSelector := defaultSelectorMap[audioSelectorKey]
			err := audioSelectorFrom(&tc.audioDownmix, &audioSelector)
			if (err != nil) != tc.wantErr {
				t.Errorf("audioSelectorsFrom() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !tc.wantErr && !reflect.DeepEqual(audioSelector, tc.want) {
				t.Errorf("got:\n%v\nwant:\n%v\ndiff:%+v", audioSelector, tc.want, cmp.Diff(audioSelector, tc.want))
			}
		})
	}
}

func getAudioSelector(cIn int64, cOut int64, tracks []int64, oc []mc.OutputChannelMapping) mc.AudioSelector {
	return mc.AudioSelector{
		DefaultSelection: mc.AudioDefaultSelectionDefault,
		Offset:           aws.Int64(0),
		ProgramSelection: aws.Int64(1),
		SelectorType:     mc.AudioSelectorTypeTrack,
		Tracks:           tracks,
		RemixSettings: &mc.RemixSettings{
			ChannelsIn:  aws.Int64(cIn),
			ChannelsOut: aws.Int64(cOut),
			ChannelMapping: &mc.ChannelMapping{
				OutputChannels: oc,
			},
		},
	}
}
