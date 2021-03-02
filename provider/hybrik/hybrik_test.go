package hybrik

import (
	"context"
	"reflect"
	"testing"

	"github.com/cbsinteractive/hybrik-sdk-go"
	"github.com/cbsinteractive/transcode-orchestrator/config"
	"github.com/cbsinteractive/transcode-orchestrator/job"
	"github.com/google/go-cmp/cmp"
)

type elementKind = string

const (
	elementKindTranscode   elementKind = "transcode"
	elementKindSource      elementKind = "source"
	elementKindPackage     elementKind = "package"
	elementKindDolbyVision elementKind = "dolby_vision"
)

var (
	testoutput = job.File{
		Name:      "file1.mp4",
		Container: "mp4",
		Video: job.Video{
			Profile:  "high",
			Level:    "4.1",
			Width:    300,
			Height:   400,
			Codec:    "h264",
			Bitrate:  job.Bitrate{BPS: 400000, Control: "CBR", TwoPass: true},
			Gop:      job.Gop{Size: 120},
			Scantype: "progressive",
		},
		Audio: job.Audio{
			Codec:   "aac",
			Bitrate: 20000,
		},
	}

	testjob = job.Job{
		ID:       "jobID",
		Provider: Name,
		Input:    job.File{Name: "s3://some/path.mp4"},
		Output: job.Dir{
			File: []job.File{testoutput},
		},
	}
)

// preset job.File, uid string, destination storageLocation, filename string,
//	execFeatures executionFeatures, computeTags map[job.ComputeClass]string
type transcodeCfg struct {
	uid                  string
	destination          storageLocation
	filename             string
	execFeatures         executionFeatures
	computeTags          map[job.ComputeClass]string
	executionEnvironment job.ExecutionEnvironment
}

// updates default preset for quick test of gop structs
func updateGopStruct(gopSize float64, gopUnit string) job.File {
	var p = defaultPreset
	p.Video.GopSize = gopSize
	p.Video.GopUnit = gopUnit

	return p
}

func TestPreset(t *testing.T) {
	tests := []struct {
		name                 string
		provider             *driver
		preset               job.File
		transcodeCfg         transcodeCfg
		wantTranscodeElement hy.TranscodePayload
		wantTags             []string
		wantErr              bool
	}{
		{
			name: "MP4/H264/AAC",
			provider: &driver{
				config: &config.Hybrik{
					PresetPath:        "some_preset_path",
					GCPCredentialsKey: "some_key",
				},
			},
			preset: defaultPreset,
			transcodeCfg: transcodeCfg{
				uid: "some_uid",
				destination: storageLocation{
					provider: storageProviderGCS,
					path:     "gs://some_bucket/encodes",
				},
				filename: "output.mp4",
				execFeatures: executionFeatures{
					segmentedRendering: &hy.SegmentedRendering{
						Duration: 60,
					},
				},
				computeTags: map[job.ComputeClass]string{
					job.ComputeClassTranscodeDefault: "transcode_default_tag",
				},
			},
			wantTranscodeElement: hy.TranscodePayload{
				SourcePipeline: hy.TranscodeSourcePipeline{
					SegmentedRendering: &hy.SegmentedRendering{
						Duration: 60,
					},
				},
				LocationTargetPayload: hy.LocationTargetPayload{
					Location: hy.TranscodeLocation{
						StorageProvider: storageProviderGCS.string(),
						Path:            "gs://some_bucket/encodes",
						Access:          &hy.StorageAccess{CredentialsKey: "some_key", MaxCrossRegionMB: -1},
					},
					Targets: []hy.TranscodeTarget{
						{
							FilePattern: "output.mp4",
							Container:   hy.TranscodeContainer{Kind: defaultPreset.Container},
							NumPasses:   2,
							Video: &hy.VideoTarget{
								Width:          intToPtr(300),
								Height:         intToPtr(400),
								Codec:          defaultPreset.Video.Codec,
								BitrateKb:      400,
								Preset:         presetSlow,
								ExactGOPFrames: 120,
								ChromaFormat:   chromaFormatYUV420P,
								BitrateMode:    "cbr",
								Profile:        "high",
								Level:          "4.1",
								InterlaceMode:  "progressive",
							},
							Audio: []hy.AudioTarget{
								{
									Codec:     defaultPreset.Audio.Codec,
									BitrateKb: 20,
									Channels:  2,
								},
							},
							ExistingFiles: "replace",
						},
					},
				},
			},
			wantTags: []string{"transcode_default_tag"},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			p, err := tt.provider.transcodeElementFromPreset(tt.preset, tt.transcodeCfg.uid, jobCfg{
				destination:       tt.transcodeCfg.destination,
				executionFeatures: tt.transcodeCfg.execFeatures,
				computeTags:       tt.transcodeCfg.computeTags,
			}, tt.transcodeCfg.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("driver.transcodeElementFromPreset() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if g, e := p.Payload, tt.wantTranscodeElement; !reflect.DeepEqual(g, e) {
				t.Fatalf("driver.transcodeElementFromPreset() wrong transcode payload\nWant %+v\nGot %+v\nDiff %s", e,
					g, cmp.Diff(e, g))
			}

			if tt.wantTags != nil {
				if g, e := p.Task.Tags, tt.wantTags; !reflect.DeepEqual(g, e) {
					t.Fatalf("driver.transcodeElementFromPreset() wrong preset request\nWant %+v\nGot %+v\nDiff %s", e,
						g, cmp.Diff(e, g))
				}
			}
		})
	}
}

