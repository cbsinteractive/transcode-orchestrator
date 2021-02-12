package mediaconvert

import (
	"context"
	"reflect"
	"strings"
	"testing"

	mc "github.com/aws/aws-sdk-go-v2/service/mediaconvert"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/cbsinteractive/transcode-orchestrator/config"
	"github.com/cbsinteractive/transcode-orchestrator/db"
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
			Width:         300,
			Height:        400,
			Codec:         "h264",
			Bitrate:       400000,
			GopSize:       120,
			GopUnit:       "frames",
			InterlaceMode: "progressive",
		},
		Audio: db.AudioPreset{
			Codec:   "aac",
			Bitrate: 20000,
		},
	}

	h265Preset = db.Preset{
		Name:        "another_preset_name",
		Description: "test_desc",
		Container:   "mp4",
		RateControl: "CBR",
		TwoPass:     false,
		Video: db.VideoPreset{
			Width:         300,
			Height:        400,
			Codec:         "h265",
			Bitrate:       400000,
			GopSize:       120,
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
			Width:         300,
			Height:        400,
			Codec:         "av1",
			Bitrate:       400000,
			GopSize:       120,
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
			Bitrate: 20000,
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
			Width:         300,
			Height:        400,
			Codec:         "h264",
			Bitrate:       400000,
			GopSize:       120,
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
			Bitrate: 20000,
		},
	}

	defaultJob = db.Job{
		ID:           "jobID",
		ProviderName: Name,
		SourceMedia:  "s3://some/path.mp4",
		SourceInfo:   db.File{ScanType: db.ScanTypeUnknown},
		Outputs: []db.TranscodeOutput{
			{Preset: db.Preset{Name: "preset_name"}, FileName: "file1.mp4"},
			{Preset: db.Preset{Name: "another_preset_name"}, FileName: "file2.mp4"},
		},
		StreamingParams: db.StreamingParams{SegmentDuration: 6},
	}
)

