package hybrik

import (
	"fmt"

	hy "github.com/cbsinteractive/hybrik-sdk-go"
	"github.com/cbsinteractive/transcode-orchestrator/job"
)

func (p *driver) audioElements(j *job.Job) (a []hy.Element) {
	tags := tag(j, job.ComputeClassTranscodeDefault)
	for i, f := range j.Output.File {
		if f.Audio.Codec == "" {
			continue
		}
		file := fmt.Sprintf("audio_output_%d.%s", i, f.Audio.Codec)
		uid := fmt.Sprintf("audio_%d", i)
		//TODO(as): make sure file name/dir is correct here
		a = append(a, label(p.audioElement(f.Kid(file)), uid, tag))
	}
	return a
}

func (p *driver) audioElement(f job.File, idx int) hy.Element {
	return hy.Element{
		Kind: "transcode",
		Task: &hy.ElementTaskOptions{
			Name: "Audio Encode",
		},
		Payload: hy.LocationTargetPayload{
			Location: p.location(f, creds),
			Targets: []hy.TranscodeTarget{{
				FilePattern:   f.Name, // TODO(as)
				ExistingFiles: "replace",
				Container: hy.TranscodeContainer{
					Kind: "elementary",
				},
				Audio: []hy.AudioTarget{{
					Codec:     a.Codec,
					BitrateKb: a.Bitrate / 1000,
					Channels:  2,
					Source:    []hy.AudioTargetSource{{TrackNum: 0}},
				}},
			}},
		},
	}
}
