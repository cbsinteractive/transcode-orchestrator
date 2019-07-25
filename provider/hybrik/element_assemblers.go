package hybrik

import (
	"fmt"
	"strconv"

	"github.com/cbsinteractive/hybrik-sdk-go"
)

func (p *hybrikProvider) defaultElementAssembler(cfg jobCfg) ([]hybrik.Element, error) {
	elements := []hybrik.Element{}

	idx := 0
	for _, outputCfg := range cfg.outputCfgs {
		e := p.mountTranscodeElement(strconv.Itoa(idx), cfg.jobID, outputCfg.filename, cfg.destBase,
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

	transcodeElements, err := transcodeElementsFromPresets(presets, cfg.destBase, cfg.executionFeatures)
	if err != nil {
		return nil, err
	}

	return []hybrik.Element{{
		UID:  dolbyVisionElementID,
		Kind: elementKindDolbyVision,
		Payload: hybrik.DolbyVisionTaskPayload{
			Module:  "profile",
			Profile: 5,
			MezzanineQC: hybrik.DoViMezzanineQC{
				Enabled: false,
				Location: hybrik.TranscodeLocation{
					StorageProvider: "s3",
					Path:            fmt.Sprintf("%s/mezzanine_qc", cfg.destBase),
				},
				FilePattern: fmt.Sprintf("%s_mezz_qc_report.txt", cfg.jobID),
				ToolVersion: hybrik.DoViMezzQCVersionDefault,
			},
			NBCPreproc: hybrik.DoViNBCPreproc{
				Task: hybrik.TaskTags{
					Tags: []string{"preproc"},
				},
				Location: hybrik.TranscodeLocation{
					StorageProvider: "s3",
					Path:            fmt.Sprintf("%s/nbc_preproc", cfg.destBase),
				},
				SDKVersion:     hybrik.DoViSDKVersionDefault,
				NumTasks:       "auto", // TODO wrap in constant
				IntervalLength: 48,
				CLIOptions: hybrik.DoViNBCPreprocCLIOptions{
					InputEDRAspect: "2",
					InputEDRPad:    "0x0x0x0",
					InputEDRCrop:   "0x0x0x0",
				},
			},
			Transcodes: transcodeElements,
			PostTranscode: hybrik.DoViPostTranscode{
				VESMux: hybrik.DoViVESMux{
					Enabled: true,
					Location: hybrik.TranscodeLocation{
						StorageProvider: "s3",
						Path:            fmt.Sprintf("%s/vesmuxer", cfg.destBase),
					},
					FilePattern: "ves.h265",
					SDKVersion:  hybrik.DoViSDKVersionDefault,
				},
				MetadataPostProc: hybrik.DoViMetadataPostProc{
					Enabled: true,
					Location: hybrik.TranscodeLocation{
						StorageProvider: "s3",
						Path:            fmt.Sprintf("%s/metadata_postproc", cfg.destBase),
					},
					FilePattern: "ves.265",
					SDKVersion:  hybrik.DoViSDKVersionDefault,
					QCSettings: hybrik.DoViQCSettings{
						Enabled:     true,
						ToolVersion: hybrik.DoViVesQCVersionDefault,
						Location: hybrik.TranscodeLocation{
							StorageProvider: "s3",
							Path:            fmt.Sprintf("%s/metadata_postproc_qc", cfg.destBase),
						},
						FilePattern: "metadata_postproc_ves_qc_report.txt",
					},
				},
				MP4Mux: hybrik.DoViMP4Mux{
					Enabled: true,
					Location: hybrik.TranscodeLocation{
						StorageProvider: "s3",
						Path:            fmt.Sprintf("%s/mp4muxer", cfg.destBase),
					},
					FilePattern:        "mux_output.mp4",
					ToolVersion:        hybrik.DoViMP4MuxerVersionDefault,
					CopySourceStartPTS: true,
					QCSettings: hybrik.DoViQCSettings{
						Enabled:     true,
						ToolVersion: hybrik.DoViMP4QCVersionDefault,
						Location: hybrik.TranscodeLocation{
							StorageProvider: "s3",
							Path:            fmt.Sprintf("%s/mp4_qc", cfg.destBase),
						},
						FilePattern: "mp4_qc_report.txt",
					},
					ElementaryStreams: []hybrik.DoViMP4MuxElementaryStream{{
						AssetURL: hybrik.AssetURL{
							StorageProvider: "s3",
							URL:             cfg.assetURL,
						},
						ExtractAudio: true,
						ExtractLocation: hybrik.TranscodeLocation{
							StorageProvider: "s3",
							Path:            fmt.Sprintf("%s/source_demux", cfg.destBase),
						},
						ExtractTask: hybrik.DoViMP4MuxExtractTask{
							RetryMethod: "retry",
							Retry: hybrik.Retry{
								Count:    3,
								DelaySec: 30,
							},
							Name: "Demux Audio",
						},
					}},
				},
			},
		},
	}}, nil
}
