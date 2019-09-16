package hybrik

import (
	"reflect"
	"testing"

	"github.com/cbsinteractive/hybrik-sdk-go"
	"github.com/cbsinteractive/video-transcoding-api/config"
	"github.com/cbsinteractive/video-transcoding-api/db"
	"github.com/cbsinteractive/video-transcoding-api/db/dbtest"
	"github.com/google/go-cmp/cmp"
)

var (
	defaultPreset = db.Preset{
		Name:        "preset_name",
		Description: "test_desc",
		Container:   "mp4",
		RateControl: "CBR",
		TwoPass:     true,
		Video: db.VideoPreset{
			Profile:       "high",
			ProfileLevel:  "4.1",
			Width:         "300",
			Height:        "400",
			Codec:         "h264",
			Bitrate:       "400000",
			GopSize:       "120",
			InterlaceMode: "progressive",
		},
		Audio: db.AudioPreset{
			Codec:   "aac",
			Bitrate: "20000",
		},
	}

	defaultJob = db.Job{
		ID:           "jobID",
		ProviderName: Name,
		SourceMedia:  "s3://some/path.mp4",
		Outputs: []db.TranscodeOutput{
			{
				Preset: db.PresetMap{
					Name: "preset_name",
					ProviderMapping: map[string]string{
						"hybrik": "preset_name",
					},
				},
				FileName: "file1.mp4",
			},
		},
	}
)

// preset db.Preset, uid string, destination storageLocation, filename string,
//	execFeatures executionFeatures, computeTags map[db.ComputeClass]string
type transcodeCfg struct {
	uid          string
	destination  storageLocation
	filename     string
	execFeatures executionFeatures
	computeTags  map[db.ComputeClass]string
}

