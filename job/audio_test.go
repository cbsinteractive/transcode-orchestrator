package job

import (
	"reflect"
	"testing"
)

func TestAudioDownmixMapping(t *testing.T) {
	tests := []struct {
		name    string
		ad      Downmix
		want    [][]bool
		wantErr bool
	}{
		{
			name: "5.1Source",
			ad: Downmix{
				Src: []AudioChannel{
					{TrackIdx: 1, ChannelIdx: 1, Layout: "L"},
					{TrackIdx: 2, ChannelIdx: 1, Layout: "R"},
					{TrackIdx: 3, ChannelIdx: 1, Layout: "C"},
					{TrackIdx: 4, ChannelIdx: 1, Layout: "LFE"},
					{TrackIdx: 5, ChannelIdx: 1, Layout: "Ls"},
					{TrackIdx: 6, ChannelIdx: 1, Layout: "Rs"},
				},
				Dst: []AudioChannel{
					{TrackIdx: 1, ChannelIdx: 1, Layout: "L"},
					{TrackIdx: 1, ChannelIdx: 2, Layout: "R"},
				},
			},
			want: [][]bool{
				{true, false, true, false, true, false},
				{false, true, true, false, false, true},
			},
		},
		{
			name: "DestinationChannelsNotStereo",
			ad: Downmix{
				Src: []AudioChannel{
					{TrackIdx: 1, ChannelIdx: 1, Layout: "L"},
					{TrackIdx: 2, ChannelIdx: 1, Layout: "R"},
					{TrackIdx: 3, ChannelIdx: 1, Layout: "C"},
					{TrackIdx: 4, ChannelIdx: 1, Layout: "LFE"},
					{TrackIdx: 5, ChannelIdx: 1, Layout: "Ls"},
					{TrackIdx: 6, ChannelIdx: 1, Layout: "Rs"},
				},
				Dst: []AudioChannel{
					{TrackIdx: 1, ChannelIdx: 1, Layout: "L"},
					{TrackIdx: 1, ChannelIdx: 2, Layout: "R"},
					{TrackIdx: 1, ChannelIdx: 2, Layout: "C"},
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ad.Map()
			if (err != nil) != tt.wantErr {
				t.Errorf("AudioDownmixMapping() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AudioDownmixMapping() = %v, want %v", got, tt.want)
			}
		})
	}
}
