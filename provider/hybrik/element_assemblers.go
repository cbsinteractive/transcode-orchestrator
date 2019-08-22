package hybrik

import (
	"fmt"
	"strconv"

	"github.com/cbsinteractive/hybrik-sdk-go"
	"github.com/cbsinteractive/video-transcoding-api/db"
	"github.com/mitchellh/hashstructure"
	"github.com/pkg/errors"
)

const (
	doViProfile5                 = 5
	doViElementKind              = "dolby_vision"
	doViMezzQCElementUID         = "mezzanine_qc"
	doViMezzQCElementModule      = "mezzanine_qc"
	doViMezzQCElementName        = "Mezzanine QC"
	doViMezzQCOutputPathTmpl     = "%s/mezzanine_qc"
	doViMezzQCReportFilenameTmpl = "%s_mezz_qc_report.txt"
	doViMP4MuxFilenameDefault    = "{source_basename}.mp4"

	computeTagPreProcDefault = "preproc"
	computeTagMezzQCDefault  = "preproc"
)

func (p *hybrikProvider) defaultElementAssembler(cfg jobCfg) ([][]hybrik.Element, error) {
	elements := []hybrik.Element{}

	idx := 0
	for _, outputCfg := range cfg.outputCfgs {
		e := p.mountTranscodeElement(strconv.Itoa(idx), cfg.jobID, outputCfg.filename, cfg.destination,
			cfg.streamingParams.SegmentDuration, outputCfg.preset, cfg.executionFeatures, cfg.computeTags)
		elements = append(elements, e)
		idx++
	}

	return [][]hybrik.Element{elements}, nil
}

func (p *hybrikProvider) dolbyVisionElementAssembler(cfg jobCfg) ([][]hybrik.Element, error) {
	elementGroups := [][]hybrik.Element{}

	presets := map[string]hybrik.Preset{}
	for _, outputCfg := range cfg.outputCfgs {
		presets[outputCfg.filename] = outputCfg.preset
	}

	mezzQCComputeTag := computeTagMezzQCDefault
	if tag, found := cfg.computeTags[db.ComputeClassDolbyVisionMezzQC]; found {
		mezzQCComputeTag = tag
	}

	// initialize our pre-transcode execution group with a mezz qc task
	preTranscodeElements := []hybrik.Element{dolbyVisionMezzQCElementFrom(mezzQCComputeTag, cfg)}

	// then add any extracted audio elements to the pre-transcode group
	audioElements, audioPresetsToFilename, err := audioElementsFrom(presets, cfg)
	if err != nil {
		return nil, err
	}
	preTranscodeElements = append(preTranscodeElements, audioElements...)

	// add pre-transcode tasks as the first element in the pipeline
	elementGroups = append(elementGroups, preTranscodeElements)

	preprocComputeTag := computeTagPreProcDefault
	if tag, found := cfg.computeTags[db.ComputeClassDolbyVisionPreprocess]; found {
		preprocComputeTag = tag
	}

	transcodeElementsWithPreset, err := transcodeElementsWithPresetsFrom(presets, cfg)
	if err != nil {
		return nil, errors.Wrap(err, "converting presets into transcode elements")
	}

	// build up our transcode tasks
	transcodeElements := []hybrik.Element{}
	for idx, transcodeWithPreset := range transcodeElementsWithPreset {
		elementaryAudioStreams := []hybrik.DoViMP4MuxElementaryStream{}
		audioCfg, found, err := audioTargetFromPreset(transcodeWithPreset.preset)
		if err != nil {
			return nil, err
		}

		if found {
			hash, err := hashstructure.Hash(audioCfg, nil)
			if err != nil {
				return nil, errors.Wrap(err, "hashing audio cfg struct")
			}

			if filename, found := audioPresetsToFilename[hash]; found {
				elementaryAudioStreams = append(elementaryAudioStreams, hybrik.DoViMP4MuxElementaryStream{
					AssetURL: hybrik.AssetURL{
						StorageProvider: cfg.destination.provider,
						URL:             fmt.Sprintf("%s/%s", cfg.destination.path, filename),
					},
				})
			}
		}

		doViElement := hybrik.Element{
			UID:  fmt.Sprintf("dolby_vision_%d", idx),
			Kind: doViElementKind,
			Task: &hybrik.ElementTaskOptions{
				Name:              fmt.Sprintf("Encode #%d", idx),
				RetryMethod:       "fail",
				SourceElementUIDs: []string{cfg.source.UID},
			},
			Payload: hybrik.DolbyVisionV2TaskPayload{
				Module:  "encoder",
				Profile: doViProfile5,
				Location: hybrik.TranscodeLocation{
					StorageProvider: cfg.destination.provider,
					Path:            cfg.destination.path,
				},
				Preprocessing: hybrik.DolbyVisionV2Preprocessing{
					Task: hybrik.TaskTags{
						Tags: []string{preprocComputeTag},
					},
				},
				Transcodes: []hybrik.Element{transcodeWithPreset.transcodeElement},
				PostTranscode: hybrik.DoViPostTranscode{
					MP4Mux: hybrik.DoViMP4Mux{
						Enabled:           true,
						FilePattern:       doViMP4MuxFilenameDefault,
						ElementaryStreams: elementaryAudioStreams,
					},
				},
			},
		}

		transcodeElements = append(transcodeElements, doViElement)
	}

	// add all transcode tasks as the second element in the pipeline
	elementGroups = append(elementGroups, transcodeElements)

	return elementGroups, nil
}

