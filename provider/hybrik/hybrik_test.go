package hybrik

import (
	"reflect"
	"testing"

	hy "github.com/cbsinteractive/hybrik-sdk-go"
	"github.com/cbsinteractive/transcode-orchestrator/client/transcoding/job"
	"github.com/cbsinteractive/transcode-orchestrator/config"
	"github.com/google/go-cmp/cmp"
)

const (
	elementKindTranscode   = "transcode"
	elementKindSource      = "source"
	elementKindPackage     = "package"
	elementKindDolbyVision = "dolby_vision"
)

var (
	defaultPreset = job.File{
		Name: "file1.mp4", Container: "mp4",
		Video: job.Video{
			Profile: "high", Level: "4.1", Width: 300, Height: 400, Codec: "h264",
			Bitrate: job.Bitrate{BPS: 400000, Control: "CBR", TwoPass: true},
			Gop:     job.Gop{Size: 120}, Scantype: "progressive",
		},
		Audio: job.Audio{Codec: "aac", Bitrate: 20000},
	}

	defaultJob    = testjob
	jobGopSeconds = testjob

	testjob = job.Job{
		ID: "jobID", Provider: Name, Input: job.File{Name: "s3://some/path.mp4"},
		Output: job.Dir{
			Path: "s3://some-dest/path",
			File: []job.File{{
				Name: "file1.mp4",
				Video: job.Video{
					Profile: "high", Level: "4.1", Width: 300, Height: 400, Codec: "h264",
					Bitrate: job.Bitrate{BPS: 400000, Control: "CBR", TwoPass: true},
					Gop:     job.Gop{Size: 120}, Scantype: "progressive",
				},
				Audio: job.Audio{Codec: "aac", Bitrate: 20000}}},
		},
	}
)

func init() {
	jobGopSeconds.Output.File[0].Video.Gop = job.Gop{Size: 2, Unit: "seconds"}
}

