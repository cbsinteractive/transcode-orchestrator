package hybrik

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/NYTimes/video-transcoding-api/config"
	"github.com/NYTimes/video-transcoding-api/db"
	"github.com/cbsinteractive/hybrik-sdk-go"
	"github.com/google/go-cmp/cmp"
)

var (
	defaultPreset = db.Preset{
		Name:        "preset_name",
		Description: "test_desc",
		Container:   "mp4",
		RateControl: "VBR",
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
					Name: "preset1",
					ProviderMapping: map[string]string{
						"hybrik": "preset1",
					},
				},
				FileName: "file1.mp4",
			},
		},
	}

	defaultHybrikPreset = hybrik.Preset{
		Key:         defaultPreset.Name,
		Name:        defaultPreset.Name,
		Description: defaultPreset.Description,
		Kind:        "transcode",
		Path:        "some_preset_path",
		Payload: hybrik.PresetPayload{
			Targets: []hybrik.PresetTarget{
				{
					FilePattern: "",
					Container:   hybrik.TranscodeContainer{Kind: defaultPreset.Container},
					Video: hybrik.VideoTarget{
						Width:         intToPtr(300),
						Codec:         defaultPreset.Video.Codec,
						BitrateKb:     400,
						MaxGOPFrames:  120,
						Profile:       "high",
						Level:         "4.1",
						InterlaceMode: "progressive",
					},
					Audio: []hybrik.AudioTarget{
						{
							Codec:     defaultPreset.Audio.Codec,
							BitrateKb: 20,
						},
					},
				},
			},
		},
	}
)

