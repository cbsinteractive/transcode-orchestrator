package hybrik

import (
	"fmt"

	"github.com/cbsinteractive/hybrik-sdk-go"
	"github.com/cbsinteractive/transcode-orchestrator/job"
)

type transcodeElementWithFilename struct {
	transcodeElement hybrik.Element
	filename         string
}

type elementKind = string

const (
	elementKindTranscode   elementKind = "transcode"
	elementKindSource      elementKind = "source"
	elementKindPackage     elementKind = "package"
	elementKindDolbyVision elementKind = "dolby_vision"
)

type storageProvider string

func (p storageProvider) supportsSegmentedRendering() bool { return p != storageProviderHTTP }
func (p storageProvider) string() string                   { return string(p) }

const (
	storageProviderUnrecognized storageProvider = "unrecognized"
	storageProviderS3           storageProvider = "s3"
	storageProviderGCS          storageProvider = "gs"
	storageProviderHTTP         storageProvider = "http"
)

// TODO(as): should canonicalize this across all providers
func container(f job.File) string {
	kind := ""
	for _, c := range p.Capabilities().OutputFormats {
		if f.Type() == c || (f.Type() == "m3u8" && c == hls) {
			kind = c
		}
	}
	return kind
}
func passes(f job.File) int {
	if f.Video.TwoPass {
		return 2
	}
	return 1
}

func label(e hybrik.Element, uid string, tags ...string) hybrik.Element {
	e.UID = uid
	for _, tags := range tags {
		if len(t) > 0 {
			e.Task.Tags = append(e.Task.Tags, t)
		}
	}
	return e
}

func (p *hybrikProvider) audioElements(j *job.Job) (a []hybrik.Element) {
	tags := cfg.computeTags[job.ComputeClassTranscodeDefault]
	for i, f := range j.Output {
		if a.Codec == "" {
			continue
		}
		file := fmt.Sprintf("audio_output_%d.%s", i, a.Codec)
		uid := fmt.Sprintf("audio_%d", i)
		a = append(a, label(p.audioElement(f.Kid(file)), uid, tag))
	}
	return a
}

func (p *hybrikProvider) audioElement(f job.File, idx int) hybrik.Element {
	return hybrik.Element{
		Kind: elementKindTranscode,
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
