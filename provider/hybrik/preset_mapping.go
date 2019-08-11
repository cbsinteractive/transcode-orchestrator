package hybrik

import (
	"fmt"

	"github.com/cbsinteractive/hybrik-sdk-go"
	"github.com/pkg/errors"
)

type elementKind = string

const (
	elementKindTranscode   elementKind = "transcode"
	elementKindSource      elementKind = "source"
	elementKindDolbyVision elementKind = "dolby_vision"
)

type storageProvider = string

const storageProviderUnrecognized storageProvider = "unrecognized"
const storageProviderS3 storageProvider = "s3"
const storageProviderGCS storageProvider = "gs"

func transcodeElementsFromPresets(presets map[string]hybrik.Preset, destination storageLocation,
	execFeatures executionFeatures) ([]hybrik.Element, error) {
	elements := []hybrik.Element{}

	idx := 0
	for filename, preset := range presets {
		element, err := transcodeElementFromPreset(preset, fmt.Sprintf("transcode_task_%d", idx),
			destination, filename, execFeatures)
		if err != nil {
			return nil, errors.Wrapf(err, "mapping hybrik preset %v into a transcode element", preset)
		}

		elements = append(elements, element)
		idx++
	}

	return elements, nil
}

func transcodeElementFromPreset(preset hybrik.Preset, uid string, destination storageLocation, filename string,
	execFeatures executionFeatures) (hybrik.Element, error) {
	if len(preset.Payload.Targets) != 1 {
		return hybrik.Element{}, errors.New("the hybrik provider only supports presets with a single target")
	}
	target := preset.Payload.Targets[0]

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

	element := hybrik.Element{
		UID:     uid,
		Kind:    elementKindTranscode,
		Payload: payload,
	}

	return element, nil
}
