package hybrik

import (
	"fmt"

	"github.com/cbsinteractive/hybrik-sdk-go"
	"github.com/cbsinteractive/transcode-orchestrator/job"
)

func (p *hybrikProvider) audioElements(j *job.Job) (a []hybrik.Element) {
	tags := cfg.computeTags[job.ComputeClassTranscodeDefault]
	for i, f := range j.Output {
		if a.Codec == "" {
			continue
		}
		file := fmt.Sprintf("audio_output_%d.%s", i, a.Codec)
		uid := fmt.Sprintf("audio_%d", i)
		//TODO(as): make sure file name/dir is correct here
		a = append(a, label(p.audioElement(f.Kid(file)), uid, tag))
	}
	return a
}

func (p *hybrikProvider) audioElement(f job.File, idx int) hybrik.Element {
	return hybrik.Element{
		Kind: "transcode",
		Task: &hybrik.ElementTaskOptions{
			Name: "Audio Encode",
		},
		Payload: hybrik.LocationTargetPayload{
			Location: p.location(f, creds),
			Targets: []hybrik.TranscodeTarget{{
				FilePattern:   file,
				ExistingFiles: "replace",
				Container: hybrik.TranscodeContainer{
					Kind: "elementary",
				},
				Audio: []hybrik.AudioTarget{{
					Codec:     a.Codec,
					BitrateKb: a.Bitrate / 1000,
					Channels:  2,
					Source:    []hybrik.AudioTargetSource{{TrackNum: 0}},
				}},
			}},
		},
	}
}