func TestPreset(t *testing.T) {
	tests := []struct {
		name                 string
		input                *job.Job
		provider             *driver
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
			input: &job.Job{
				ID: "some_uid", Provider: Name,
				Input: job.File{Name: "s3://some/path.mp4"},
				Output: job.Dir{
					Path: "gs://some_bucket/encodes",
					File: []job.File{{
						Name: "output.mp4", Container: "mp4",
						Video: job.Video{
							Profile: "high", Level: "4.1", Width: 300, Height: 400, Codec: "h264",
							Bitrate: job.Bitrate{BPS: 400000, Control: "CBR", TwoPass: true},
							Gop:     job.Gop{Size: 120}, Scantype: "progressive",
						},
						Audio: job.Audio{Codec: "aac", Bitrate: 20000},
					}},
				},
				Features: map[string]interface{}{
					"segmentedRendering": &SegmentedRendering{Duration: 60},
				},
				Env: job.Env{
					Tags: map[string]string{job.TagTranscodeDefault: "transcode_default_tag"},
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
						Path:            "gs://some_bucket/encodes/some_uid",
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
			p := tt.provider.transcodeElems(tt.input)[0]
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
		name       string
		input      *job.Job
		assertion  func(hy.TranscodePayload, *testing.T)
		wantErrMsg string
	}{
		{
			name: "HDR10",
			input: &job.Job{
				ID: "jobID", Provider: Name, Input: job.File{Name: "s3://some/path.mp4"},
				Output: job.Dir{
					Path: "gs://some_bucket/encodes",
					File: []job.File{{
						Name: "file1.mp4",
						Video: job.Video{
							Profile: "high", Level: "4.1", Width: 300, Height: 400, Codec: "h265",
							Bitrate: job.Bitrate{BPS: 400000, Control: "CBR", TwoPass: true},
							Gop:     job.Gop{Size: 120}, Scantype: "progressive",
							HDR10: job.HDR10{
								Enabled: true, MaxCLL: 10000, MaxFALL: 400,
								MasterDisplay: "G(13250,34500)B(7500,3000)R(34000,16000)WP(15635,16450)L(10000000,1)",
							},
						},
						Audio: job.Audio{Codec: "aac", Bitrate: 20000}}},
				},
			},
			assertion: func(input hy.TranscodePayload, t *testing.T) {
				transcodeTargets, ok := input.Targets.([]hy.TranscodeTarget)
				if !ok {
					t.Errorf("targets are not TranscodeTargets")
				}
				t0 := transcodeTargets[0]

				tests := []struct {
					name      string
					got, want interface{}
				}{
					{"masterdisplay", t0.Video.HDR10.MasterDisplay, "G(13250,34500)B(7500,3000)R(34000,16000)WP(15635,16450)L(10000000,1)"},
					{"maxcll", t0.Video.HDR10.MaxCLL, 10000},
					{"maxfall", t0.Video.HDR10.MaxFALL, 400},
					{"colortrc", t0.Video.ColorTRC, colorTRCSMPTE2084},
					{"colormatrix", t0.Video.ColorMatrix, colorMatrixBT2020NC},
					{"colorformat", t0.Video.ChromaFormat, chromaFormatYUV420P10LE},
					{"colorprimaries", t0.Video.ColorPrimaries, colorPrimaryBT2020},
					{"codecprofile", t0.Video.Profile, "main10"},
					{"vtag", t0.Video.VTag, "hvc1"},
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
			input: &job.Job{
				ID: "jobID", Provider: Name, Input: job.File{Name: "s3://some/in.mxf"},
				Output: job.Dir{
					File: []job.File{{
						Name: "file1.mp4",
						Video: job.Video{
							Profile: "high", Level: "4.1", Width: 300, Height: 400, Codec: "h265",
							Bitrate: job.Bitrate{BPS: 400000, Control: "CBR", TwoPass: false},
							Gop:     job.Gop{Size: 120}, Scantype: "progressive",
							HDR10: job.HDR10{
								Enabled: true, MaxCLL: 10000, MaxFALL: 400,
								MasterDisplay: "G(13250,34500)B(7500,3000)R(34000,16000)WP(15635,16450)L(10000000,1)",
							},
						},
						Audio: job.Audio{Codec: "aac", Bitrate: 20000}}},
				},
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
			input: &job.Job{
				ID: "jobID", Provider: Name, Input: job.File{Name: "s3://some/in.mxf"},
				Output: job.Dir{
					File: []job.File{{
						Name: "file1.mp4",
						Video: job.Video{
							Profile: "high", Level: "4.1", Width: 300, Height: 400, Codec: "h265",
							Bitrate: job.Bitrate{BPS: 10000000, Control: "vbr", TwoPass: true},
							Gop:     job.Gop{Size: 120}, Scantype: "progressive",
							HDR10: job.HDR10{
								Enabled: true, MaxCLL: 10000, MaxFALL: 400,
								MasterDisplay: "G(13250,34500)B(7500,3000)R(34000,16000)WP(15635,16450)L(10000000,1)",
							},
						},
						Audio: job.Audio{Codec: "aac", Bitrate: 20000}}},
				},
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
					{"ratecontrol", firstTarget.Video.BitrateMode, "vbr"},
					{"bitrate", firstTarget.Video.BitrateKb, 10000},
					{"min", firstTarget.Video.MinBitrateKb, 9000},
					{"max", firstTarget.Video.MaxBitrateKb, 11000},
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &driver{
				config: &config.Hybrik{
					PresetPath: "some_preset_path",
				},
			}

			gotElement := p.transcodeElems(tt.input)[0]

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
			preset: jobGopSeconds.Output.File[0],
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
				Name: defaultPreset.Name, Container: "mp4",
				Video: job.Video{
					Codec: "h265", Profile: "main10",
					Width: 300, Scantype: "progressive",
					Bitrate: job.Bitrate{BPS: 12000}, Gop: job.Gop{Size: 120, Mode: "fixed"},
					DolbyVision: job.DolbyVision{Enabled: true},
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
								Tags: []string{"preproc"},
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

			tt.job.Output.File[0] = tt.preset
			got, err := p.jobRequest(&tt.job)
			if err != nil {
				if tt.wantErr != err.Error() {
					t.Errorf("driver.presetsToTranscodeJob() error = %v, wantErr %q", err, tt.wantErr)
				}

				return
			}

			if g, e := *got, tt.wantJob; !reflect.DeepEqual(g, e) {
				t.Fatalf("driver.presetsToTranscodeJob() wrong job request\nWant %+v\nGot %+v\nDiff %s", e,
					g, cmp.Diff(e, g))
			}
		})
	}
}

func lastPayload(t *testing.T, j hy.CreateJob) hy.TranscodePayload {
	t.Helper()
	p := j.Payload.Elements
	return p[len(p)-1].Payload.(hy.TranscodePayload)
}

func TestDolbyVisionMetadata(t *testing.T) {
	j := job.Job{
		ID: "jobID", Provider: Name, Input: job.File{Name: "s3://some/path.mp4"},
		ExtraFiles: map[string]string{
			job.TagDolbyVisionMetadata: "s3://test_sidecar_location/path/file.xml",
		},
		Output: job.Dir{
			Path: "s3://some-dest/path",
			File: []job.File{{
				Name: "file1.mp4",
				Video: job.Video{
					Profile: "high", Level: "4.1", Width: 300, Height: 400, Codec: "h264",
					Bitrate: job.Bitrate{BPS: 400000, Control: "CBR", TwoPass: true},
					Gop:     job.Gop{Size: 120}, Scantype: "progressive",
				},
				Audio: job.Audio{Codec: "aac", Bitrate: 20000}}},
		},
	}

	want := hy.Element{
		UID: "source_file", Kind: "source",
		Payload: hy.ElementPayload{
			Kind: "asset_urls",
			Payload: []hy.AssetPayload{
				{StorageProvider: "s3", URL: "s3://some/path.mp4"},
				{StorageProvider: "s3", URL: "s3://test_sidecar_location/path/file.xml",
					Contents: []hy.AssetContents{{Kind: "metadata", Payload: hy.AssetContentsPayload{Standard: "dolbyvision_metadata"}}},
				},
			},
		},
	}

	p := &driver{
		config: &config.Hybrik{
			Destination: "s3://some-dest/path",
			PresetPath:  "some_preset_path",
		},
	}
	jr, err := p.jobRequest(&j)
	if err != nil {
		t.Fatal(err)
	}
	have := jr.Payload.Elements[0]
	if !reflect.DeepEqual(have, want) {
		t.Fatalf("\n\t\thave: +%v\n\t\twant: +%v", have, want)
	}
}

func TestSegmentedRendering(t *testing.T) {
	j := job.Job{
		Features: job.Features{"segmentedRendering": SegmentedRendering{Duration: 50}},
		Output:   job.Dir{Path: "s3://path", File: []job.File{{Name: "1.mp4", Video: job.Video{Codec: "h264"}}}}}
	t.Log(features(&j))

	p := &driver{config: &config.Hybrik{}}
	for _, tc := range []struct {
		input string
		want  *hy.SegmentedRendering
	}{
		{"s3://file.mp4", &hy.SegmentedRendering{Duration: 50}},
		{"gs://file.mp4", &hy.SegmentedRendering{Duration: 50}},
		{"http://file.mp4", nil},
	} {
		j.Input.Name = tc.input
		jr, _ := p.jobRequest(&j)
		have := lastPayload(t, *jr).SourcePipeline.SegmentedRendering
		if !reflect.DeepEqual(have, tc.want) {
			t.Fatalf("%q: \n\t\thave: %+v\n\t\twant: %+v", tc.input, have, tc.want)
		}
	}

	/*
		{
			name: "segmentedRenderingS3",
			jobModifier: func(j job.Job) job.Job {
				j.Input.Name = "s3://bucket/path/file.mp4"
				j.Features = job.Features{
					"segmentedRendering": SegmentedRendering{Duration: 50},
				}

				return j
			},
			assertion: func(j hy.CreateJob, t *testing.T) {
				r := lastPayload(t, j).SourcePipeline.SegmentedRendering
				if g, e := r.Duration, 50; g != e {
					t.Fatalf("duration:\nGot %d\nWant %d", g, e)
				}
			},
		},
		{
			name: "segmentedRenderingGCS",
			jobModifier: func(j job.Job) job.Job {
				j.Input.Name = "gs://bucket/path/file.mp4"
				j.Features = job.Features{
					"segmentedRendering": SegmentedRendering{Duration: 50},
				}
				return j
			},
			assertion: func(j hy.CreateJob, t *testing.T) {
				r := lastPayload(t, j).SourcePipeline.SegmentedRendering
				if g, e := r.Duration, 50; g != e {
					t.Fatalf("duration:\nGot %d\nWant %d", g, e)
				}
			},
		},
		{
			name: "segmentedRenderingHTTP",
			jobModifier: func(j job.Job) job.Job {
				j.Input.Name = "http://example.com/path/file.mp4"
				j.Features = job.Features{
					"segmentedRendering": SegmentedRendering{Duration: 50},
				}
				return j
			},
			assertion: func(j hy.CreateJob, t *testing.T) {
				r := lastPayload(t, j).SourcePipeline.SegmentedRendering
				if r != nil {
					t.Fatalf("segmented rendering set erroneously: %v", r)
				}
			},
		},
	*/
}

func TestTranscodeJobFields(t *testing.T) {
	tests := []struct {
		name        string
		jobModifier func(job job.Job) job.Job
		assertion   func(hy.CreateJob, *testing.T)
		wantErrMsg  string
	}{
		{
			name: "pathOverride",
			jobModifier: func(job job.Job) job.Job {
				job.Output.Path = "s3://per-job-defined-bucket/some/base/path"
				return job
			},
			assertion: func(c hy.CreateJob, t *testing.T) {
				if g, e := c.Payload.Elements[1].Payload.(hy.TranscodePayload).Location.Path, "s3://per-job-defined-bucket/some/base/path/jobID"; g != e {
					t.Errorf("destination location path: got %q, expected %q", g, e)
				}
			},
		},
		{
			name: "tags",
			jobModifier: func(j job.Job) job.Job {
				j.Env.Tags = map[string]string{
					job.TagTranscodeDefault: "custom_tag",
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
			j.Output.File[0] = defaultPreset
			j = tt.jobModifier(j)
			got, err := p.jobRequest(&j)
			if err != nil && tt.wantErrMsg != err.Error() {
				t.Fatalf("error = %v, wantErr %q", err, tt.wantErrMsg)
			}

			if tt.assertion != nil {
				tt.assertion(*got, t)
			}
		})
	}
}

func intToPtr(i int) *int {
	return &i
}
