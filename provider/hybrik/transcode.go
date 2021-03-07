package hybrik

import (
	"fmt"

	hy "github.com/cbsinteractive/hybrik-sdk-go"
	"github.com/cbsinteractive/transcode-orchestrator/job"
)

const SourceUID = "source_file"

func (p *driver) auth(j *Job) (a struct{ Read, Write string }) {
	a.Read = j.Env.InputAlias
	a.Write = j.Env.OutputAlias
	return
}

func (p *driver) validate(j *Job) error {
	n := countDolbyVision(&j.Output)
	if n > 0 && n != j.Output.Len() {
		return ErrMixedPresets
	}
	return nil
}

const LegacyDolbyVision = true

// assemble converts the job into a matrix of elements
// callers should ensure the job was already validated
// using p.validate
func (p *driver) assemble(j *Job) [][]hy.Element {
	if countDolbyVision(&j.Output) == 0 {
		return [][]hy.Element{p.transcodeElems(j)}
	}
	if LegacyDolbyVision {
		// NOTE(as): the original comment reads:
		// "switch back to this once Hybrik fixes bug with GCP jobs hanging"
		return p.dolbyVisionLegacy(j)
	}
	return p.dolbyVisionJob(j)
}

func (p *driver) srcFrom(j *Job) hy.Element {
	creds := j.Env.InputAlias
	assets := []hy.AssetPayload{p.asset(&j.Input, creds)}

	if dolby := j.Asset(job.TagDolbyVisionMetadata); dolby != nil {
		assets = append(assets, p.asset(dolby, creds, hy.AssetContents{
			Kind:    "metadata",
			Payload: hy.AssetContentsPayload{Standard: "dolbyvision_metadata"},
		}))
	}
	return hy.Element{
		UID:  SourceUID,
		Kind: "source",
		Payload: hy.ElementPayload{
			Kind:    "asset_urls",
			Payload: assets,
		},
	}
}

func (p *driver) transcodeElems(j *Job) (e []hy.Element) {
	for i, f := range j.Output.File {
		f = j.Abs(f)
		target := hy.TranscodeTarget{
			FilePattern:   f.Base(),
			ExistingFiles: "replace",
			Container:     hy.TranscodeContainer{Kind: p.container(f)}, //TODO(as): validation
			NumPasses:     passes(f),
			Video:         videoTarget(f.Video),
			Audio:         audioTarget(f.Audio),
		}
		var opts *hy.TranscodeTaskOptions
		if applyHDR(&target, f) {
			opts = &hy.TranscodeTaskOptions{Pipeline: &hy.PipelineOptions{EncoderVersion: hy.EncoderVersion4_10bit}}
		}
		if j.Input.Type() == "mxf" {
			applyMXF(&target, f)
		}
		e = append(e, hy.Element{
			Kind: "transcode",
			UID:  fmt.Sprintf("transcode_task_%d", i),
			Task: &hy.ElementTaskOptions{
				Name: fmt.Sprintf("Transcode - %s", f.Base()),
				Tags: tag(j, job.TagTranscodeDefault),
			},
			Payload: hy.TranscodePayload{
				Options:        opts,
				SourcePipeline: hy.TranscodeSourcePipeline{SegmentedRendering: features(j)},
				LocationTargetPayload: hy.LocationTargetPayload{
					Location: p.location(f, p.auth(j).Write),
					Targets:  []hy.TranscodeTarget{target},
				},
			},
		})
	}
	return e
}

func videoTarget(v job.Video) *hy.VideoTarget {
	if !v.On() {
		return nil
	}

	var frames, seconds int
	if v.Gop.Seconds() {
		seconds = int(v.Gop.Size)
	} else {
		frames = int(v.Gop.Size)
	}

	// TODO: Understand video-transcoding-api profile + level settings in relation to vp8
	// For now, we will omit and leave to encoder defaults
	if canon(v.Codec) == "vp8" {
		v.Profile = ""
		v.Level = ""
	}

	w, h := &v.Width, &v.Height
	if *w == 0 {
		w = nil
	}
	if *h == 0 {
		h = nil
	}
	vbrON := RateControl[canon(v.Bitrate.Control)]
	return &hy.VideoTarget{
		Codec:             v.Codec,
		Width:             w,
		Height:            h,
		BitrateKb:         v.Bitrate.Kbps(),
		MinBitrateKb:      vbrON * v.Bitrate.Percent(-vbrVariability).Kbps(),
		MaxBitrateKb:      vbrON * v.Bitrate.Percent(+vbrVariability).Kbps(),
		BitrateMode:       canon(v.Bitrate.Control),
		Profile:           canon(v.Profile),
		Level:             canon(v.Level),
		ExactGOPFrames:    frames,
		ExactKeyFrame:     seconds,
		InterlaceMode:     v.Scantype,
		Preset:            "slow",
		ChromaFormat:      chromaFormatYUV420P,
		UseSceneDetection: false,
	}
}

func audioTarget(a job.Audio) []hy.AudioTarget {
	if (a == job.Audio{}) {
		return []hy.AudioTarget{}
	}
	return []hy.AudioTarget{
		{
			Codec:     a.Codec,
			Channels:  2,
			BitrateKb: a.Bitrate / 1000,
		},
	}
}

func mute(j Job) *Job {
	for i := range j.Output.File {
		j.Output.File[i].Audio = job.Audio{}
	}
	return &j
}

// TODO(as): should canonicalize this across all providers
func (p *driver) container(f job.File) string {
	for _, c := range p.Capabilities().OutputFormats {
		if c == f.Type() {
			return c
		}
	}
	return ""
}
func passes(f job.File) int {
	if f.Video.Bitrate.TwoPass {
		return 2
	}
	return 1
}

func label(e hy.Element, uid string, tags ...string) hy.Element {
	e.UID = uid
	for _, t := range tags {
		if len(t) > 0 {
			e.Task.Tags = append(e.Task.Tags, t)
		}
	}
	return e
}
