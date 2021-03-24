package hybrik

import (
	"fmt"

	hy "github.com/cbsinteractive/hybrik-sdk-go"
	"github.com/cbsinteractive/transcode-orchestrator/av"
)

func (p *driver) audioElements(j *av.Job) (a []hy.Element) {
	tags := tag(j, av.TagTranscodeDefault)
	for i, f := range j.Output.File {
		if f.Audio.Codec == "" {
			continue
		}
		uid := fmt.Sprintf("audio_%d", i)
		//TODO(as): make sure file name/dir is correct here
		a = append(a, label(hy.Element{
			Kind: "transcode",
			Task: &hy.ElementTaskOptions{
				Name: "Audio Encode",
			},
			Payload: hy.LocationTargetPayload{
				Location: p.location(f, p.auth(j).Write),
				Targets: []hy.TranscodeTarget{{
					FilePattern:   fmt.Sprintf("audio_output_%d.%s", i, f.Audio.Codec),
					ExistingFiles: "replace",
					Container: hy.TranscodeContainer{
						Kind: "elementary",
					},
					Audio: []hy.AudioTarget{{
						Codec:     f.Audio.Codec,
						BitrateKb: f.Audio.Bitrate / 1000,
						Channels:  2,
						Source:    []hy.AudioTargetSource{{TrackNum: 0}},
					}},
				}},
			},
		}, uid, tags...))
	}
	return a
}
