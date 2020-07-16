package mediaconvert

import (
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/cbsinteractive/transcode-orchestrator/db"
)

const (
	muteChannel   int64  = 0
	setChannel    int64  = -60
	lfeChannel    string = "LFE"
	centerChannel string = "C"
	leftChannel   string = "L"
	rightChannel  string = "R"
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
			ChannelMapping: p.audioChannelMappingFrom(audioDownmix),
		}
	}

	return map[string]mediaconvert.AudioSelector{
		"Audio Selector 1": audioSelector,
	}
}

func (p *mcProvider) audioChannelMappingFrom(audioDownmix db.AudioDownmix) *mediaconvert.ChannelMapping {
	var outputChannelMapping []mediaconvert.OutputChannelMapping

	for _, dest := range audioDownmix.DestChannels {
		var outputChannel []int64

		for _, src := range audioDownmix.SrcChannels {
			if src.Layout == lfeChannel {
				outputChannel = append(outputChannel, muteChannel)
				continue
			}

			for _, l := range src.Layout {
				layout := string(l)

				if layout != centerChannel && layout != rightChannel && layout != leftChannel {
					continue
				}

				if layout == centerChannel || layout == dest.Layout {
					outputChannel = append(outputChannel, setChannel)
					continue
				}

				outputChannel = append(outputChannel, muteChannel)
			}
		}

		outputChannelMapping = append(outputChannelMapping,
			mediaconvert.OutputChannelMapping{InputChannels: outputChannel})
	}

	return &mediaconvert.ChannelMapping{
		OutputChannels: outputChannelMapping,
	}
}
