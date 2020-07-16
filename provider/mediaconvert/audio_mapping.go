package mediaconvert

import (
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/cbsinteractive/transcode-orchestrator/db"
)

const (
	muteChannel int64 = 0
	setChannel  int64 = -60
)

func (p *mcProvider) audioSelectorsFrom(audioDownmix db.AudioDownmix) map[string]mediaconvert.AudioSelector {
	audioSelector := mediaconvert.AudioSelector{
		DefaultSelection: mediaconvert.AudioDefaultSelectionDefault,
	}

	if audioDownmix.IsSet() {
		var offset int64
		var programSelection int64 = 1
		var channelsIn int64 = int64(len(audioDownmix.SrcChannels))
		var channelsOut int64 = int64(len(audioDownmix.DestChannels))

		audioSelector.Offset = &offset
		audioSelector.SelectorType = mediaconvert.AudioSelectorTypeTrack
		audioSelector.ProgramSelection = &programSelection

		for i := 0; i < len(audioDownmix.SrcChannels); i++ {
			audioSelector.Tracks = append(audioSelector.Tracks, int64(i+1))
		}

		audioSelector.RemixSettings = &mediaconvert.RemixSettings{
			ChannelsIn:     &channelsIn,
			ChannelsOut:    &channelsOut,
			ChannelMapping: p.stereoAudioChannelMappingFrom(audioDownmix.SrcChannels),
		}
	}

	return map[string]mediaconvert.AudioSelector{
		"Audio Selector 1": audioSelector,
	}
}

func (p *mcProvider) stereoAudioChannelMappingFrom(audioChannels []db.AudioChannel) *mediaconvert.ChannelMapping {
	var leftChannel, rightChannel []int64

	for _, ac := range audioChannels {
		layout := db.ChannelLayout(ac.Layout)
		switch layout {
		case db.LFE:
			leftChannel = append(leftChannel, muteChannel)
			rightChannel = append(rightChannel, muteChannel)
		case db.Center:
			leftChannel = append(leftChannel, setChannel)
			rightChannel = append(rightChannel, setChannel)
		case db.Left, db.LeftBack, db.LeftSurround, db.LeftTotal:
			leftChannel = append(leftChannel, setChannel)
			rightChannel = append(rightChannel, muteChannel)
		case db.Right, db.RightBack, db.RightSurround, db.RightTotal:
			leftChannel = append(leftChannel, muteChannel)
			rightChannel = append(rightChannel, setChannel)
		}
	}

	return &mediaconvert.ChannelMapping{
		OutputChannels: []mediaconvert.OutputChannelMapping{
			mediaconvert.OutputChannelMapping{InputChannels: leftChannel},
			mediaconvert.OutputChannelMapping{InputChannels: rightChannel},
		},
	}
}
