package mediaconvert

import (
	"context"
	"reflect"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/cbsinteractive/transcode-orchestrator/config"
	"github.com/cbsinteractive/transcode-orchestrator/db"
	"github.com/cbsinteractive/transcode-orchestrator/db/dbtest"
	"github.com/cbsinteractive/transcode-orchestrator/provider"
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
			GopUnit:       "frames",
			InterlaceMode: "progressive",
		},
		Audio: db.AudioPreset{
			Codec:   "aac",
			Bitrate: "20000",
		},
	}

	h265Preset = db.Preset{
		Name:        "another_preset_name",
		Description: "test_desc",
		Container:   "mp4",
		RateControl: "CBR",
		TwoPass:     false,
		Video: db.VideoPreset{
			Width:         "300",
			Height:        "400",
			Codec:         "h265",
			Bitrate:       "400000",
			GopSize:       "120",
			GopUnit:       "frames",
			InterlaceMode: "progressive",
		},
	}

	av1Preset = db.Preset{
		Name:        "yet_another_preset_name",
		Description: "test_desc",
		Container:   "mp4",
		TwoPass:     false,
		Video: db.VideoPreset{
			Width:         "300",
			Height:        "400",
			Codec:         "av1",
			Bitrate:       "400000",
			GopSize:       "120",
			GopUnit:       "frames",
			InterlaceMode: "progressive",
		},
	}

	audioOnlyPreset = db.Preset{
		Name:        "preset_name",
		Description: "test_desc",
		Container:   "mp4",
		Audio: db.AudioPreset{
			Codec:   "aac",
			Bitrate: "20000",
		},
	}

	tcBurninPreset = db.Preset{
		Name:        "preset_name",
		Description: "test_desc",
		Container:   "mov",
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
			GopUnit:       "frames",
			InterlaceMode: "progressive",
			Overlays: &db.Overlays{
				TimecodeBurnin: &db.TimecodeBurnin{
					Enabled:  true,
					FontSize: 12,
					Position: 7,
				},
			},
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
		SourceInfo: db.File{
			ScanType: db.ScanTypeUnknown,
		},
		Outputs: []db.TranscodeOutput{
			{
				Preset: db.PresetMap{
					Name: "preset_name",
					ProviderMapping: map[string]string{
						"mediaconvert": "preset1",
					},
				},
				FileName: "file1.mp4",
			},
			{
				Preset: db.PresetMap{
					Name: "another_preset_name",
					ProviderMapping: map[string]string{
						"mediaconvert": "preset2",
					},
				},
				FileName: "file2.mp4",
			},
		},
		StreamingParams: db.StreamingParams{
			SegmentDuration: 6,
		},
	}
)

func Test_mcProvider_CreatePreset(t *testing.T) {
	client := &testMediaConvertClient{t: t}
	p := &mcProvider{client: client, repository: dbtest.NewFakeRepository(false)}
	presetName, err := p.CreatePreset(context.Background(), defaultPreset)
	if err != nil {
		t.Errorf("mcProvider.CreatePreset() did not expect an error, got %+v", err)
		return
	}

	preset, err := p.GetPreset(context.Background(), presetName)
	if err != nil {
		t.Errorf("didn't expect GetPreset to return an error, got %+v", err)
		return
	}

	if g, e := preset.(*db.LocalPreset).Preset, defaultPreset; !reflect.DeepEqual(g, e) {
		t.Fatalf("CreatePreset(): wrong preset \nWant %+v\nGot %+v\nDiff %s", e,
			g, cmp.Diff(e, g))
	}
}

