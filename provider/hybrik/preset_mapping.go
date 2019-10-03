package hybrik

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cbsinteractive/hybrik-sdk-go"
	"github.com/cbsinteractive/video-transcoding-api/db"
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

type storageProvider = string

const (
	storageProviderUnrecognized storageProvider = "unrecognized"
	storageProviderS3           storageProvider = "s3"
	storageProviderGCS          storageProvider = "gs"
	storageProviderHTTP         storageProvider = "http"
)

func (p *hybrikProvider) transcodeElementsWithPresetsFrom(presets map[string]db.Preset, cfg jobCfg) ([]transcodeElementWithFilename, error) {
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

func (p *hybrikProvider) transcodeElementFromPreset(preset db.Preset, uid string, cfg jobCfg, filename string) (hybrik.Element, error) {
	var minGOPFrames, maxGOPFrames, gopSize int

	gopSize, err := strconv.Atoi(preset.Video.GopSize)
	if err != nil {
		return hybrik.Element{}, err
	}

	minGOPFrames = gopSize
	maxGOPFrames = gopSize

	container := ""
	for _, c := range p.Capabilities().OutputFormats {
		if preset.Container == c || (preset.Container == "m3u8" && c == hls) {
			container = c
		}
	}

	if container == "" {
		return hybrik.Element{}, ErrUnsupportedContainer
	}

	bitrate, err := strconv.Atoi(preset.Video.Bitrate)
	if err != nil {
		return hybrik.Element{}, ErrBitrateNan
	}

	var videoWidth *int
	var videoHeight *int

	if preset.Video.Width != "" {
		var presetWidth int
		presetWidth, err = strconv.Atoi(preset.Video.Width)
		if err != nil {
			return hybrik.Element{}, ErrVideoWidthNan
		}
		videoWidth = &presetWidth
	}

	if preset.Video.Height != "" {
		var presetHeight int
		presetHeight, err = strconv.Atoi(preset.Video.Height)
		if err != nil {
			return hybrik.Element{}, ErrVideoHeightNan
		}
		videoHeight = &presetHeight
	}

	videoProfile := strings.ToLower(preset.Video.Profile)
	videoLevel := preset.Video.ProfileLevel

	// TODO: Understand video-transcoding-api profile + level settings in relation to vp8
	// For now, we will omit and leave to encoder defaults
	if preset.Video.Codec == "vp8" {
		videoProfile = ""
		videoLevel = ""
	}

	audioTargets, err := audioTargetsFrom(preset.Audio)
	if err != nil {
		return hybrik.Element{}, errors.Wrap(err, "building audio targets")
	}

	numPasses := 1
	if preset.TwoPass {
		numPasses = 2
	}

	payload := hybrik.TranscodePayload{
		LocationTargetPayload: hybrik.LocationTargetPayload{
			Location: p.transcodeLocationFrom(cfg.destination, cfg.executionEnvironment),
			Targets: []hybrik.TranscodeTarget{{
				FilePattern:   filename,
				ExistingFiles: "replace",
				Container: hybrik.TranscodeContainer{
					Kind: container,
				},
				NumPasses: numPasses,
				Video: &hybrik.VideoTarget{
					Width:             videoWidth,
					Height:            videoHeight,
					BitrateMode:       strings.ToLower(preset.RateControl),
					BitrateKb:         bitrate / 1000,
					Preset:            presetSlow,
					Codec:             preset.Video.Codec,
					ChromaFormat:      chromaFormatYUV420P,
					Profile:           videoProfile,
					Level:             videoLevel,
					MinGOPFrames:      minGOPFrames,
					MaxGOPFrames:      maxGOPFrames,
					ExactGOPFrames:    maxGOPFrames,
					InterlaceMode:     preset.Video.InterlaceMode,
					UseSceneDetection: false,
				},
				Audio: audioTargets,
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
	runFunc func(hybrikPreset hybrik.TranscodePayload, preset db.Preset) (hybrik.TranscodePayload, error)
}

func transcodePayloadModifiersFor(preset db.Preset) []transcodePayloadModifier {
	modifiers := []transcodePayloadModifier{}

	// Rate control
	if preset.RateControl != "" {
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

func (p *hybrikProvider) audioElementsFrom(presets map[string]db.Preset, cfg jobCfg) ([]hybrik.Element, map[uint64]string, error) {
	audioConfigurations := map[uint64]db.AudioPreset{}
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
			return nil, nil, errors.Wrap(err, "generating audio element from db.Preset")
		}
		audioElements = append(audioElements, audioElement)

		audioPresetsToFilename[hash] = outputFilename
		idx++
	}
	return audioElements, audioPresetsToFilename, nil
}

func (p *hybrikProvider) transcodeAudioElementFromPreset(target db.AudioPreset, outputFilename string, idx int,
	cfg jobCfg, container string) (hybrik.Element, error) {
	transcodeComputeTags := []string{}
	if tag, found := cfg.computeTags[db.ComputeClassTranscodeDefault]; found {
		transcodeComputeTags = append(transcodeComputeTags, tag)
	}

	bitrate, err := strconv.Atoi(target.Bitrate)
	if err != nil {
		return hybrik.Element{}, ErrBitrateNan
	}

	return hybrik.Element{
		UID:  fmt.Sprintf("audio_%d", idx),
		Kind: elementKindTranscode,
		Task: &hybrik.ElementTaskOptions{
			Tags: transcodeComputeTags,
			Name: "Audio Encode",
		},
		Payload: hybrik.LocationTargetPayload{
			Location: p.transcodeLocationFrom(cfg.destination, cfg.executionEnvironment),
			Targets: []hybrik.TranscodeTarget{{
				FilePattern:   outputFilename,
				ExistingFiles: "replace",
				Container: hybrik.TranscodeContainer{
					Kind: container,
				},
				Audio: []hybrik.AudioTarget{{
					Codec:     target.Codec,
					BitrateKb: bitrate / 1000,
					Channels:  2,
					Source:    []hybrik.AudioTargetSource{{TrackNum: 0}},
				}},
			}},
		},
	}, nil
}
