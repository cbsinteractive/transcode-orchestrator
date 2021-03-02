package hybrik

/*
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

func (p *driver) dolbyVisionLegacyElementAssembler(cfg jobCfg) ([][]hy.Element, error) {
	presetsWithoutAudio := map[string]job.File{}
	for _, outputCfg := range cfg.outputCfgs {
		preset := outputCfg.Preset

		// removing audio so we can processing this separately
		preset.Audio = job.Audio{}
		presetsWithoutAudio[outputCfg.FileName] = preset
	}

	transcodeElementsWithFilenames, err := p.transcodeElementsWithPresetsFrom(presetsWithoutAudio, cfg)
	if err != nil {
		return nil, err
	}

	transcodeElements := []hy.Element{}
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
	if tag, found := cfg.computeTags[job.ComputeClassDolbyVisionPreprocess]; found {
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

	return [][]hy.Element{{{
		UID:  dolbyVisionElementID,
		Kind: elementKindDolbyVision,
		Task: &hy.ElementTaskOptions{
			Tags: []string{preprocComputeTag},
		},
		Payload: hy.DolbyVisionTaskPayload{
			Module:  doViModuleProfile,
			Profile: doViProfile5,
			MezzanineQC: hy.DoViMezzanineQC{
				Enabled: false,
				Location: p.transcodeLocationFrom(storageLocation{
					provider: cfg.destination.provider,
					path:     fmt.Sprintf(doViLegacyMezzQCOutputPathTmpl, cfg.destination.path),
				}, cfg.executionEnvironment.OutputAlias),
				Task: hy.TaskTags{
					Tags: []string{preprocComputeTag},
				},
				FilePattern: fmt.Sprintf(doViMezzQCReportFilenameTmpl, cfg.jobID),
				ToolVersion: hy.DoViMezzQCVersionDefault,
			},
			NBCPreproc: hy.DoViNBCPreproc{
				Task: hy.TaskTags{
					Tags: []string{preprocComputeTag},
				},
				Location: p.transcodeLocationFrom(storageLocation{
					provider: cfg.destination.provider,
					path:     fmt.Sprintf(doViNBCPreProcOutputPathTmpl, cfg.destination.path),
				}, cfg.executionEnvironment.OutputAlias),
				SDKVersion:     hy.DoViSDKVersionDefault,
				NumTasks:       doViPreProcNumTasks,
				IntervalLength: doViPreProcIntervalLength,
				CLIOptions: hy.DoViNBCPreprocCLIOptions{
					InputEDRAspect: doViInputEDRAspectDefault,
					InputEDRPad:    doViInputEDRCropDefault,
					InputEDRCrop:   doViInputEDRPadDefault,
				},
			},
			Transcodes: transcodeElements,
			PostTranscode: hy.DoViPostTranscode{
				Task: &hy.TaskTags{Tags: []string{preprocComputeTag}},
				VESMux: &hy.DoViVESMux{
					Enabled: true,
					Location: p.transcodeLocationFrom(storageLocation{
						provider: cfg.destination.provider,
						path:     fmt.Sprintf(doViVESMuxOutputPathTmpl, cfg.destination.path),
					}, cfg.executionEnvironment.OutputAlias),
					FilePattern: doViVESMuxFilenameDefault,
					SDKVersion:  hy.DoViSDKVersionDefault,
				},
				MetadataPostProc: &hy.DoViMetadataPostProc{
					Enabled: true,
					Location: p.transcodeLocationFrom(storageLocation{
						provider: cfg.destination.provider,
						path:     fmt.Sprintf(doViMetadataPostProcOutputPathTmpl, cfg.destination.path),
					}, cfg.executionEnvironment.OutputAlias),
					FilePattern: doViMetadataPostProcFilenameDefault,
					SDKVersion:  hy.DoViSDKVersionDefault,
					QCSettings: hy.DoViQCSettings{
						Enabled:     true,
						ToolVersion: hy.DoViVesQCVersionDefault,
						Location: p.transcodeLocationFrom(storageLocation{
							provider: cfg.destination.provider,
							path:     fmt.Sprintf(doViMetadataPostProcQCOutputPathTmpl, cfg.destination.path),
						}, cfg.executionEnvironment.OutputAlias),
						FilePattern: doViMetadataPostProcQCFilenameDefault,
					},
				},
				MP4Mux: hy.DoViMP4Mux{
					Enabled:            true,
					Location:           &destStorageLocation,
					FilePattern:        doViMP4MuxFilenameDefault,
					ToolVersion:        hy.DoViMP4MuxerVersionDefault,
					CopySourceStartPTS: true,
					CLIOptions:         map[string]string{doViMP4MuxDVH1FlagKey: ""},
					QCSettings: &hy.DoViQCSettings{
						Enabled:     true,
						ToolVersion: hy.DoViMP4QCVersionDefault,
						Location: p.transcodeLocationFrom(storageLocation{
							provider: cfg.destination.provider,
							path:     fmt.Sprintf(doViMP4MuxQCOutputPathTmpl, cfg.destination.path),
						}, cfg.executionEnvironment.OutputAlias),
						FilePattern: doViMP4MuxQCFilenameDefault,
					},
					ElementaryStreams: []hy.DoViMP4MuxElementaryStream{{
						AssetURL: p.assetURLFrom(storageLocation{
							provider: cfg.sourceLocation.provider,
							path:     cfg.sourceLocation.path,
						}, cfg.executionEnvironment.OutputAlias),
						ExtractAudio:    true,
						ExtractLocation: &demuxOutputStorageLocation,
						ExtractTask: &hy.DoViMP4MuxExtractTask{
							RetryMethod: retryMethodRetry,
							Retry: hy.Retry{
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

*/
