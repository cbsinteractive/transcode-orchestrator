package hybrik

import (
	"fmt"

	"github.com/cbsinteractive/hybrik-sdk-go"
	"github.com/cbsinteractive/video-transcoding-api/db"
	"github.com/pkg/errors"
)

type transcodeElementWithPreset struct {
	transcodeElement hybrik.Element
	preset           hybrik.Preset
}

type elementKind = string

const (
	elementKindTranscode elementKind = "transcode"
	elementKindSource    elementKind = "source"
	elementKindPackage   elementKind = "package"
)

type storageProvider = string

const (
	storageProviderUnrecognized storageProvider = "unrecognized"
	storageProviderS3           storageProvider = "s3"
	storageProviderGCS          storageProvider = "gs"
)

func transcodeElementFromPreset(preset hybrik.Preset, uid string, destination storageLocation, filename string,
	execFeatures executionFeatures, computeTags map[db.ComputeClass]string) (hybrik.Element, error) {
	if len(preset.Payload.Targets) != 1 {
		return hybrik.Element{}, errors.New("the hybrik provider only supports presets with a single target")
	}
	target := preset.Payload.Targets[0]

	// default video preset to slow
	target.Video.Preset = presetSlow

	payload := hybrik.TranscodePayload{
		LocationTargetPayload: hybrik.LocationTargetPayload{
			Location: hybrik.TranscodeLocation{
				StorageProvider: destination.provider,
				Path:            fmt.Sprintf("%s/elementary", destination.path),
			},
			Targets: []hybrik.PresetTarget{{
				FilePattern:   filename,
				ExistingFiles: target.ExistingFiles,
				Container: hybrik.TranscodeContainer{
					Kind: target.Container.Kind,
				},
				NumPasses: target.NumPasses,
				Video:     target.Video,
				Audio:     target.Audio,
			}},
		},
		Options: preset.Payload.Options,
	}

	if execFeatures.segmentedRendering != nil {
		payload.SourcePipeline = hybrik.TranscodeSourcePipeline{SegmentedRendering: execFeatures.segmentedRendering}
	}

	transcodeComputeTags := []string{}
	if tag, found := computeTags[db.ComputeClassDolbyVisionTranscode]; found {
		transcodeComputeTags = append(transcodeComputeTags, tag)
	}

	element := hybrik.Element{
		UID:  uid,
		Kind: elementKindTranscode,
		Task: &hybrik.ElementTaskOptions{
			Tags: transcodeComputeTags,
		},
		Payload: payload,
	}

	return element, nil
}

func transcodeAudioElementFromPreset(target hybrik.AudioTarget, outputFilename string, idx int,
	computeTags map[db.ComputeClass]string, destination storageLocation, container string) hybrik.Element {
	transcodeComputeTags := []string{}
	if tag, found := computeTags[db.ComputeClassTranscodeDefault]; found {
		transcodeComputeTags = append(transcodeComputeTags, tag)
	}

	return hybrik.Element{
		UID:  fmt.Sprintf("audio_%d", idx),
		Kind: elementKindTranscode,
		Task: &hybrik.ElementTaskOptions{
			Tags: transcodeComputeTags,
			Name: "Audio Encode",
		},
		Payload: hybrik.LocationTargetPayload{
			Location: hybrik.TranscodeLocation{
				StorageProvider: destination.provider,
				Path:            destination.path,
			},
			Targets: []hybrik.TranscodeTarget{{
				FilePattern:   outputFilename,
				ExistingFiles: "replace",
				Container: hybrik.TranscodeContainer{
					Kind: container,
				},
				Audio: []map[string]interface{}{
					{
						"codec":      target.Codec,
						"bitrate_kb": target.BitrateKb,
						"channels":   2,
						"source": []map[string]interface{}{
							{
								"track": 1,
							},
						},
					},
				},
			}},
		},
	}
}