func Test_mcProvider_CreatePreset_fields(t *testing.T) {
	tests := []struct {
		name           string
		presetModifier func(preset db.Preset) db.Preset
		assertion      func(mediaconvert.Output, *testing.T)
		wantErrMsg     string
	}{
		{
			name: "hls presets are set correctly",
			presetModifier: func(p db.Preset) db.Preset {
				p.Container = "m3u8"
				return p
			},
			assertion: func(output mediaconvert.Output, t *testing.T) {
				if g, e := output.ContainerSettings.Container, mediaconvert.ContainerTypeM3u8; g != e {
					t.Fatalf("got %q, expected %q", g, e)
				}
			},
		},
		{
			name: "cmaf presets are set correctly",
			presetModifier: func(p db.Preset) db.Preset {
				p.Container = "cmaf"
				return p
			},
			assertion: func(output mediaconvert.Output, t *testing.T) {
				if g, e := output.ContainerSettings.Container, mediaconvert.ContainerTypeCmfc; g != e {
					t.Fatalf("got %q, expected %q", g, e)
				}
			},
		},
		{
			name: "hdr10 values in presets are set correctly",
			presetModifier: func(p db.Preset) db.Preset {
				p.Video.HDR10Settings.Enabled = true
				p.Video.HDR10Settings.MaxCLL = 10000
				p.Video.HDR10Settings.MaxFALL = 400
				p.Video.HDR10Settings.MasterDisplay = "G(8500,39850)B(6550,2300)R(35400,14600)WP(15635,16450)L(100000000000,0)"
				return p
			},
			assertion: func(output mediaconvert.Output, t *testing.T) {
				colorSpaceConversion := output.VideoDescription.VideoPreprocessors.ColorCorrector.ColorSpaceConversion
				if g, e := colorSpaceConversion, mediaconvert.ColorSpaceConversionForceHdr10; g != e {
					t.Fatalf("got %q, expected %q", g, e)
				}

				wantHDR10Metadata := &mediaconvert.Hdr10Metadata{
					BluePrimaryX:              aws.Int64(6550),
					BluePrimaryY:              aws.Int64(2300),
					GreenPrimaryX:             aws.Int64(8500),
					GreenPrimaryY:             aws.Int64(39850),
					RedPrimaryX:               aws.Int64(35400),
					RedPrimaryY:               aws.Int64(14600),
					WhitePointX:               aws.Int64(15635),
					WhitePointY:               aws.Int64(16450),
					MaxLuminance:              aws.Int64(100000000000),
					MinLuminance:              aws.Int64(0),
					MaxContentLightLevel:      aws.Int64(10000),
					MaxFrameAverageLightLevel: aws.Int64(400),
				}

				gotHDR10Metadata := output.VideoDescription.VideoPreprocessors.ColorCorrector.Hdr10Metadata
				if g, e := gotHDR10Metadata, wantHDR10Metadata; !reflect.DeepEqual(g, e) {
					t.Fatalf("got %q, expected %q", g, e)
				}
			},
		},
		{
			name: "unrecognized containers return an error",
			presetModifier: func(p db.Preset) db.Preset {
				p.Container = "unrecognized"
				return p
			},
			wantErrMsg: `mapping preset container to MediaConvert container: container "unrecognized" not supported with mediaconvert`,
		},
		{
			name: "unrecognized h264 codec returns an error",
			presetModifier: func(p db.Preset) db.Preset {
				p.Video.Codec = "vp9001"
				return p
			},
			wantErrMsg: `generating video preset: video codec "vp9001" is not yet supported with mediaconvert`,
		},
		{
			name: "unrecognized h264 codec profile returns an error",
			presetModifier: func(p db.Preset) db.Preset {
				p.Video.Profile = "8000"
				return p
			},
			wantErrMsg: `generating video preset: building h264 codec settings: h264 profile "8000" is not supported with mediaconvert`,
		},
		{
			name: "unrecognized h264 codec level returns an error",
			presetModifier: func(p db.Preset) db.Preset {
				p.Video.ProfileLevel = "9001"
				return p
			},
			wantErrMsg: `generating video preset: building h264 codec settings: h264 level "9001" is not supported with mediaconvert`,
		},
		{
			name: "bad video width returns an error",
			presetModifier: func(p db.Preset) db.Preset {
				p.Video.Width = "s"
				return p
			},
			wantErrMsg: `generating video preset: parsing video width "s" to int64: strconv.ParseInt: parsing "s": invalid syntax`,
		},
		{
			name: "bad video height returns an error",
			presetModifier: func(p db.Preset) db.Preset {
				p.Video.Height = "h"
				return p
			},
			wantErrMsg: `generating video preset: parsing video height "h" to int64: strconv.ParseInt: parsing "h": invalid syntax`,
		},
		{
			name: "bad video bitrate returns an error",
			presetModifier: func(p db.Preset) db.Preset {
				p.Video.Bitrate = "bitrate"
				return p
			},
			wantErrMsg: `generating video preset: building h264 codec settings: parsing video bitrate "bitrate" to int64: strconv.ParseInt: parsing "bitrate": invalid syntax`,
		},
		{
			name: "bad video gop size returns an error",
			presetModifier: func(p db.Preset) db.Preset {
				p.Video.GopSize = "gop"
				return p
			},
			wantErrMsg: `generating video preset: building h264 codec settings: parsing gop size "gop" to float64: strconv.ParseFloat: parsing "gop": invalid syntax`,
		},
		{
			name: "unrecognized rate control mode returns an error",
			presetModifier: func(p db.Preset) db.Preset {
				p.RateControl = "not supported"
				return p
			},
			wantErrMsg: `generating video preset: building h264 codec settings: rate control mode "not supported" is not supported with mediaconvert`,
		},
		{
			name: "unrecognized interlace modes return an error",
			presetModifier: func(p db.Preset) db.Preset {
				p.Video.InterlaceMode = "unsupported mode"
				return p
			},
			wantErrMsg: `generating video preset: building h264 codec settings: h264 interlace mode "unsupported mode" is not supported with mediaconvert`,
		},
		{
			name: "unrecognized audio bitrate returns an error",
			presetModifier: func(p db.Preset) db.Preset {
				p.Audio.Bitrate = "aud_bitrate"
				return p
			},
			wantErrMsg: `generating audio preset: parsing audio bitrate "aud_bitrate" to int64: strconv.ParseInt: parsing "aud_bitrate": invalid syntax`,
		},
		{
			name: "unrecognized audio codec returns an error",
			presetModifier: func(p db.Preset) db.Preset {
				p.Audio.Codec = "aab"
				return p
			},
			wantErrMsg: `generating audio preset: audio codec "aab" is not yet supported with mediaconvert`,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			client := &testMediaConvertClient{t: t}

			repo, err := fakeDBWithPresets(tt.presetModifier(defaultPreset))
			if err != nil {
				t.Error(err)
				return
			}

			p := &mcProvider{client: client, repository: repo, cfg: &config.MediaConvert{Destination: "s3://some_dest"}}

			_, err = p.Transcode(context.Background(), &db.Job{
				ID: "jobID", ProviderName: Name, SourceMedia: "s3://some/path.mp4",
				Outputs: []db.TranscodeOutput{{Preset: db.PresetMap{Name: defaultPreset.Name}, FileName: "file1.mp4"}},
			})
			if err != nil && !strings.Contains(err.Error(), tt.wantErrMsg) {
				t.Errorf("mcProvider.Transcode() error = %v, wantErr %q", err, tt.wantErrMsg)
				return
			}

			if tt.assertion != nil {
				tt.assertion(client.createJobCalledWith.Settings.OutputGroups[0].Outputs[0], t)
			}
		})
	}
}

func Test_mcProvider_GetPreset(t *testing.T) {
	client := &testMediaConvertClient{t: t}

	fakeDB, err := fakeDBWithPresets(defaultPreset)
	if err != nil {
		t.Errorf("mcProvider.DeletePreset() error = %v", err)
		return
	}

	p := &mcProvider{client: client, repository: fakeDB}

	_, err = p.GetPreset(context.Background(), defaultPreset.Name)
	if err != nil {
		t.Fatalf("expected GetPreset() not to return an error, got: %v", err)
	}
}

func Test_mcProvider_DeletePreset(t *testing.T) {
	client := &testMediaConvertClient{t: t}

	fakeDB, err := fakeDBWithPresets(defaultPreset)
	if err != nil {
		t.Errorf("mcProvider.DeletePreset() error = %v", err)
		return
	}

	p := &mcProvider{client: client, repository: fakeDB}

	_, err = p.GetPreset(context.Background(), defaultPreset.Name)
	if err != nil {
		t.Fatalf("did not expect GetPreset() to return an error, got %+v", err)
	}

	err = p.DeletePreset(context.Background(), defaultPreset.Name)
	if err != nil {
		t.Fatalf("expected DeletePreset() not to return an error, got: %v", err)
	}

	_, err = p.GetPreset(context.Background(), defaultPreset.Name)
	if err == nil || err.Error() != "local preset not found" {
		t.Fatal("expected GetPreset() to return an error, not nil")
	}
}

