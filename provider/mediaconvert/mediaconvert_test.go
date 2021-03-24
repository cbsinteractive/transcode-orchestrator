package mediaconvert

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"

	mc "github.com/aws/aws-sdk-go-v2/service/mediaconvert"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/cbsinteractive/transcode-orchestrator/av"
	"github.com/cbsinteractive/transcode-orchestrator/config"
	"github.com/google/go-cmp/cmp"
)

var (
	defaultPreset = av.File{
		Name:      "file1.mp4",
		Container: "mp4",
		Video: av.Video{
			Codec: "h264", Profile: "high", Level: "4.1",
			Width: 300, Height: 400, Scantype: "progressive",
			Bitrate: av.Bitrate{BPS: 400000, Control: "VBR"},
			TwoPass: true,
			Gop:     av.Gop{Size: 120, Unit: "frames"},
		},
		Audio: av.Audio{Codec: "aac", Bitrate: 20000},
	}

	h265Preset = av.File{
		Name:      "file1.mp4",
		Container: "mp4",
		Video: av.Video{
			Codec: "h265",
			Width: 300, Height: 400, Scantype: "progressive",
			Bitrate: av.Bitrate{BPS: 400000, Control: "CBR"},
			Gop:     av.Gop{Size: 120, Unit: "frames"},
		},
	}

	av1Preset = av.File{
		Name:      "file1.mp4",
		Container: "mp4",
		Video: av.Video{
			Codec: "av1",
			Width: 300, Height: 400, Scantype: "progressive",
			Bitrate: av.Bitrate{BPS: 400000},
			Gop:     av.Gop{Size: 120, Unit: "frames"},
		},
	}

	audioOnlyPreset = av.File{
		Name:      "file1.mp4",
		Container: "mp4",
		Audio:     av.Audio{Codec: "aac", Bitrate: 20000},
	}

	defaultJob = av.Job{
		ID: "jobID", Provider: Name,
		Input: av.File{Name: "s3://some/path.mp4"},
		Output: av.Dir{File: []av.File{
			{Name: "file1.mp4"},
			{Name: "file2.mp4"},
		}},
	}
)

func TestHDR10(t *testing.T) {
	i := aws.Int64
	display := "G(8500,39850)B(6550,2300)R(35400,14600)WP(15635,16450)L(100000000000,0)"
	want := &mc.Hdr10Metadata{
		GreenPrimaryX: i(8500), GreenPrimaryY: i(39850),
		BluePrimaryX: i(6550), BluePrimaryY: i(2300),
		RedPrimaryX: i(35400), RedPrimaryY: i(14600),
		WhitePointX: i(15635), WhitePointY: i(16450),
		MinLuminance: i(0), MaxLuminance: i(100000000000),
		MaxContentLightLevel:      i(10000),
		MaxFrameAverageLightLevel: i(400),
	}

	p := defaultPreset
	p.Video.HDR10.Enabled = true
	p.Video.HDR10.MaxCLL = 10000
	p.Video.HDR10.MaxFALL = 400
	p.Video.HDR10.MasterDisplay = display

	d := &driver{cfg: config.MediaConvert{Destination: "s3://some_dest"}}
	req, err := d.createRequest(nil, &av.Job{
		Output: av.Dir{File: []av.File{p}},
	})
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	cc := req.Settings.OutputGroups[0].Outputs[0].VideoDescription.VideoPreprocessors.ColorCorrector
	if g, e := cc.ColorSpaceConversion, mc.ColorSpaceConversionForceHdr10; g != e {
		t.Fatalf("force hdr10: have %v, want %v", g, e)
	}

	if g, e := cc.Hdr10Metadata, want; !reflect.DeepEqual(g, e) {
		t.Fatalf("metadata: have %v, want %v", g, e)
	}
}

