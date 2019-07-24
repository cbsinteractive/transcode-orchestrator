package hybrik

import (
	"fmt"

	"github.com/NYTimes/video-transcoding-api/db"
	hwrapper "github.com/cbsinteractive/hybrik-sdk-go"
)

const (
	h265Codec              = "h265"
	h265CodecProfileMain10 = "main10"
	h265VideoTagValueHVC1  = "hvc1"

	chromaFormatYUV420P10LE = "yuv420p10le"
	colorPrimaryBT2020      = "bt2020"
	colorMatrixBT2020NC     = "bt2020nc"
	colorTRCSMPTE2084       = "st2084"
)

func enrichPresetWithHDRMetadata(hybrikPreset hwrapper.Preset, preset db.Preset) (hwrapper.Preset, error) {
	hdr10Enabled := preset.Video.HDR10Settings.Enabled

	for idx, target := range hybrikPreset.Payload.Targets {
		if hdr10Enabled {
			enrichedVideoTarget, err := enrichVideoTargetWithHDR10Metadata(target.Video, preset.Video.HDR10Settings)
			if err != nil {
				return hybrikPreset, err
			}

			hybrikPreset.Payload.Targets[idx].Video = enrichedVideoTarget

			// tell Hybrik to use an encoder that supports 10bit color
			hybrikPreset.Payload.Options = &hwrapper.TranscodeTaskOptions{
				Pipeline: &hwrapper.PipelineOptions{
					EncoderVersion: hwrapper.EncoderVersion4_10bit,
				},
			}
		}
	}

	return hybrikPreset, nil
}

func enrichVideoTargetWithHDR10Metadata(video hwrapper.VideoTarget, hdr10 db.HDR10Settings) (hwrapper.VideoTarget, error) {
	// ensure main10 is used if using HEVC
	if video.Codec == h265Codec {
		if video.Profile == "" {
			video.Profile = h265CodecProfileMain10
		} else if video.Profile != h265CodecProfileMain10 {
			return hwrapper.VideoTarget{}, fmt.Errorf("for HDR content encoded with h265, " +
				"the codec profile must be main10")
		}

		video.FFMPEGArgs = fmt.Sprintf("-tag:v %s", h265VideoTagValueHVC1)
	}

	hdr10Metadata := &hwrapper.HDR10Settings{}

	hdr10Metadata.Source = "media"

	if hdr10.MaxCLL != 0 {
		hdr10Metadata.MaxCLL = int(hdr10.MaxCLL)
	}

	if hdr10.MaxFALL != 0 {
		hdr10Metadata.MaxFALL = int(hdr10.MaxFALL)
	}

	if hdr10.MasterDisplay != "" {
		hdr10Metadata.MasterDisplay = hdr10.MasterDisplay
	}

	video.HDR10 = hdr10Metadata
	video.ChromaFormat = chromaFormatYUV420P10LE
	video.ColorPrimaries = colorPrimaryBT2020
	video.ColorMatrix = colorMatrixBT2020NC
	video.ColorTRC = colorTRCSMPTE2084

	return video, nil
}
