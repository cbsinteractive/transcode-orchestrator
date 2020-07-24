package mediaconvert

import (
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/cbsinteractive/transcode-orchestrator/db"
	"github.com/cbsinteractive/transcode-orchestrator/service"
)

type mapSettings map[bool]int64

var (
	channelSet = mapSettings{
		true:  0,
		false: -60,
	}
)

func audioSelectorFrom(audioDownmix *db.AudioDownmix, audioSelector *mediaconvert.AudioSelector) error {
	audioSelector.Offset = aws.Int64(0)
	audioSelector.ProgramSelection = aws.Int64(1)
	audioSelector.SelectorType = mediaconvert.AudioSelectorTypeTrack
	audioSelector.Tracks = trackListFrom(*audioDownmix)

	channelMapping, err := audioChannelMappingFrom(*audioDownmix)
	if err != nil {
		return err
	}

	audioSelector.RemixSettings = &mediaconvert.RemixSettings{
		ChannelsIn:     aws.Int64(int64(len(audioDownmix.SrcChannels))),
		ChannelsOut:    aws.Int64(int64(len(audioDownmix.DestChannels))),
		ChannelMapping: channelMapping,
	}

	return nil
}

func trackListFrom(audioDownmix db.AudioDownmix) (tracks []int64) {
	uniqueTracks := make(map[int]struct{})

	for _, channel := range audioDownmix.SrcChannels {
		if _, found := uniqueTracks[channel.TrackIdx]; !found {
			tracks = append(tracks, int64(channel.TrackIdx))
			uniqueTracks[channel.TrackIdx] = struct{}{}
		}
	}

	return tracks
}

func audioChannelMappingFrom(audioDownmix db.AudioDownmix) (*mediaconvert.ChannelMapping, error) {
	var outputChannelMapping []mediaconvert.OutputChannelMapping

	mapping, err := service.AudioDownmixMapping(audioDownmix)
	if err != nil {
		return &mediaconvert.ChannelMapping{}, err
	}

	for _, channel := range mapping {
		var outputChannel []int64

		for _, setting := range channel {
			outputChannel = append(outputChannel, channelSet[setting])
		}

		outputChannelMapping = append(outputChannelMapping,
			mediaconvert.OutputChannelMapping{InputChannels: outputChannel})
	}

	return &mediaconvert.ChannelMapping{
		OutputChannels: outputChannelMapping,
	}, nil
}
