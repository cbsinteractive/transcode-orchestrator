package hybrik

import (
	"fmt"

	hy "github.com/cbsinteractive/hybrik-sdk-go"
	"github.com/cbsinteractive/transcode-orchestrator/client/transcoding/job"
)

func (p *driver) dolbyVisionLegacy(j *Job) [][]hy.Element {
	txcode := p.transcodeElems(mute(*j))

	var preproc = struct {
		Tasks    string `json:"doViPreProcNumTasks,omitempty"`
		Interval int    `json:"doViPreProcIntervalLength,omitempty"`
	}{"auto", 48}

	features0(j, &preproc)

	tag := tag(j, job.TagDolbyVisionPreprocess, "preproc")
	src := p.assetURL(&j.Input, p.auth(j).Write)
	dst := p.location(j.Dir(), p.auth(j).Write)
	tmp := p.location(j.Dir().Join("tmp"), p.auth(j).Write)

	// src := p.assetURLFrom(storageLocation{provider: cfg.sourceLocation.provider, path: cfg.sourceLocation.path}, cfg.executionEnvironment.OutputAlias)
	//dst := p.transcodeLocationFrom(storageLocation{provider: cfg.destination.provider,path:     cfg.destination.path,}, cfg.executionEnvironment.OutputAlias)
	// tmp := p.transcodeLocationFrom(storageLocation{provider: cfg.destination.provider, path: fmt.Sprintf("%s/tmp", cfg.destination.path)}, cfg.executionEnvironment.OutputAlias)

	return [][]hy.Element{{{
		UID:  "dolby_vision_task",
		Kind: "dolby_vision",
		Task: &hy.ElementTaskOptions{Tags: tag},
		Payload: hy.DolbyVisionTaskPayload{
			Module: "profile", Profile: 5,
			MezzanineQC: hy.DoViMezzanineQC{
				Enabled:  false, // NOTE(as): why is this thing even set then?
				Location: tmp, FilePattern: fmt.Sprintf("%s_mezz_qc_report.txt", j.ID),
				Task:        hy.TaskTags{Tags: tag},
				ToolVersion: hy.DoViMezzQCVersionDefault,
			},
			NBCPreproc: hy.DoViNBCPreproc{
				Task:           hy.TaskTags{Tags: tag},
				Location:       tmp,
				SDKVersion:     hy.DoViSDKVersionDefault,
				NumTasks:       preproc.Tasks,
				IntervalLength: preproc.Interval,
				CLIOptions: hy.DoViNBCPreprocCLIOptions{
					InputEDRAspect: "2",
					InputEDRPad:    "0x0x0x0",
					InputEDRCrop:   "0x0x0x0",
				},
			},
			Transcodes: txcode,
			PostTranscode: hy.DoViPostTranscode{
				Task: &hy.TaskTags{Tags: tag},
				VESMux: &hy.DoViVESMux{
					Location: tmp, FilePattern: "ves.h265",
					SDKVersion: hy.DoViSDKVersionDefault, Enabled: true,
				},
				MetadataPostProc: &hy.DoViMetadataPostProc{
					Location: tmp, FilePattern: "postproc.265",
					QCSettings: hy.DoViQCSettings{
						Location: tmp, FilePattern: "metadata_postproc_qc_report.txt",
						ToolVersion: hy.DoViVesQCVersionDefault,
						Enabled:     true,
					},
					Enabled: true, SDKVersion: hy.DoViSDKVersionDefault,
				},
				MP4Mux: hy.DoViMP4Mux{
					Enabled:            true,
					FilePattern:        "{source_basename}.mp4",
					Location:           &dst,
					ToolVersion:        hy.DoViMP4MuxerVersionDefault,
					CopySourceStartPTS: true,
					CLIOptions:         map[string]string{"dvh1flag": ""},
					QCSettings: &hy.DoViQCSettings{
						Enabled:     true,
						FilePattern: "mp4_qc_report.txt",
						Location:    tmp,
						ToolVersion: hy.DoViMP4QCVersionDefault,
					},
					ElementaryStreams: []hy.DoViMP4MuxElementaryStream{{
						AssetURL:        src,
						ExtractAudio:    true,
						ExtractLocation: &tmp,
						ExtractTask: &hy.DoViMP4MuxExtractTask{
							RetryMethod: "retry",
							Retry:       hy.Retry{Count: 3, DelaySec: 30},
							Name:        "Demux Audio",
						},
					}},
				},
			},
		},
	}}}
}