func TestHybrikProvider_transcodeElementFromPreset(t *testing.T) {
	tests := []struct {
		name                 string
		provider             *hybrikProvider
		preset               db.Preset
		transcodeCfg         transcodeCfg
		wantTranscodeElement hybrik.TranscodePayload
		wantTags             []string
		wantErr              bool
	}{
		{
			name: "a valid h264/aac mp4 preset results in the expected preset sent to the Hybrik API",
			provider: &hybrikProvider{
				config: &config.Hybrik{
					PresetPath:        "some_preset_path",
					GCPCredentialsKey: "some_key",
				},
				repository: dbtest.NewFakeRepository(false),
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
					segmentedRendering: &hybrik.SegmentedRendering{
						Duration: 60,
					},
				},
				computeTags: map[db.ComputeClass]string{
					db.ComputeClassTranscodeDefault: "transcode_default_tag",
				},
			},
			wantTranscodeElement: hybrik.TranscodePayload{
				SourcePipeline: hybrik.TranscodeSourcePipeline{
					SegmentedRendering: &hybrik.SegmentedRendering{
						Duration: 60,
					},
				},
				LocationTargetPayload: hybrik.LocationTargetPayload{
					Location: hybrik.TranscodeLocation{
						StorageProvider: storageProviderGCS,
						Path:            "gs://some_bucket/encodes",
						Access:          &hybrik.StorageAccess{CredentialsKey: "some_key"},
					},
					Targets: []hybrik.TranscodeTarget{
						{
							FilePattern: "output.mp4",
							Container:   hybrik.TranscodeContainer{Kind: defaultPreset.Container},
							NumPasses:   2,
							Video: &hybrik.VideoTarget{
								Width:          intToPtr(300),
								Height:         intToPtr(400),
								Codec:          defaultPreset.Video.Codec,
								BitrateKb:      400,
								Preset:         presetSlow,
								MaxGOPFrames:   120,
								ExactGOPFrames: 120,
								ChromaFormat:   chromaFormatYUV420P,
								MinGOPFrames:   120,
								BitrateMode:    "cbr",
								Profile:        "high",
								Level:          "4.1",
								InterlaceMode:  "progressive",
							},
							Audio: []hybrik.AudioTarget{
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
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			p, err := tt.provider.transcodeElementFromPreset(tt.preset, tt.transcodeCfg.uid, jobCfg{
				destination:       tt.transcodeCfg.destination,
				executionFeatures: tt.transcodeCfg.execFeatures,
				computeTags:       tt.transcodeCfg.computeTags,
			}, tt.transcodeCfg.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("hybrikProvider.transcodeElementFromPreset() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if g, e := p.Payload, tt.wantTranscodeElement; !reflect.DeepEqual(g, e) {
				t.Fatalf("hybrikProvider.transcodeElementFromPreset() wrong transcode payload\nWant %+v\nGot %+v\nDiff %s", e,
					g, cmp.Diff(e, g))
			}

			if tt.wantTags != nil {
				if g, e := p.Task.Tags, tt.wantTags; !reflect.DeepEqual(g, e) {
					t.Fatalf("hybrikProvider.transcodeElementFromPreset() wrong preset request\nWant %+v\nGot %+v\nDiff %s", e,
						g, cmp.Diff(e, g))
				}
			}
		})
	}
}

func TestHybrikProvider_transcodeElementFromPreset_fields(t *testing.T) {
	tests := []struct {
		name           string
		presetModifier func(preset db.Preset) db.Preset
		transcodeCfg   transcodeCfg
		assertion      func(hybrik.TranscodePayload, *testing.T)
		wantErrMsg     string
	}{
		{
			name: "hevc/hdr10 presets are set correctly",
			presetModifier: func(p db.Preset) db.Preset {
				p.Video.Codec = "h265"
				p.Video.Profile = ""

				p.Video.HDR10Settings = db.HDR10Settings{
					Enabled:       true,
					MaxCLL:        10000,
					MaxFALL:       400,
					MasterDisplay: "G(13250,34500)B(7500,3000)R(34000,16000)WP(15635,16450)L(10000000,1)",
				}
				return p
			},
			assertion: func(input hybrik.TranscodePayload, t *testing.T) {
				transcodeTargets, ok := input.Targets.([]hybrik.TranscodeTarget)
				if !ok {
					t.Errorf("targets are not TranscodeTargets")
				}
				firstTarget := transcodeTargets[0]

				tests := []struct {
					name      string
					got, want interface{}
				}{
					{
						name: "hdr10 master display",
						got:  firstTarget.Video.HDR10.MasterDisplay,
						want: "G(13250,34500)B(7500,3000)R(34000,16000)WP(15635,16450)L(10000000,1)",
					},
					{
						name: "hdr10 max cll",
						got:  firstTarget.Video.HDR10.MaxCLL,
						want: 10000,
					},
					{
						name: "hdr10 max fall",
						got:  firstTarget.Video.HDR10.MaxFALL,
						want: 400,
					},
					{
						name: "hdr10 color trc",
						got:  firstTarget.Video.ColorTRC,
						want: colorTRCSMPTE2084,
					},
					{
						name: "hdr10 color matrix",
						got:  firstTarget.Video.ColorMatrix,
						want: colorMatrixBT2020NC,
					},
					{
						name: "hdr10 color format",
						got:  firstTarget.Video.ChromaFormat,
						want: chromaFormatYUV420P10LE,
					},
					{
						name: "hdr10 color primaries",
						got:  firstTarget.Video.ColorPrimaries,
						want: colorPrimaryBT2020,
					},
					{
						name: "codec profile",
						got:  firstTarget.Video.Profile,
						want: "main10",
					},
					{
						name: "vtag",
						got:  firstTarget.Video.VTag,
						want: "hvc1",
					},
				}

				for _, tt := range tests {
					tt := tt
					t.Run(tt.name, func(t *testing.T) {
						if g, e := tt.got, tt.want; !reflect.DeepEqual(g, e) {
							t.Fatalf("%s: got %q, expected %q", tt.name, g, e)
						}
					})
				}
			},
		},
		{
			name: "hevc/hdr10 presets with mxf sources are set correctly",
			presetModifier: func(p db.Preset) db.Preset {
				p.Video.Codec = "h265"
				p.Video.Profile = ""
				p.SourceContainer = "mxf"
				p.Video.HDR10Settings = db.HDR10Settings{
					Enabled:       true,
					MaxCLL:        10000,
					MaxFALL:       400,
					MasterDisplay: "G(13250,34500)B(7500,3000)R(34000,16000)WP(15635,16450)L(10000000,1)",
				}

				// setting twoPass to false to ensure it's forced to true
				p.TwoPass = false

				return p
			},
			assertion: func(input hybrik.TranscodePayload, t *testing.T) {
				transcodeTargets, ok := input.Targets.([]hybrik.TranscodeTarget)
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
					tt := tt
					t.Run(tt.name, func(t *testing.T) {
						if g, e := tt.got, tt.want; !reflect.DeepEqual(g, e) {
							t.Fatalf("%s: got %q, expected %q", tt.name, g, e)
						}
					})
				}
			},
		},
		{
			name: "vbr transcodes are constrained to a fixed percent of variability",
			presetModifier: func(preset db.Preset) db.Preset {
				preset.RateControl = "vbr"
				preset.Video.Bitrate = "10000000"
				return preset
			},
			assertion: func(input hybrik.TranscodePayload, t *testing.T) {
				transcodeTargets, ok := input.Targets.([]hybrik.TranscodeTarget)
				if !ok {
					t.Errorf("targets are not TranscodeTargets")
				}
				firstTarget := transcodeTargets[0]

				tests := []struct {
					name      string
					got, want interface{}
				}{
					{"bitrate mode", firstTarget.Video.BitrateMode, rateControlModeVBR},
					{"bitrate", firstTarget.Video.BitrateKb, 10000},
					{"min bitrate", firstTarget.Video.MinBitrateKb, 10000 * (100 - vbrVariabilityPercent) / 100},
					{"max bitrate", firstTarget.Video.MaxBitrateKb, 10000 * (100 + vbrVariabilityPercent) / 100},
				}

				for _, tt := range tests {
					tt := tt
					t.Run(tt.name, func(t *testing.T) {
						if g, e := tt.got, tt.want; !reflect.DeepEqual(g, e) {
							t.Fatalf("%s: got %v, expected %v", tt.name, g, e)
						}
					})
				}
			},
		},
		{
			name: "transcodes sent with unsupported rate control modes result in an error",
			presetModifier: func(preset db.Preset) db.Preset {
				preset.RateControl = "fake_mode"
				return preset
			},
			wantErrMsg: `running "rateControl" transcode payload modifier: rate control mode "fake_mode" is not ` +
				`supported in hybrik, the currently supported modes are map[cbr:{} vbr:{}]`,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			p := &hybrikProvider{
				config: &config.Hybrik{
					PresetPath: "some_preset_path",
				},
				repository: dbtest.NewFakeRepository(false),
			}

			gotElement, err := p.transcodeElementFromPreset(tt.presetModifier(defaultPreset), tt.transcodeCfg.uid, jobCfg{
				destination:       tt.transcodeCfg.destination,
				executionFeatures: tt.transcodeCfg.execFeatures,
				computeTags:       tt.transcodeCfg.computeTags,
			}, tt.transcodeCfg.filename)
			if err != nil && tt.wantErrMsg != err.Error() {
				t.Errorf("hybrikProvider.transcodeElementFromPreset()error = %v, wantErr %q", err, tt.wantErrMsg)
				return
			}

			if tt.assertion != nil {
				tt.assertion(gotElement.Payload.(hybrik.TranscodePayload), t)
			}
		})
	}
}

func TestHybrikProvider_presetsToTranscodeJob(t *testing.T) {
	tests := []struct {
		name    string
		job     *db.Job
		preset  db.Preset
		wantJob hybrik.CreateJob
		wantErr string
	}{
		{
			name:   "a valid mp4 transcode job is mapped correctly to a hybrik job input",
			job:    &defaultJob,
			preset: defaultPreset,
			wantJob: hybrik.CreateJob{
				Name: "Job jobID [path.mp4]",
				Payload: hybrik.CreateJobPayload{
					Elements: []hybrik.Element{
						{
							UID:  "source_file",
							Kind: "source",
							Payload: hybrik.ElementPayload{
								Kind:    "asset_urls",
								Payload: []hybrik.AssetPayload{{StorageProvider: "s3", URL: "s3://some/path.mp4"}},
							},
						},
						{
							UID:  "transcode_task_0",
							Kind: "transcode",
							Task: &hybrik.ElementTaskOptions{Name: "Transcode - file1.mp4", Tags: []string{}},
							Payload: hybrik.TranscodePayload{
								LocationTargetPayload: hybrik.LocationTargetPayload{
									Location: hybrik.TranscodeLocation{
										Path:            "s3://some-dest/path/jobID",
										StorageProvider: storageProviderS3,
									},
									Targets: []hybrik.TranscodeTarget{{
										Audio: []hybrik.AudioTarget{{
											BitrateKb: 20,
											Channels:  2,
											Codec:     "aac",
										}},
										Container: hybrik.TranscodeContainer{
											Kind: "mp4",
										},
										ExistingFiles: "replace",
										FilePattern:   "file1.mp4",
										NumPasses:     2,
										Video: &hybrik.VideoTarget{
											Width:          intToPtr(300),
											Height:         intToPtr(400),
											BitrateMode:    "cbr",
											BitrateKb:      400,
											Preset:         presetSlow,
											ChromaFormat:   chromaFormatYUV420P,
											Codec:          "h264",
											Profile:        "high",
											Level:          "4.1",
											MinGOPFrames:   120,
											MaxGOPFrames:   120,
											ExactGOPFrames: 120,
											InterlaceMode:  "progressive",
										},
									}},
								},
							},
						},
					},
					Connections: []hybrik.Connection{
						{
							From: []hybrik.ConnectionFrom{{Element: "source_file"}},
							To: hybrik.ConnectionTo{
								Success: []hybrik.ToSuccess{{Element: "transcode_task_0"}},
							},
						},
					},
				},
			},
		},
		// TODO uncomment once Hybrik fixes bug and we can re-enable the new structure
		//{
		//	name: "a valid mp4 hevc dolbyVision transcode job is mapped correctly to a hybrik job input",
		//	job:  &defaultJob,
		//	preset: db.Preset{
		//		Name:        defaultPreset.Name,
		//		Description: defaultPreset.Description,
		//		Container:   "mp4",
		//		Video: db.VideoPreset{
		//			Profile:       "main10",
		//			Width:         "300",
		//			Codec:         "h265",
		//			Bitrate:       "12000",
		//			GopSize:       "120",
		//			GopMode:       "fixed",
		//			InterlaceMode: "progressive",
		//			DolbyVisionSettings: db.DolbyVisionSettings{
		//				Enabled: true,
		//			},
		//		},
		//	},
		//	wantJob: hybrik.CreateJob{
		//		Name: "Job jobID [path.mp4]",
		//		Payload: hybrik.CreateJobPayload{
		//			Elements: []hybrik.Element{
		//				{
		//					UID:  "source_file",
		//					Kind: "source",
		//					Payload: hybrik.ElementPayload{
		//						Kind:    "asset_urls",
		//						Payload: []hybrik.AssetPayload{{StorageProvider: "s3", URL: "s3://some/path.mp4"}},
		//					},
		//				},
		//				{
		//					UID:  "mezzanine_qc",
		//					Kind: "dolby_vision",
		//					Task: &hybrik.ElementTaskOptions{Name: "Mezzanine QC", Tags: []string{"preproc"}},
		//					Payload: hybrik.DoViV2MezzanineQCPayload{
		//						Module: "mezzanine_qc",
		//						Params: hybrik.DoViV2MezzanineQCPayloadParams{
		//							Location:    hybrik.TranscodeLocation{StorageProvider: "s3", Path: "s3://some-dest/path/jobID/mezzanine_qc"},
		//							FilePattern: "jobID_mezz_qc_report.txt",
		//						},
		//					},
		//				},
		//				{
		//					UID:  "dolby_vision_0",
		//					Kind: "dolby_vision",
		//					Task: &hybrik.ElementTaskOptions{
		//						Name:              "Encode #0",
		//						RetryMethod:       "fail",
		//						Tags:              []string{computeTagPreProcDefault},
		//						SourceElementUIDs: []string{"source_file"},
		//					},
		//					Payload: hybrik.DolbyVisionV2TaskPayload{
		//						Module:        "encoder",
		//						Profile:       5,
		//						Location:      hybrik.TranscodeLocation{StorageProvider: "s3", Path: "s3://some-dest/path/jobID"},
		//						Preprocessing: hybrik.DolbyVisionV2Preprocessing{Task: hybrik.TaskTags{Tags: []string{"preproc"}}},
		//						Transcodes: []hybrik.Element{
		//							{
		//								UID:  "transcode_task_0",
		//								Kind: "transcode",
		//								Task: &hybrik.ElementTaskOptions{Name: "Transcode - file1.mp4", Tags: []string{}},
		//								Payload: hybrik.TranscodePayload{
		//									LocationTargetPayload: hybrik.LocationTargetPayload{
		//										Location: hybrik.TranscodeLocation{StorageProvider: "s3", Path: "s3://some-dest/path/jobID"},
		//										Targets: []hybrik.TranscodeTarget{
		//											{
		//												FilePattern:   "file1.mp4",
		//												ExistingFiles: "replace",
		//												Container:     hybrik.TranscodeContainer{Kind: "elementary"},
		//												NumPasses:     1,
		//												Video: &hybrik.VideoTarget{
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
		//												Audio: []hybrik.AudioTarget{},
		//											},
		//										},
		//									},
		//									Options: &hybrik.TranscodeTaskOptions{Pipeline: &hybrik.PipelineOptions{EncoderVersion: "hybrik_4.0_10bit"}},
		//								},
		//							},
		//						},
		//						PostTranscode: hybrik.DoViPostTranscode{
		//							Task: &hybrik.TaskTags{Tags: []string{computeTagPreProcDefault}},
		//							MP4Mux: hybrik.DoViMP4Mux{
		//								Enabled:           true,
		//								FilePattern:       "{source_basename}.mp4",
		//								ElementaryStreams: []hybrik.DoViMP4MuxElementaryStream{},
		//								CLIOptions: map[string]string{
		//									doViMP4MuxDVH1FlagKey: "",
		//								},
		//							},
		//						},
		//					},
		//				},
		//			},
		//			Connections: []hybrik.Connection{
		//				{
		//					From: []hybrik.ConnectionFrom{{Element: "source_file"}},
		//					To:   hybrik.ConnectionTo{Success: []hybrik.ToSuccess{{Element: "mezzanine_qc"}}},
		//				},
		//				{
		//					From: []hybrik.ConnectionFrom{{Element: "mezzanine_qc"}},
		//					To:   hybrik.ConnectionTo{Success: []hybrik.ToSuccess{{Element: "dolby_vision_0"}}},
		//				},
		//			},
		//		},
		//	},
		//},
		// TODO remove once Hybrik fixes bug and we can re-enable the new structure
		{
			name: "a valid mp4 hevc dolbyVision transcode job is mapped correctly to a hybrik job input",
			job:  &defaultJob,
			preset: db.Preset{
				Name:        defaultPreset.Name,
				Description: defaultPreset.Description,
				Container:   "mp4",
				Video: db.VideoPreset{
					Profile:       "main10",
					Width:         "300",
					Codec:         "h265",
					Bitrate:       "12000",
					GopSize:       "120",
					GopMode:       "fixed",
					InterlaceMode: "progressive",
					DolbyVisionSettings: db.DolbyVisionSettings{
						Enabled: true,
					},
				},
			},
			wantJob: hybrik.CreateJob{
				Name: "Job jobID [path.mp4]",
				Payload: hybrik.CreateJobPayload{
					Elements: []hybrik.Element{
						{
							UID:  "source_file",
							Kind: "source",
							Payload: hybrik.ElementPayload{
								Kind:    "asset_urls",
								Payload: []hybrik.AssetPayload{{StorageProvider: "s3", URL: "s3://some/path.mp4"}},
							},
						},
						{
							UID:  "dolby_vision_task",
							Kind: "dolby_vision",
							Task: &hybrik.ElementTaskOptions{
								Tags: []string{computeTagPreProcDefault},
							},
							Payload: hybrik.DolbyVisionTaskPayload{
								Module:  "profile",
								Profile: 5,
								MezzanineQC: hybrik.DoViMezzanineQC{
									Location:    hybrik.TranscodeLocation{StorageProvider: "s3", Path: "s3://some-dest/path/jobID/tmp"},
									Task:        hybrik.TaskTags{Tags: []string{"preproc"}},
									FilePattern: "jobID_mezz_qc_report.txt",
									ToolVersion: "2.6.2",
								},
								NBCPreproc: hybrik.DoViNBCPreproc{
									Task:           hybrik.TaskTags{Tags: []string{"preproc"}},
									Location:       hybrik.TranscodeLocation{StorageProvider: "s3", Path: "s3://some-dest/path/jobID/tmp"},
									SDKVersion:     "4.2.1_ga",
									NumTasks:       "auto",
									IntervalLength: 48,
									CLIOptions:     hybrik.DoViNBCPreprocCLIOptions{InputEDRAspect: "2", InputEDRPad: "0x0x0x0", InputEDRCrop: "0x0x0x0"},
								},
								Transcodes: []hybrik.Element{
									{
										UID:  "transcode_task_0",
										Kind: "transcode",
										Task: &hybrik.ElementTaskOptions{Name: "Transcode - file1.mp4", Tags: []string{}},
										Payload: hybrik.TranscodePayload{
											LocationTargetPayload: hybrik.LocationTargetPayload{
												Location: hybrik.TranscodeLocation{StorageProvider: "s3", Path: "s3://some-dest/path/jobID"},
												Targets: []hybrik.TranscodeTarget{
													{
														FilePattern:   "file1.mp4",
														ExistingFiles: "replace",
														Container:     hybrik.TranscodeContainer{Kind: "elementary"},
														NumPasses:     1,
														Video: &hybrik.VideoTarget{
															Width:          intToPtr(300),
															BitrateKb:      12,
															Preset:         "slow",
															Codec:          "h265",
															Profile:        "main10",
															Tune:           "grain",
															MinGOPFrames:   120,
															MaxGOPFrames:   120,
															ExactGOPFrames: 120,
															InterlaceMode:  "progressive",
															ChromaFormat:   "yuv420p10le",
															X265Options:    "concatenation={auto_concatenation_flag}:vbv-init=0.6:vbv-end=0.6:annexb=1:hrd=1:aud=1:videoformat=5:range=full:colorprim=2:transfer=2:colormatrix=2:rc-lookahead=48:qg-size=32:scenecut=0:no-open-gop=1:frame-threads=0:repeat-headers=1:nr-inter=400:nr-intra=100:psy-rd=0:cbqpoffs=0:crqpoffs=3",
															VTag:           "hvc1",
															FFMPEGArgs:     " -strict experimental",
														},
														Audio: []hybrik.AudioTarget{},
													},
												},
											},
											Options: &hybrik.TranscodeTaskOptions{Pipeline: &hybrik.PipelineOptions{EncoderVersion: "hybrik_4.0_10bit"}},
										},
									},
								},
								PostTranscode: hybrik.DoViPostTranscode{
									Task: &hybrik.TaskTags{Tags: []string{"preproc"}},
									VESMux: &hybrik.DoViVESMux{
										Enabled:     true,
										Location:    hybrik.TranscodeLocation{StorageProvider: "s3", Path: "s3://some-dest/path/jobID/tmp"},
										FilePattern: "ves.h265",
										SDKVersion:  "4.2.1_ga",
									},
									MetadataPostProc: &hybrik.DoViMetadataPostProc{
										Enabled:     true,
										Location:    hybrik.TranscodeLocation{StorageProvider: "s3", Path: "s3://some-dest/path/jobID/tmp"},
										FilePattern: "postproc.265",
										SDKVersion:  "4.2.1_ga",
										QCSettings: hybrik.DoViQCSettings{
											Enabled:     true,
											ToolVersion: "0.9.0.9",
											Location:    hybrik.TranscodeLocation{StorageProvider: "s3", Path: "s3://some-dest/path/jobID/tmp"},
											FilePattern: "metadata_postproc_qc_report.txt",
										},
									},
									MP4Mux: hybrik.DoViMP4Mux{
										Enabled:            true,
										Location:           &hybrik.TranscodeLocation{StorageProvider: "s3", Path: "s3://some-dest/path/jobID"},
										FilePattern:        "{source_basename}.mp4",
										ToolVersion:        "1.2.8",
										CopySourceStartPTS: true,
										QCSettings: &hybrik.DoViQCSettings{
											Enabled:     true,
											ToolVersion: "1.1.4",
											Location:    hybrik.TranscodeLocation{StorageProvider: "s3", Path: "s3://some-dest/path/jobID/tmp"},
											FilePattern: "mp4_qc_report.txt",
										},
										CLIOptions: map[string]string{"dvh1flag": ""},
										ElementaryStreams: []hybrik.DoViMP4MuxElementaryStream{
											{
												AssetURL:        hybrik.AssetURL{StorageProvider: "s3", URL: "s3://some/path.mp4"},
												ExtractAudio:    true,
												ExtractLocation: &hybrik.TranscodeLocation{StorageProvider: "s3", Path: "s3://some-dest/path/jobID/tmp"},
												ExtractTask: &hybrik.DoViMP4MuxExtractTask{
													RetryMethod: "retry",
													Retry:       hybrik.Retry{Count: 3, DelaySec: 30},
													Name:        "Demux Audio",
												},
											},
										},
									},
								},
							},
						},
					},
					Connections: []hybrik.Connection{
						{
							From: []hybrik.ConnectionFrom{{Element: "source_file"}},
							To:   hybrik.ConnectionTo{Success: []hybrik.ToSuccess{{Element: "dolby_vision_task"}}},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			fakeDB, err := fakeDBWithPreset(tt.preset)
			if err != nil {
				t.Errorf("hybrikProvider.presetsToTranscodeJob() error = %v", err)
				return
			}

			p := &hybrikProvider{
				config: &config.Hybrik{
					Destination: "s3://some-dest/path",
					PresetPath:  "some_preset_path",
				},
				repository: fakeDB,
			}

			got, err := p.createJobReqFrom(tt.job)
			if err != nil {
				if tt.wantErr != err.Error() {
					t.Errorf("hybrikProvider.presetsToTranscodeJob() error = %v, wantErr %q", err, tt.wantErr)
				}

				return
			}

			if g, e := got, tt.wantJob; !reflect.DeepEqual(g, e) {
				t.Fatalf("hybrikProvider.presetsToTranscodeJob() wrong job request\nWant %+v\nGot %+v\nDiff %s", e,
					g, cmp.Diff(e, g))
			}
		})
	}
}

func TestHybrikProvider_presetsToTranscodeJob_fields(t *testing.T) {
	tests := []struct {
		name        string
		jobModifier func(job db.Job) db.Job
		assertion   func(hybrik.CreateJob, *testing.T)
		wantErrMsg  string
	}{
		{
			name: "when a dolby vision sidecar asset is included, it is correctly added to the source element",
			jobModifier: func(job db.Job) db.Job {
				job.SidecarAssets = map[db.SidecarAssetKind]string{
					db.SidecarAssetKindDolbyVisionMetadata: "s3://test_sidecar_location/path/file.xml",
				}

				return job
			},
			assertion: func(createJob hybrik.CreateJob, t *testing.T) {
				gotSource := createJob.Payload.Elements[0]

				expectSource := hybrik.Element{
					UID:  "source_file",
					Kind: "source",
					Payload: hybrik.ElementPayload{
						Kind: "asset_urls",
						Payload: []hybrik.AssetPayload{
							{StorageProvider: "s3", URL: "s3://some/path.mp4"},
							{
								StorageProvider: "s3",
								URL:             "s3://test_sidecar_location/path/file.xml",
								Contents: []hybrik.AssetContents{
									{
										Kind:    "metadata",
										Payload: hybrik.AssetContentsPayload{Standard: "dolbyvision_metadata"},
									},
								},
							},
						},
					},
				}

				if g, e := gotSource, expectSource; !reflect.DeepEqual(g, e) {
					t.Fatalf("hybrikProvider.presetsToTranscodeJob() wrong job request\nWant %+v\nGot %+v\nDiff %s", e,
						g, cmp.Diff(e, g))
				}
			},
		},
		{
			name: "when a destination base path is defined, the defined destination is used instead of the" +
				" globally configured one",
			jobModifier: func(job db.Job) db.Job {
				job.DestinationBasePath = "s3://per-job-defined-bucket/some/base/path"
				return job
			},
			assertion: func(createJob hybrik.CreateJob, t *testing.T) {
				if len(createJob.Payload.Elements) < 2 {
					t.Error("job has less than two elements, tried to pull the second element (transcode)")
					return
				}
				gotTranscode := createJob.Payload.Elements[1]

				payload, ok := gotTranscode.Payload.(hybrik.TranscodePayload)
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
			name: "when custom compute tags are specified, the right tags are added to the output",
			jobModifier: func(job db.Job) db.Job {
				job.ExecutionEnv.ComputeTags = map[db.ComputeClass]string{
					db.ComputeClassTranscodeDefault: "custom_tag",
				}

				return job
			},
			assertion: func(createJob hybrik.CreateJob, t *testing.T) {
				gotTask := createJob.Payload.Elements[1].Task

				expectTask := &hybrik.ElementTaskOptions{Name: "Transcode - file1.mp4", Tags: []string{"custom_tag"}}

				if g, e := gotTask, expectTask; !reflect.DeepEqual(g, e) {
					t.Fatalf("hybrikProvider.presetsToTranscodeJob() wrong job request\nWant %+v\nGot %+v\nDiff %s", e,
						g, cmp.Diff(e, g))
				}
			},
		},
		{
			name: "when HLS packaging is specified, a package task is added with the correct values",
			jobModifier: func(job db.Job) db.Job {
				job.StreamingParams = db.StreamingParams{
					SegmentDuration: 4,
					Protocol:        "hls",
				}

				job.ExecutionEnv.ComputeTags = map[db.ComputeClass]string{
					db.ComputeClassTranscodeDefault: "default_transcode_class",
				}

				return job
			},
			assertion: func(createJob hybrik.CreateJob, t *testing.T) {
				gotTask := createJob.Payload.Elements[len(createJob.Payload.Elements)-1]

				expectTask := hybrik.Element{
					UID:  "packager",
					Kind: elementKindPackage,
					Task: &hybrik.ElementTaskOptions{
						Tags: []string{"default_transcode_class"},
					},
					Payload: hybrik.PackagePayload{
						Location: hybrik.TranscodeLocation{
							StorageProvider: "s3",
							Path:            "s3://some-dest/path/jobID/hls",
						},
						FilePattern:        "master.m3u8",
						Kind:               "hls",
						SegmentationMode:   "segmented_mp4",
						SegmentDurationSec: 4,
						HLS: &hybrik.HLSPackagingSettings{
							IncludeIFRAMEManifests: true,
							HEVCCodecIDPrefix:      "hvc1",
						},
					},
				}

				if g, e := gotTask, expectTask; !reflect.DeepEqual(g, e) {
					t.Fatalf("hybrikProvider.presetsToTranscodeJob() wrong package task\nWant %+v\nGot %+v\nDiff %s", e,
						g, cmp.Diff(e, g))
				}
			},
		},
		{
			name: "when DASH packaging is specified, a package task is added with the correct values",
			jobModifier: func(job db.Job) db.Job {
				job.StreamingParams = db.StreamingParams{
					SegmentDuration: 4,
					Protocol:        "dash",
				}

				return job
			},
			assertion: func(createJob hybrik.CreateJob, t *testing.T) {
				gotTask := createJob.Payload.Elements[len(createJob.Payload.Elements)-1]

				expectTask := hybrik.Element{
					UID:  "packager",
					Kind: elementKindPackage,
					Payload: hybrik.PackagePayload{
						Location: hybrik.TranscodeLocation{
							StorageProvider: "s3",
							Path:            "s3://some-dest/path/jobID/dash",
						},
						FilePattern:        "master.mpd",
						Kind:               "dash",
						SegmentationMode:   "segmented_mp4",
						SegmentDurationSec: 4,
						DASH: &hybrik.DASHPackagingSettings{
							SegmentationMode:   "segmented_mp4",
							SegmentDurationSec: "4",
						},
					},
				}

				if g, e := gotTask, expectTask; !reflect.DeepEqual(g, e) {
					t.Fatalf("hybrikProvider.presetsToTranscodeJob() wrong package task\nWant %+v\nGot %+v\nDiff %s", e,
						g, cmp.Diff(e, g))
				}
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			fakeDB, err := fakeDBWithPreset(defaultPreset)
			if err != nil {
				t.Errorf("hybrikProvider.presetsToTranscodeJob() error = %v", err)
				return
			}

			p := &hybrikProvider{
				config: &config.Hybrik{
					Destination: "s3://some-dest/path",
					PresetPath:  "some_preset_path",
				},
				repository: fakeDB,
			}

			modifiedJob := tt.jobModifier(defaultJob)
			got, err := p.createJobReqFrom(&modifiedJob)
			if err != nil && tt.wantErrMsg != err.Error() {
				t.Errorf("hybrikProvider.presetsToTranscodeJob() error = %v, wantErr %q", err, tt.wantErrMsg)
				return
			}

			if tt.assertion != nil {
				tt.assertion(got, t)
			}
		})
	}
}

func fakeDBWithPreset(preset db.Preset) (db.Repository, error) {
	fakeDB := dbtest.NewFakeRepository(false)
	err := fakeDB.CreateLocalPreset(&db.LocalPreset{Name: preset.Name, Preset: preset})
	if err != nil {
		return nil, err
	}

	return fakeDB, nil
}

func intToPtr(i int) *int {
	return &i
}
