package av

import (
	"fmt"
)

//ChannelLayout describes layout of an audio channel
type ChannelLayout string

const (
	ChannelLayoutCenter        ChannelLayout = "C"
	ChannelLayoutLeft          ChannelLayout = "L"
	ChannelLayoutRight         ChannelLayout = "R"
	ChannelLayoutLeftSurround  ChannelLayout = "Ls"
	ChannelLayoutRightSurround ChannelLayout = "Rs"
	ChannelLayoutLeftBack      ChannelLayout = "Lb"
	ChannelLayoutRightBack     ChannelLayout = "Rb"
	ChannelLayoutLeftTotal     ChannelLayout = "Lt"
	ChannelLayoutRightTotal    ChannelLayout = "Rt"
	ChannelLayoutLFE           ChannelLayout = "LFE"
)

// Audio defines audio transcoding parameters
type Audio struct {
	Codec     string `json:"codec,omitempty"`
	Bitrate   int    `json:"bitrate,omitempty"`
	Normalize bool   `json:"normalize,omitempty"`
	Discrete  bool   `json:"discrete,omitempty"`
}

func (a *Audio) On() bool {
	return a != nil && *a != (Audio{})
}

// AudioChannel describes the position and attributes of a
// single channel of audio inside a container
type AudioChannel struct {
	TrackIdx, ChannelIdx int
	Layout               string
}

// Downmix contains two sets of audio channels. The goal is to map
// the channels in Src to the channels in Dst. Currently, we only use
// this for stereo downmixing, meaning, Dst will look like a common
// stereo track.
//
// This object is a bit complex because it stores both Src and Dst
// channel lists. It would probably make more sense to have the
// Audio object hold the channel list, and create the mapping by
// modifying the Map (which we would move to AudioChannel) to
// take a destination AudioChannel.
//
// i.e., AudioChannel.Map(dst AudioChannel) ([][]bool, error)
//
// with the calling convention
//
// src := j.Input.Audio.Channel
// dst := j.Output[0].Audio.Channel
// src.Map(dst)
//
type Downmix struct {
	Src []AudioChannel
	Dst []AudioChannel
}

// Map converts
func (ad *Downmix) Map() ([][]bool, error) {
	dm, found := downmixMappings[len(ad.Dst)]
	if !found {
		return nil, fmt.Errorf("no downmix config found when converting %d src channels to %d destination channels",
			len(ad.Src), len(ad.Dst))
	}

	m := make([][]bool, len(ad.Dst))
	for i := range m {
		m[i] = make([]bool, len(ad.Src))
	}

	for srcIdx, src := range ad.Src {
		for destIdx, enabled := range dm[ChannelLayout(src.Layout)] {
			m[destIdx][srcIdx] = enabled
		}
	}

	return m, nil
}

var downmixMappings = map[int]downmixMap{
	2: {
		ChannelLayoutLFE:           nil,
		ChannelLayoutCenter:        {0: true, 1: true},
		ChannelLayoutLeft:          {0: true},
		ChannelLayoutRight:         {1: true},
		ChannelLayoutLeftSurround:  {0: true},
		ChannelLayoutRightSurround: {1: true},
		ChannelLayoutLeftBack:      {0: true},
		ChannelLayoutRightBack:     {1: true},
		ChannelLayoutLeftTotal:     {0: true},
		ChannelLayoutRightTotal:    {1: true},
	},
}

type downmixMap map[ChannelLayout]map[int]bool
