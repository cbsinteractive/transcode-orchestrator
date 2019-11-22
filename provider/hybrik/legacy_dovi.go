package hybrik

import (
	"fmt"

	"github.com/cbsinteractive/hybrik-sdk-go"
	"github.com/cbsinteractive/video-transcoding-api/db"
)

// TODO remove once Hybrik fixes bugs with the new Dolby Vision job structure

const (
	dolbyVisionElementID                  = "dolby_vision_task"
	doViModuleProfile                     = "profile"
	doViNBCPreProcOutputPathTmpl          = "%s/tmp"
	doViVESMuxOutputPathTmpl              = "%s/tmp"
	doViVESMuxFilenameDefault             = "ves.h265"
	doViMetadataPostProcOutputPathTmpl    = "%s/tmp"
	doViMetadataPostProcFilenameDefault   = "postproc.265"
	doViMetadataPostProcQCOutputPathTmpl  = "%s/tmp"
	doViMetadataPostProcQCFilenameDefault = "metadata_postproc_qc_report.txt"
	doViMP4MuxQCOutputPathTmpl            = "%s/tmp"
	doViMP4MuxQCFilenameDefault           = "mp4_qc_report.txt"
	doViSourceDemuxOutputPathTmpl         = "%s/tmp"
	doViLegacyMezzQCOutputPathTmpl        = "%s/tmp"
	doViPreProcNumTasksAuto               = "auto"
	doViPreProcIntervalLengthDefault      = 48
	doViInputEDRAspectDefault             = "2"
	doViInputEDRCropDefault               = "0x0x0x0"
	doViInputEDRPadDefault                = "0x0x0x0"

	retryMethodRetry  = "retry"
	retryCountDefault = 3
	retryDelayDefault = 30
)

