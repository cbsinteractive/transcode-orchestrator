package hybrik

import (
	"errors"
	"fmt"
	"strings"

	hwrapper "github.com/cbsinteractive/hybrik-sdk-go"
	"github.com/cbsinteractive/transcode-orchestrator/job"
)

const (
	h265Codec                  = "h265"
	h265CodecProfileMain10     = "main10"
	h265VideoTagValueHVC1      = "hvc1"
	h265DolbyVisionArgsDefualt = "concatenation={auto_concatenation_flag}:vbv-init=0.6:vbv-end=0.6:annexb=1:hrd=1:" +
		"aud=1:videoformat=5:range=full:colorprim=2:transfer=2:colormatrix=2:rc-lookahead=48:qg-size=32:scenecut=0:" +
		"no-open-gop=1:frame-threads=0:repeat-headers=1:nr-inter=400:nr-intra=100:psy-rd=0:cbqpoffs=0:crqpoffs=3"

	containerKindElementary = "elementary"

	chromaFormatYUV420P10LE = "yuv420p10le"
	chromaFormatYUV420P     = "yuv420p"
	colorPrimaryBT2020      = "bt2020"
	colorMatrixBT2020NC     = "bt2020nc"
	colorTRCSMPTE2084       = "st2084"

	presetSlow = "slow"

	tuneGrain = "grain"
)

var canon = strings.ToLower

func checkHDR(f job.File) error {
	if !hasHDR(f) || canon(f.Video.Codec) != h265Codec {
		return nil
	}
	p := canon(f.Video.Profile)
	if p == "" || p == "main10" {
		return nil
	}
	return fmt.Errorf("hdr: h265: profile must be main10")
}

func hasHDR(f job.File) bool {
	return f.Video.HDR10.Enabled || f.Video.DolbyVision.Enabled
}

/*
// HDR only
		payload.Options = &hwrapper.TranscodeTaskOptions{
			Pipeline: &hwrapper.PipelineOptions{
				EncoderVersion: hwrapper.EncoderVersion4_10bit,
			},
		}
*/

func applyHDR(v *hwrapper.VideoTarget, f job.File) bool {
	return applyDoVi(v, f.Video.DolbyVision) || applyHDR10(v, f.Video.HDR10)
}

func applyDoVi(v *hwrapper.VideoTarget, h job.DolbyVision) bool {
	if !h.Enabled {
		return false
	}
	if v.Codec == h265Codec {
		v.Profile = h265CodecProfileMain10
		v.Tune = tuneGrain
		v.VTag = h265VideoTagValueHVC1
		v.X265Options = h265DolbyVisionArgsDefualt
	}
	v.FFMPEGArgs = fmt.Sprintf("%s -strict experimental", v.FFMPEGArgs)
	v.ChromaFormat = chromaFormatYUV420P10LE
	// hybrik needs this format to feed into the DoVi mp4 muxer
	t.Container.Kind = containerKindElementary
	return true
}

func applyHDR10(t *hwrapper.TranscodeTarget, h job.HDR10) bool {
	if !h.Enabled {
		return false
	}
	if t.Video.Codec == h265Codec {
		t.Video.Profile = h265CodecProfileMain10
		t.Video.Tune = tuneGrain
		t.Video.VTag = h265VideoTagValueHVC1
	}
	t.Video.ChromaFormat = chromaFormatYUV420P10LE
	t.Video.ColorPrimaries = colorPrimaryBT2020
	t.Video.ColorMatrix = colorMatrixBT2020NC
	t.Video.ColorTRC = colorTRCSMPTE2084
	t.Video.HDR10 = &hwrapper.HDR10Settings{
		Source:        "media",
		MaxCLL:        h.MaxCLL,
		MaxFALL:       h.MaxFALL,
		MasterDisplay: h.MasterDisplay,
	}
}

func countDolbyVision(d *job.Dir) (enabled int) {
	for _, f := range d.File {
		if f.Video.DolbyVision.Enabled {
			enabled++
		}
	}
	return enabled
}

var ErrMixedPresets = errors.New("job contains inconsistent DolbyVision outputs")