func TestSupport(t *testing.T) {
	d := &driver{}
	run := func(p av.File) (*mc.CreateJobInput, error) {

		return d.createRequest(nil, &av.Job{
			Output: av.Dir{File: []av.File{p}},
		})
	}
	warn := false
	ck := func(ctx string, err error, want ...error) {
		t.Helper()
		w := error(nil)
		if len(want) > 0 {
			w = want[0]
		}
		if !errors.Is(err, w) {
			if !warn {
				t.Fatalf("%s: %v (want %v)", ctx, err, w)
			}
			t.Logf("%s: %v (want %v)", ctx, err, w)
		}
	}

	t.Run("Container", func(t *testing.T) {
		for _, box := range []string{"mp4", "mxf", "mov", "webm"} {
			mc, err := run(av.File{Container: box})
			ck(box, err)
			have := string(mc.Settings.OutputGroups[0].Outputs[0].ContainerSettings.Container)
			have = strings.ToLower(have)
			if have != box {
				t.Logf("container type mismatch: %v != %v", have, box)
			}
		}

		for _, box := range []string{"mp4", "mxf", "mov", "webm"} {
			for _, codec := range []string{"", "h264", "h265", "av1", "xdcam", "vp8"} {
				_, err := run(av.File{Container: box, Video: av.Video{Codec: codec}})
				ck(box+"/"+codec, err)
			}
		}
	})

	t.Run("Audio", func(t *testing.T) {
		for i, tt := range []struct {
			p   av.Audio
			err error
		}{
			{av.Audio{Codec: "aac"}, nil},
			{av.Audio{Codec: "aad"}, ErrUnsupported},
		} {
			_, err := run(av.File{Container: "mp4", Audio: tt.p})
			ck(fmt.Sprintf("%d", i), err, tt.err)
		}
	})

	t.Run("Video", func(t *testing.T) {
		for i, tt := range []struct {
			p   av.Video
			err error
		}{
			{av.Video{Codec: "vp9001"}, ErrUnsupported},
			{av.Video{Codec: "h264", Profile: "8000"}, ErrUnsupported},
			{av.Video{Codec: "h264", Profile: "main"}, nil},
			{av.Video{Codec: "h264", Profile: "high"}, nil},
			{av.Video{Codec: "h264", Profile: "main", Level: "1812"}, nil},
			{av.Video{Codec: "h264", Profile: "main", Level: "@@@@"}, nil},
			{av.Video{Codec: "h265", Profile: "main"}, nil},
			{av.Video{Codec: "h265", Profile: "main", Level: "1812"}, ErrUnsupported},
			{av.Video{Codec: "h265", Profile: "main", Level: "@@@@"}, ErrUnsupported},

			// Below: flaky tests or behavior
			// NOTE(as): we seem to have special logic for this because of HDR support
			//		{av.Video{Codec: "h265", Profile: "9000"}, ErrUnsupported},
			{av.Video{Codec: "h264", Scantype: "efas"}, nil},
			{av.Video{Codec: "h265", Scantype: "?"}, nil},
			{av.Video{Codec: "av1", Profile: "f"}, nil},
		} {
			_, err := run(av.File{Container: "mp4", Video: tt.p})
			ck(fmt.Sprintf("video%d", i), err, tt.err)
		}
	})

	t.Run("FlakyValidation", func(t *testing.T) {
		warn = true
		for _, codec := range []string{"h264", "h265", "av1", "xdcam", "vp8"} {
			_, err := run(av.File{Container: "mp4", Video: av.Video{Codec: codec, Bitrate: av.Bitrate{Control: "?"}}})
			ck(codec+"/ratecontrol", err, ErrUnsupported)

			_, err = run(av.File{Container: "mp4", Video: av.Video{Codec: codec, Profile: "?"}})
			ck(codec+"/profile", err, ErrUnsupported)

			_, err = run(av.File{Container: "mp4", Video: av.Video{Codec: codec, Level: "?"}})
			ck(codec+"/profilelevel", err, ErrUnsupported)

			_, err = run(av.File{Container: "mp4", Video: av.Video{Codec: codec, Scantype: "?"}})
			ck(codec+"/interlace", err, ErrUnsupported)
		}
	})

}