func TestTranscodePreset(t *testing.T) {
	tests := []struct {
		name           string
		presetModifier func(preset job.File) job.File
		transcodeCfg   transcodeCfg
		assertion      func(hy.TranscodePayload, *testing.T)
		wantErrMsg     string
	}{
		{
			name: "HDR10",
			presetModifier: func(p job.File) job.File {
				p.Video.Codec = "h265"
				p.Video.Profile = ""

				p.Video.HDR10 = job.HDR10{
					Enabled:       true,
					MaxCLL:        10000,
					MaxFALL:       400,
					MasterDisplay: "G(13250,34500)B(7500,3000)R(34000,16000)WP(15635,16450)L(10000000,1)",
				}
				return p
			},
			assertion: func(input hy.TranscodePayload, t *testing.T) {
				transcodeTargets, ok := input.Targets.([]hy.TranscodeTarget)
				if !ok {
					t.Errorf("targets are not TranscodeTargets")
				}
				firstTarget := transcodeTargets[0]

				tests := []struct {
					name      string
					got, want interface{}
				}{
					{"masterdisplay", firstTarget.Video.HDR10.MasterDisplay, "G(13250,34500)B(7500,3000)R(34000,16000)WP(15635,16450)L(10000000,1)"},
					{"maxcll", firstTarget.Video.HDR10.MaxCLL, 10000},
					{"maxfall", firstTarget.Video.HDR10.MaxFALL, 400},
					{"colortrc", firstTarget.Video.ColorTRC, colorTRCSMPTE2084},
					{"colormatrix", firstTarget.Video.ColorMatrix, colorMatrixBT2020NC},
					{"colorformat", firstTarget.Video.ChromaFormat, chromaFormatYUV420P10LE},
					{"colorprimaries", firstTarget.Video.ColorPrimaries, colorPrimaryBT2020},
					{"codecprofile", firstTarget.Video.Profile, "main10"},
					{"vtag", firstTarget.Video.VTag, "hvc1"},
				}

				for _, tt := range tests {

					t.Run(tt.name, func(t *testing.T) {
						if g, e := tt.got, tt.want; !reflect.DeepEqual(g, e) {
							t.Fatalf("%s: got %q, expected %q", tt.name, g, e)
						}
					})
				}
			},
		},
		{
			name: "hevc/hdr10/mxf",
			presetModifier: func(p job.File) job.File {
				p.Video.Codec = "h265"
				p.Video.Profile = ""
				p.SourceContainer = "mxf"
				p.Video.HDR10 = job.HDR10{
					Enabled:       true,
					MaxCLL:        10000,
					MaxFALL:       400,
					MasterDisplay: "G(13250,34500)B(7500,3000)R(34000,16000)WP(15635,16450)L(10000000,1)",
				}

				// setting twoPass to false to ensure it's forced to true
				p.TwoPass = false

				return p
			},
			assertion: func(input hy.TranscodePayload, t *testing.T) {
				transcodeTargets, ok := input.Targets.([]hy.TranscodeTarget)
				if !ok {
					t.Errorf("targets are not TranscodeTargets")
				}
				firstTarget := transcodeTargets[0]

				tests := []struct {
					name      string
					got, want interface{}
				}{
					{"ffmpeg params", firstTarget.Video.FFMPEGArgs, ""},
					{"number of passes", firstTarget.NumPasses, 2},
					{"hdr10 input type", firstTarget.Video.HDR10.Source, "source_metadata"},
				}

				for _, tt := range tests {

					t.Run(tt.name, func(t *testing.T) {
						if g, e := tt.got, tt.want; !reflect.DeepEqual(g, e) {
							t.Fatalf("%s: got %q, expected %q", tt.name, g, e)
						}
					})
				}
			},
		},
		{
			name: "vbr",
			presetModifier: func(preset job.File) job.File {
				preset.RateControl = "vbr"
				preset.Video.Bitrate = 10000000
				return preset
			},
			assertion: func(input hy.TranscodePayload, t *testing.T) {
				transcodeTargets, ok := input.Targets.([]hy.TranscodeTarget)
				if !ok {
					t.Errorf("targets are not TranscodeTargets")
				}
				firstTarget := transcodeTargets[0]

				tests := []struct {
					name      string
					got, want interface{}
				}{
					{"ratecontrol", firstTarget.Video.BitrateMode, rateControlModeVBR},
					{"bitrate", firstTarget.Video.BitrateKb, 10000},
					{"min", firstTarget.Video.MinBitrateKb, 10000 * (100 - vbrVariabilityPercent) / 100},
					{"max", firstTarget.Video.MaxBitrateKb, 10000 * (100 + vbrVariabilityPercent) / 100},
				}

				for _, tt := range tests {

					t.Run(tt.name, func(t *testing.T) {
						if g, e := tt.got, tt.want; !reflect.DeepEqual(g, e) {
							t.Fatalf("%s: got %v, expected %v", tt.name, g, e)
						}
					})
				}
			},
		},
		{
			name: "ratecontrolErr",
			presetModifier: func(preset job.File) job.File {
				preset.RateControl = "fake_mode"
				return preset
			},
			wantErrMsg: `running "rateControl" transcode payload modifier: rate control mode "fake_mode" is not ` +
				`supported in hybrik, the currently supported modes are map[cbr:{} vbr:{}]`,
		},
		{
			name: "aws/maxCrossRegionMB/unset",
			transcodeCfg: transcodeCfg{
				destination: storageLocation{
					provider: storageProviderS3,
					path:     "s3://some_bucket/encodes",
				},
				executionEnvironment: job.ExecutionEnvironment{
					OutputAlias: "test_alias",
				},
			},
			presetModifier: func(p job.File) job.File {
				return p
			},
			assertion: func(payload hy.TranscodePayload, t *testing.T) {
				if maxCrossRegionMB := payload.Location.Access.MaxCrossRegionMB; maxCrossRegionMB != 0 {
					t.Errorf("maxCrossRegionMB was %d, expected it to be 0", maxCrossRegionMB)
				}
			},
		},
		{
			name: "gcs/maxCrossRegionMB/unlimited",
			transcodeCfg: transcodeCfg{
				uid: "some_uid",
				destination: storageLocation{
					provider: storageProviderGCS,
					path:     "gs://some_bucket/encodes",
				},
				filename: "output.mp4",
				executionEnvironment: job.ExecutionEnvironment{
					OutputAlias: "test_alias",
				},
			},
			presetModifier: func(p job.File) job.File {
				return p
			},
			assertion: func(payload hy.TranscodePayload, t *testing.T) {
				if maxCrossRegionMB := payload.Location.Access.MaxCrossRegionMB; maxCrossRegionMB != -1 {
					t.Errorf("maxCrossRegionMB was %d, expected it to be -1", maxCrossRegionMB)
				}
			},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			p := &driver{
				config: &config.Hybrik{
					PresetPath: "some_preset_path",
				},
			}

			gotElement, err := p.transcodeElementFromPreset(tt.presetModifier(defaultPreset), tt.transcodeCfg.uid, jobCfg{
				destination:          tt.transcodeCfg.destination,
				executionFeatures:    tt.transcodeCfg.execFeatures,
				computeTags:          tt.transcodeCfg.computeTags,
				executionEnvironment: tt.transcodeCfg.executionEnvironment,
			}, tt.transcodeCfg.filename)
			if err != nil && tt.wantErrMsg != err.Error() {
				t.Errorf("driver.transcodeElementFromPreset()error = %v, wantErr %q", err, tt.wantErrMsg)
				return
			}

			if tt.assertion != nil {
				tt.assertion(gotElement.Payload.(hy.TranscodePayload), t)
			}
		})
	}
}

