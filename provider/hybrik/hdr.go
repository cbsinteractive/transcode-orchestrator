package hybrik

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cbsinteractive/hybrik-sdk-go"
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

	tuneGrain               = "grain"
	computeTagMezzQCDefault = "preproc"
)

var ErrMixedPresets = errors.New("job contains inconsistent DolbyVision outputs")
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

func applyHDR(t *hwrapper.TranscodeTarget, f job.File) bool {
	if t.Video == nil {
		return false
	}
	if applyHDR10(t.Video, f.Video.HDR10) {
		return true
	}
	if applyDoVi(t.Video, f.Video.DolbyVision) {
		// hybrik needs this format to feed into the DoVi mp4 muxer
		t.Container.Kind = containerKindElementary
		return true
	}
	return false
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
	return true
}

func applyHDR10(v *hwrapper.VideoTarget, h job.HDR10) bool {
	if !h.Enabled {
		return false
	}
	if v.Codec == h265Codec {
		v.Profile = h265CodecProfileMain10
		v.Tune = tuneGrain
		v.VTag = h265VideoTagValueHVC1
	}
	v.ChromaFormat = chromaFormatYUV420P10LE
	v.ColorPrimaries = colorPrimaryBT2020
	v.ColorMatrix = colorMatrixBT2020NC
	v.ColorTRC = colorTRCSMPTE2084
	v.HDR10 = &hwrapper.HDR10Settings{
		Source:        "media",
		MaxCLL:        h.MaxCLL,
		MaxFALL:       h.MaxFALL,
		MasterDisplay: h.MasterDisplay,
	}
	return true
}

func (p *hybrikProvider) dolbyVisionJob(j *Job) (e [][]hybrik.Element) {
	// initialize our pre-transcode execution group with a mezz qc task
	// then add any extracted audio elements to the pre-transcode group
	// and add pre-transcode tasks as the first element in the pipeline
	// add all transcode tasks as the second element in the pipeline
	return [][]hybrik.Element{
		{p.dolbyVisionMezzQC(j)},
		p.audioElements(j),
		p.dolbyVisionTranscode(j),
	}
}

func (p *hybrikProvider) dolbyVisionMezzQC(j *Job) hybrik.Element {
	tag := tag(j, job.ComputeClassDolbyVisionPreprocess, "preproc")
	return hybrik.Element{
		UID: "mezzanine_qc", Kind: "dolby_vision",
		Task: &hybrik.ElementTaskOptions{Name: "Mezzanine QC", Tags: tag},
		Payload: hybrik.DoViV2MezzanineQCPayload{
			Module: "mezzanine_qc",
			Params: hybrik.DoViV2MezzanineQCPayloadParams{
				Location:    p.location(job.File{Name: j.Location("mezzanine_qc")}, auth(j).Write),
				FilePattern: fmt.Sprintf("%s_mezz_qc_report.txt", j.ID),
			},
		},
	}
}

func (p *hybrikProvider) dolbyVisionTranscode(j *Job) (e []hybrik.Element) {
	txcode := p.transcodeElems(mute(*j))
	tag := tag(j, job.ComputeClassDolbyVisionPreprocess, "preproc")

	for i, f := range j.Output.File {
		a := []hybrik.DoViMP4MuxElementaryStream{}
		if (f.Audio != job.Audio{}) {
			a = append(a, hybrik.DoViMP4MuxElementaryStream{
				AssetURL: p.assetURL(f, j.auth.write),
			})
		}

		e = append(e, hybrik.Element{
			UID:  fmt.Sprintf("dolby_vision_%d", i),
			Kind: "dolby_vision",
			Task: &hybrik.ElementTaskOptions{
				Name:              fmt.Sprintf("Encode #%d", i),
				Tags:              tag,
				SourceElementUIDs: []string{SourceUID},
				RetryMethod:       "fail",
			},
			Payload: hybrik.DolbyVisionV2TaskPayload{
				Module: "encoder", Profile: 5,
				Location: p.location(f, j.auth.write),
				Preprocessing: hybrik.DolbyVisionV2Preprocessing{
					Task: hybrik.TaskTags{Tags: tag},
				},
				Transcodes: []hybrik.Element{txcode[i]},
				PostTranscode: hybrik.DoViPostTranscode{
					Task: &hybrik.TaskTags{Tags: tag},
					MP4Mux: hybrik.DoViMP4Mux{
						Enabled:           true,
						FilePattern:       "{source_basename}.mp4",
						ElementaryStreams: a,
						CLIOptions:        map[string]string{"dvh1flag": ""},
					},
				},
			},
		})
	}

	return e
}

func countDolbyVision(d *job.Dir) (enabled int) {
	for _, f := range d.File {
		if f.Video.DolbyVision.Enabled {
			enabled++
		}
	}
	return enabled
}
