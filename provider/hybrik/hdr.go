package hybrik

import (
	"errors"
	"fmt"

	hwrapper "github.com/cbsinteractive/hybrik-sdk-go"
	"github.com/cbsinteractive/transcode-orchestrator/db"
)

type hdrType = string

const (
	hdrTypeHDR10       hdrType = "hdr10"
	hdrTypeDolbyVision hdrType = "dolbyVision"

	h265Codec                  = "h265"
	h265CodecProfileMain10     = "main10"
	h265VideoTagValueHVC1      = "hvc1"
	h265DolbyVisionArgsDefualt = "concatenation={auto_concatenation_flag}:vbv-init=0.6:vbv-end=0.6:annexb=1:hrd=1:" +
		"aud=1:videoformat=5:range=full:colorprim=2:transfer=2:colormatrix=2:rc-lookahead=48:qg-size=32:scenecut=0:" +
		"no-open-gop=1:frame-threads=0:repeat-headers=1:nr-inter=400:nr-intra=100:psy-rd=0:cbqpoffs=0:crqpoffs=3"

	ffmpegStrictTypeExperimental = "experimental"

	containerKindElementary = "elementary"

	chromaFormatYUV420P10LE = "yuv420p10le"
	chromaFormatYUV420P     = "yuv420p"
	colorPrimaryBT2020      = "bt2020"
	colorMatrixBT2020NC     = "bt2020nc"
	colorTRCSMPTE2084       = "st2084"

	presetSlow = "slow"

	tuneGrain = "grain"
)

func enrichTranscodePayloadWithHDRMetadata(payload hwrapper.TranscodePayload, preset db.Preset) (hwrapper.TranscodePayload, error) {
	hdr, hdrEnabled := hdrTypeFromPreset(preset)
	if !hdrEnabled {
		return payload, nil
	}

	transcodeTargets, ok := payload.Targets.([]hwrapper.TranscodeTarget)
	if !ok {
		return hwrapper.TranscodePayload{}, fmt.Errorf("targets are not TranscodeTargets: %v", payload.LocationTargetPayload.Targets)
	}

	for idx, target := range transcodeTargets {
		if target.Video.Codec == h265Codec {
			if target.Video.Profile == "" {
				target.Video.Profile = h265CodecProfileMain10
			} else if target.Video.Profile != h265CodecProfileMain10 {
				return hwrapper.TranscodePayload{}, fmt.Errorf("for HDR content encoded with h265, " +
					"the codec profile must be main10")
			}

			target.Video.Tune = tuneGrain
			target.Video.VTag = h265VideoTagValueHVC1
		}

		payload.Options = &hwrapper.TranscodeTaskOptions{
			Pipeline: &hwrapper.PipelineOptions{
				EncoderVersion: hwrapper.EncoderVersion4_10bit,
			},
		}

		switch hdr {
		case hdrTypeHDR10:
			enrichVideoTargetWithHDR10Metadata(target.Video, preset.Video.HDR10Settings)
		case hdrTypeDolbyVision:
			// append ffmpeg `-strict` arg
			target.Video.FFMPEGArgs = fmt.Sprintf("%s -strict %s", target.Video.FFMPEGArgs,
				ffmpegStrictTypeExperimental)

			// add dolby vision x265 options if using h.265
			if target.Video.Codec == h265Codec {
				target.Video.X265Options = h265DolbyVisionArgsDefualt
			}

			target.Video.ChromaFormat = chromaFormatYUV420P10LE

			// hybrik needs this format to feed into the DoVi mp4 muxer
			target.Container.Kind = containerKindElementary

			// set the enriched target back onto the preset
			transcodeTargets[idx] = target
		}
	}
	payload.Targets = transcodeTargets

	return payload, nil
}

func hdrTypeFromPreset(preset db.Preset) (hdrType, bool) {
	if preset.Video.HDR10Settings.Enabled {
		return hdrTypeHDR10, true
	} else if preset.Video.DolbyVisionSettings.Enabled {
		return hdrTypeDolbyVision, true
	}

	return "", false
}

func enrichVideoTargetWithHDR10Metadata(video *hwrapper.VideoTarget, hdr10 db.HDR10Settings) {
	hdr10Metadata := &hwrapper.HDR10Settings{}

	// signal that the HDR metadata is being pulled from the source media
	hdr10Metadata.Source = "media"

	// if max content light level is set, add to hybrik HDR10 metadata
	if hdr10.MaxCLL != 0 {
		hdr10Metadata.MaxCLL = int(hdr10.MaxCLL)
	}

	// if max frame average light level is set, add to hybrik HDR10 metadata
	if hdr10.MaxFALL != 0 {
		hdr10Metadata.MaxFALL = int(hdr10.MaxFALL)
	}

	// if a mastering display config string is supplied, add to hybrik HDR10 metadata
	if hdr10.MasterDisplay != "" {
		hdr10Metadata.MasterDisplay = hdr10.MasterDisplay
	}

	video.HDR10 = hdr10Metadata

	video.ChromaFormat = chromaFormatYUV420P10LE
	video.ColorPrimaries = colorPrimaryBT2020
	video.ColorMatrix = colorMatrixBT2020NC
	video.ColorTRC = colorTRCSMPTE2084
}

func dolbyVisionEnabledOnAllPresets(cfgs map[string]outputCfg) (bool, error) {
	record := struct{ doViPresetFound, nonDoViPresetFound bool }{}

	for _, cfg := range cfgs {
		if enabled := cfg.localPreset.Video.DolbyVisionSettings.Enabled; enabled {
			record.doViPresetFound = true
		} else {
			record.nonDoViPresetFound = true
		}
	}

	var mixedPresetsErr error
	if record.doViPresetFound && record.nonDoViPresetFound {
		mixedPresetsErr = errors.New("found presets containing a mix of DolbyVision and non DolbyVision presets, " +
			"this is not supported at this time")
	}

	return record.doViPresetFound && !record.nonDoViPresetFound, mixedPresetsErr
}