func TestPresetConversion(t *testing.T) {
	tests := []struct {
		name    string
		job     job.Job
		preset  job.File
		wantJob hy.CreateJob
		wantErr string
	}{
		{
			name:   "MP4",
			job:    defaultJob,
			preset: defaultPreset,
			wantJob: hy.CreateJob{
				Name: "Job jobID [path.mp4]",
				Payload: hy.CreateJobPayload{
					Elements: []hy.Element{
						{
							UID:  "source_file",
							Kind: "source",
							Payload: hy.ElementPayload{
								Kind:    "asset_urls",
								Payload: []hy.AssetPayload{{StorageProvider: "s3", URL: "s3://some/path.mp4"}},
							},
						},
						{
							UID:  "transcode_task_0",
							Kind: "transcode",
							Task: &hy.ElementTaskOptions{Name: "Transcode - file1.mp4", Tags: []string{}},
							Payload: hy.TranscodePayload{
								LocationTargetPayload: hy.LocationTargetPayload{
									Location: hy.TranscodeLocation{
										Path:            "s3://some-dest/path/jobID",
										StorageProvider: storageProviderS3.string(),
									},
									Targets: []hy.TranscodeTarget{{
										Audio: []hy.AudioTarget{{
											BitrateKb: 20,
											Channels:  2,
											Codec:     "aac",
										}},
										Container: hy.TranscodeContainer{
											Kind: "mp4",
										},
										ExistingFiles: "replace",
										FilePattern:   "file1.mp4",
										NumPasses:     2,
										Video: &hy.VideoTarget{
											Width:          intToPtr(300),
											Height:         intToPtr(400),
											BitrateMode:    "cbr",
											BitrateKb:      400,
											Preset:         presetSlow,
											ChromaFormat:   chromaFormatYUV420P,
											Codec:          "h264",
											Profile:        "high",
											Level:          "4.1",
											ExactGOPFrames: 120,
											InterlaceMode:  "progressive",
										},
									}},
								},
							},
						},
					},
					Connections: []hy.Connection{
						{
							From: []hy.ConnectionFrom{{Element: "source_file"}},
							To: hy.ConnectionTo{
								Success: []hy.ToSuccess{{Element: "transcode_task_0"}},
							},
						},
					},
				},
			},
		},
		{
			name:   "gopSeconds",
			job:    defaultJob,
			preset: updateGopStruct(2, "seconds"),
			wantJob: hy.CreateJob{
				Name: "Job jobID [path.mp4]",
				Payload: hy.CreateJobPayload{
					Elements: []hy.Element{
						{
							UID:  "source_file",
							Kind: "source",
							Payload: hy.ElementPayload{
								Kind:    "asset_urls",
								Payload: []hy.AssetPayload{{StorageProvider: "s3", URL: "s3://some/path.mp4"}},
							},
						},
						{
							UID:  "transcode_task_0",
							Kind: "transcode",
							Task: &hy.ElementTaskOptions{Name: "Transcode - file1.mp4", Tags: []string{}},
							Payload: hy.TranscodePayload{
								LocationTargetPayload: hy.LocationTargetPayload{
									Location: hy.TranscodeLocation{
										Path:            "s3://some-dest/path/jobID",
										StorageProvider: storageProviderS3.string(),
									},
									Targets: []hy.TranscodeTarget{{
										Audio: []hy.AudioTarget{{
											BitrateKb: 20,
											Channels:  2,
											Codec:     "aac",
										}},
										Container: hy.TranscodeContainer{
											Kind: "mp4",
										},
										ExistingFiles: "replace",
										FilePattern:   "file1.mp4",
										NumPasses:     2,
										Video: &hy.VideoTarget{
											Width:         intToPtr(300),
											Height:        intToPtr(400),
											BitrateMode:   "cbr",
											BitrateKb:     400,
											Preset:        presetSlow,
											ChromaFormat:  chromaFormatYUV420P,
											Codec:         "h264",
											Profile:       "high",
											Level:         "4.1",
											ExactKeyFrame: 2,
											InterlaceMode: "progressive",
										},
									}},
								},
							},
						},
					},
					Connections: []hy.Connection{
						{
							From: []hy.ConnectionFrom{{Element: "source_file"}},
							To: hy.ConnectionTo{
								Success: []hy.ToSuccess{{Element: "transcode_task_0"}},
							},
						},
					},
				},
			},
		},
		// TODO uncomment once Hybrik fixes bug and we can re-enable the new structure
		//{
		//	name: "dolbyVision",
		//	job:  &defaultJob,
		//	preset: job.File{
		//		Name:        defaultPreset.Name,
		//		Description: defaultPreset.Description,
		//		Container:   "mp4",
		//		Video: job.Video{
		//			Profile:       "main10",
		//			Width:         "300",
		//			Codec:         "h265",
		//			Bitrate:       "12000",
		//			GopSize:       "120",
		//			GopMode:       "fixed",
		//			InterlaceMode: "progressive",
		//			DolbyVisionSettings: job.DolbyVisionSettings{
		//				Enabled: true,
		//			},
		//		},
		//	},
		//	wantJob: hy.CreateJob{
		//		Name: "Job jobID [path.mp4]",
		//		Payload: hy.CreateJobPayload{
		//			Elements: []hy.Element{
		//				{
		//					UID:  "source_file",
		//					Kind: "source",
		//					Payload: hy.ElementPayload{
		//						Kind:    "asset_urls",
		//						Payload: []hy.AssetPayload{{StorageProvider: "s3", URL: "s3://some/path.mp4"}},
		//					},
		//				},
		//				{
		//					UID:  "mezzanine_qc",
		//					Kind: "dolby_vision",
		//					Task: &hy.ElementTaskOptions{Name: "Mezzanine QC", Tags: []string{"preproc"}},
		//					Payload: hy.DoViV2MezzanineQCPayload{
		//						Module: "mezzanine_qc",
		//						Params: hy.DoViV2MezzanineQCPayloadParams{
		//							Location:    hy.TranscodeLocation{StorageProvider: "s3", Path: "s3://some-dest/path/jobID/mezzanine_qc"},
		//							FilePattern: "jobID_mezz_qc_report.txt",
		//						},
		//					},
		//				},
		//				{
		//					UID:  "dolby_vision_0",
		//					Kind: "dolby_vision",
		//					Task: &hy.ElementTaskOptions{
		//						Name:              "Encode #0",
		//						RetryMethod:       "fail",
		//						Tags:              []string{computeTagPreProcDefault},
		//						SourceElementUIDs: []string{"source_file"},
		//					},
		//					Payload: hy.DolbyVisionV2TaskPayload{
		//						Module:        "encoder",
		//						Profile:       5,
		//						Location:      hy.TranscodeLocation{StorageProvider: "s3", Path: "s3://some-dest/path/jobID"},
		//						Preprocessing: hy.DolbyVisionV2Preprocessing{Task: hy.TaskTags{Tags: []string{"preproc"}}},
		//						Transcodes: []hy.Element{
		//							{
		//								UID:  "transcode_task_0",
		//								Kind: "transcode",
		//								Task: &hy.ElementTaskOptions{Name: "Transcode - file1.mp4", Tags: []string{}},
		//								Payload: hy.TranscodePayload{
		//									LocationTargetPayload: hy.LocationTargetPayload{
		//										Location: hy.TranscodeLocation{StorageProvider: "s3", Path: "s3://some-dest/path/jobID"},
		//										Targets: []hy.TranscodeTarget{
		//											{
		//												FilePattern:   "file1.mp4",
		//												ExistingFiles: "replace",
		//												Container:     hy.TranscodeContainer{Kind: "elementary"},
		//												NumPasses:     1,
		//												Video: &hy.VideoTarget{
		//													Width:          intToPtr(300),
		//													BitrateKb:      12,
		//													Preset:         "slow",
		//													Codec:          "h265",
		//													Profile:        "main10",
		//													MinGOPFrames:   120,
		//													Tune:           "grain",
		//													ChromaFormat:   chromaFormatYUV420P10LE,
		//													MaxGOPFrames:   120,
		//													ExactGOPFrames: 120,
		//													InterlaceMode:  "progressive",
		//													X265Options:    "concatenation={auto_concatenation_flag}:vbv-init=0.6:vbv-end=0.6:annexb=1:hrd=1:aud=1:videoformat=5:range=full:colorprim=2:transfer=2:colormatrix=2:rc-lookahead=48:qg-size=32:scenecut=0:no-open-gop=1:frame-threads=0:repeat-headers=1:nr-inter=400:nr-intra=100:psy-rd=0:cbqpoffs=0:crqpoffs=3",
		//													VTag:           "hvc1",
		//													FFMPEGArgs:     " -strict experimental",
		//												},
		//												Audio: []hy.AudioTarget{},
		//											},
		//										},
		//									},
		//									Options: &hy.TranscodeTaskOptions{Pipeline: &hy.PipelineOptions{EncoderVersion: "hybrik_4.0_10bit"}},
		//								},
		//							},
		//						},
		//						PostTranscode: hy.DoViPostTranscode{
		//							Task: &hy.TaskTags{Tags: []string{computeTagPreProcDefault}},
		//							MP4Mux: hy.DoViMP4Mux{
		//								Enabled:           true,
		//								FilePattern:       "{source_basename}.mp4",
		//								ElementaryStreams: []hy.DoViMP4MuxElementaryStream{},
		//								CLIOptions: map[string]string{
		//									doViMP4MuxDVH1FlagKey: "",
		//								},
		//							},
		//						},
		//					},
		//				},
		//			},
		//			Connections: []hy.Connection{
		//				{
		//					From: []hy.ConnectionFrom{{Element: "source_file"}},
		//					To:   hy.ConnectionTo{Success: []hy.ToSuccess{{Element: "mezzanine_qc"}}},
		//				},
		//				{
		//					From: []hy.ConnectionFrom{{Element: "mezzanine_qc"}},
		//					To:   hy.ConnectionTo{Success: []hy.ToSuccess{{Element: "dolby_vision_0"}}},
		//				},
		//			},
		//		},
		//	},
		//},
		// TODO remove once Hybrik fixes bug and we can re-enable the new structure
		{
			name: "mp4dolbyVision",
			job:  defaultJob,
			preset: job.File{
				Name:        defaultPreset.Name,
				Description: defaultPreset.Description,
				Container:   "mp4",
				Video: job.Video{
					Profile:       "main10",
					Width:         300,
					Codec:         "h265",
					Bitrate:       12000,
					GopSize:       120,
					GopMode:       "fixed",
					InterlaceMode: "progressive",
					DolbyVision: job.DolbyVision{
						Enabled: true,
					},
				},
			},
			wantJob: hy.CreateJob{
				Name: "Job jobID [path.mp4]",
				Payload: hy.CreateJobPayload{
					Elements: []hy.Element{
						{
							UID:  "source_file",
							Kind: "source",
							Payload: hy.ElementPayload{
								Kind:    "asset_urls",
								Payload: []hy.AssetPayload{{StorageProvider: "s3", URL: "s3://some/path.mp4"}},
							},
						},
						{
							UID:  "dolby_vision_task",
							Kind: "dolby_vision",
							Task: &hy.ElementTaskOptions{
								Tags: []string{computeTagPreProcDefault},
							},
							Payload: hy.DolbyVisionTaskPayload{
								Module:  "profile",
								Profile: 5,
								MezzanineQC: hy.DoViMezzanineQC{
									Location:    hy.TranscodeLocation{StorageProvider: "s3", Path: "s3://some-dest/path/jobID/tmp"},
									Task:        hy.TaskTags{Tags: []string{"preproc"}},
									FilePattern: "jobID_mezz_qc_report.txt",
									ToolVersion: "2.6.2",
								},
								NBCPreproc: hy.DoViNBCPreproc{
									Task:           hy.TaskTags{Tags: []string{"preproc"}},
									Location:       hy.TranscodeLocation{StorageProvider: "s3", Path: "s3://some-dest/path/jobID/tmp"},
									SDKVersion:     "4.2.1_ga",
									NumTasks:       "auto",
									IntervalLength: 48,
									CLIOptions:     hy.DoViNBCPreprocCLIOptions{InputEDRAspect: "2", InputEDRPad: "0x0x0x0", InputEDRCrop: "0x0x0x0"},
								},
								Transcodes: []hy.Element{
									{
										UID:  "transcode_task_0",
										Kind: "transcode",
										Task: &hy.ElementTaskOptions{Name: "Transcode - file1.mp4", Tags: []string{}},
										Payload: hy.TranscodePayload{
											LocationTargetPayload: hy.LocationTargetPayload{
												Location: hy.TranscodeLocation{StorageProvider: "s3", Path: "s3://some-dest/path/jobID"},
												Targets: []hy.TranscodeTarget{
													{
														FilePattern:   "file1.mp4",
														ExistingFiles: "replace",
														Container:     hy.TranscodeContainer{Kind: "elementary"},
														NumPasses:     1,
														Video: &hy.VideoTarget{
															Width:          intToPtr(300),
															BitrateKb:      12,
															Preset:         "slow",
															Codec:          "h265",
															Profile:        "main10",
															Tune:           "grain",
															ExactGOPFrames: 120,
															InterlaceMode:  "progressive",
															ChromaFormat:   "yuv420p10le",
															X265Options:    "concatenation={auto_concatenation_flag}:vbv-init=0.6:vbv-end=0.6:annexb=1:hrd=1:aud=1:videoformat=5:range=full:colorprim=2:transfer=2:colormatrix=2:rc-lookahead=48:qg-size=32:scenecut=0:no-open-gop=1:frame-threads=0:repeat-headers=1:nr-inter=400:nr-intra=100:psy-rd=0:cbqpoffs=0:crqpoffs=3",
															VTag:           "hvc1",
															FFMPEGArgs:     " -strict experimental",
														},
														Audio: []hy.AudioTarget{},
													},
												},
											},
											Options: &hy.TranscodeTaskOptions{Pipeline: &hy.PipelineOptions{EncoderVersion: "hybrik_4.0_10bit"}},
										},
									},
								},
								PostTranscode: hy.DoViPostTranscode{
									Task: &hy.TaskTags{Tags: []string{"preproc"}},
									VESMux: &hy.DoViVESMux{
										Enabled:     true,
										Location:    hy.TranscodeLocation{StorageProvider: "s3", Path: "s3://some-dest/path/jobID/tmp"},
										FilePattern: "ves.h265",
										SDKVersion:  "4.2.1_ga",
									},
									MetadataPostProc: &hy.DoViMetadataPostProc{
										Enabled:     true,
										Location:    hy.TranscodeLocation{StorageProvider: "s3", Path: "s3://some-dest/path/jobID/tmp"},
										FilePattern: "postproc.265",
										SDKVersion:  "4.2.1_ga",
										QCSettings: hy.DoViQCSettings{
											Enabled:     true,
											ToolVersion: "0.9.0.9",
											Location:    hy.TranscodeLocation{StorageProvider: "s3", Path: "s3://some-dest/path/jobID/tmp"},
											FilePattern: "metadata_postproc_qc_report.txt",
										},
									},
									MP4Mux: hy.DoViMP4Mux{
										Enabled:            true,
										Location:           &hy.TranscodeLocation{StorageProvider: "s3", Path: "s3://some-dest/path/jobID"},
										FilePattern:        "{source_basename}.mp4",
										ToolVersion:        "1.2.8",
										CopySourceStartPTS: true,
										QCSettings: &hy.DoViQCSettings{
											Enabled:     true,
											ToolVersion: "1.1.4",
											Location:    hy.TranscodeLocation{StorageProvider: "s3", Path: "s3://some-dest/path/jobID/tmp"},
											FilePattern: "mp4_qc_report.txt",
										},
										CLIOptions: map[string]string{"dvh1flag": ""},
										ElementaryStreams: []hy.DoViMP4MuxElementaryStream{
											{
												AssetURL:        hy.AssetURL{StorageProvider: "s3", URL: "s3://some/path.mp4"},
												ExtractAudio:    true,
												ExtractLocation: &hy.TranscodeLocation{StorageProvider: "s3", Path: "s3://some-dest/path/jobID/tmp"},
												ExtractTask: &hy.DoViMP4MuxExtractTask{
													RetryMethod: "retry",
													Retry:       hy.Retry{Count: 3, DelaySec: 30},
													Name:        "Demux Audio",
												},
											},
										},
									},
								},
							},
						},
					},
					Connections: []hy.Connection{
						{
							From: []hy.ConnectionFrom{{Element: "source_file"}},
							To:   hy.ConnectionTo{Success: []hy.ToSuccess{{Element: "dolby_vision_task"}}},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			p := &driver{
				config: &config.Hybrik{
					Destination: "s3://some-dest/path",
					PresetPath:  "some_preset_path",
				},
			}

			tt.job.Outputs[0].Preset = tt.preset
			got, err := p.createJobReqFrom(context.Background(), &tt.job)
			if err != nil {
				if tt.wantErr != err.Error() {
					t.Errorf("driver.presetsToTranscodeJob() error = %v, wantErr %q", err, tt.wantErr)
				}

				return
			}

			if g, e := got, tt.wantJob; !reflect.DeepEqual(g, e) {
				t.Fatalf("driver.presetsToTranscodeJob() wrong job request\nWant %+v\nGot %+v\nDiff %s", e,
					g, cmp.Diff(e, g))
			}
		})
	}
}

func TestTranscodeJobFields(t *testing.T) {
	tests := []struct {
		name        string
		jobModifier func(job job.Job) job.Job
		assertion   func(hy.CreateJob, *testing.T)
		wantErrMsg  string
	}{
		{
			name: "dolbyVision",
			jobModifier: func(j job.Job) job.Job {
				j.SidecarAssets = map[job.SidecarAssetKind]string{
					job.SidecarAssetKindDolbyVisionMetadata: "s3://test_sidecar_location/path/file.xml",
				}

				return j
			},
			assertion: func(createJob hy.CreateJob, t *testing.T) {
				gotSource := createJob.Payload.Elements[0]

				expectSource := hy.Element{
					UID:  "source_file",
					Kind: "source",
					Payload: hy.ElementPayload{
						Kind: "asset_urls",
						Payload: []hy.AssetPayload{
							{StorageProvider: "s3", URL: "s3://some/path.mp4"},
							{
								StorageProvider: "s3",
								URL:             "s3://test_sidecar_location/path/file.xml",
								Contents: []hy.AssetContents{
									{
										Kind:    "metadata",
										Payload: hy.AssetContentsPayload{Standard: "dolbyvision_metadata"},
									},
								},
							},
						},
					},
				}

				if g, e := gotSource, expectSource; !reflect.DeepEqual(g, e) {
					t.Fatalf("driver.presetsToTranscodeJob() wrong job request\nWant %+v\nGot %+v\nDiff %s", e,
						g, cmp.Diff(e, g))
				}
			},
		},
		{
			name: "pathOverride",
			jobModifier: func(job job.Job) job.Job {
				job.DestinationBasePath = "s3://per-job-defined-bucket/some/base/path"
				return job
			},
			assertion: func(createJob hy.CreateJob, t *testing.T) {
				if len(createJob.Payload.Elements) < 2 {
					t.Error("job has less than two elements, tried to pull the second element (transcode)")
					return
				}
				gotTranscode := createJob.Payload.Elements[1]

				payload, ok := gotTranscode.Payload.(hy.TranscodePayload)
				if !ok {
					t.Error("transcode payload was not a map of string to map[string]interface{}")
					return
				}

				if g, e := payload.Location.Path, "s3://per-job-defined-bucket/some/base/path/jobID"; g != e {
					t.Errorf("destination location path: got %q, expected %q", g, e)
				}
			},
		},
		{
			name: "tags",
			jobModifier: func(j job.Job) job.Job {
				j.ExecutionEnv.ComputeTags = map[job.ComputeClass]string{
					job.ComputeClassTranscodeDefault: "custom_tag",
				}

				return j
			},
			assertion: func(createJob hy.CreateJob, t *testing.T) {
				gotTask := createJob.Payload.Elements[1].Task

				expectTask := &hy.ElementTaskOptions{Name: "Transcode - file1.mp4", Tags: []string{"custom_tag"}}

				if g, e := gotTask, expectTask; !reflect.DeepEqual(g, e) {
					t.Fatalf("driver.presetsToTranscodeJob() wrong job request\nWant %+v\nGot %+v\nDiff %s", e,
						g, cmp.Diff(e, g))
				}
			},
		},
		{
			name: "HLS",
			jobModifier: func(j job.Job) job.Job {
				j.StreamingParams = job.StreamingParams{
					SegmentDuration: 4,
					Protocol:        "hls",
				}

				j.ExecutionEnv.ComputeTags = map[job.ComputeClass]string{
					job.ComputeClassTranscodeDefault: "default_transcode_class",
				}

				return j
			},
			assertion: func(createJob hy.CreateJob, t *testing.T) {
				gotTask := createJob.Payload.Elements[len(createJob.Payload.Elements)-1]

				expectTask := hy.Element{
					UID:  "packager",
					Kind: elementKindPackage,
					Task: &hy.ElementTaskOptions{
						Tags: []string{"default_transcode_class"},
					},
					Payload: hy.PackagePayload{
						Location: hy.TranscodeLocation{
							StorageProvider: "s3",
							Path:            "s3://some-dest/path/jobID/hls",
						},
						FilePattern:        "master.m3u8",
						Kind:               "hls",
						SegmentationMode:   "segmented_mp4",
						SegmentDurationSec: 4,
						HLS: &hy.HLSPackagingSettings{
							IncludeIFRAMEManifests: true,
							HEVCCodecIDPrefix:      "hvc1",
						},
					},
				}

				if g, e := gotTask, expectTask; !reflect.DeepEqual(g, e) {
					t.Fatalf("driver.presetsToTranscodeJob() wrong package task\nWant %+v\nGot %+v\nDiff %s", e,
						g, cmp.Diff(e, g))
				}
			},
		},
		{
			name: "DASH",
			jobModifier: func(j job.Job) job.Job {
				j.StreamingParams = job.StreamingParams{
					SegmentDuration: 4,
					Protocol:        "dash",
				}

				return j
			},
			assertion: func(createJob hy.CreateJob, t *testing.T) {
				gotTask := createJob.Payload.Elements[len(createJob.Payload.Elements)-1]

				expectTask := hy.Element{
					UID:  "packager",
					Kind: elementKindPackage,
					Payload: hy.PackagePayload{
						Location: hy.TranscodeLocation{
							StorageProvider: "s3",
							Path:            "s3://some-dest/path/jobID/dash",
						},
						FilePattern:        "master.mpd",
						Kind:               "dash",
						SegmentationMode:   "segmented_mp4",
						SegmentDurationSec: 4,
						DASH: &hy.DASHPackagingSettings{
							SegmentationMode:   "segmented_mp4",
							SegmentDurationSec: "4",
						},
					},
				}

				if g, e := gotTask, expectTask; !reflect.DeepEqual(g, e) {
					t.Fatalf("driver.presetsToTranscodeJob() wrong package task\nWant %+v\nGot %+v\nDiff %s", e,
						g, cmp.Diff(e, g))
				}
			},
		},
		{
			name: "segmentedRenderingS3",
			jobModifier: func(j job.Job) job.Job {
				j.SourceMedia = "s3://bucket/path/file.mp4"
				j.ExecutionFeatures = job.ExecutionFeatures{
					featureSegmentedRendering: SegmentedRendering{Duration: 50},
				}

				return j
			},
			assertion: func(createJob hy.CreateJob, t *testing.T) {
				elements := createJob.Payload.Elements
				transcode, ok := elements[len(elements)-1].Payload.(hy.TranscodePayload)
				if !ok {
					t.Error("could not find a transcode payload in the job")
					return
				}

				segRendering := transcode.SourcePipeline.SegmentedRendering
				if segRendering == nil {
					t.Error("segmented rendering was nil, expected segmented rendering to be set")
					return
				}

				if g, e := segRendering.Duration, 50; g != e {
					t.Fatalf("driver.presetsToTranscodeJob() wrong segmented rendering du"+
						"ration:\nGot %d\nWant %d", g, e)
				}
			},
		},
		{
			name: "segmentedRenderingGCS",
			jobModifier: func(j job.Job) job.Job {
				j.SourceMedia = "gs://bucket/path/file.mp4"
				j.ExecutionFeatures = job.ExecutionFeatures{
					featureSegmentedRendering: SegmentedRendering{Duration: 50},
				}
				return j
			},
			assertion: func(createJob hy.CreateJob, t *testing.T) {
				elements := createJob.Payload.Elements
				transcode, ok := elements[len(elements)-1].Payload.(hy.TranscodePayload)
				if !ok {
					t.Error("could not find a transcode payload in the job")
					return
				}

				segRendering := transcode.SourcePipeline.SegmentedRendering
				if segRendering == nil {
					t.Error("segmented rendering was nil, expected segmented rendering to be set")
					return
				}

				if g, e := segRendering.Duration, 50; g != e {
					t.Fatalf("driver.presetsToTranscodeJob() wrong segmented rendering du"+
						"ration:\nGot %d\nWant %d", g, e)
				}
			},
		},
		{
			name: "segmentedRenderingHTTP",
			jobModifier: func(j job.Job) job.Job {
				j.SourceMedia = "http://example.com/path/file.mp4"
				j.ExecutionFeatures = job.ExecutionFeatures{
					featureSegmentedRendering: SegmentedRendering{Duration: 50},
				}

				return j
			},
			assertion: func(createJob hy.CreateJob, t *testing.T) {
				elements := createJob.Payload.Elements
				transcode, ok := elements[len(elements)-1].Payload.(hy.TranscodePayload)
				if !ok {
					t.Error("could not find a transcode payload in the job")
					return
				}

				segRendering := transcode.SourcePipeline.SegmentedRendering
				if segRendering != nil {
					t.Errorf("segmented rendering was %+v, expected segmented rendering to be nil", segRendering)
					return
				}
			},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			p := &driver{
				config: &config.Hybrik{
					Destination: "s3://some-dest/path",
					PresetPath:  "some_preset_path",
				},
			}
			j := defaultJob
			j.Outputs[0].Preset = defaultPreset
			j = tt.jobModifier(j)
			got, err := p.createJobReqFrom(context.Background(), &j)
			if err != nil && tt.wantErrMsg != err.Error() {
				t.Errorf("driver.presetsToTranscodeJob() error = %v, wantErr %q", err, tt.wantErrMsg)
				return
			}

			if tt.assertion != nil {
				tt.assertion(got, t)
			}
		})
	}
}

func intToPtr(i int) *int {
	return &i
}