func Test_mcProvider_Transcode(t *testing.T) {
	tests := []struct {
		cfg         *config.MediaConvert
		name        string
		job         *db.Job
		preset      db.Preset
		destination string
		wantJobReq  mediaconvert.CreateJobInput
		wantErr     bool
	}{
		{
			name: "a valid h264/aac mp4 transcode job is mapped correctly to a mediaconvert job input",
			cfg: &config.MediaConvert{
				Role:            "some-role",
				DefaultQueueARN: "some:default:queue:arn",
			},
			job: &db.Job{
				ID:           "jobID",
				ProviderName: Name,
				SourceMedia:  "s3://some/path.mp4",
				Outputs:      []db.TranscodeOutput{{Preset: db.PresetMap{Name: defaultPreset.Name}, FileName: "file1.mp4"}},
				Labels:       []string{"bill:some-bu"},
			},
			preset:      defaultPreset,
			destination: "s3://some/destination",
			wantJobReq: mediaconvert.CreateJobInput{
				Role:  aws.String("some-role"),
				Queue: aws.String("some:default:queue:arn"),
				Tags: map[string]string{
					"bill:some-bu": "true",
				},
				Settings: &mediaconvert.JobSettings{
					Inputs: []mediaconvert.Input{
						{
							AudioSelectors: map[string]mediaconvert.AudioSelector{
								"Audio Selector 1": {
									DefaultSelection: mediaconvert.AudioDefaultSelectionDefault,
								},
							},
							FileInput: aws.String("s3://some/path.mp4"),
							VideoSelector: &mediaconvert.VideoSelector{
								ColorSpace: mediaconvert.ColorSpaceFollow,
							},
							TimecodeSource: mediaconvert.InputTimecodeSourceZerobased,
						},
					},
					OutputGroups: []mediaconvert.OutputGroup{
						{
							OutputGroupSettings: &mediaconvert.OutputGroupSettings{
								Type: mediaconvert.OutputGroupTypeFileGroupSettings,
								FileGroupSettings: &mediaconvert.FileGroupSettings{
									Destination: aws.String("s3://some/destination/jobID/m"),
								},
							},
							Outputs: []mediaconvert.Output{
								{
									NameModifier: aws.String("file1"),
									ContainerSettings: &mediaconvert.ContainerSettings{
										Container: mediaconvert.ContainerTypeMp4,
									},
									VideoDescription: &mediaconvert.VideoDescription{
										Height:            aws.Int64(400),
										Width:             aws.Int64(300),
										RespondToAfd:      mediaconvert.RespondToAfdNone,
										ScalingBehavior:   mediaconvert.ScalingBehaviorDefault,
										TimecodeInsertion: mediaconvert.VideoTimecodeInsertionDisabled,
										AntiAlias:         mediaconvert.AntiAliasEnabled,
										VideoPreprocessors: &mediaconvert.VideoPreprocessor{
											Deinterlacer: &mediaconvert.Deinterlacer{
												Algorithm: mediaconvert.DeinterlaceAlgorithmInterpolate,
												Control:   mediaconvert.DeinterlacerControlNormal,
												Mode:      mediaconvert.DeinterlacerModeAdaptive,
											},
										},
										CodecSettings: &mediaconvert.VideoCodecSettings{
											Codec: mediaconvert.VideoCodecH264,
											H264Settings: &mediaconvert.H264Settings{
												Bitrate:            aws.Int64(400000),
												CodecLevel:         mediaconvert.H264CodecLevelAuto,
												CodecProfile:       mediaconvert.H264CodecProfileHigh,
												InterlaceMode:      mediaconvert.H264InterlaceModeProgressive,
												ParControl:         mediaconvert.H264ParControlSpecified,
												ParNumerator:       aws.Int64(1),
												ParDenominator:     aws.Int64(1),
												QualityTuningLevel: mediaconvert.H264QualityTuningLevelMultiPassHq,
												RateControlMode:    mediaconvert.H264RateControlModeVbr,
												GopSize:            aws.Float64(120),
												GopSizeUnits:       "FRAMES",
											},
										},
									},
									AudioDescriptions: []mediaconvert.AudioDescription{
										{
											CodecSettings: &mediaconvert.AudioCodecSettings{
												Codec: mediaconvert.AudioCodecAac,
												AacSettings: &mediaconvert.AacSettings{
													Bitrate:         aws.Int64(20000),
													CodecProfile:    mediaconvert.AacCodecProfileLc,
													CodingMode:      mediaconvert.AacCodingModeCodingMode20,
													RateControlMode: mediaconvert.AacRateControlModeCbr,
													SampleRate:      aws.Int64(defaultAudioSampleRate),
												},
											},
										},
									},
									Extension: aws.String("mp4"),
								},
							},
						},
					},
					TimecodeConfig: &mediaconvert.TimecodeConfig{
						Source: mediaconvert.TimecodeSourceZerobased,
					},
				},
			},
		},
		{
			name: "a valid h264/aac mp4 transcode job with interlaced input is mapped correctly to a mediaconvert job input",
			job: &db.Job{
				ID:           "jobID",
				ProviderName: Name,
				SourceMedia:  "s3://some/path.mp4",
				SourceInfo: db.File{
					ScanType: db.ScanTypeInterlaced,
				},
				Outputs: []db.TranscodeOutput{{Preset: db.PresetMap{Name: defaultPreset.Name}, FileName: "file1.mp4"}},
			},
			preset:      defaultPreset,
			destination: "s3://some/destination",
			wantJobReq: mediaconvert.CreateJobInput{
				Role:  aws.String(""),
				Queue: aws.String(""),
				Tags:  map[string]string{},
				Settings: &mediaconvert.JobSettings{
					Inputs: []mediaconvert.Input{
						{
							AudioSelectors: map[string]mediaconvert.AudioSelector{
								"Audio Selector 1": {
									DefaultSelection: mediaconvert.AudioDefaultSelectionDefault,
								},
							},
							FileInput: aws.String("s3://some/path.mp4"),
							VideoSelector: &mediaconvert.VideoSelector{
								ColorSpace: mediaconvert.ColorSpaceFollow,
							},
							TimecodeSource: mediaconvert.InputTimecodeSourceZerobased,
						},
					},
					OutputGroups: []mediaconvert.OutputGroup{
						{
							OutputGroupSettings: &mediaconvert.OutputGroupSettings{
								Type: mediaconvert.OutputGroupTypeFileGroupSettings,
								FileGroupSettings: &mediaconvert.FileGroupSettings{
									Destination: aws.String("s3://some/destination/jobID/m"),
								},
							},
							Outputs: []mediaconvert.Output{
								{
									NameModifier: aws.String("file1"),
									ContainerSettings: &mediaconvert.ContainerSettings{
										Container: mediaconvert.ContainerTypeMp4,
									},
									VideoDescription: &mediaconvert.VideoDescription{
										Height:            aws.Int64(400),
										Width:             aws.Int64(300),
										RespondToAfd:      mediaconvert.RespondToAfdNone,
										ScalingBehavior:   mediaconvert.ScalingBehaviorDefault,
										TimecodeInsertion: mediaconvert.VideoTimecodeInsertionDisabled,
										AntiAlias:         mediaconvert.AntiAliasEnabled,
										VideoPreprocessors: &mediaconvert.VideoPreprocessor{
											Deinterlacer: &mediaconvert.Deinterlacer{
												Algorithm: mediaconvert.DeinterlaceAlgorithmInterpolate,
												Control:   mediaconvert.DeinterlacerControlNormal,
												Mode:      mediaconvert.DeinterlacerModeDeinterlace,
											},
										},
										CodecSettings: &mediaconvert.VideoCodecSettings{
											Codec: mediaconvert.VideoCodecH264,
											H264Settings: &mediaconvert.H264Settings{
												Bitrate:            aws.Int64(400000),
												CodecLevel:         mediaconvert.H264CodecLevelAuto,
												CodecProfile:       mediaconvert.H264CodecProfileHigh,
												InterlaceMode:      mediaconvert.H264InterlaceModeProgressive,
												ParControl:         mediaconvert.H264ParControlSpecified,
												ParNumerator:       aws.Int64(1),
												ParDenominator:     aws.Int64(1),
												QualityTuningLevel: mediaconvert.H264QualityTuningLevelMultiPassHq,
												RateControlMode:    mediaconvert.H264RateControlModeVbr,
												GopSize:            aws.Float64(120),
												GopSizeUnits:       "FRAMES",
											},
										},
									},
									AudioDescriptions: []mediaconvert.AudioDescription{
										{
											CodecSettings: &mediaconvert.AudioCodecSettings{
												Codec: mediaconvert.AudioCodecAac,
												AacSettings: &mediaconvert.AacSettings{
													Bitrate:         aws.Int64(20000),
													CodecProfile:    mediaconvert.AacCodecProfileLc,
													CodingMode:      mediaconvert.AacCodingModeCodingMode20,
													RateControlMode: mediaconvert.AacRateControlModeCbr,
													SampleRate:      aws.Int64(defaultAudioSampleRate),
												},
											},
										},
									},
									Extension: aws.String("mp4"),
								},
							},
						},
					},
					TimecodeConfig: &mediaconvert.TimecodeConfig{
						Source: mediaconvert.TimecodeSourceZerobased,
					},
				},
			},
		},
		{
			name: "a valid h264/aac mp4 transcode job with progressive input is mapped correctly to a mediaconvert job input",
			job: &db.Job{
				ID:           "jobID",
				ProviderName: Name,
				SourceMedia:  "s3://some/path.mp4",
				SourceInfo: db.File{
					ScanType: db.ScanTypeProgressive,
				},
				Outputs: []db.TranscodeOutput{{Preset: db.PresetMap{Name: defaultPreset.Name}, FileName: "file1.mp4"}},
			},
			preset:      defaultPreset,
			destination: "s3://some/destination",
			wantJobReq: mediaconvert.CreateJobInput{
				Role:  aws.String(""),
				Queue: aws.String(""),
				Tags:  map[string]string{},
				Settings: &mediaconvert.JobSettings{
					Inputs: []mediaconvert.Input{
						{
							AudioSelectors: map[string]mediaconvert.AudioSelector{
								"Audio Selector 1": {
									DefaultSelection: mediaconvert.AudioDefaultSelectionDefault,
								},
							},
							FileInput: aws.String("s3://some/path.mp4"),
							VideoSelector: &mediaconvert.VideoSelector{
								ColorSpace: mediaconvert.ColorSpaceFollow,
							},
							TimecodeSource: mediaconvert.InputTimecodeSourceZerobased,
						},
					},
					OutputGroups: []mediaconvert.OutputGroup{
						{
							OutputGroupSettings: &mediaconvert.OutputGroupSettings{
								Type: mediaconvert.OutputGroupTypeFileGroupSettings,
								FileGroupSettings: &mediaconvert.FileGroupSettings{
									Destination: aws.String("s3://some/destination/jobID/m"),
								},
							},
							Outputs: []mediaconvert.Output{
								{
									NameModifier: aws.String("file1"),
									ContainerSettings: &mediaconvert.ContainerSettings{
										Container: mediaconvert.ContainerTypeMp4,
									},
									VideoDescription: &mediaconvert.VideoDescription{
										Height:             aws.Int64(400),
										Width:              aws.Int64(300),
										RespondToAfd:       mediaconvert.RespondToAfdNone,
										ScalingBehavior:    mediaconvert.ScalingBehaviorDefault,
										TimecodeInsertion:  mediaconvert.VideoTimecodeInsertionDisabled,
										AntiAlias:          mediaconvert.AntiAliasEnabled,
										VideoPreprocessors: &mediaconvert.VideoPreprocessor{},
										CodecSettings: &mediaconvert.VideoCodecSettings{
											Codec: mediaconvert.VideoCodecH264,
											H264Settings: &mediaconvert.H264Settings{
												Bitrate:            aws.Int64(400000),
												CodecLevel:         mediaconvert.H264CodecLevelAuto,
												CodecProfile:       mediaconvert.H264CodecProfileHigh,
												InterlaceMode:      mediaconvert.H264InterlaceModeProgressive,
												ParControl:         mediaconvert.H264ParControlSpecified,
												ParNumerator:       aws.Int64(1),
												ParDenominator:     aws.Int64(1),
												QualityTuningLevel: mediaconvert.H264QualityTuningLevelMultiPassHq,
												RateControlMode:    mediaconvert.H264RateControlModeVbr,
												GopSize:            aws.Float64(120),
												GopSizeUnits:       "FRAMES",
											},
										},
									},
									AudioDescriptions: []mediaconvert.AudioDescription{
										{
											CodecSettings: &mediaconvert.AudioCodecSettings{
												Codec: mediaconvert.AudioCodecAac,
												AacSettings: &mediaconvert.AacSettings{
													Bitrate:         aws.Int64(20000),
													CodecProfile:    mediaconvert.AacCodecProfileLc,
													CodingMode:      mediaconvert.AacCodingModeCodingMode20,
													RateControlMode: mediaconvert.AacRateControlModeCbr,
													SampleRate:      aws.Int64(defaultAudioSampleRate),
												},
											},
										},
									},
									Extension: aws.String("mp4"),
								},
							},
						},
					},
					TimecodeConfig: &mediaconvert.TimecodeConfig{
						Source: mediaconvert.TimecodeSourceZerobased,
					},
				},
			},
		},
		{
			name: "a valid h265 video-only mp4 transcode job is mapped correctly to a mediaconvert job input",
			job: &db.Job{
				ID:           "jobID",
				ProviderName: Name,
				SourceMedia:  "s3://some/path.mp4",
				Outputs:      []db.TranscodeOutput{{Preset: db.PresetMap{Name: h265Preset.Name}, FileName: "file1.mp4"}},
			},
			preset:      h265Preset,
			destination: "s3://some/destination",
			wantJobReq: mediaconvert.CreateJobInput{
				Role:  aws.String(""),
				Queue: aws.String(""),
				Tags:  map[string]string{},
				Settings: &mediaconvert.JobSettings{
					Inputs: []mediaconvert.Input{
						{
							AudioSelectors: map[string]mediaconvert.AudioSelector{
								"Audio Selector 1": {
									DefaultSelection: mediaconvert.AudioDefaultSelectionDefault,
								},
							},
							FileInput: aws.String("s3://some/path.mp4"),
							VideoSelector: &mediaconvert.VideoSelector{
								ColorSpace: mediaconvert.ColorSpaceFollow,
							},
							TimecodeSource: mediaconvert.InputTimecodeSourceZerobased,
						},
					},
					OutputGroups: []mediaconvert.OutputGroup{
						{
							OutputGroupSettings: &mediaconvert.OutputGroupSettings{
								Type: mediaconvert.OutputGroupTypeFileGroupSettings,
								FileGroupSettings: &mediaconvert.FileGroupSettings{
									Destination: aws.String("s3://some/destination/jobID/m"),
								},
							},
							Outputs: []mediaconvert.Output{
								{
									NameModifier: aws.String("file1"),
									ContainerSettings: &mediaconvert.ContainerSettings{
										Container: mediaconvert.ContainerTypeMp4,
									},
									VideoDescription: &mediaconvert.VideoDescription{
										Height:            aws.Int64(400),
										Width:             aws.Int64(300),
										RespondToAfd:      mediaconvert.RespondToAfdNone,
										ScalingBehavior:   mediaconvert.ScalingBehaviorDefault,
										TimecodeInsertion: mediaconvert.VideoTimecodeInsertionDisabled,
										AntiAlias:         mediaconvert.AntiAliasEnabled,
										VideoPreprocessors: &mediaconvert.VideoPreprocessor{
											Deinterlacer: &mediaconvert.Deinterlacer{
												Algorithm: mediaconvert.DeinterlaceAlgorithmInterpolate,
												Control:   mediaconvert.DeinterlacerControlNormal,
												Mode:      mediaconvert.DeinterlacerModeAdaptive,
											},
										},
										CodecSettings: &mediaconvert.VideoCodecSettings{
											Codec: mediaconvert.VideoCodecH265,
											H265Settings: &mediaconvert.H265Settings{
												Bitrate:                        aws.Int64(400000),
												GopSize:                        aws.Float64(120),
												GopSizeUnits:                   "FRAMES",
												CodecLevel:                     mediaconvert.H265CodecLevelAuto,
												CodecProfile:                   mediaconvert.H265CodecProfileMainMain,
												InterlaceMode:                  mediaconvert.H265InterlaceModeProgressive,
												ParControl:                     mediaconvert.H265ParControlSpecified,
												ParNumerator:                   aws.Int64(1),
												ParDenominator:                 aws.Int64(1),
												QualityTuningLevel:             mediaconvert.H265QualityTuningLevelSinglePassHq,
												RateControlMode:                mediaconvert.H265RateControlModeCbr,
												WriteMp4PackagingType:          mediaconvert.H265WriteMp4PackagingTypeHvc1,
												AlternateTransferFunctionSei:   mediaconvert.H265AlternateTransferFunctionSeiDisabled,
												SpatialAdaptiveQuantization:    mediaconvert.H265SpatialAdaptiveQuantizationEnabled,
												TemporalAdaptiveQuantization:   mediaconvert.H265TemporalAdaptiveQuantizationEnabled,
												FlickerAdaptiveQuantization:    mediaconvert.H265FlickerAdaptiveQuantizationEnabled,
												SceneChangeDetect:              mediaconvert.H265SceneChangeDetectEnabled,
												UnregisteredSeiTimecode:        mediaconvert.H265UnregisteredSeiTimecodeDisabled,
												SampleAdaptiveOffsetFilterMode: mediaconvert.H265SampleAdaptiveOffsetFilterModeAdaptive,
											},
										},
									},
									Extension: aws.String("mp4"),
								},
							},
						},
					},
					TimecodeConfig: &mediaconvert.TimecodeConfig{
						Source: mediaconvert.TimecodeSourceZerobased,
					},
				},
			},
		},
		{
			name: "a valid av1 video-only mp4 transcode job is mapped correctly to a mediaconvert job input",
			job: &db.Job{
				ID:           "jobID",
				ProviderName: Name,
				SourceMedia:  "s3://some/path.mp4",
				Outputs:      []db.TranscodeOutput{{Preset: db.PresetMap{Name: av1Preset.Name}, FileName: "file1.mp4"}},
			},
			preset:      av1Preset,
			destination: "s3://some/destination",
			wantJobReq: mediaconvert.CreateJobInput{
				Role:  aws.String(""),
				Queue: aws.String(""),
				Tags:  map[string]string{},
				Settings: &mediaconvert.JobSettings{
					Inputs: []mediaconvert.Input{
						{
							AudioSelectors: map[string]mediaconvert.AudioSelector{
								"Audio Selector 1": {
									DefaultSelection: mediaconvert.AudioDefaultSelectionDefault,
								},
							},
							FileInput: aws.String("s3://some/path.mp4"),
							VideoSelector: &mediaconvert.VideoSelector{
								ColorSpace: mediaconvert.ColorSpaceFollow,
							},
							TimecodeSource: mediaconvert.InputTimecodeSourceZerobased,
						},
					},
					OutputGroups: []mediaconvert.OutputGroup{
						{
							OutputGroupSettings: &mediaconvert.OutputGroupSettings{
								Type: mediaconvert.OutputGroupTypeFileGroupSettings,
								FileGroupSettings: &mediaconvert.FileGroupSettings{
									Destination: aws.String("s3://some/destination/jobID/m"),
								},
							},
							Outputs: []mediaconvert.Output{
								{
									NameModifier: aws.String("file1"),
									ContainerSettings: &mediaconvert.ContainerSettings{
										Container: mediaconvert.ContainerTypeMp4,
									},
									VideoDescription: &mediaconvert.VideoDescription{
										Height:            aws.Int64(400),
										Width:             aws.Int64(300),
										RespondToAfd:      mediaconvert.RespondToAfdNone,
										ScalingBehavior:   mediaconvert.ScalingBehaviorDefault,
										TimecodeInsertion: mediaconvert.VideoTimecodeInsertionDisabled,
										AntiAlias:         mediaconvert.AntiAliasEnabled,
										VideoPreprocessors: &mediaconvert.VideoPreprocessor{
											Deinterlacer: &mediaconvert.Deinterlacer{
												Algorithm: mediaconvert.DeinterlaceAlgorithmInterpolate,
												Control:   mediaconvert.DeinterlacerControlNormal,
												Mode:      mediaconvert.DeinterlacerModeAdaptive,
											},
										},
										CodecSettings: &mediaconvert.VideoCodecSettings{
											Codec: mediaconvert.VideoCodecAv1,
											Av1Settings: &mediaconvert.Av1Settings{
												MaxBitrate: aws.Int64(400000),
												GopSize:    aws.Float64(120),
												QvbrSettings: &mediaconvert.Av1QvbrSettings{
													QvbrQualityLevel:         aws.Int64(7),
													QvbrQualityLevelFineTune: aws.Float64(0),
												},
												RateControlMode: mediaconvert.Av1RateControlModeQvbr,
											},
										},
									},
									Extension: aws.String("mp4"),
								},
							},
						},
					},
					TimecodeConfig: &mediaconvert.TimecodeConfig{
						Source: mediaconvert.TimecodeSourceZerobased,
					},
				},
			},
		},
		{
			name: "a job is mapped correctly to a mediaconvert job input when a preferred queue is defined",
			cfg: &config.MediaConvert{
				DefaultQueueARN:   "some:default:queue:arn",
				PreferredQueueARN: "some:preferred:queue:arn",
			},
			job: &db.Job{
				ID:           "jobID",
				ProviderName: Name,
				SourceMedia:  "s3://some/path.mp4",
				Outputs:      []db.TranscodeOutput{{Preset: db.PresetMap{Name: defaultPreset.Name}, FileName: "file1.mp4"}},
			},
			preset:      defaultPreset,
			destination: "s3://some/destination",
			wantJobReq: mediaconvert.CreateJobInput{
				Role:            aws.String(""),
				Queue:           aws.String("some:preferred:queue:arn"),
				HopDestinations: []mediaconvert.HopDestination{{WaitMinutes: aws.Int64(defaultQueueHopTimeoutMins)}},
				Tags:            map[string]string{},
				Settings: &mediaconvert.JobSettings{
					Inputs: []mediaconvert.Input{
						{
							AudioSelectors: map[string]mediaconvert.AudioSelector{
								"Audio Selector 1": {
									DefaultSelection: mediaconvert.AudioDefaultSelectionDefault,
								},
							},
							FileInput: aws.String("s3://some/path.mp4"),
							VideoSelector: &mediaconvert.VideoSelector{
								ColorSpace: mediaconvert.ColorSpaceFollow,
							},
							TimecodeSource: mediaconvert.InputTimecodeSourceZerobased,
						},
					},
					OutputGroups: []mediaconvert.OutputGroup{
						{
							OutputGroupSettings: &mediaconvert.OutputGroupSettings{
								Type: mediaconvert.OutputGroupTypeFileGroupSettings,
								FileGroupSettings: &mediaconvert.FileGroupSettings{
									Destination: aws.String("s3://some/destination/jobID/m"),
								},
							},
							Outputs: []mediaconvert.Output{
								{
									NameModifier: aws.String("file1"),
									ContainerSettings: &mediaconvert.ContainerSettings{
										Container: mediaconvert.ContainerTypeMp4,
									},
									VideoDescription: &mediaconvert.VideoDescription{
										Height:            aws.Int64(400),
										Width:             aws.Int64(300),
										RespondToAfd:      mediaconvert.RespondToAfdNone,
										ScalingBehavior:   mediaconvert.ScalingBehaviorDefault,
										TimecodeInsertion: mediaconvert.VideoTimecodeInsertionDisabled,
										AntiAlias:         mediaconvert.AntiAliasEnabled,
										VideoPreprocessors: &mediaconvert.VideoPreprocessor{
											Deinterlacer: &mediaconvert.Deinterlacer{
												Algorithm: mediaconvert.DeinterlaceAlgorithmInterpolate,
												Control:   mediaconvert.DeinterlacerControlNormal,
												Mode:      mediaconvert.DeinterlacerModeAdaptive,
											},
										},
										CodecSettings: &mediaconvert.VideoCodecSettings{
											Codec: mediaconvert.VideoCodecH264,
											H264Settings: &mediaconvert.H264Settings{
												Bitrate:            aws.Int64(400000),
												CodecLevel:         mediaconvert.H264CodecLevelAuto,
												CodecProfile:       mediaconvert.H264CodecProfileHigh,
												InterlaceMode:      mediaconvert.H264InterlaceModeProgressive,
												QualityTuningLevel: mediaconvert.H264QualityTuningLevelMultiPassHq,
												RateControlMode:    mediaconvert.H264RateControlModeVbr,
												GopSize:            aws.Float64(120),
												GopSizeUnits:       mediaconvert.H264GopSizeUnitsFrames,
												ParControl:         mediaconvert.H264ParControlSpecified,
												ParNumerator:       aws.Int64(1),
												ParDenominator:     aws.Int64(1),
											},
										},
									},
									AudioDescriptions: []mediaconvert.AudioDescription{
										{
											CodecSettings: &mediaconvert.AudioCodecSettings{
												Codec: mediaconvert.AudioCodecAac,
												AacSettings: &mediaconvert.AacSettings{
													Bitrate:         aws.Int64(20000),
													CodecProfile:    mediaconvert.AacCodecProfileLc,
													CodingMode:      mediaconvert.AacCodingModeCodingMode20,
													RateControlMode: mediaconvert.AacRateControlModeCbr,
													SampleRate:      aws.Int64(defaultAudioSampleRate),
												},
											},
										},
									},
									Extension: aws.String("mp4"),
								},
							},
						},
					},
					TimecodeConfig: &mediaconvert.TimecodeConfig{
						Source: mediaconvert.TimecodeSourceZerobased,
					},
				},
			},
		},
		{
			name: "JobWithAudioDownmixAndTimeCodeBurninForMovOutput",
			cfg: &config.MediaConvert{
				DefaultQueueARN:   "some:default:queue:arn",
				PreferredQueueARN: "some:preferred:queue:arn",
			},
			job: &db.Job{
				ID:           "jobID",
				ProviderName: Name,
				SourceMedia:  "s3://some/path.mov",
				Outputs:      []db.TranscodeOutput{{Preset: db.PresetMap{Name: tcBurninPreset.Name}, FileName: "file1.mov"}},
				AudioDownmix: &db.AudioDownmix{
					SrcChannels: []db.AudioChannel{
						{TrackIdx: 1, ChannelIdx: 1, Layout: "L"},
						{TrackIdx: 1, ChannelIdx: 2, Layout: "R"},
						{TrackIdx: 1, ChannelIdx: 3, Layout: "C"},
						{TrackIdx: 1, ChannelIdx: 4, Layout: "LFE"},
						{TrackIdx: 1, ChannelIdx: 5, Layout: "Ls"},
						{TrackIdx: 1, ChannelIdx: 6, Layout: "Rs"},
					},
					DestChannels: []db.AudioChannel{
						{TrackIdx: 1, ChannelIdx: 1, Layout: "L"},
						{TrackIdx: 1, ChannelIdx: 2, Layout: "R"},
					},
				},
			},
			preset:      tcBurninPreset,
			destination: "s3://some/destination",
			wantJobReq: mediaconvert.CreateJobInput{
				Role:            aws.String(""),
				Queue:           aws.String("some:preferred:queue:arn"),
				HopDestinations: []mediaconvert.HopDestination{{WaitMinutes: aws.Int64(defaultQueueHopTimeoutMins)}},
				Tags:            map[string]string{},
				Settings: &mediaconvert.JobSettings{
					Inputs: []mediaconvert.Input{
						{
							AudioSelectors: map[string]mediaconvert.AudioSelector{
								"Audio Selector 1": getAudioSelector(6, 2, []int64{1}, []mediaconvert.OutputChannelMapping{
									{InputChannels: []int64{0, -60, 0, -60, 0, -60}},
									{InputChannels: []int64{-60, 0, 0, -60, -60, 0}},
								}),
							},
							FileInput: aws.String("s3://some/path.mov"),
							VideoSelector: &mediaconvert.VideoSelector{
								ColorSpace: mediaconvert.ColorSpaceFollow,
							},
							TimecodeSource: mediaconvert.InputTimecodeSourceZerobased,
						},
					},
					OutputGroups: []mediaconvert.OutputGroup{
						{
							OutputGroupSettings: &mediaconvert.OutputGroupSettings{
								Type: mediaconvert.OutputGroupTypeFileGroupSettings,
								FileGroupSettings: &mediaconvert.FileGroupSettings{
									Destination: aws.String("s3://some/destination/jobID/m"),
								},
							},
							Outputs: []mediaconvert.Output{
								{
									NameModifier: aws.String("file1"),
									ContainerSettings: &mediaconvert.ContainerSettings{
										Container: mediaconvert.ContainerTypeMov,
										MovSettings: &mediaconvert.MovSettings{
											ClapAtom:           mediaconvert.MovClapAtomExclude,
											CslgAtom:           mediaconvert.MovCslgAtomInclude,
											PaddingControl:     mediaconvert.MovPaddingControlOmneon,
											Reference:          mediaconvert.MovReferenceSelfContained,
											Mpeg2FourCCControl: mediaconvert.MovMpeg2FourCCControlMpeg,
										},
									},
									VideoDescription: &mediaconvert.VideoDescription{
										Height:            aws.Int64(400),
										Width:             aws.Int64(300),
										RespondToAfd:      mediaconvert.RespondToAfdNone,
										ScalingBehavior:   mediaconvert.ScalingBehaviorDefault,
										TimecodeInsertion: mediaconvert.VideoTimecodeInsertionDisabled,
										AntiAlias:         mediaconvert.AntiAliasEnabled,
										VideoPreprocessors: &mediaconvert.VideoPreprocessor{
											Deinterlacer: &mediaconvert.Deinterlacer{
												Algorithm: mediaconvert.DeinterlaceAlgorithmInterpolate,
												Control:   mediaconvert.DeinterlacerControlNormal,
												Mode:      mediaconvert.DeinterlacerModeAdaptive,
											},
											TimecodeBurnin: &mediaconvert.TimecodeBurnin{
												FontSize: aws.Int64(12),
												Position: mediaconvert.TimecodeBurninPositionBottomLeft,
												Prefix:   aws.String(""),
											},
										},
										CodecSettings: &mediaconvert.VideoCodecSettings{
											Codec: mediaconvert.VideoCodecH264,
											H264Settings: &mediaconvert.H264Settings{
												Bitrate:            aws.Int64(400000),
												CodecLevel:         mediaconvert.H264CodecLevelAuto,
												CodecProfile:       mediaconvert.H264CodecProfileHigh,
												InterlaceMode:      mediaconvert.H264InterlaceModeProgressive,
												QualityTuningLevel: mediaconvert.H264QualityTuningLevelMultiPassHq,
												RateControlMode:    mediaconvert.H264RateControlModeVbr,
												GopSize:            aws.Float64(120),
												GopSizeUnits:       mediaconvert.H264GopSizeUnitsFrames,
												ParControl:         mediaconvert.H264ParControlSpecified,
												ParNumerator:       aws.Int64(1),
												ParDenominator:     aws.Int64(1),
											},
										},
									},
									AudioDescriptions: []mediaconvert.AudioDescription{
										{
											CodecSettings: &mediaconvert.AudioCodecSettings{
												Codec: mediaconvert.AudioCodecAac,
												AacSettings: &mediaconvert.AacSettings{
													Bitrate:         aws.Int64(20000),
													CodecProfile:    mediaconvert.AacCodecProfileLc,
													CodingMode:      mediaconvert.AacCodingModeCodingMode20,
													RateControlMode: mediaconvert.AacRateControlModeCbr,
													SampleRate:      aws.Int64(defaultAudioSampleRate),
												},
											},
										},
									},
									Extension: aws.String("mov"),
								},
							},
						},
					},
					TimecodeConfig: &mediaconvert.TimecodeConfig{
						Source: mediaconvert.TimecodeSourceZerobased,
					},
				},
			},
		},
		//{
		//	name: "acceleration is enabled and the default queue is used when a source has a large filesize",
		//	cfg: &config.MediaConvert{
		//		DefaultQueueARN:   "some:default:queue:arn",
		//		PreferredQueueARN: "some:preferred:queue:arn",
		//	},
		//	job: &db.Job{
		//		ID:           "jobID",
		//		ProviderName: Name,
		//		SourceMedia:  "s3://some/path.mp4",
		//		SourceInfo:   db.File{FileSize: 1_000_000_000},
		//		Outputs:      []db.TranscodeOutput{{Preset: db.PresetMap{Name: audioOnlyPreset.Name}, FileName: "file1.mp4"}},
		//	},
		//	preset:      audioOnlyPreset,
		//	destination: "s3://some/destination",
		//	wantJobReq: mediaconvert.CreateJobInput{
		//		AccelerationSettings: &mediaconvert.AccelerationSettings{
		//			Mode: mediaconvert.AccelerationModePreferred,
		//		},
		//		Role:  aws.String(""),
		//		Queue: aws.String("some:default:queue:arn"),
		//		Tags: map[string]string{},
		//		Settings: &mediaconvert.JobSettings{
		//			Inputs: []mediaconvert.Input{
		//				{
		//					AudioSelectors: map[string]mediaconvert.AudioSelector{
		//						"Audio Selector 1": {
		//							DefaultSelection: mediaconvert.AudioDefaultSelectionDefault,
		//						},
		//					},
		//					FileInput: aws.String("s3://some/path.mp4"),
		//					VideoSelector: &mediaconvert.VideoSelector{
		//						ColorSpace: mediaconvert.ColorSpaceFollow,
		//					},
		//					TimecodeSource: mediaconvert.InputTimecodeSourceZerobased,
		//				},
		//			},
		//			OutputGroups: []mediaconvert.OutputGroup{
		//				{
		//					OutputGroupSettings: &mediaconvert.OutputGroupSettings{
		//						Type: mediaconvert.OutputGroupTypeFileGroupSettings,
		//						FileGroupSettings: &mediaconvert.FileGroupSettings{
		//							Destination: aws.String("s3://some/destination/jobID/m"),
		//						},
		//					},
		//					Outputs: []mediaconvert.Output{
		//						{
		//							NameModifier: aws.String("file1"),
		//							ContainerSettings: &mediaconvert.ContainerSettings{
		//								Container: mediaconvert.ContainerTypeMp4,
		//							},
		//							AudioDescriptions: []mediaconvert.AudioDescription{
		//								{
		//									CodecSettings: &mediaconvert.AudioCodecSettings{
		//										Codec: mediaconvert.AudioCodecAac,
		//										AacSettings: &mediaconvert.AacSettings{
		//											Bitrate:         aws.Int64(20000),
		//											CodecProfile:    mediaconvert.AacCodecProfileLc,
		//											CodingMode:      mediaconvert.AacCodingModeCodingMode20,
		//											RateControlMode: mediaconvert.AacRateControlModeCbr,
		//											SampleRate:      aws.Int64(defaultAudioSampleRate),
		//										},
		//									},
		//								},
		//							},
		//							Extension: aws.String("mp4"),
		//						},
		//					},
		//				},
		//			},
		//			TimecodeConfig: &mediaconvert.TimecodeConfig{
		//				Source: mediaconvert.TimecodeSourceZerobased,
		//			},
		//		},
		//	},
		//},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			repo, err := fakeDBWithPresets(tt.preset)
			if err != nil {
				t.Error(err)
				return
			}

			if tt.cfg == nil {
				tt.cfg = &config.MediaConvert{Destination: tt.destination}
			} else {
				tt.cfg.Destination = tt.destination
			}

			client := &testMediaConvertClient{t: t}
			p := &mcProvider{
				client:     client,
				cfg:        tt.cfg,
				repository: repo,
			}
			_, err = p.Transcode(context.Background(), tt.job)
			if (err != nil) != tt.wantErr {
				t.Errorf("mcProvider.Transcode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if g, e := client.createJobCalledWith, tt.wantJobReq; !reflect.DeepEqual(g, e) {
				t.Fatalf("Transcode(): wrong job request\nWant %+v\nGot %+v\nDiff %s", e,
					g, cmp.Diff(e, g))
			}
		})
	}
}