func audioElementsFrom(presets map[string]hybrik.Preset, cfg jobCfg) ([]hybrik.Element, map[uint64]string, error) {
	audioConfigurations := map[uint64]hybrik.AudioTarget{}
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

		audioElement := transcodeAudioElementFromPreset(audioTarget, outputFilename, idx, cfg.computeTags,
			cfg.destination, containerKindElementary)
		audioElements = append(audioElements, audioElement)

		audioPresetsToFilename[hash] = outputFilename
		idx++
	}
	return audioElements, audioPresetsToFilename, nil
}

func transcodeElementsWithPresetsFrom(presets map[string]hybrik.Preset, cfg jobCfg) ([]transcodeElementWithPreset, error) {
	transcodeElementsWithPreset := []transcodeElementWithPreset{}
	idx := 0
	for filename, preset := range presets {
		// removing audio as we're processing this separately
		presetWithoutAudio := preset
		for idx, target := range preset.Payload.Targets {
			target.Audio = []hybrik.AudioTarget{}
			presetWithoutAudio.Payload.Targets[idx] = target
		}
		element, err := transcodeElementFromPreset(presetWithoutAudio, fmt.Sprintf("transcode_task_%d", idx),
			cfg.destination, filename, cfg.executionFeatures, cfg.computeTags)
		if err != nil {
			return nil, errors.Wrapf(err, "mapping hybrik preset %v into a transcode element", preset)
		}

		transcodeElementsWithPreset = append(transcodeElementsWithPreset, transcodeElementWithPreset{
			transcodeElement: element,
			preset:           preset,
		})
		idx++
	}

	return transcodeElementsWithPreset, nil
}

func dolbyVisionMezzQCElementFrom(mezzQCComputeTag string, cfg jobCfg) hybrik.Element {
	mezzQCElement := hybrik.Element{
		UID:  doViMezzQCElementUID,
		Kind: doViElementKind,
		Task: &hybrik.ElementTaskOptions{
			Name: doViMezzQCElementName,
			Tags: []string{mezzQCComputeTag},
		},
		Payload: hybrik.DoViV2MezzanineQCPayload{
			Module: doViMezzQCElementModule,
			Params: hybrik.DoViV2MezzanineQCPayloadParams{
				Location: hybrik.TranscodeLocation{
					StorageProvider: cfg.destination.provider,
					Path:            fmt.Sprintf(doViMezzQCOutputPathTmpl, cfg.destination.path),
				},
				FilePattern: fmt.Sprintf(doViMezzQCReportFilenameTmpl, cfg.jobID),
			},
		},
	}
	return mezzQCElement
}

func audioTargetFromPreset(preset hybrik.Preset) (hybrik.AudioTarget, bool, error) {
	if len(preset.Payload.Targets) == 0 {
		return hybrik.AudioTarget{}, false, fmt.Errorf("preset has no targets: %v", preset)
	}

	target := preset.Payload.Targets[0]
	if len(target.Audio) == 0 {
		return hybrik.AudioTarget{}, false, nil
	}

	return target.Audio[0], true, nil
}