func (p *hybrikProvider) dolbyVisionLegacyElementAssembler(cfg jobCfg) ([][]hybrik.Element, error) {
	presetsWithoutAudio := map[string]db.Preset{}
	for _, outputCfg := range cfg.outputCfgs {
		preset := outputCfg.localPreset

		// removing audio so we can processing this separately
		preset.Audio = db.AudioPreset{}
		presetsWithoutAudio[outputCfg.filename] = preset
	}

	transcodeElementsWithFilenames, err := p.transcodeElementsWithPresetsFrom(presetsWithoutAudio, cfg)
	if err != nil {
		return nil, err
	}

	transcodeElements := []hybrik.Element{}
	for _, transcodeWithFilename := range transcodeElementsWithFilenames {
		transcodeElements = append(transcodeElements, transcodeWithFilename.transcodeElement)
	}

	doViPreProcNumTasks := doViPreProcNumTasksAuto
	if numTasks := cfg.executionFeatures.doViPreProcSegmentation.numTasks; numTasks != "" {
		doViPreProcNumTasks = numTasks
	}

	doViPreProcIntervalLength := doViPreProcIntervalLengthDefault
	if intervalLength := cfg.executionFeatures.doViPreProcSegmentation.intervalLength; intervalLength != 0 {
		doViPreProcIntervalLength = intervalLength
	}

	preprocComputeTag := computeTagPreProcDefault
	if tag, found := cfg.computeTags[db.ComputeClassDolbyVisionPreprocess]; found {
		preprocComputeTag = tag
	}

	destStorageLocation := p.transcodeLocationFrom(storageLocation{
		provider: cfg.destination.provider,
		path:     cfg.destination.path,
	}, cfg.executionEnvironment.OutputAlias)

	demuxOutputStorageLocation := p.transcodeLocationFrom(storageLocation{
		provider: cfg.destination.provider,
		path:     fmt.Sprintf(doViSourceDemuxOutputPathTmpl, cfg.destination.path),
	}, cfg.executionEnvironment.OutputAlias)

	return [][]hybrik.Element{{{
		UID:  dolbyVisionElementID,
		Kind: elementKindDolbyVision,
		Task: &hybrik.ElementTaskOptions{
			Tags: []string{preprocComputeTag},
		},
		Payload: hybrik.DolbyVisionTaskPayload{
			Module:  doViModuleProfile,
			Profile: doViProfile5,
			MezzanineQC: hybrik.DoViMezzanineQC{
				Enabled: false,
				Location: p.transcodeLocationFrom(storageLocation{
					provider: cfg.destination.provider,
					path:     fmt.Sprintf(doViLegacyMezzQCOutputPathTmpl, cfg.destination.path),
				}, cfg.executionEnvironment.OutputAlias),
				Task: hybrik.TaskTags{
					Tags: []string{preprocComputeTag},
				},
				FilePattern: fmt.Sprintf(doViMezzQCReportFilenameTmpl, cfg.jobID),
				ToolVersion: hybrik.DoViMezzQCVersionDefault,
			},
			NBCPreproc: hybrik.DoViNBCPreproc{
				Task: hybrik.TaskTags{
					Tags: []string{preprocComputeTag},
				},
				Location: p.transcodeLocationFrom(storageLocation{
					provider: cfg.destination.provider,
					path:     fmt.Sprintf(doViNBCPreProcOutputPathTmpl, cfg.destination.path),
				}, cfg.executionEnvironment.OutputAlias),
				SDKVersion:     hybrik.DoViSDKVersionDefault,
				NumTasks:       doViPreProcNumTasks,
				IntervalLength: doViPreProcIntervalLength,
				CLIOptions: hybrik.DoViNBCPreprocCLIOptions{
					InputEDRAspect: doViInputEDRAspectDefault,
					InputEDRPad:    doViInputEDRCropDefault,
					InputEDRCrop:   doViInputEDRPadDefault,
				},
			},
			Transcodes: transcodeElements,
			PostTranscode: hybrik.DoViPostTranscode{
				Task: &hybrik.TaskTags{Tags: []string{preprocComputeTag}},
				VESMux: &hybrik.DoViVESMux{
					Enabled: true,
					Location: p.transcodeLocationFrom(storageLocation{
						provider: cfg.destination.provider,
						path:     fmt.Sprintf(doViVESMuxOutputPathTmpl, cfg.destination.path),
					}, cfg.executionEnvironment.OutputAlias),
					FilePattern: doViVESMuxFilenameDefault,
					SDKVersion:  hybrik.DoViSDKVersionDefault,
				},
				MetadataPostProc: &hybrik.DoViMetadataPostProc{
					Enabled: true,
					Location: p.transcodeLocationFrom(storageLocation{
						provider: cfg.destination.provider,
						path:     fmt.Sprintf(doViMetadataPostProcOutputPathTmpl, cfg.destination.path),
					}, cfg.executionEnvironment.OutputAlias),
					FilePattern: doViMetadataPostProcFilenameDefault,
					SDKVersion:  hybrik.DoViSDKVersionDefault,
					QCSettings: hybrik.DoViQCSettings{
						Enabled:     true,
						ToolVersion: hybrik.DoViVesQCVersionDefault,
						Location: p.transcodeLocationFrom(storageLocation{
							provider: cfg.destination.provider,
							path:     fmt.Sprintf(doViMetadataPostProcQCOutputPathTmpl, cfg.destination.path),
						}, cfg.executionEnvironment.OutputAlias),
						FilePattern: doViMetadataPostProcQCFilenameDefault,
					},
				},
				MP4Mux: hybrik.DoViMP4Mux{
					Enabled:            true,
					Location:           &destStorageLocation,
					FilePattern:        doViMP4MuxFilenameDefault,
					ToolVersion:        hybrik.DoViMP4MuxerVersionDefault,
					CopySourceStartPTS: true,
					CLIOptions:         map[string]string{doViMP4MuxDVH1FlagKey: ""},
					QCSettings: &hybrik.DoViQCSettings{
						Enabled:     true,
						ToolVersion: hybrik.DoViMP4QCVersionDefault,
						Location: p.transcodeLocationFrom(storageLocation{
							provider: cfg.destination.provider,
							path:     fmt.Sprintf(doViMP4MuxQCOutputPathTmpl, cfg.destination.path),
						}, cfg.executionEnvironment.OutputAlias),
						FilePattern: doViMP4MuxQCFilenameDefault,
					},
					ElementaryStreams: []hybrik.DoViMP4MuxElementaryStream{{
						AssetURL: p.assetURLFrom(storageLocation{
							provider: cfg.sourceLocation.provider,
							path:     cfg.sourceLocation.path,
						}, cfg.executionEnvironment.OutputAlias),
						ExtractAudio:    true,
						ExtractLocation: &demuxOutputStorageLocation,
						ExtractTask: &hybrik.DoViMP4MuxExtractTask{
							RetryMethod: retryMethodRetry,
							Retry: hybrik.Retry{
								Count:    retryCountDefault,
								DelaySec: retryDelayDefault,
							},
							Name: "Demux Audio",
						},
					}},
				},
			},
		},
	}}}, nil
}
