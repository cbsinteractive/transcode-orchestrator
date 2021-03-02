package hybrik

import (
	"fmt"

	"github.com/cbsinteractive/hybrik-sdk-go"
	"github.com/cbsinteractive/transcode-orchestrator/job"
)

func mute(j Job) *Job {
	for i := range j.Output.File {
		j.Output.File[i].Audio = job.Audio{}
	}
	return &j
}

func (p *hybrikProvider) transcodeElems(j *Job) (e []hybrik.Element) {
	for i, f := range j.Output.File {
		target := hybrik.TranscodeTarget{
			FilePattern:   f.Name,
			ExistingFiles: "replace",
			Container:     hybrik.TranscodeContainer{Kind: container(f)},
			NumPasses:     passes(f),
			Video:         videoTarget(f.Video),
			Audio:         audioTarget(f.Audio),
		}
		applyHDR(target.Video, f)
		applyMXF(&target, f)
		e = append(e, hybrik.Element{
			Kind: elementKindTranscode,
			UID:  fmt.Sprintf(transcodeElementIDTemplate, i),
			Task: &hybrik.ElementTaskOptions{
				Name: fmt.Sprintf("Transcode - %s", f.Name),
				Tags: tag(j, job.ComputeClassTranscodeDefault),
			},
			Payload: hybrik.TranscodePayload{
				SourcePipeline: hybrik.TranscodeSourcePipeline{SegmentedRendering: features(j)},
				LocationTargetPayload: hybrik.LocationTargetPayload{
					Location: p.location(f, auth(j).Write),
					Targets:  []hybrik.TranscodeTarget{target},
				},
			},
		})
	}
	return e
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

const (
	computeTagMezzQCDefault = "preproc"
)

func (p *hybrikProvider) dolbyVisionMezzQC(j *Job) hybrik.Element {
	tag := tag(j, "job.ComputeClassDolbyVisionPreprocess", "preproc")
	return hybrik.Element{
		UID: "mezzanine_qc", Kind: "dolby_vision",
		Task: &hybrik.ElementTaskOptions{Name: "Mezzanine QC", Tags: tag},
		Payload: hybrik.DoViV2MezzanineQCPayload{
			Module: "mezzanine_qc",
			Params: hybrik.DoViV2MezzanineQCPayloadParams{
				Location:    p.location(f.Location("mezzanine_qc"), j.auth.Write),
				FilePattern: fmt.Sprintf("%s_mezz_qc_report.txt", j.ID),
			},
		},
	}
}
