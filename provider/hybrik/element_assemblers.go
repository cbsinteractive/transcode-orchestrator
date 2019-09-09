package hybrik

import (
	"fmt"

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
	doViMP4MuxDVH1FlagKey        = "dvh1flag"

	computeTagPreProcDefault = "preproc"
	computeTagMezzQCDefault  = "preproc"
)

func (p *hybrikProvider) defaultElementAssembler(cfg jobCfg) ([][]hybrik.Element, error) {
	elements := []hybrik.Element{}

	idx := 0
	for _, outputCfg := range cfg.outputCfgs {
		element, err := p.transcodeElementFromPreset(outputCfg.localPreset, fmt.Sprintf(transcodeElementIDTemplate, idx),
			cfg, outputCfg.filename)
		if err != nil {
			return nil, err
		}

		elements = append(elements, element)
		idx++
	}

	return [][]hybrik.Element{elements}, nil
}

func (p *hybrikProvider) dolbyVisionElementAssembler(cfg jobCfg) ([][]hybrik.Element, error) {
	elementGroups := [][]hybrik.Element{}

	presets := map[string]db.Preset{}
	presetsWithoutAudio := map[string]db.Preset{}
	for _, outputCfg := range cfg.outputCfgs {
		preset := outputCfg.localPreset
		presets[outputCfg.filename] = preset

		// removing audio so we can processing this separately
		preset.Audio = db.AudioPreset{}
		presetsWithoutAudio[outputCfg.filename] = preset
	}

	mezzQCComputeTag := computeTagMezzQCDefault
	if tag, found := cfg.computeTags[db.ComputeClassDolbyVisionPreprocess]; found {
		mezzQCComputeTag = tag
	}

	// initialize our pre-transcode execution group with a mezz qc task
	preTranscodeElements := []hybrik.Element{p.dolbyVisionMezzQCElementFrom(mezzQCComputeTag, cfg)}

	// then add any extracted audio elements to the pre-transcode group
	audioElements, audioPresetsToFilename, err := p.audioElementsFrom(presets, cfg)
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

	transcodeElementsWithFilenames, err := p.transcodeElementsWithPresetsFrom(presetsWithoutAudio, cfg)
	if err != nil {
		return nil, errors.Wrap(err, "converting presets into transcode elements")
	}

	// build up our transcode tasks
	transcodeElements := []hybrik.Element{}
	for idx, transcodeWithFilename := range transcodeElementsWithFilenames {
		elementaryAudioStreams := []hybrik.DoViMP4MuxElementaryStream{}
		audioCfg, found, err := audioTargetFromPreset(presets[transcodeWithFilename.filename])
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
					AssetURL: p.assetURLFrom(storageLocation{
						provider: cfg.destination.provider,
						path:     fmt.Sprintf("%s/%s", cfg.destination.path, filename),
					}, cfg.executionEnvironment),
				})
			}
		}

		doViElement := hybrik.Element{
			UID:  fmt.Sprintf("dolby_vision_%d", idx),
			Kind: doViElementKind,
			Task: &hybrik.ElementTaskOptions{
				Name:              fmt.Sprintf("Encode #%d", idx),
				RetryMethod:       "fail",
				Tags:              []string{preprocComputeTag},
				SourceElementUIDs: []string{cfg.source.UID},
			},
			Payload: hybrik.DolbyVisionV2TaskPayload{
				Module:   "encoder",
				Profile:  doViProfile5,
				Location: p.transcodeLocationFrom(cfg.destination, cfg.executionEnvironment),
				Preprocessing: hybrik.DolbyVisionV2Preprocessing{
					Task: hybrik.TaskTags{
						Tags: []string{preprocComputeTag},
					},
				},
				Transcodes: []hybrik.Element{transcodeWithFilename.transcodeElement},
				PostTranscode: hybrik.DoViPostTranscode{
					Task: &hybrik.TaskTags{Tags: []string{preprocComputeTag}},
					MP4Mux: hybrik.DoViMP4Mux{
						Enabled:           true,
						FilePattern:       doViMP4MuxFilenameDefault,
						ElementaryStreams: elementaryAudioStreams,
						CLIOptions:        map[string]string{doViMP4MuxDVH1FlagKey: ""},
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

func (p *hybrikProvider) dolbyVisionMezzQCElementFrom(mezzQCComputeTag string, cfg jobCfg) hybrik.Element {
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
				Location: p.transcodeLocationFrom(storageLocation{
					provider: cfg.destination.provider,
					path:     fmt.Sprintf(doViMezzQCOutputPathTmpl, cfg.destination.path),
				}, cfg.executionEnvironment),
				FilePattern: fmt.Sprintf(doViMezzQCReportFilenameTmpl, cfg.jobID),
			},
		},
	}
	return mezzQCElement
}

func audioTargetFromPreset(preset db.Preset) (db.AudioPreset, bool, error) {
	audioCfg := preset.Audio
	return audioCfg, audioCfg != (db.AudioPreset{}), nil
}
