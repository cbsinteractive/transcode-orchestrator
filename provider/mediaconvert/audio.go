package mediaconvert

import (
	mc "github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/cbsinteractive/transcode-orchestrator/client/transcoding/job"
)

type mapSettings map[bool]int64

var (
	channelSet = mapSettings{
		true:  0,
		false: -60,
	}
)

func audioSelectorFrom(ad *job.Downmix, s *mc.AudioSelector) error {
	if ad == nil {
		return nil
	}
	s.Offset = aws.Int64(0)
	s.ProgramSelection = aws.Int64(1)
	s.SelectorType = mc.AudioSelectorTypeTrack
	s.Tracks = uniqueTracks(ad)

	cmap, err := audioChannelMapping(ad)
	if err != nil {
		return err
	}

	s.RemixSettings = &mc.RemixSettings{
		ChannelsIn:     aws.Int64(int64(len(ad.Src))),
		ChannelsOut:    aws.Int64(int64(len(ad.Dst))),
		ChannelMapping: cmap,
	}

	return nil
}

func uniqueTracks(ad *job.Downmix) (tracks []int64) {
	seen := map[int64]bool{}

	for _, channel := range ad.Src {
		idx := int64(channel.TrackIdx)
		if !seen[idx] {
			seen[idx] = true
			tracks = append(tracks, idx)
		}
	}

	return tracks
}

func audioChannelMapping(ad *job.Downmix) (*mc.ChannelMapping, error) {
	var out []mc.OutputChannelMapping

	mapping, err := ad.Map()
	if err != nil {
		return &mc.ChannelMapping{}, err
	}

	for _, channel := range mapping {
		var gain []int64
		for _, on := range channel {
			gain = append(gain, channelSet[on])
		}
		out = append(out, mc.OutputChannelMapping{InputChannels: gain})
	}

	return &mc.ChannelMapping{
		OutputChannels: out,
	}, nil
}