func TestDriverCreate(t *testing.T) {
	vp8Preset := func(audioCodec string) av.File {
		return av.File{
			Name: "file1.webm",
			Video: av.Video{
				Codec: "vp8",
				Width: 300, Height: 400, Scantype: "progressive",
				Bitrate: av.Bitrate{BPS: 400000, Control: "VBR"},
				Gop:     av.Gop{Size: 120, Unit: "frames"},
			},
			Audio: av.Audio{
				Codec:     audioCodec,
				Bitrate:   96000,
				Normalize: true,
			},
		}
	}

	tests := []struct {
		cfg         config.MediaConvert
		name        string
		job         *av.Job
		preset      av.File
		destination string
		wantJobReq  mc.CreateJobInput
		wantErr     bool
	}{
		{
			name: "H264/AAC/MP4",
			cfg:  config.MediaConvert{Role: "some-role", DefaultQueueARN: "some:default:queue:arn"},
			job: &av.Job{
				ID: "jobID", Provider: Name,
				Input:  av.File{Name: "s3://some/path.mp4"},
				Output: av.Dir{File: []av.File{{Name: "file1.mp4"}}},
				Labels: []string{"bill:some-bu", "some-more-labels"},
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
			job: &av.Job{
				ID:       "jobID",
				Provider: Name,
				Input:    av.File{Name: "s3://some/path.mp4", Video: av.Video{Scantype: av.ScanInterlaced}},
				Output:   av.Dir{File: []av.File{{Name: "file1.mp4"}}},
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
			job: &av.Job{
				ID:       "jobID",
				Provider: Name,
				Input:    av.File{Name: "s3://some/path.mp4", Video: av.Video{Scantype: av.ScanProgressive}},
				Output:   av.Dir{File: []av.File{{Name: "file1.mp4"}}},
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
			job: &av.Job{
				ID:       "jobID",
				Provider: Name,
				Input:    av.File{Name: "s3://some/path.mp4"},
				Output:   av.Dir{File: []av.File{{Name: "file1.mp4"}}},
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
			job: &av.Job{
				ID:       "jobID",
				Provider: Name,
				Input:    av.File{Name: "s3://some/path.mp4"},
				Output:   av.Dir{File: []av.File{{Name: "file1.mp4"}}},
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
			job: &av.Job{
				ID:       "jobID",
				Provider: Name,
				Input:    av.File{Name: "s3://some/path.webm"},
				Output:   av.Dir{File: []av.File{{Name: "file1.webm"}}},
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
			job: &av.Job{
				ID:       "jobID",
				Provider: Name,
				Input:    av.File{Name: "s3://some/path.webm"},
				Output:   av.Dir{File: []av.File{{Name: "file1.webm"}}},
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
			name: "MapPreferredQueue",
			cfg: config.MediaConvert{
				DefaultQueueARN:   "some:default:queue:arn",
				PreferredQueueARN: "some:preferred:queue:arn",
			},
			job: &av.Job{
				ID:       "jobID",
				Provider: Name,
				Input:    av.File{Name: "s3://some/path.mp4"},
				Output:   av.Dir{File: []av.File{{Name: "file1.mp4"}}},
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
		//{
		//	name: "acceleration is enabled and the default queue is used when a source has a large filesize",
		//	cfg: config.MediaConvert{
		//		DefaultQueueARN:   "some:default:queue:arn",
		//		PreferredQueueARN: "some:preferred:queue:arn",
		//	},
		//	job: &av.Job{
		//		ID:           "jobID",
		//		Provider: Name,
		//		SourceMedia:  "s3://some/path.mp4",
		//		SourceInfo:   av.File{FileSize: 1_000_000_000},
		//		Outputs:      []job.TranscodeOutput{{Preset: av.File{Name: audioOnlyPreset.Name}, FileName: "file1.mp4"}},
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
			tt.cfg.Destination = tt.destination

			p := &driver{
				cfg: tt.cfg,
			}
			tt.job.Output.File[0] = tt.preset
			input, err := p.createRequest(context.Background(), tt.job)
			if err != nil {
				t.Fatal(err)
			}

			//t.Logf("want: `%s`,\n", readable(tt.wantJobReq))
			//req := mc.CreateJobInput{}
			//json.Unmarshal([]byte(readable(tt.wantJobReq)), &req)
			if g, e := input, &tt.wantJobReq; !reflect.DeepEqual(g, e) {
				t.Fatalf("translation:\n\t\thave: %s\n\t\twant: %s", readable(g), readable(e))
			}
		})
	}
}

func TestAudioMixTimecodeBurnin(t *testing.T) {

	input := &av.Job{
		Input: av.File{
			Name: "s3://some/path.mov",
			// TODO(as): can this be split across the source
			// and the destination? Is Downmix the right name
			// should it just be a Mix? It doesn't seem very
			// general in its current state.
			Downmix: &av.Downmix{
				Src: []av.AudioChannel{
					{TrackIdx: 1, ChannelIdx: 1, Layout: "L"},
					{TrackIdx: 1, ChannelIdx: 2, Layout: "R"},
					{TrackIdx: 1, ChannelIdx: 3, Layout: "C"},
					{TrackIdx: 1, ChannelIdx: 4, Layout: "LFE"},
					{TrackIdx: 1, ChannelIdx: 5, Layout: "Ls"},
					{TrackIdx: 1, ChannelIdx: 6, Layout: "Rs"},
				},
				Dst: []av.AudioChannel{
					{TrackIdx: 1, ChannelIdx: 1, Layout: "L"},
					{TrackIdx: 1, ChannelIdx: 2, Layout: "R"},
				},
			}},
		Output: av.Dir{
			Path: "s3://some/destination/jobID",
			File: []av.File{{
				Name: "file1.mov",
				Video: av.Video{
					Codec: "h264", Profile: "high", Level: "4.1",
					Width: 300, Height: 400, Scantype: "progressive",
					Bitrate:  av.Bitrate{BPS: 400000, Control: "VBR"},
					TwoPass:  true,
					Gop:      av.Gop{Size: 120, Unit: "frames"},
					Overlays: av.Overlays{TimecodeBurnin: &av.Timecode{FontSize: 12, Position: 7}},
				},
				Audio: av.Audio{Codec: "aac", Bitrate: 20000},
			}},
		},
	}
	p := &driver{cfg: config.MediaConvert{PreferredQueueARN: "some:preferred:queue:arn", DefaultQueueARN: "some:default:queue:arn"}}
	j, err := p.createRequest(context.Background(), input)
	if err != nil {
		t.Fatal(err)
	}
	if g, e := j, &mcBurninDownmixJob; !reflect.DeepEqual(g, e) {
		t.Fatalf("translation:\n\t\thave: %s\n\t\twant: %s", readable(g), readable(e))
	}
}

var mcBurninDownmixJob = mc.CreateJobInput{
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
}

func TestDriverStatus(t *testing.T) {
	// NOTE(as): I don't see the value of preserving trailing slashes here
	// so I deleted them from the expected output where they were present
	//
	// BEFORE: "s3://some/destination/jobID"
	// AFTER: "s3://some/destination/jobID/"
	//
	tests := []struct {
		name        string
		destination string
		mcJob       mc.Job
		want        av.Status
	}{
		{
			name:        "Submitted",
			destination: "s3://some/destination",
			want:        av.Status{State: av.StateQueued, Provider: Name, Output: av.Dir{Path: "s3://some/destination/jobID"}},
			mcJob:       mc.Job{Status: mc.JobStatusSubmitted},
		},
		{
			name:        "Progressing",
			destination: "s3://some/destination",
			want:        av.Status{State: av.StateStarted, Provider: Name, Progress: 42, Output: av.Dir{Path: "s3://some/destination/jobID"}},
			mcJob:       mc.Job{Status: mc.JobStatusProgressing, JobPercentComplete: aws.Int64(42)},
		},
		{
			name:        "Complete",
			destination: "s3://some/destination",
			want: av.Status{
				State: av.StateFinished, Provider: Name, Progress: 100,
				Output: av.Dir{
					Path: "s3://some/destination/jobID",
					File: []av.File{
						{Name: "s3://some/destination/jobID/m_modifier.mp4", Container: "mp4", Video: av.Video{Height: 102, Width: 324}},
						{Name: "s3://some/destination/jobID/m_another_modifier.mp4", Container: "mp4"},
					},
				},
			},
			mcJob: mc.Job{
				Status: mc.JobStatusComplete,
				Settings: &mc.JobSettings{
					OutputGroups: []mc.OutputGroup{
						{
							OutputGroupSettings: &mc.OutputGroupSettings{Type: mc.OutputGroupTypeFileGroupSettings, FileGroupSettings: &mc.FileGroupSettings{Destination: aws.String("s3://some/destination/jobID/m")}},
							Outputs: []mc.Output{
								{NameModifier: aws.String("_modifier"), VideoDescription: &mc.VideoDescription{Height: aws.Int64(102), Width: aws.Int64(324)},
									ContainerSettings: &mc.ContainerSettings{Container: mc.ContainerTypeMp4, Mp4Settings: &mc.Mp4Settings{Mp4MajorBrand: aws.String("isom")}}},
								{NameModifier: aws.String("_another_modifier"),
									ContainerSettings: &mc.ContainerSettings{Container: mc.ContainerTypeMp4, Mp4Settings: &mc.Mp4Settings{Mp4MajorBrand: aws.String("isom")}}},
								{NameModifier: aws.String(""),
									ContainerSettings: &mc.ContainerSettings{Container: mc.ContainerTypeM2ts}},
								{ContainerSettings: &mc.ContainerSettings{Container: mc.ContainerTypeM2ts}},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			p := &driver{cfg: config.MediaConvert{
				Destination: tt.destination,
			}}

			status := p.status(&defaultJob, &tt.mcJob)
			if g, e := status, &tt.want; !reflect.DeepEqual(g, e) {
				t.Fatalf("driver.JobStatus(): wrong job request\nWant %+v\nGot %+v\nDiff %s", e, g, cmp.Diff(e, g))
			}
		})
	}
}

func readable(i interface{}) string {
	data, _ := json.Marshal(i)
	var v interface{}
	json.Unmarshal(data, &v)
	filterjunk(v)
	data, _ = json.Marshal(v)
	return string(data)
}
func filterjunk(v interface{}) {
	switch t := v.(type) {
	case map[string]interface{}:
		for k, v := range t {
			if v == nil || len(fmt.Sprint(v)) == 0 {
				delete(t, k)
			} else {
				filterjunk(v)
			}
		}
	case []interface{}:
		for _, v := range t {
			filterjunk(v)
		}
	}
}