func TestHybrikProvider_hybrikPresetFrom(t *testing.T) {
	tests := []struct {
		name             string
		provider         *hybrikProvider
		preset           db.Preset
		wantHybrikPreset hybrik.Preset
		wantErr          bool
	}{
		{
			name: "a valid h264/aac mp4 preset results in the expected mediaconvert preset sent to the Hybrik API",
			provider: &hybrikProvider{
				config: &config.Hybrik{
					PresetPath: "some_preset_path",
				},
			},
			preset: defaultPreset,
			wantHybrikPreset: hybrik.Preset{
				Key:         defaultPreset.Name,
				Name:        defaultPreset.Name,
				Description: defaultPreset.Description,
				Kind:        "transcode",
				Path:        "some_preset_path",
				Payload: hybrik.PresetPayload{
					Targets: []hybrik.PresetTarget{
						{
							FilePattern: "",
							Container:   hybrik.TranscodeContainer{Kind: defaultPreset.Container},
							Video: hybrik.VideoTarget{
								Width:         intToPtr(300),
								Height:        intToPtr(400),
								Codec:         defaultPreset.Video.Codec,
								BitrateKb:     400,
								MaxGOPFrames:  120,
								Profile:       "high",
								Level:         "4.1",
								InterlaceMode: "progressive",
							},
							Audio: []hybrik.AudioTarget{
								{
									Codec:     defaultPreset.Audio.Codec,
									BitrateKb: 20,
								},
							},
							ExistingFiles: "replace",
							UID:           "target",
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			p, err := tt.provider.hybrikPresetFrom(tt.preset)
			if (err != nil) != tt.wantErr {
				t.Errorf("hybrikProvider.hybrikPresetFrom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if g, e := p, tt.wantHybrikPreset; !reflect.DeepEqual(g, e) {
				t.Fatalf("hybrikProvider.hybrikPresetFrom() wrong preset request\nWant %+v\nGot %+v\nDiff %s", e,
					g, cmp.Diff(e, g))
			}
		})
	}
}

func TestHybrikProvider_hybrikPresetFrom_fields(t *testing.T) {
	tests := []struct {
		name           string
		presetModifier func(preset db.Preset) db.Preset
		assertion      func(hybrik.Preset, *testing.T)
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
			assertion: func(input hybrik.Preset, t *testing.T) {
				firstTarget := input.Payload.Targets[0]

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
						name: "ffmpeg params",
						got:  firstTarget.Video.FFMPEGArgs,
						want: "-tag:v hvc1",
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
			name: "hevc/dolby vision presets are set correctly",
			presetModifier: func(p db.Preset) db.Preset {
				p.Video.DolbyVisionSettings = db.DolbyVisionSettings{
					Enabled: true,
				}
				return p
			},
			assertion: func(input hybrik.Preset, t *testing.T) {
				if g, e := input.UserData, `{"dolbyVision":true}`; g != e {
					t.Fatalf("user data: got %q, expected %q", g, e)
				}
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			p := &hybrikProvider{
				config: &config.Hybrik{
					PresetPath: "some_preset_path",
				},
			}
			gotPreset, err := p.hybrikPresetFrom(tt.presetModifier(defaultPreset))
			if err != nil && tt.wantErrMsg != err.Error() {
				t.Errorf("hybrikProvider.hybrikPresetFrom()error = %v, wantErr %q", err, tt.wantErrMsg)
				return
			}

			if tt.assertion != nil {
				tt.assertion(gotPreset, t)
			}
		})
	}
}

func TestHybrikProvider_presetsToTranscodeJob(t *testing.T) {
	customPresetDataDolbyVisionEnabled, err := json.Marshal(customPresetData{DolbyVisionEnabled: true})
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		job     *db.Job
		preset  hybrik.Preset
		wantJob hybrik.CreateJob
		wantErr string
	}{
		{
			name:   "a valid mp4 transcode job is mapped correctly to a hybrik job input",
			job:    &defaultJob,
			preset: defaultHybrikPreset,
			wantJob: hybrik.CreateJob{
				Name: "Job jobID [path.mp4]",
				Payload: hybrik.CreateJobPayload{
					Elements: []hybrik.Element{
						{
							UID:  "source_file",
							Kind: "source",
							Payload: map[string]interface{}{
								"kind": "asset_urls",
								"payload": []interface{}{
									map[string]interface{}{
										"storage_provider": "s3",
										"url":              "s3://some/path.mp4",
									},
								},
							},
						},
						{
							UID:    "transcode_task_0",
							Kind:   "transcode",
							Task:   &hybrik.ElementTaskOptions{Name: "Transcode - preset_name"},
							Preset: &hybrik.TranscodePreset{Key: "preset_name"},
							Payload: map[string]interface{}{
								"location": map[string]interface{}{
									"path":             "/jobID",
									"storage_provider": "s3",
								},
								"source_pipeline": map[string]interface{}{
									"options": map[string]interface{}{},
									"scaler":  map[string]interface{}{},
								},
								"targets": []interface{}{
									map[string]interface{}{
										"container":    map[string]interface{}{},
										"file_pattern": "file1.mp4",
									},
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
		{
			name: "a valid mp4 hevc dolbyVision transcode job is mapped correctly to a hybrik job input",
			job:  &defaultJob,
			preset: hybrik.Preset{
				Key:         defaultPreset.Name,
				Name:        defaultPreset.Name,
				Description: defaultPreset.Description,
				Kind:        "transcode",
				Path:        "some_preset_path",
				UserData:    string(customPresetDataDolbyVisionEnabled),
				Payload: hybrik.PresetPayload{
					Targets: []hybrik.PresetTarget{
						{
							FilePattern: "",
							Container:   hybrik.TranscodeContainer{Kind: defaultPreset.Container},
							Video: hybrik.VideoTarget{
								Width:         intToPtr(300),
								Codec:         "h265",
								BitrateKb:     12000,
								MaxGOPFrames:  120,
								Profile:       "main10",
								InterlaceMode: "progressive",
							},
							Audio: []hybrik.AudioTarget{
								{
									Codec:     defaultPreset.Audio.Codec,
									BitrateKb: 20,
								},
							},
						},
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
							Payload: map[string]interface{}{
								"kind": "asset_urls",
								"payload": []interface{}{
									map[string]interface{}{
										"storage_provider": "s3",
										"url":              "s3://some/path.mp4",
									},
								},
							},
						},
						{
							UID:  "dolby_vision_task",
							Kind: "dolby_vision",
							Payload: map[string]interface{}{
								"mezzanine_qc": map[string]interface{}{
									"enabled":      false,
									"file_pattern": "jobID_mezz_qc_report.txt",
									"location": map[string]interface{}{
										"path":             "/jobID/mezzanine_qc",
										"storage_provider": "s3",
									},
									"task":         map[string]interface{}{},
									"tool_version": "2.6.2",
								},
								"module": "profile",
								"nbc_preproc": map[string]interface{}{
									"cli_options": map[string]interface{}{
										"inputEDRAspect": "2",
										"inputEDRCrop":   "0x0x0x0",
										"inputEDRPad":    "0x0x0x0",
									},
									"dovi_sdk_version": "4.2.1_ga",
									"interval_length":  float64(48),
									"location": map[string]interface{}{
										"path":             "/jobID/nbc_preproc",
										"storage_provider": "s3",
									},
									"num_tasks": "auto",
									"task": map[string]interface{}{
										"tags": []interface{}{"preproc"},
									},
								},
								"post_transcode": map[string]interface{}{
									"metadata_postproc": map[string]interface{}{
										"dovi_sdk_version": "4.2.1_ga",
										"enabled":          true,
										"file_pattern":     "postproc.265",
										"location": map[string]interface{}{
											"path":             "/jobID/metadata_postproc",
											"storage_provider": "s3",
										},
										"qc": map[string]interface{}{
											"enabled":      true,
											"file_pattern": "metadata_postproc_qc_report.txt",
											"location": map[string]interface{}{
												"path":             "/jobID/metadata_postproc_qc",
												"storage_provider": "s3",
											},
											"tool_version": "0.9.0.9",
										},
									},
									"mp4_mux": map[string]interface{}{
										"copy_source_start_pts": true,
										"elementary_streams": []interface{}{
											map[string]interface{}{
												"asset_url": map[string]interface{}{
													"storage_provider": "s3",
													"url":              "s3://some/path.mp4",
												},
												"extract_audio": true,
												"extract_location": map[string]interface{}{
													"path":             "/jobID/source_demux",
													"storage_provider": "s3",
												},
												"extract_task": map[string]interface{}{
													"name": "Demux Audio",
													"retry": map[string]interface{}{
														"count":     float64(3),
														"delay_sec": float64(30),
													},
													"retry_method": "retry",
												},
											},
										},
										"enabled":      true,
										"file_pattern": "{source_basename}.mp4",
										"location": map[string]interface{}{
											"path":             "/jobID",
											"storage_provider": "s3",
										},
										"qc": map[string]interface{}{
											"enabled":      true,
											"file_pattern": "mp4_qc_report.txt",
											"location": map[string]interface{}{
												"path":             "/jobID/mp4_qc",
												"storage_provider": "s3",
											},
											"tool_version": "1.1.4",
										},
										"tool_version": "1.2.8",
									},
									"task": map[string]interface{}{},
									"ves_mux": map[string]interface{}{
										"dovi_sdk_version": "4.2.1_ga",
										"enabled":          true,
										"file_pattern":     "ves.h265",
										"location": map[string]interface{}{
											"path":             "/jobID/vesmuxer",
											"storage_provider": "s3",
										},
									},
								},
								"profile": float64(5),
								"transcodes": []interface{}{
									map[string]interface{}{
										"kind": "transcode",
										"payload": map[string]interface{}{
											"location": map[string]interface{}{
												"path":             "/jobID/elementary",
												"storage_provider": "s3",
											},
											"source_pipeline": map[string]interface{}{
												"options": map[string]interface{}{},
												"scaler":  map[string]interface{}{},
											},
											"targets": []interface{}{
												map[string]interface{}{
													"audio": []interface{}{map[string]interface{}{
														"bitrate_kb": float64(20),
														"codec":      "aac",
													}},
													"container":    map[string]interface{}{"kind": "mp4"},
													"file_pattern": "file1.mp4",
													"video": map[string]interface{}{
														"bitrate_kb":     float64(12000),
														"codec":          "h265",
														"interlace_mode": "progressive",
														"max_gop_frames": float64(120),
														"profile":        "main10",
														"width":          float64(300),
													},
												},
											},
										},
										"uid": "transcode_task_0",
									},
								},
							},
						},
					},
					Connections: []hybrik.Connection{
						{
							From: []hybrik.ConnectionFrom{{Element: "source_file"}},
							To: hybrik.ConnectionTo{
								Success: []hybrik.ToSuccess{{Element: "dolby_vision_task"}},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			p := &hybrikProvider{
				config: &config.Hybrik{
					PresetPath: "some_preset_path",
				},
				c: &testClient{getPresetReturn: tt.preset},
			}

			got, err := p.presetsToTranscodeJob(tt.job)
			if err != nil && tt.wantErr != err.Error() {
				t.Errorf("hybrikProvider.presetsToTranscodeJob() error = %v, wantErr %q", err, tt.wantErr)
				return
			}

			var createJobInput hybrik.CreateJob
			err = json.Unmarshal([]byte(got), &createJobInput)
			if err != nil {
				t.Errorf("hybrikProvider.presetsToTranscodeJob() error while unmarshalling create job input")
			}

			if g, e := createJobInput, tt.wantJob; !reflect.DeepEqual(g, e) {
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
					db.SidecarAssetKindDolbyVisionMetadata: "test_sidecar_location",
				}

				return job
			},
			assertion: func(createJob hybrik.CreateJob, t *testing.T) {
				gotSource := createJob.Payload.Elements[0]

				expectSource := hybrik.Element{
					UID:  "source_file",
					Kind: "source",
					Payload: map[string]interface{}{
						"kind": "asset_urls",
						"payload": []interface{}{
							map[string]interface{}{
								"storage_provider": "s3",
								"url":              "s3://some/path.mp4",
							},
							map[string]interface{}{
								"storage_provider": "s3",
								"url":              "test_sidecar_location",
								"contents": []interface{}{
									map[string]interface{}{
										"kind": assetContentsKindMetadata,
										"payload": map[string]interface{}{
											"standard": assetContentsStandardDolbyVisionMetadata,
										},
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

				payload, ok := gotTranscode.Payload.(map[string]interface{})
				if !ok {
					t.Error("transcode payload was not a map of string to map[string]interface{}")
					return
				}

				location, ok := payload["location"].(map[string]interface{})
				if !ok {
					t.Error("transcode payload location was not a map of string to map[string]interface{}")
					return
				}

				if g, e := location["path"].(string), "s3://per-job-defined-bucket/some/base/path/jobID"; g != e {
					t.Errorf("destination location path: got %q, expected %q", g, e)
				}
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			p := &hybrikProvider{
				config: &config.Hybrik{
					PresetPath: "some_preset_path",
				},
				c: &testClient{getPresetReturn: defaultHybrikPreset},
			}

			modifiedJob := tt.jobModifier(defaultJob)
			got, err := p.presetsToTranscodeJob(&modifiedJob)
			if err != nil && tt.wantErrMsg != err.Error() {
				t.Errorf("hybrikProvider.presetsToTranscodeJob() error = %v, wantErr %q", err, tt.wantErrMsg)
				return
			}

			var createJobInput hybrik.CreateJob
			err = json.Unmarshal([]byte(got), &createJobInput)
			if err != nil {
				t.Errorf("hybrikProvider.presetsToTranscodeJob() error while unmarshalling create job input")
			}

			if tt.assertion != nil {
				tt.assertion(createJobInput, t)
			}
		})
	}
}

func intToPtr(i int) *int {
	return &i
}