func Test_driver_CreatePreset_fields(t *testing.T) {
	tests := []struct {
		name           string
		presetModifier func(preset db.Preset) db.Preset
		assertion      func(mc.Output, *testing.T)
		wantErrMsg     string
	}{
		{
			name: "hls presets are set correctly",
			presetModifier: func(p db.Preset) db.Preset {
				p.Container = "m3u8"
				return p
			},
			assertion: func(output mc.Output, t *testing.T) {
				if g, e := output.ContainerSettings.Container, mc.ContainerTypeM3u8; g != e {
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
			assertion: func(output mc.Output, t *testing.T) {
				if g, e := output.ContainerSettings.Container, mc.ContainerTypeCmfc; g != e {
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
			assertion: func(output mc.Output, t *testing.T) {
				colorSpaceConversion := output.VideoDescription.VideoPreprocessors.ColorCorrector.ColorSpaceConversion
				if g, e := colorSpaceConversion, mc.ColorSpaceConversionForceHdr10; g != e {
					t.Fatalf("got %q, expected %q", g, e)
				}

				wantHDR10Metadata := &mc.Hdr10Metadata{
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
		{},
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

		t.Run(tt.name, func(t *testing.T) {
			p := &driver{cfg: &config.MediaConvert{Destination: "s3://some_dest"}}

			input, err := p.createRequest(context.Background(), &db.Job{
				ID: "jobID", ProviderName: Name, SourceMedia: "s3://some/path.mp4",
				Outputs: []db.TranscodeOutput{{Preset: db.Preset{Name: defaultPreset.Name}, FileName: "file1.mp4"}},
			})
			if err != nil && !strings.Contains(err.Error(), tt.wantErrMsg) {
				t.Fatalf("driver.Transcode() error = %v, wantErr %q", err, tt.wantErrMsg)
			}
			if tt.assertion != nil {
				t.Logf("input: %#v", err)
				tt.assertion(input.Settings.OutputGroups[0].Outputs[0], t)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	vp8Preset := func(audioCodec string) db.Preset {
		return db.Preset{
			Name:        "preset_name",
			Description: "test_desc",
			Container:   "webm",
			RateControl: "VBR",
			Video: db.VideoPreset{
				Width:   300,
				Height:  400,
				Codec:   "vp8",
				Bitrate: 400000,
				GopSize: 120,
				GopUnit: "frames",
			},
			Audio: db.AudioPreset{
				Codec:         audioCodec,
				Bitrate:       96000,
				Normalization: true,
			},
		}
	}

	tests := []struct {
		cfg         *config.MediaConvert
		name        string
		job         *db.Job
		preset      db.Preset
		destination string
		wantJobReq  mc.CreateJobInput
		wantErr     bool
	}{
		{
			name: "H264/AAC/MP4",
			cfg: &config.MediaConvert{
				Role:            "some-role",
				DefaultQueueARN: "some:default:queue:arn",
			},
			job: &db.Job{
				ID:           "jobID",
				ProviderName: Name,
				SourceMedia:  "s3://some/path.mp4",
				Outputs:      []db.TranscodeOutput{{Preset: db.Preset{Name: defaultPreset.Name}, FileName: "file1.mp4"}},
				Labels:       []string{"bill:some-bu", "some-more-labels"},
			},
			preset:      defaultPreset,
			destination: "s3://some/destination",
			wantJobReq: mc.CreateJobInput{
				Role:  aws.String("some-role"),
				Queue: aws.String("some:default:queue:arn"),
				Tags: map[string]string{
					"bill:some-bu":     "true",
					"some-more-labels": "true",
				},
				Settings: &mc.JobSettings{
					Inputs: []mc.Input{
						{
							AudioSelectors: map[string]mc.AudioSelector{
								"Audio Selector 1": {
									DefaultSelection: mc.AudioDefaultSelectionDefault,
								},
							},
							FileInput: aws.String("s3://some/path.mp4"),
							VideoSelector: &mc.VideoSelector{
								ColorSpace: mc.ColorSpaceFollow,
							},
							TimecodeSource: mc.InputTimecodeSourceZerobased,
						},
					},
					OutputGroups: []mc.OutputGroup{
						{
							OutputGroupSettings: &mc.OutputGroupSettings{
								Type: mc.OutputGroupTypeFileGroupSettings,
								FileGroupSettings: &mc.FileGroupSettings{
									Destination: aws.String("s3://some/destination/jobID/m"),
								},
							},
							Outputs: []mc.Output{
								{
									NameModifier: aws.String("file1"),
									ContainerSettings: &mc.ContainerSettings{
										Container: mc.ContainerTypeMp4,
										Mp4Settings: &mc.Mp4Settings{
											Mp4MajorBrand: aws.String("isom"),
										},
									},
									VideoDescription: &mc.VideoDescription{
										Height:            aws.Int64(400),
										Width:             aws.Int64(300),
										RespondToAfd:      mc.RespondToAfdNone,
										ScalingBehavior:   mc.ScalingBehaviorDefault,
										TimecodeInsertion: mc.VideoTimecodeInsertionDisabled,
										AntiAlias:         mc.AntiAliasEnabled,
										VideoPreprocessors: &mc.VideoPreprocessor{
											Deinterlacer: &mc.Deinterlacer{
												Algorithm: mc.DeinterlaceAlgorithmInterpolate,
												Control:   mc.DeinterlacerControlNormal,
												Mode:      mc.DeinterlacerModeAdaptive,
											},
										},
										CodecSettings: &mc.VideoCodecSettings{
											Codec: mc.VideoCodecH264,
											H264Settings: &mc.H264Settings{
												Bitrate:            aws.Int64(400000),
												CodecLevel:         mc.H264CodecLevelAuto,
												CodecProfile:       mc.H264CodecProfileHigh,
												InterlaceMode:      mc.H264InterlaceModeProgressive,
												ParControl:         mc.H264ParControlSpecified,
												ParNumerator:       aws.Int64(1),
												ParDenominator:     aws.Int64(1),
												QualityTuningLevel: mc.H264QualityTuningLevelMultiPassHq,
												RateControlMode:    mc.H264RateControlModeVbr,
												GopSize:            aws.Float64(120),
												GopSizeUnits:       "FRAMES",
											},
										},
									},
									AudioDescriptions: []mc.AudioDescription{
										{
											CodecSettings: &mc.AudioCodecSettings{
												Codec: mc.AudioCodecAac,
												AacSettings: &mc.AacSettings{
													Bitrate:         aws.Int64(20000),
													CodecProfile:    mc.AacCodecProfileLc,
													CodingMode:      mc.AacCodingModeCodingMode20,
													RateControlMode: mc.AacRateControlModeCbr,
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
					TimecodeConfig: &mc.TimecodeConfig{
						Source: mc.TimecodeSourceZerobased,
					},
				},
			},
		},
		{
			name: "H264/AAC/MP4-Interlaced",
			job: &db.Job{
				ID:           "jobID",
				ProviderName: Name,
				SourceMedia:  "s3://some/path.mp4",
				SourceInfo: db.File{
					ScanType: db.ScanTypeInterlaced,
				},
				Outputs: []db.TranscodeOutput{{Preset: db.Preset{Name: defaultPreset.Name}, FileName: "file1.mp4"}},
			},
			preset:      defaultPreset,
			destination: "s3://some/destination",
			wantJobReq: mc.CreateJobInput{
				Role:  aws.String(""),
				Queue: aws.String(""),
				Tags:  map[string]string{},
				Settings: &mc.JobSettings{
					Inputs: []mc.Input{
						{
							AudioSelectors: map[string]mc.AudioSelector{
								"Audio Selector 1": {
									DefaultSelection: mc.AudioDefaultSelectionDefault,
								},
							},
							FileInput: aws.String("s3://some/path.mp4"),
							VideoSelector: &mc.VideoSelector{
								ColorSpace: mc.ColorSpaceFollow,
							},
							TimecodeSource: mc.InputTimecodeSourceZerobased,
						},
					},
					OutputGroups: []mc.OutputGroup{
						{
							OutputGroupSettings: &mc.OutputGroupSettings{
								Type: mc.OutputGroupTypeFileGroupSettings,
								FileGroupSettings: &mc.FileGroupSettings{
									Destination: aws.String("s3://some/destination/jobID/m"),
								},
							},
							Outputs: []mc.Output{
								{
									NameModifier: aws.String("file1"),
									ContainerSettings: &mc.ContainerSettings{
										Container: mc.ContainerTypeMp4,
										Mp4Settings: &mc.Mp4Settings{
											Mp4MajorBrand: aws.String("isom"),
										},
									},
									VideoDescription: &mc.VideoDescription{
										Height:            aws.Int64(400),
										Width:             aws.Int64(300),
										RespondToAfd:      mc.RespondToAfdNone,
										ScalingBehavior:   mc.ScalingBehaviorDefault,
										TimecodeInsertion: mc.VideoTimecodeInsertionDisabled,
										AntiAlias:         mc.AntiAliasEnabled,
										VideoPreprocessors: &mc.VideoPreprocessor{
											Deinterlacer: &mc.Deinterlacer{
												Algorithm: mc.DeinterlaceAlgorithmInterpolate,
												Control:   mc.DeinterlacerControlNormal,
												Mode:      mc.DeinterlacerModeDeinterlace,
											},
										},
										CodecSettings: &mc.VideoCodecSettings{
											Codec: mc.VideoCodecH264,
											H264Settings: &mc.H264Settings{
												Bitrate:            aws.Int64(400000),
												CodecLevel:         mc.H264CodecLevelAuto,
												CodecProfile:       mc.H264CodecProfileHigh,
												InterlaceMode:      mc.H264InterlaceModeProgressive,
												ParControl:         mc.H264ParControlSpecified,
												ParNumerator:       aws.Int64(1),
												ParDenominator:     aws.Int64(1),
												QualityTuningLevel: mc.H264QualityTuningLevelMultiPassHq,
												RateControlMode:    mc.H264RateControlModeVbr,
												GopSize:            aws.Float64(120),
												GopSizeUnits:       "FRAMES",
											},
										},
									},
									AudioDescriptions: []mc.AudioDescription{
										{
											CodecSettings: &mc.AudioCodecSettings{
												Codec: mc.AudioCodecAac,
												AacSettings: &mc.AacSettings{
													Bitrate:         aws.Int64(20000),
													CodecProfile:    mc.AacCodecProfileLc,
													CodingMode:      mc.AacCodingModeCodingMode20,
													RateControlMode: mc.AacRateControlModeCbr,
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
					TimecodeConfig: &mc.TimecodeConfig{
						Source: mc.TimecodeSourceZerobased,
					},
				},
			},
		},
		{
			name: "H264/AAC/MP4-Progressive",
			job: &db.Job{
				ID:           "jobID",
				ProviderName: Name,
				SourceMedia:  "s3://some/path.mp4",
				SourceInfo: db.File{
					ScanType: db.ScanTypeProgressive,
				},
				Outputs: []db.TranscodeOutput{{Preset: db.Preset{Name: defaultPreset.Name}, FileName: "file1.mp4"}},
			},
			preset:      defaultPreset,
			destination: "s3://some/destination",
			wantJobReq: mc.CreateJobInput{
				Role:  aws.String(""),
				Queue: aws.String(""),
				Tags:  map[string]string{},
				Settings: &mc.JobSettings{
					Inputs: []mc.Input{
						{
							AudioSelectors: map[string]mc.AudioSelector{
								"Audio Selector 1": {
									DefaultSelection: mc.AudioDefaultSelectionDefault,
								},
							},
							FileInput: aws.String("s3://some/path.mp4"),
							VideoSelector: &mc.VideoSelector{
								ColorSpace: mc.ColorSpaceFollow,
							},
							TimecodeSource: mc.InputTimecodeSourceZerobased,
						},
					},
					OutputGroups: []mc.OutputGroup{
						{
							OutputGroupSettings: &mc.OutputGroupSettings{
								Type: mc.OutputGroupTypeFileGroupSettings,
								FileGroupSettings: &mc.FileGroupSettings{
									Destination: aws.String("s3://some/destination/jobID/m"),
								},
							},
							Outputs: []mc.Output{
								{
									NameModifier: aws.String("file1"),
									ContainerSettings: &mc.ContainerSettings{
										Container: mc.ContainerTypeMp4,
										Mp4Settings: &mc.Mp4Settings{
											Mp4MajorBrand: aws.String("isom"),
										},
									},
									VideoDescription: &mc.VideoDescription{
										Height:             aws.Int64(400),
										Width:              aws.Int64(300),
										RespondToAfd:       mc.RespondToAfdNone,
										ScalingBehavior:    mc.ScalingBehaviorDefault,
										TimecodeInsertion:  mc.VideoTimecodeInsertionDisabled,
										AntiAlias:          mc.AntiAliasEnabled,
										VideoPreprocessors: &mc.VideoPreprocessor{},
										CodecSettings: &mc.VideoCodecSettings{
											Codec: mc.VideoCodecH264,
											H264Settings: &mc.H264Settings{
												Bitrate:            aws.Int64(400000),
												CodecLevel:         mc.H264CodecLevelAuto,
												CodecProfile:       mc.H264CodecProfileHigh,
												InterlaceMode:      mc.H264InterlaceModeProgressive,
												ParControl:         mc.H264ParControlSpecified,
												ParNumerator:       aws.Int64(1),
												ParDenominator:     aws.Int64(1),
												QualityTuningLevel: mc.H264QualityTuningLevelMultiPassHq,
												RateControlMode:    mc.H264RateControlModeVbr,
												GopSize:            aws.Float64(120),
												GopSizeUnits:       "FRAMES",
											},
										},
									},
									AudioDescriptions: []mc.AudioDescription{
										{
											CodecSettings: &mc.AudioCodecSettings{
												Codec: mc.AudioCodecAac,
												AacSettings: &mc.AacSettings{
													Bitrate:         aws.Int64(20000),
													CodecProfile:    mc.AacCodecProfileLc,
													CodingMode:      mc.AacCodingModeCodingMode20,
													RateControlMode: mc.AacRateControlModeCbr,
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
					TimecodeConfig: &mc.TimecodeConfig{
						Source: mc.TimecodeSourceZerobased,
					},
				},
			},
		},
		{
			name: "H265/MP4-VideoOnly",
			job: &db.Job{
				ID:           "jobID",
				ProviderName: Name,
				SourceMedia:  "s3://some/path.mp4",
				Outputs:      []db.TranscodeOutput{{Preset: db.Preset{Name: h265Preset.Name}, FileName: "file1.mp4"}},
			},
			preset:      h265Preset,
			destination: "s3://some/destination",
			wantJobReq: mc.CreateJobInput{
				Role:  aws.String(""),
				Queue: aws.String(""),
				Tags:  map[string]string{},
				Settings: &mc.JobSettings{
					Inputs: []mc.Input{
						{
							AudioSelectors: map[string]mc.AudioSelector{
								"Audio Selector 1": {
									DefaultSelection: mc.AudioDefaultSelectionDefault,
								},
							},
							FileInput: aws.String("s3://some/path.mp4"),
							VideoSelector: &mc.VideoSelector{
								ColorSpace: mc.ColorSpaceFollow,
							},
							TimecodeSource: mc.InputTimecodeSourceZerobased,
						},
					},
					OutputGroups: []mc.OutputGroup{
						{
							OutputGroupSettings: &mc.OutputGroupSettings{
								Type: mc.OutputGroupTypeFileGroupSettings,
								FileGroupSettings: &mc.FileGroupSettings{
									Destination: aws.String("s3://some/destination/jobID/m"),
								},
							},
							Outputs: []mc.Output{
								{
									NameModifier: aws.String("file1"),
									ContainerSettings: &mc.ContainerSettings{
										Container: mc.ContainerTypeMp4,
										Mp4Settings: &mc.Mp4Settings{
											Mp4MajorBrand: aws.String("isom"),
										},
									},
									VideoDescription: &mc.VideoDescription{
										Height:            aws.Int64(400),
										Width:             aws.Int64(300),
										RespondToAfd:      mc.RespondToAfdNone,
										ScalingBehavior:   mc.ScalingBehaviorDefault,
										TimecodeInsertion: mc.VideoTimecodeInsertionDisabled,
										AntiAlias:         mc.AntiAliasEnabled,
										VideoPreprocessors: &mc.VideoPreprocessor{
											Deinterlacer: &mc.Deinterlacer{
												Algorithm: mc.DeinterlaceAlgorithmInterpolate,
												Control:   mc.DeinterlacerControlNormal,
												Mode:      mc.DeinterlacerModeAdaptive,
											},
										},
										CodecSettings: &mc.VideoCodecSettings{
											Codec: mc.VideoCodecH265,
											H265Settings: &mc.H265Settings{
												Bitrate:                        aws.Int64(400000),
												GopSize:                        aws.Float64(120),
												GopSizeUnits:                   "FRAMES",
												CodecLevel:                     mc.H265CodecLevelAuto,
												CodecProfile:                   mc.H265CodecProfileMainMain,
												InterlaceMode:                  mc.H265InterlaceModeProgressive,
												ParControl:                     mc.H265ParControlSpecified,
												ParNumerator:                   aws.Int64(1),
												ParDenominator:                 aws.Int64(1),
												QualityTuningLevel:             mc.H265QualityTuningLevelSinglePassHq,
												RateControlMode:                mc.H265RateControlModeCbr,
												WriteMp4PackagingType:          mc.H265WriteMp4PackagingTypeHvc1,
												AlternateTransferFunctionSei:   mc.H265AlternateTransferFunctionSeiDisabled,
												SpatialAdaptiveQuantization:    mc.H265SpatialAdaptiveQuantizationEnabled,
												TemporalAdaptiveQuantization:   mc.H265TemporalAdaptiveQuantizationEnabled,
												FlickerAdaptiveQuantization:    mc.H265FlickerAdaptiveQuantizationEnabled,
												SceneChangeDetect:              mc.H265SceneChangeDetectEnabled,
												UnregisteredSeiTimecode:        mc.H265UnregisteredSeiTimecodeDisabled,
												SampleAdaptiveOffsetFilterMode: mc.H265SampleAdaptiveOffsetFilterModeAdaptive,
											},
										},
									},
									Extension: aws.String("mp4"),
								},
							},
						},
					},
					TimecodeConfig: &mc.TimecodeConfig{
						Source: mc.TimecodeSourceZerobased,
					},
				},
			},
		},
		{
			name: "AV1/MP4",
			job: &db.Job{
				ID:           "jobID",
				ProviderName: Name,
				SourceMedia:  "s3://some/path.mp4",
				Outputs:      []db.TranscodeOutput{{Preset: db.Preset{Name: av1Preset.Name}, FileName: "file1.mp4"}},
			},
			preset:      av1Preset,
			destination: "s3://some/destination",
			wantJobReq: mc.CreateJobInput{
				Role:  aws.String(""),
				Queue: aws.String(""),
				Tags:  map[string]string{},
				Settings: &mc.JobSettings{
					Inputs: []mc.Input{
						{
							AudioSelectors: map[string]mc.AudioSelector{
								"Audio Selector 1": {
									DefaultSelection: mc.AudioDefaultSelectionDefault,
								},
							},
							FileInput: aws.String("s3://some/path.mp4"),
							VideoSelector: &mc.VideoSelector{
								ColorSpace: mc.ColorSpaceFollow,
							},
							TimecodeSource: mc.InputTimecodeSourceZerobased,
						},
					},
					OutputGroups: []mc.OutputGroup{
						{
							OutputGroupSettings: &mc.OutputGroupSettings{
								Type: mc.OutputGroupTypeFileGroupSettings,
								FileGroupSettings: &mc.FileGroupSettings{
									Destination: aws.String("s3://some/destination/jobID/m"),
								},
							},
							Outputs: []mc.Output{
								{
									NameModifier: aws.String("file1"),
									ContainerSettings: &mc.ContainerSettings{
										Container: mc.ContainerTypeMp4,
										Mp4Settings: &mc.Mp4Settings{
											Mp4MajorBrand: aws.String("isom"),
										},
									},
									VideoDescription: &mc.VideoDescription{
										Height:            aws.Int64(400),
										Width:             aws.Int64(300),
										RespondToAfd:      mc.RespondToAfdNone,
										ScalingBehavior:   mc.ScalingBehaviorDefault,
										TimecodeInsertion: mc.VideoTimecodeInsertionDisabled,
										AntiAlias:         mc.AntiAliasEnabled,
										VideoPreprocessors: &mc.VideoPreprocessor{
											Deinterlacer: &mc.Deinterlacer{
												Algorithm: mc.DeinterlaceAlgorithmInterpolate,
												Control:   mc.DeinterlacerControlNormal,
												Mode:      mc.DeinterlacerModeAdaptive,
											},
										},
										CodecSettings: &mc.VideoCodecSettings{
											Codec: mc.VideoCodecAv1,
											Av1Settings: &mc.Av1Settings{
												MaxBitrate: aws.Int64(400000),
												GopSize:    aws.Float64(120),
												QvbrSettings: &mc.Av1QvbrSettings{
													QvbrQualityLevel:         aws.Int64(7),
													QvbrQualityLevelFineTune: aws.Float64(0),
												},
												RateControlMode: mc.Av1RateControlModeQvbr,
											},
										},
									},
									Extension: aws.String("mp4"),
								},
							},
						},
					},
					TimecodeConfig: &mc.TimecodeConfig{
						Source: mc.TimecodeSourceZerobased,
					},
				},
			},
		},
		{
			name: "VP8/Vorbis/Webm",
			job: &db.Job{
				ID:           "jobID",
				ProviderName: Name,
				SourceMedia:  "s3://some/path.webm",
				Outputs:      []db.TranscodeOutput{{Preset: db.Preset{Name: defaultPreset.Name}, FileName: "file1.webm"}},
			},
			preset:      vp8Preset("vorbis"),
			destination: "s3://some/destination",
			wantJobReq: mc.CreateJobInput{
				Role:  aws.String(""),
				Queue: aws.String(""),
				Tags:  map[string]string{},
				Settings: &mc.JobSettings{
					Inputs: []mc.Input{
						{
							AudioSelectors: map[string]mc.AudioSelector{
								"Audio Selector 1": {
									DefaultSelection: mc.AudioDefaultSelectionDefault,
								},
							},
							FileInput: aws.String("s3://some/path.webm"),
							VideoSelector: &mc.VideoSelector{
								ColorSpace: mc.ColorSpaceFollow,
							},
							TimecodeSource: mc.InputTimecodeSourceZerobased,
						},
					},
					OutputGroups: []mc.OutputGroup{
						{
							OutputGroupSettings: &mc.OutputGroupSettings{
								Type: mc.OutputGroupTypeFileGroupSettings,
								FileGroupSettings: &mc.FileGroupSettings{
									Destination: aws.String("s3://some/destination/jobID/m"),
								},
							},
							Outputs: []mc.Output{
								{
									NameModifier: aws.String("file1"),
									ContainerSettings: &mc.ContainerSettings{
										Container: mc.ContainerTypeWebm,
									},
									VideoDescription: &mc.VideoDescription{
										Height:            aws.Int64(400),
										Width:             aws.Int64(300),
										RespondToAfd:      mc.RespondToAfdNone,
										ScalingBehavior:   mc.ScalingBehaviorDefault,
										TimecodeInsertion: mc.VideoTimecodeInsertionDisabled,
										AntiAlias:         mc.AntiAliasEnabled,
										VideoPreprocessors: &mc.VideoPreprocessor{
											Deinterlacer: &mc.Deinterlacer{
												Algorithm: mc.DeinterlaceAlgorithmInterpolate,
												Control:   mc.DeinterlacerControlNormal,
												Mode:      mc.DeinterlacerModeAdaptive,
											},
										},
										CodecSettings: &mc.VideoCodecSettings{
											Codec: mc.VideoCodecVp8,
											Vp8Settings: &mc.Vp8Settings{
												Bitrate:          aws.Int64(400000),
												GopSize:          aws.Float64(120),
												RateControlMode:  mc.Vp8RateControlModeVbr,
												FramerateControl: mc.Vp8FramerateControlInitializeFromSource,
												ParControl:       mc.Vp8ParControlSpecified,
												ParNumerator:     aws.Int64(1),
												ParDenominator:   aws.Int64(1),
											},
										},
									},
									AudioDescriptions: []mc.AudioDescription{
										{
											AudioNormalizationSettings: &mc.AudioNormalizationSettings{
												Algorithm:        mc.AudioNormalizationAlgorithmItuBs17703,
												AlgorithmControl: mc.AudioNormalizationAlgorithmControlCorrectAudio,
											},
											CodecSettings: &mc.AudioCodecSettings{
												Codec: mc.AudioCodecVorbis,
												VorbisSettings: &mc.VorbisSettings{
													Channels:   aws.Int64(2),
													SampleRate: aws.Int64(defaultAudioSampleRate),
													VbrQuality: aws.Int64(2),
												},
											},
										},
									},
									Extension: aws.String("webm"),
								},
							},
						},
					},
					TimecodeConfig: &mc.TimecodeConfig{
						Source: mc.TimecodeSourceZerobased,
					},
				},
			},
		},
		{
			name: "VP8/Opus/Webm",
			job: &db.Job{
				ID:           "jobID",
				ProviderName: Name,
				SourceMedia:  "s3://some/path.webm",
				Outputs:      []db.TranscodeOutput{{Preset: db.Preset{Name: defaultPreset.Name}, FileName: "file1.webm"}},
			},
			preset:      vp8Preset("opus"),
			destination: "s3://some/destination",
			wantJobReq: mc.CreateJobInput{
				Role:  aws.String(""),
				Queue: aws.String(""),
				Tags:  map[string]string{},
				Settings: &mc.JobSettings{
					Inputs: []mc.Input{
						{
							AudioSelectors: map[string]mc.AudioSelector{
								"Audio Selector 1": {
									DefaultSelection: mc.AudioDefaultSelectionDefault,
								},
							},
							FileInput: aws.String("s3://some/path.webm"),
							VideoSelector: &mc.VideoSelector{
								ColorSpace: mc.ColorSpaceFollow,
							},
							TimecodeSource: mc.InputTimecodeSourceZerobased,
						},
					},
					OutputGroups: []mc.OutputGroup{
						{
							OutputGroupSettings: &mc.OutputGroupSettings{
								Type: mc.OutputGroupTypeFileGroupSettings,
								FileGroupSettings: &mc.FileGroupSettings{
									Destination: aws.String("s3://some/destination/jobID/m"),
								},
							},
							Outputs: []mc.Output{
								{
									NameModifier: aws.String("file1"),
									ContainerSettings: &mc.ContainerSettings{
										Container: mc.ContainerTypeWebm,
									},
									VideoDescription: &mc.VideoDescription{
										Height:            aws.Int64(400),
										Width:             aws.Int64(300),
										RespondToAfd:      mc.RespondToAfdNone,
										ScalingBehavior:   mc.ScalingBehaviorDefault,
										TimecodeInsertion: mc.VideoTimecodeInsertionDisabled,
										AntiAlias:         mc.AntiAliasEnabled,
										VideoPreprocessors: &mc.VideoPreprocessor{
											Deinterlacer: &mc.Deinterlacer{
												Algorithm: mc.DeinterlaceAlgorithmInterpolate,
												Control:   mc.DeinterlacerControlNormal,
												Mode:      mc.DeinterlacerModeAdaptive,
											},
										},
										CodecSettings: &mc.VideoCodecSettings{
											Codec: mc.VideoCodecVp8,
											Vp8Settings: &mc.Vp8Settings{
												Bitrate:          aws.Int64(400000),
												GopSize:          aws.Float64(120),
												RateControlMode:  mc.Vp8RateControlModeVbr,
												FramerateControl: mc.Vp8FramerateControlInitializeFromSource,
												ParControl:       mc.Vp8ParControlSpecified,
												ParNumerator:     aws.Int64(1),
												ParDenominator:   aws.Int64(1),
											},
										},
									},
									AudioDescriptions: []mc.AudioDescription{
										{
											AudioNormalizationSettings: &mc.AudioNormalizationSettings{
												Algorithm:        mc.AudioNormalizationAlgorithmItuBs17703,
												AlgorithmControl: mc.AudioNormalizationAlgorithmControlCorrectAudio,
											},
											CodecSettings: &mc.AudioCodecSettings{
												Codec: mc.AudioCodecOpus,
												OpusSettings: &mc.OpusSettings{
													Channels:   aws.Int64(2),
													Bitrate:    aws.Int64(96000),
													SampleRate: aws.Int64(defaultAudioSampleRate),
												},
											},
										},
									},
									Extension: aws.String("webm"),
								},
							},
						},
					},
					TimecodeConfig: &mc.TimecodeConfig{
						Source: mc.TimecodeSourceZerobased,
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
				Outputs:      []db.TranscodeOutput{{Preset: db.Preset{Name: defaultPreset.Name}, FileName: "file1.mp4"}},
			},
			preset:      defaultPreset,
			destination: "s3://some/destination",
			wantJobReq: mc.CreateJobInput{
				Role:            aws.String(""),
				Queue:           aws.String("some:preferred:queue:arn"),
				HopDestinations: []mc.HopDestination{{WaitMinutes: aws.Int64(defaultQueueHopTimeoutMins)}},
				Tags:            map[string]string{},
				Settings: &mc.JobSettings{
					Inputs: []mc.Input{
						{
							AudioSelectors: map[string]mc.AudioSelector{
								"Audio Selector 1": {
									DefaultSelection: mc.AudioDefaultSelectionDefault,
								},
							},
							FileInput: aws.String("s3://some/path.mp4"),
							VideoSelector: &mc.VideoSelector{
								ColorSpace: mc.ColorSpaceFollow,
							},
							TimecodeSource: mc.InputTimecodeSourceZerobased,
						},
					},
					OutputGroups: []mc.OutputGroup{
						{
							OutputGroupSettings: &mc.OutputGroupSettings{
								Type: mc.OutputGroupTypeFileGroupSettings,
								FileGroupSettings: &mc.FileGroupSettings{
									Destination: aws.String("s3://some/destination/jobID/m"),
								},
							},
							Outputs: []mc.Output{
								{
									NameModifier: aws.String("file1"),
									ContainerSettings: &mc.ContainerSettings{
										Container: mc.ContainerTypeMp4,
										Mp4Settings: &mc.Mp4Settings{
											Mp4MajorBrand: aws.String("isom"),
										},
									},
									VideoDescription: &mc.VideoDescription{
										Height:            aws.Int64(400),
										Width:             aws.Int64(300),
										RespondToAfd:      mc.RespondToAfdNone,
										ScalingBehavior:   mc.ScalingBehaviorDefault,
										TimecodeInsertion: mc.VideoTimecodeInsertionDisabled,
										AntiAlias:         mc.AntiAliasEnabled,
										VideoPreprocessors: &mc.VideoPreprocessor{
											Deinterlacer: &mc.Deinterlacer{
												Algorithm: mc.DeinterlaceAlgorithmInterpolate,
												Control:   mc.DeinterlacerControlNormal,
												Mode:      mc.DeinterlacerModeAdaptive,
											},
										},
										CodecSettings: &mc.VideoCodecSettings{
											Codec: mc.VideoCodecH264,
											H264Settings: &mc.H264Settings{
												Bitrate:            aws.Int64(400000),
												CodecLevel:         mc.H264CodecLevelAuto,
												CodecProfile:       mc.H264CodecProfileHigh,
												InterlaceMode:      mc.H264InterlaceModeProgressive,
												QualityTuningLevel: mc.H264QualityTuningLevelMultiPassHq,
												RateControlMode:    mc.H264RateControlModeVbr,
												GopSize:            aws.Float64(120),
												GopSizeUnits:       mc.H264GopSizeUnitsFrames,
												ParControl:         mc.H264ParControlSpecified,
												ParNumerator:       aws.Int64(1),
												ParDenominator:     aws.Int64(1),
											},
										},
									},
									AudioDescriptions: []mc.AudioDescription{
										{
											CodecSettings: &mc.AudioCodecSettings{
												Codec: mc.AudioCodecAac,
												AacSettings: &mc.AacSettings{
													Bitrate:         aws.Int64(20000),
													CodecProfile:    mc.AacCodecProfileLc,
													CodingMode:      mc.AacCodingModeCodingMode20,
													RateControlMode: mc.AacRateControlModeCbr,
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
					TimecodeConfig: &mc.TimecodeConfig{
						Source: mc.TimecodeSourceZerobased,
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
				Outputs:      []db.TranscodeOutput{{Preset: db.Preset{Name: tcBurninPreset.Name}, FileName: "file1.mov"}},
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
			wantJobReq: mc.CreateJobInput{
				Role:            aws.String(""),
				Queue:           aws.String("some:preferred:queue:arn"),
				HopDestinations: []mc.HopDestination{{WaitMinutes: aws.Int64(defaultQueueHopTimeoutMins)}},
				Tags:            map[string]string{},
				Settings: &mc.JobSettings{
					Inputs: []mc.Input{
						{
							AudioSelectors: map[string]mc.AudioSelector{
								"Audio Selector 1": getAudioSelector(6, 2, []int64{1}, []mc.OutputChannelMapping{
									{InputChannels: []int64{0, -60, 0, -60, 0, -60}},
									{InputChannels: []int64{-60, 0, 0, -60, -60, 0}},
								}),
							},
							FileInput: aws.String("s3://some/path.mov"),
							VideoSelector: &mc.VideoSelector{
								ColorSpace: mc.ColorSpaceFollow,
							},
							TimecodeSource: mc.InputTimecodeSourceZerobased,
						},
					},
					OutputGroups: []mc.OutputGroup{
						{
							OutputGroupSettings: &mc.OutputGroupSettings{
								Type: mc.OutputGroupTypeFileGroupSettings,
								FileGroupSettings: &mc.FileGroupSettings{
									Destination: aws.String("s3://some/destination/jobID/m"),
								},
							},
							Outputs: []mc.Output{
								{
									NameModifier: aws.String("file1"),
									ContainerSettings: &mc.ContainerSettings{
										Container: mc.ContainerTypeMov,
										MovSettings: &mc.MovSettings{
											ClapAtom:           mc.MovClapAtomExclude,
											CslgAtom:           mc.MovCslgAtomInclude,
											PaddingControl:     mc.MovPaddingControlOmneon,
											Reference:          mc.MovReferenceSelfContained,
											Mpeg2FourCCControl: mc.MovMpeg2FourCCControlMpeg,
										},
									},
									VideoDescription: &mc.VideoDescription{
										Height:            aws.Int64(400),
										Width:             aws.Int64(300),
										RespondToAfd:      mc.RespondToAfdNone,
										ScalingBehavior:   mc.ScalingBehaviorDefault,
										TimecodeInsertion: mc.VideoTimecodeInsertionDisabled,
										AntiAlias:         mc.AntiAliasEnabled,
										VideoPreprocessors: &mc.VideoPreprocessor{
											Deinterlacer: &mc.Deinterlacer{
												Algorithm: mc.DeinterlaceAlgorithmInterpolate,
												Control:   mc.DeinterlacerControlNormal,
												Mode:      mc.DeinterlacerModeAdaptive,
											},
											TimecodeBurnin: &mc.TimecodeBurnin{
												FontSize: aws.Int64(12),
												Position: mc.TimecodeBurninPositionBottomLeft,
												Prefix:   aws.String(""),
											},
										},
										CodecSettings: &mc.VideoCodecSettings{
											Codec: mc.VideoCodecH264,
											H264Settings: &mc.H264Settings{
												Bitrate:            aws.Int64(400000),
												CodecLevel:         mc.H264CodecLevelAuto,
												CodecProfile:       mc.H264CodecProfileHigh,
												InterlaceMode:      mc.H264InterlaceModeProgressive,
												QualityTuningLevel: mc.H264QualityTuningLevelMultiPassHq,
												RateControlMode:    mc.H264RateControlModeVbr,
												GopSize:            aws.Float64(120),
												GopSizeUnits:       mc.H264GopSizeUnitsFrames,
												ParControl:         mc.H264ParControlSpecified,
												ParNumerator:       aws.Int64(1),
												ParDenominator:     aws.Int64(1),
											},
										},
									},
									AudioDescriptions: []mc.AudioDescription{
										{
											CodecSettings: &mc.AudioCodecSettings{
												Codec: mc.AudioCodecAac,
												AacSettings: &mc.AacSettings{
													Bitrate:         aws.Int64(20000),
													CodecProfile:    mc.AacCodecProfileLc,
													CodingMode:      mc.AacCodingModeCodingMode20,
													RateControlMode: mc.AacRateControlModeCbr,
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
					TimecodeConfig: &mc.TimecodeConfig{
						Source: mc.TimecodeSourceZerobased,
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
		//		Outputs:      []db.TranscodeOutput{{Preset: db.Preset{Name: audioOnlyPreset.Name}, FileName: "file1.mp4"}},
		//	},
		//	preset:      audioOnlyPreset,
		//	destination: "s3://some/destination",
		//	wantJobReq: mc.CreateJobInput{
		//		AccelerationSettings: &mc.AccelerationSettings{
		//			Mode: mc.AccelerationModePreferred,
		//		},
		//		Role:  aws.String(""),
		//		Queue: aws.String("some:default:queue:arn"),
		//		Tags: map[string]string{},
		//		Settings: &mc.JobSettings{
		//			Inputs: []mc.Input{
		//				{
		//					AudioSelectors: map[string]mc.AudioSelector{
		//						"Audio Selector 1": {
		//							DefaultSelection: mc.AudioDefaultSelectionDefault,
		//						},
		//					},
		//					FileInput: aws.String("s3://some/path.mp4"),
		//					VideoSelector: &mc.VideoSelector{
		//						ColorSpace: mc.ColorSpaceFollow,
		//					},
		//					TimecodeSource: mc.InputTimecodeSourceZerobased,
		//				},
		//			},
		//			OutputGroups: []mc.OutputGroup{
		//				{
		//					OutputGroupSettings: &mc.OutputGroupSettings{
		//						Type: mc.OutputGroupTypeFileGroupSettings,
		//						FileGroupSettings: &mc.FileGroupSettings{
		//							Destination: aws.String("s3://some/destination/jobID/m"),
		//						},
		//					},
		//					Outputs: []mc.Output{
		//						{
		//							NameModifier: aws.String("file1"),
		//							ContainerSettings: &mc.ContainerSettings{
		//								Container: mc.ContainerTypeMp4,
		//								Mp4Settings: &mc.Mp4Settings{
		//									Mp4MajorBrand: aws.String("isom"),
		//								},
		//							},
		//							AudioDescriptions: []mc.AudioDescription{
		//								{
		//									CodecSettings: &mc.AudioCodecSettings{
		//										Codec: mc.AudioCodecAac,
		//										AacSettings: &mc.AacSettings{
		//											Bitrate:         aws.Int64(20000),
		//											CodecProfile:    mc.AacCodecProfileLc,
		//											CodingMode:      mc.AacCodingModeCodingMode20,
		//											RateControlMode: mc.AacRateControlModeCbr,
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
		//			TimecodeConfig: &mc.TimecodeConfig{
		//				Source: mc.TimecodeSourceZerobased,
		//			},
		//		},
		//	},
		//},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			if tt.cfg == nil {
				tt.cfg = &config.MediaConvert{Destination: tt.destination}
			} else {
				tt.cfg.Destination = tt.destination
			}

			p := &driver{
				cfg: tt.cfg,
			}
			input, err := p.createRequest(context.Background(), tt.job)
			if (err != nil) != tt.wantErr {
				t.Fatalf("driver.Transcode() error = %v, wantErr %v", err, tt.wantErr)
			}

			if g, e := input, tt.wantJobReq; !reflect.DeepEqual(g, e) {
				t.Fatalf("Transcode(): wrong job request\nWant %+v\nGot %+v\nDiff %s", e,
					g, cmp.Diff(e, g))
			}
		})
	}
}

func Test_driver_JobStatus(t *testing.T) {
	tests := []struct {
		name        string
		destination string
		mcJob       mc.Job
		wantStatus  provider.Status
		wantErr     bool
	}{
		{
			name:        "a job that has been queued returns the correct status",
			destination: "s3://some/destination",
			mcJob: mc.Job{
				Status: mc.JobStatusSubmitted,
			},
			wantStatus: provider.Status{
				State:        provider.StateQueued,
				ProviderName: Name,
				Output: provider.Output{
					Destination: "s3://some/destination/jobID/",
				},
			},
		},
		{
			name:        "a job that is currently transcoding returns the correct status",
			destination: "s3://some/destination",
			mcJob: mc.Job{
				Status:             mc.JobStatusProgressing,
				JobPercentComplete: aws.Int64(42),
			},
			wantStatus: provider.Status{
				State:        provider.StateStarted,
				ProviderName: Name,
				Progress:     42,
				Output: provider.Output{
					Destination: "s3://some/destination/jobID/",
				},
			},
		},
		{
			name:        "a job that has finished transcoding returns the correct status",
			destination: "s3://some/destination",
			mcJob: mc.Job{
				Status: mc.JobStatusComplete,
				Settings: &mc.JobSettings{
					OutputGroups: []mc.OutputGroup{
						{
							OutputGroupSettings: &mc.OutputGroupSettings{
								Type: mc.OutputGroupTypeFileGroupSettings,
								FileGroupSettings: &mc.FileGroupSettings{
									Destination: aws.String("s3://some/destination/jobID/m"),
								},
							},
							Outputs: []mc.Output{
								{
									NameModifier: aws.String("_modifier"),
									VideoDescription: &mc.VideoDescription{
										Height: aws.Int64(102),
										Width:  aws.Int64(324),
									},
									ContainerSettings: &mc.ContainerSettings{
										Container: mc.ContainerTypeMp4,
										Mp4Settings: &mc.Mp4Settings{
											Mp4MajorBrand: aws.String("isom"),
										},
									},
								},
								{
									NameModifier: aws.String("_another_modifier"),
									ContainerSettings: &mc.ContainerSettings{
										Container: mc.ContainerTypeMp4,
										Mp4Settings: &mc.Mp4Settings{
											Mp4MajorBrand: aws.String("isom"),
										},
									},
								},
								{
									NameModifier: aws.String("_123"),
									ContainerSettings: &mc.ContainerSettings{
										Container: mc.ContainerTypeM2ts,
									},
								},
								{
									ContainerSettings: &mc.ContainerSettings{
										Container: mc.ContainerTypeM2ts,
									},
								},
							},
						},
					},
				},
			},
			wantStatus: provider.Status{
				State:        provider.StateFinished,
				ProviderName: Name,
				Progress:     100,
				Output: provider.Output{
					Destination: "s3://some/destination/jobID/",
					Files: []provider.File{
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

		t.Run(tt.name, func(t *testing.T) {
			p := &driver{cfg: &config.MediaConvert{
				Destination: tt.destination,
			}}

			status := p.status(&defaultJob, &tt.mcJob)
			if g, e := status, &tt.wantStatus; !reflect.DeepEqual(g, e) {
				t.Fatalf("driver.JobStatus(): wrong job request\nWant %+v\nGot %+v\nDiff %s", e,
					g, cmp.Diff(e, g))
			}
		})
	}
}
