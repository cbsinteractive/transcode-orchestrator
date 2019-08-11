package hybrik

import (
	"fmt"
	"strconv"

	"github.com/cbsinteractive/hybrik-sdk-go"
)

const (
	doViModuleProfile                     = "profile"
	doViProfile5                          = 5
	doViMezzQCOutputPathTmpl              = "%s/mezzanine_qc"
	doViMezzQCReportFilenameTmpl          = "%s_mezz_qc_report.txt"
	doViNBCPreProcOutputPathTmpl          = "%s/nbc_preproc"
	doViVESMuxOutputPathTmpl              = "%s/vesmuxer"
	doViVESMuxFilenameDefault             = "ves.h265"
	doViMetadataPostProcOutputPathTmpl    = "%s/metadata_postproc"
	doViMetadataPostProcFilenameDefault   = "postproc.265"
	doViMetadataPostProcQCOutputPathTmpl  = "%s/metadata_postproc_qc"
	doViMetadataPostProcQCFilenameDefault = "metadata_postproc_qc_report.txt"
	doViMP4MuxFilenameDefault             = "{source_basename}.mp4"
	doViMP4MuxQCOutputPathTmpl            = "%s/mp4_qc"
	doViMP4MuxQCFilenameDefault           = "mp4_qc_report.txt"
	doViSourceDemuxOutputPathTmpl         = "%s/source_demux"
	doViPreProcNumTasksAuto               = "auto"
	doViPreProcIntervalLengthDefault      = 48
	doViInputEDRAspectDefault             = "2"
	doViInputEDRCropDefault               = "0x0x0x0"
	doViInputEDRPadDefault                = "0x0x0x0"

	retryMethodRetry  = "retry"
	retryCountDefault = 3
	retryDelayDefault = 30

	computeTagPreProc = "preproc"
)

func (p *hybrikProvider) defaultElementAssembler(cfg jobCfg) ([]hybrik.Element, error) {
	elements := []hybrik.Element{}

	idx := 0
	for _, outputCfg := range cfg.outputCfgs {
		e := p.mountTranscodeElement(strconv.Itoa(idx), cfg.jobID, outputCfg.filename, cfg.destination,
			cfg.streamingParams.SegmentDuration, outputCfg.preset, cfg.executionFeatures)
		elements = append(elements, e)
		idx++
	}

	return elements, nil
}

func (p *hybrikProvider) dolbyVisionElementAssembler(cfg jobCfg) ([]hybrik.Element, error) {
	presets := map[string]hybrik.Preset{}
	for _, outputCfg := range cfg.outputCfgs {
		presets[outputCfg.filename] = outputCfg.preset
	}

	transcodeElements, err := transcodeElementsFromPresets(presets, cfg.destination, cfg.executionFeatures)
	if err != nil {
		return nil, err
	}

	doViPreProcNumTasks := doViPreProcNumTasksAuto
	if numTasks := cfg.executionFeatures.doViPreProcSegmentation.numTasks; numTasks != "" {
		doViPreProcNumTasks = numTasks
	}

	doViPreProcIntervalLength := doViPreProcIntervalLengthDefault
	if intervalLength := cfg.executionFeatures.doViPreProcSegmentation.intervalLength; intervalLength != 0 {
		doViPreProcIntervalLength = intervalLength
	}

	return []hybrik.Element{{
		UID:  dolbyVisionElementID,
		Kind: elementKindDolbyVision,
		Payload: hybrik.DolbyVisionTaskPayload{
			Module:  doViModuleProfile,
			Profile: doViProfile5,
			MezzanineQC: hybrik.DoViMezzanineQC{
				Enabled: false,
				Location: hybrik.TranscodeLocation{
					StorageProvider: cfg.destination.provider,
					Path:            fmt.Sprintf(doViMezzQCOutputPathTmpl, cfg.destination.path),
				},
				FilePattern: fmt.Sprintf(doViMezzQCReportFilenameTmpl, cfg.jobID),
				ToolVersion: hybrik.DoViMezzQCVersionDefault,
			},
			NBCPreproc: hybrik.DoViNBCPreproc{
				Task: hybrik.TaskTags{
					Tags: []string{computeTagPreProc},
				},
				Location: hybrik.TranscodeLocation{
					StorageProvider: cfg.destination.provider,
					Path:            fmt.Sprintf(doViNBCPreProcOutputPathTmpl, cfg.destination.path),
				},
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
				VESMux: hybrik.DoViVESMux{
					Enabled: true,
					Location: hybrik.TranscodeLocation{
						StorageProvider: cfg.destination.provider,
						Path:            fmt.Sprintf(doViVESMuxOutputPathTmpl, cfg.destination.path),
					},
					FilePattern: doViVESMuxFilenameDefault,
					SDKVersion:  hybrik.DoViSDKVersionDefault,
				},
				MetadataPostProc: hybrik.DoViMetadataPostProc{
					Enabled: true,
					Location: hybrik.TranscodeLocation{
						StorageProvider: cfg.destination.provider,
						Path:            fmt.Sprintf(doViMetadataPostProcOutputPathTmpl, cfg.destination.path),
					},
					FilePattern: doViMetadataPostProcFilenameDefault,
					SDKVersion:  hybrik.DoViSDKVersionDefault,
					QCSettings: hybrik.DoViQCSettings{
						Enabled:     true,
						ToolVersion: hybrik.DoViVesQCVersionDefault,
						Location: hybrik.TranscodeLocation{
							StorageProvider: cfg.destination.provider,
							Path:            fmt.Sprintf(doViMetadataPostProcQCOutputPathTmpl, cfg.destination.path),
						},
						FilePattern: doViMetadataPostProcQCFilenameDefault,
					},
				},
				MP4Mux: hybrik.DoViMP4Mux{
					Enabled: true,
					Location: hybrik.TranscodeLocation{
						StorageProvider: cfg.destination.provider,
						Path:            cfg.destination.path,
					},
					FilePattern:        doViMP4MuxFilenameDefault,
					ToolVersion:        hybrik.DoViMP4MuxerVersionDefault,
					CopySourceStartPTS: true,
					QCSettings: hybrik.DoViQCSettings{
						Enabled:     true,
						ToolVersion: hybrik.DoViMP4QCVersionDefault,
						Location: hybrik.TranscodeLocation{
							StorageProvider: cfg.destination.provider,
							Path:            fmt.Sprintf(doViMP4MuxQCOutputPathTmpl, cfg.destination.path),
						},
						FilePattern: doViMP4MuxQCFilenameDefault,
					},
					ElementaryStreams: []hybrik.DoViMP4MuxElementaryStream{{
						AssetURL: hybrik.AssetURL{
							StorageProvider: cfg.sourceLocation.provider,
							URL:             cfg.sourceLocation.path,
						},
						ExtractAudio: true,
						ExtractLocation: hybrik.TranscodeLocation{
							StorageProvider: cfg.destination.provider,
							Path:            fmt.Sprintf(doViSourceDemuxOutputPathTmpl, cfg.destination.path),
						},
						ExtractTask: hybrik.DoViMP4MuxExtractTask{
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
	}}, nil
}
