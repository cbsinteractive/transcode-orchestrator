package hybrik

import (
	"fmt"

	"github.com/cbsinteractive/hybrik-sdk-go"
	"github.com/cbsinteractive/transcode-orchestrator/db"
	"github.com/cbsinteractive/transcode-orchestrator/job"
	"github.com/mitchellh/hashstructure"
	"github.com/pkg/errors"
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

func (p *hybrikProvider) transcodeElementsWithPresetsFrom(presets map[string]job.Preset, cfg jobCfg) ([]transcodeElementWithFilename, error) {
	transcodeElementsWithFilename := []transcodeElementWithFilename{}
	idx := 0
	for filename, preset := range presets {
		element, err := p.transcodeElementFromPreset(preset, fmt.Sprintf(transcodeElementIDTemplate, idx), cfg, filename)
		if err != nil {
			return nil, errors.Wrapf(err, "mapping hybrik preset %v into a transcode element", preset)
		}

		transcodeElementsWithFilename = append(transcodeElementsWithFilename, transcodeElementWithFilename{
			transcodeElement: element,
			filename:         filename,
		})
		idx++
	}

	return transcodeElementsWithFilename, nil
}

func (p *hybrikProvider) transcodeElementFromPreset(preset job.Preset, uid string, cfg jobCfg, filename string) (hybrik.Element, error) {
	container := ""
	for _, c := range p.Capabilities().OutputFormats {
		if preset.Container == c || (preset.Container == "m3u8" && c == hls) {
			container = c
		}
	}

	if container == "" {
		return hybrik.Element{}, ErrUnsupportedContainer
	}

	videoTarget, err := videoTargetFrom(preset.Video, preset.RateControl)
	if err != nil {
		return hybrik.Element{}, errors.Wrap(err, "building video targets")
	}

	audioTarget, err := audioTargetFrom(preset.Audio)
	if err != nil {
		return hybrik.Element{}, errors.Wrap(err, "building audio targets")
	}

	numPasses := 1
	if preset.TwoPass {
		numPasses = 2
	}

	payload := hybrik.TranscodePayload{
		LocationTargetPayload: hybrik.LocationTargetPayload{
			Location: p.transcodeLocationFrom(cfg.destination, cfg.executionEnvironment.OutputAlias),
			Targets: []hybrik.TranscodeTarget{{
				FilePattern:   filename,
				ExistingFiles: "replace",
				Container: hybrik.TranscodeContainer{
					Kind: container,
				},
				NumPasses: numPasses,
				Video:     videoTarget,
				Audio:     audioTarget,
			}},
		},
	}

	if cfg.executionFeatures.segmentedRendering != nil {
		payload.SourcePipeline = hybrik.TranscodeSourcePipeline{SegmentedRendering: cfg.executionFeatures.segmentedRendering}
	}

	for _, modifier := range transcodePayloadModifiersFor(preset) {
		payload, err = modifier.runFunc(payload, preset)
		if err != nil {
			return hybrik.Element{}, errors.Wrapf(err, "running %q transcode payload modifier", modifier.name)
		}
	}

	transcodeComputeTags := []string{}
	if tag, found := cfg.computeTags[db.ComputeClassTranscodeDefault]; found {
		transcodeComputeTags = append(transcodeComputeTags, tag)
	}

	element := hybrik.Element{
		UID:  uid,
		Kind: elementKindTranscode,
		Task: &hybrik.ElementTaskOptions{
			Tags: transcodeComputeTags,
			Name: fmt.Sprintf("Transcode - %s", filename),
		},
		Payload: payload,
	}

	return element, nil
}

type transcodePayloadModifier struct {
	name    string
	runFunc func(hybrikPreset hybrik.TranscodePayload, preset job.Preset) (hybrik.TranscodePayload, error)
}

func transcodePayloadModifiersFor(preset job.Preset) []transcodePayloadModifier {
	modifiers := []transcodePayloadModifier{}

	// Rate control
	if preset.RateControl != "" && preset.Video != (db.Video{}) {
		modifiers = append(modifiers, transcodePayloadModifier{name: "rateControl", runFunc: enrichTranscodePayloadWithRateControl})
	}

	// HDR
	if _, hdrEnabled := hdrTypeFromPreset(preset); hdrEnabled {
		modifiers = append(modifiers, transcodePayloadModifier{name: "hdr", runFunc: enrichTranscodePayloadWithHDRMetadata})
	}

	// MXF sources
	if preset.SourceContainer == "mxf" {
		modifiers = append(modifiers, transcodePayloadModifier{name: "mxf", runFunc: modifyTranscodePayloadForMXFSources})
	}

	return modifiers
}

func (p *hybrikProvider) audioElementsFrom(presets map[string]job.Preset, cfg jobCfg) ([]hybrik.Element, map[uint64]string, error) {
	audioConfigurations := map[uint64]job.Audio{}
	for _, preset := range presets {
		audioCfg, found, err := audioTargetFromPreset(preset)
		if err != nil {
			return nil, nil, err
		} else if !found {
			continue
		}

		cfgHash, err := hashstructure.Hash(audioCfg, nil)
		if err != nil {
			return nil, nil, errors.Wrap(err, "hashing a preset audio cfg")
		}

		audioConfigurations[cfgHash] = audioCfg
	}

	audioElements := []hybrik.Element{}
	audioPresetsToFilename := map[uint64]string{}
	idx := 0
	for hash, audioTarget := range audioConfigurations {
		outputFilename := fmt.Sprintf("audio_output_%d.%s", idx, audioTarget.Codec)

		audioElement, err := p.transcodeAudioElementFromPreset(audioTarget, outputFilename, idx, cfg,
			containerKindElementary)
		if err != nil {
			return nil, nil, errors.Wrap(err, "generating audio element from job.Preset")
		}
		audioElements = append(audioElements, audioElement)

		audioPresetsToFilename[hash] = outputFilename
		idx++
	}
	return audioElements, audioPresetsToFilename, nil
}

func (p *hybrikProvider) transcodeAudioElementFromPreset(target job.Audio, outputFilename string, idx int,
	cfg jobCfg, container string) (hybrik.Element, error) {
	transcodeComputeTags := []string{}
	if tag, found := cfg.computeTags[db.ComputeClassTranscodeDefault]; found {
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
			Location: p.transcodeLocationFrom(cfg.destination, cfg.executionEnvironment.OutputAlias),
			Targets: []hybrik.TranscodeTarget{{
				FilePattern:   outputFilename,
				ExistingFiles: "replace",
				Container: hybrik.TranscodeContainer{
					Kind: container,
				},
				Audio: []hybrik.AudioTarget{{
					Codec:     target.Codec,
					BitrateKb: target.Bitrate / 1000,
					Channels:  2,
					Source:    []hybrik.AudioTargetSource{{TrackNum: 0}},
				}},
			}},
		},
	}, nil
}