func Test_mcProvider_CancelJob(t *testing.T) {
	jobID := "some_job_id"
	client := &testMediaConvertClient{t: t}
	p := &mcProvider{client: client}
	err := p.CancelJob(context.Background(), jobID)
	if err != nil {
		t.Fatalf("expected CancelJob() not to return an error, got: %v", err)
	}

	if g, e := client.cancelJobCalledWith, jobID; g != e {
		t.Fatalf("got %q, expected %q", g, e)
	}
}

func Test_mcProvider_Healthcheck(t *testing.T) {
	client := &testMediaConvertClient{t: t}
	p := &mcProvider{client: client}

	err := p.Healthcheck()
	if err != nil {
		t.Fatalf("expected Healthcheck() not to return an error, got: %v", err)
	}

	if !client.listJobsCalled {
		t.Fatal("expected Healthcheck() to call ListJobs")
	}
}

func Test_mcProvider_JobStatus(t *testing.T) {
	tests := []struct {
		name        string
		destination string
		mcJob       mediaconvert.Job
		wantStatus  provider.JobStatus
		wantErr     bool
	}{
		{
			name:        "a job that has been queued returns the correct status",
			destination: "s3://some/destination",
			mcJob: mediaconvert.Job{
				Status: mediaconvert.JobStatusSubmitted,
			},
			wantStatus: provider.JobStatus{
				Status:       provider.StatusQueued,
				ProviderName: Name,
				Output: provider.JobOutput{
					Destination: "s3://some/destination/jobID/",
				},
			},
		},
		{
			name:        "a job that is currently transcoding returns the correct status",
			destination: "s3://some/destination",
			mcJob: mediaconvert.Job{
				Status:             mediaconvert.JobStatusProgressing,
				JobPercentComplete: aws.Int64(42),
			},
			wantStatus: provider.JobStatus{
				Status:       provider.StatusStarted,
				ProviderName: Name,
				Progress:     42,
				Output: provider.JobOutput{
					Destination: "s3://some/destination/jobID/",
				},
			},
		},
		{
			name:        "a job that has finished transcoding returns the correct status",
			destination: "s3://some/destination",
			mcJob: mediaconvert.Job{
				Status: mediaconvert.JobStatusComplete,
				Settings: &mediaconvert.JobSettings{
					OutputGroups: []mediaconvert.OutputGroup{
						{
							OutputGroupSettings: &mediaconvert.OutputGroupSettings{
								Type: mediaconvert.OutputGroupTypeFileGroupSettings,
								FileGroupSettings: &mediaconvert.FileGroupSettings{
									Destination: aws.String("s3://some/destination/jobID/m"),
								},
							},
							Outputs: []mediaconvert.Output{
								{
									NameModifier: aws.String("_modifier"),
									VideoDescription: &mediaconvert.VideoDescription{
										Height: aws.Int64(102),
										Width:  aws.Int64(324),
									},
									ContainerSettings: &mediaconvert.ContainerSettings{
										Container: mediaconvert.ContainerTypeMp4,
									},
								},
								{
									NameModifier: aws.String("_another_modifier"),
									ContainerSettings: &mediaconvert.ContainerSettings{
										Container: mediaconvert.ContainerTypeMp4,
									},
								},
								{
									NameModifier: aws.String("_123"),
									ContainerSettings: &mediaconvert.ContainerSettings{
										Container: mediaconvert.ContainerTypeM2ts,
									},
								},
								{
									ContainerSettings: &mediaconvert.ContainerSettings{
										Container: mediaconvert.ContainerTypeM2ts,
									},
								},
							},
						},
					},
				},
			},
			wantStatus: provider.JobStatus{
				Status:       provider.StatusFinished,
				ProviderName: Name,
				Progress:     100,
				Output: provider.JobOutput{
					Destination: "s3://some/destination/jobID/",
					Files: []provider.OutputFile{
						{
							Path:      "s3://some/destination/jobID/m_modifier.mp4",
							Container: "mp4",
							Height:    102,
							Width:     324,
						},
						{
							Path:      "s3://some/destination/jobID/m_another_modifier.mp4",
							Container: "mp4",
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			client := &testMediaConvertClient{
				t:                   t,
				jobReturnedByGetJob: tt.mcJob,
			}

			p := &mcProvider{client: client, cfg: &config.MediaConvert{
				Destination: tt.destination,
			}}

			status, err := p.JobStatus(context.Background(), &defaultJob)
			if (err != nil) != tt.wantErr {
				t.Errorf("mcProvider.JobStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if g, e := status, &tt.wantStatus; !reflect.DeepEqual(g, e) {
				t.Fatalf("mcProvider.JobStatus(): wrong job request\nWant %+v\nGot %+v\nDiff %s", e,
					g, cmp.Diff(e, g))
			}
		})
	}
}

func fakeDBWithPresets(presets ...db.Preset) (db.Repository, error) {
	fakeDB := dbtest.NewFakeRepository(false)

	for _, preset := range presets {
		err := fakeDB.CreateLocalPreset(&db.LocalPreset{Name: preset.Name, Preset: preset})
		if err != nil {
			return nil, err
		}

		err = fakeDB.CreatePresetMap(&db.PresetMap{
			Name: preset.Name,
			ProviderMapping: map[string]string{
				Name: preset.Name,
			},
		})
		if err != nil {
			return nil, err
		}
	}

	return fakeDB, nil
}
