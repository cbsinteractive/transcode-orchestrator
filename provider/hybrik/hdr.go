package hybrik

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/NYTimes/video-transcoding-api/db"
	hwrapper "github.com/cbsinteractive/hybrik-sdk-go"
)

type hdrType = string

const (
	hdrTypeHDR10       hdrType = "hdr10"
	hdrTypeDolbyVision hdrType = "dolbyVision"

	h265Codec                  = "h265"
	h265CodecProfileMain10     = "main10"
	h265VideoTagValueHVC1      = "hvc1"
	h265DolbyVisionArgsDefualt = "vbv-init=0.6:vbv-end=0.6:annexb=1:hrd=1:aud=1:videoformat=5:range=full" +
		":colorprim=2:transfer=2:colormatrix=2:rc-lookahead=48:qg-size=32:scenecut=0:no-open-gop=1:" +
		"frame-threads=0:repeat-headers=1:nr-inter=400:nr-intra=100:psy-rd=0:cbqpoffs=0:crqpoffs=3"

	ffmpegStrictTypeExperimental = "experimental"

	containerKindElementary = "elementary"

	chromaFormatYUV420P10LE = "yuv420p10le"
	colorPrimaryBT2020      = "bt2020"
	colorMatrixBT2020NC     = "bt2020nc"
	colorTRCSMPTE2084       = "st2084"
)

func enrichPresetWithHDRMetadata(hybrikPreset hwrapper.Preset, preset db.Preset) (hwrapper.Preset, error) {
	hdr, hdrEnabled := hdrTypeFromPreset(preset)
	if !hdrEnabled {
		return hybrikPreset, nil
	}

	for idx, target := range hybrikPreset.Payload.Targets {
		if target.Video.Codec == h265Codec {
			if target.Video.Profile == "" {
				target.Video.Profile = h265CodecProfileMain10
			} else if target.Video.Profile != h265CodecProfileMain10 {
				return hwrapper.Preset{}, fmt.Errorf("for HDR content encoded with h265, " +
					"the codec profile must be main10")
			}

			target.Video.FFMPEGArgs = fmt.Sprintf("-tag:v %s", h265VideoTagValueHVC1)
		}

		hybrikPreset.Payload.Options = &hwrapper.TranscodeTaskOptions{
			Pipeline: &hwrapper.PipelineOptions{
				EncoderVersion: hwrapper.EncoderVersion4_10bit,
			},
		}

		switch hdr {
		case hdrTypeHDR10:
			enrichedVideoTarget, err := enrichVideoTargetWithHDR10Metadata(target.Video, preset.Video.HDR10Settings)
			if err != nil {
				return hybrikPreset, err
			}

			hybrikPreset.Payload.Targets[idx].Video = enrichedVideoTarget
		case hdrTypeDolbyVision:
			// signal this is a dolby vision preset inside the custom user data
			b, err := json.Marshal(customPresetData{DolbyVisionEnabled: true})
			if err != nil {
				return hwrapper.Preset{}, err
			}
			hybrikPreset.UserData = string(b)

			// append ffmpeg `-strict` arg
			target.Video.FFMPEGArgs = fmt.Sprintf("%s -strict %s", target.Video.FFMPEGArgs,
				ffmpegStrictTypeExperimental)

			// add dolby vision x265 options if using h.265
			if target.Video.Codec == h265Codec {
				target.Video.X265Options = h265DolbyVisionArgsDefualt
			}

			// hybrik needs this format to feed into the DoVi mp4 muxer
			target.Container.Kind = containerKindElementary

			// zero out audio targets, we're processing video only
			target.Audio = []hwrapper.AudioTarget{}

			// set the enriched target back onto the preset
			hybrikPreset.Payload.Targets[idx] = target
		}
	}

	return hybrikPreset, nil
}

func hdrTypeFromPreset(preset db.Preset) (hdrType, bool) {
	if preset.Video.HDR10Settings.Enabled {
		return hdrTypeHDR10, true
	} else if preset.Video.DolbyVisionSettings.Enabled {
		return hdrTypeDolbyVision, true
	}

	return "", false
}

func enrichVideoTargetWithHDR10Metadata(video hwrapper.VideoTarget, hdr10 db.HDR10Settings) (hwrapper.VideoTarget, error) {
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

	return video, nil
}

func dolbyVisionEnabledOnAllPresets(cfgs map[string]outputCfg) (bool, error) {
	record := struct{ doViPresetFound, nonDoViPresetFound bool }{}

	for _, cfg := range cfgs {
		if cfg.preset.UserData == "" {
			record.nonDoViPresetFound = true
			continue
		}

		var customData customPresetData
		err := json.Unmarshal([]byte(cfg.preset.UserData), &customData)
		if err != nil {
			record.nonDoViPresetFound = true
			continue
		}

		if !customData.DolbyVisionEnabled {
			record.nonDoViPresetFound = true
		} else {
			record.doViPresetFound = true
		}
	}

	var mixedPresetsErr error
	if record.doViPresetFound && record.nonDoViPresetFound {
		mixedPresetsErr = errors.New("found presets containing a mix of DolbyVision and non DolbyVision presets, " +
			"this is not supported at this time")
	}

	return record.doViPresetFound && !record.nonDoViPresetFound, mixedPresetsErr
}

func hdr10TranscodePayloadModifier(transcodePayload hwrapper.TranscodePayload) hwrapper.TranscodePayload {
	// hybrik has a bug where HDR10 jobs break when run with segmented rendering
	// this disables configurations from enabling this feature until they fix it TODO (TS)
	transcodePayload.SourcePipeline.SegmentedRendering = nil

	return transcodePayload
}
