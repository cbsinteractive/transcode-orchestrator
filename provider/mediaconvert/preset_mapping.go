package mediaconvert

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	mc "github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/cbsinteractive/transcode-orchestrator/job"
	"github.com/pkg/errors"
)

var timecodePositionMap = map[int]mc.TimecodeBurninPosition{
	0: mc.TimecodeBurninPositionTopCenter,
	1: mc.TimecodeBurninPositionTopLeft,
	2: mc.TimecodeBurninPositionTopRight,
	3: mc.TimecodeBurninPositionMiddleCenter,
	4: mc.TimecodeBurninPositionMiddleLeft,
	5: mc.TimecodeBurninPositionMiddleRight,
	6: mc.TimecodeBurninPositionBottomCenter,
	7: mc.TimecodeBurninPositionBottomLeft,
	8: mc.TimecodeBurninPositionBottomRight,
}

func outputFrom(f job.File, input job.File) (mc.Output, error) {
	container, err := containerFrom(f.Container)
	if err != nil {
		return mc.Output{}, fmt.Errorf("container: %w", err)
	}

	var mvideo *mc.VideoDescription
	if f.Video.On() {
		mvideo, err = videoPresetFrom(f, input)
		if err != nil {
			return mc.Output{}, fmt.Errorf("video: %w", err)
		}
	}

	var maudio []mc.AudioDescription
	if f.Audio.On() {
		audio, err := audioPresetFrom(f)
		if err != nil {
			return mc.Output{}, fmt.Errorf("audio: %w", err)
		}
		if f.Audio.Discrete {
			maudio = audioSplit(audio)
		} else {
			maudio = append(maudio)
		}
	}

	output := mc.Output{
		ContainerSettings: containerSettingsFrom(container),
		VideoDescription:  mvideo,
		AudioDescriptions: maudio,
	}

	return output, nil
}

func state(status mc.JobStatus) job.State {
	switch status {
	case mc.JobStatusSubmitted:
		return job.StateQueued
	case mc.JobStatusProgressing:
		return job.StateStarted
	case mc.JobStatusComplete:
		return job.StateFinished
	case mc.JobStatusCanceled:
		return job.StateCanceled
	case mc.JobStatusError:
		return job.StateFailed
	default:
		return job.StateUnknown
	}
}

func containerFrom(v string) (mc.ContainerType, error) {
	switch strings.ToLower(v) {
	case "mxf":
		return mc.ContainerTypeMxf, nil
	case "m3u8":
		return mc.ContainerTypeM3u8, nil
	case "cmaf":
		return mc.ContainerTypeCmfc, nil
	case "mp4":
		return mc.ContainerTypeMp4, nil
	case "mov":
		return mc.ContainerTypeMov, nil
	case "webm":
		return mc.ContainerTypeWebm, nil
	default:
		return "", fmt.Errorf("%w: %q", ErrUnsupported, v)
	}
}

func containerSettingsFrom(container mc.ContainerType) *mc.ContainerSettings {
	cs := &mc.ContainerSettings{
		Container: container,
	}

	switch container {
	case mc.ContainerTypeMxf:
		// NOTE(as): AWS claims to auto-detect profile
	case mc.ContainerTypeMp4:
		cs.Mp4Settings = &mc.Mp4Settings{
			//ISO specification for base media file format
			Mp4MajorBrand: aws.String("isom"),
		}
	case mc.ContainerTypeMov:
		cs.MovSettings = &mc.MovSettings{
			ClapAtom:           mc.MovClapAtomExclude,
			CslgAtom:           mc.MovCslgAtomInclude,
			PaddingControl:     mc.MovPaddingControlOmneon,
			Reference:          mc.MovReferenceSelfContained,
			Mpeg2FourCCControl: mc.MovMpeg2FourCCControlMpeg,
		}
	}

	return cs
}

func videoPresetFrom(f job.File, input job.File) (*mc.VideoDescription, error) {
	mv := mc.VideoDescription{
		ScalingBehavior:   mc.ScalingBehaviorDefault,
		TimecodeInsertion: mc.VideoTimecodeInsertionDisabled,
		AntiAlias:         mc.AntiAliasEnabled,
		RespondToAfd:      mc.RespondToAfdNone,
	}

	if f.Video.Width != 0 {
		mv.Width = aws.Int64(int64(f.Video.Width))
	}
	if f.Video.Height != 0 {
		mv.Height = aws.Int64(int64(f.Video.Height))
	}

	var s *mc.VideoCodecSettings
	var err error

	codec := strings.ToLower(f.Video.Codec)
	switch codec {
	case "xdcam":
		s, err = mpeg2XDCAM.generate(f)
		defer func() {
			if mv.VideoPreprocessors != nil {
				mv.VideoPreprocessors.Deinterlacer = nil
			}
		}()
	case "h264":
		s, err = h264CodecSettingsFrom(f)
	case "h265":
		s, err = h265CodecSettingsFrom(f)
	case "vp8":
		s, err = vp8CodecSettingsFrom(f)
	case "av1":
		s, err = av1CodecSettingsFrom(f)
	default:
		return nil, fmt.Errorf("video: codec: %w: %q", ErrUnsupported, codec)
	}
	if err != nil {
		return nil, fmt.Errorf("building %s codec settings: %w", codec, err)
	}
	mv.CodecSettings = s

	videoPreprocessors, err := videoPreprocessorsFrom(f.Video)
	if err != nil {
		return nil, errors.Wrap(err, "building videoPreprocessors")
	}
	mv.VideoPreprocessors = videoPreprocessors

	if f.Video.Scantype != "progressive" {
		return &mv, nil
	}
	switch input.Video.Scantype {
	case job.ScanProgressive:
	case job.ScanInterlaced:
		mv.VideoPreprocessors.Deinterlacer = &mc.Deinterlacer{
			Algorithm: mc.DeinterlaceAlgorithmInterpolate,
			Control:   mc.DeinterlacerControlNormal,
			Mode:      mc.DeinterlacerModeDeinterlace,
		}
	default:
		mv.VideoPreprocessors.Deinterlacer = &mc.Deinterlacer{
			Algorithm: mc.DeinterlaceAlgorithmInterpolate,
			Control:   mc.DeinterlacerControlNormal,
			Mode:      mc.DeinterlacerModeAdaptive,
		}
	}

	return &mv, nil
}

var (
	deinterlacerStandard = mc.Deinterlacer{
		Algorithm: mc.DeinterlaceAlgorithmInterpolate,
		Control:   mc.DeinterlacerControlNormal,
		Mode:      mc.DeinterlacerModeDeinterlace,
	}
	deinterlacerAdaptive = mc.Deinterlacer{
		Algorithm: mc.DeinterlaceAlgorithmInterpolate,
		Control:   mc.DeinterlacerControlNormal,
		Mode:      mc.DeinterlacerModeAdaptive,
	}
)

type setter struct {
	dst job.File
	src job.File
}

func (s setter) Scan(v *mc.VideoDescription) *mc.VideoDescription {
	const (
		// constants have same value for src/dst, but different types...
		progressive = string(job.ScanProgressive)
		interlaced  = string(job.ScanInterlaced)
	)
	if v == nil {
		v = &mc.VideoDescription{}
	}
	if v.VideoPreprocessors == nil {
		v.VideoPreprocessors = &mc.VideoPreprocessor{}
	}

	switch s.dst.Video.Scantype {
	case interlaced:
	case progressive:
		fallthrough
	default: // progressive
		switch string(s.src.Video.Scantype) {
		case progressive:
		case interlaced:
			v.VideoPreprocessors.Deinterlacer = &deinterlacerStandard
		default:
			v.VideoPreprocessors.Deinterlacer = &deinterlacerAdaptive
		}
	}
	return v
}

func (s setter) Crop(v *mc.VideoDescription) *mc.VideoDescription {
	if v == nil {
		v = &mc.VideoDescription{}
	}

	var (
		crop = s.dst.Video.Crop
		h, w = int(s.src.Video.Height), int(s.src.Video.Width)
	)
	if crop.Empty() || h <= 0 || w <= 0 {
		return v
	}

	roundEven := func(i, mod int) *int64 {
		if i%2 != 0 {
			i += mod
		}
		return aws.Int64(int64(i))
	}
	v.Crop = &mc.Rectangle{
		Height: roundEven(h-crop.Top-crop.Bottom, -1),
		Width:  roundEven(w-crop.Left-crop.Right, -1),
		X:      roundEven(crop.Left, 1),
		Y:      roundEven(crop.Top, 1),
	}
	return v
}

func videoPreprocessorsFrom(v job.Video) (*mc.VideoPreprocessor, error) {
	mvpp := &mc.VideoPreprocessor{}

	if tcb := v.Overlays.TimecodeBurnin; tcb != nil {
		mvpp.TimecodeBurnin = &mc.TimecodeBurnin{
			Prefix:   &tcb.Prefix,
			FontSize: aws.Int64(int64(tcb.FontSize)),
			Position: timecodePositionMap[tcb.Position],
		}
	}

	if hdr10 := v.HDR10; hdr10.Enabled {
		md := &mc.Hdr10Metadata{}
		if hdr10.MasterDisplay != "" {
			d, err := parseMasterDisplay(hdr10.MasterDisplay)
			if err != nil {
				return mvpp, fmt.Errorf("parsing master display string: %w", err)
			}
			md.BluePrimaryX = aws.Int64(d.bluePrimaryX)
			md.BluePrimaryY = aws.Int64(d.bluePrimaryY)
			md.GreenPrimaryX = aws.Int64(d.greenPrimaryX)
			md.GreenPrimaryY = aws.Int64(d.greenPrimaryY)
			md.RedPrimaryX = aws.Int64(d.redPrimaryX)
			md.RedPrimaryY = aws.Int64(d.redPrimaryY)
			md.WhitePointX = aws.Int64(d.whitePointX)
			md.WhitePointY = aws.Int64(d.whitePointY)
			md.MaxLuminance = aws.Int64(d.maxLuminance)
			md.MinLuminance = aws.Int64(d.minLuminance)
		}

		if hdr10.MaxCLL != 0 {
			md.MaxContentLightLevel = aws.Int64(int64(hdr10.MaxCLL))
		}

		if hdr10.MaxFALL != 0 {
			md.MaxFrameAverageLightLevel = aws.Int64(int64(hdr10.MaxFALL))
		}

		mvpp.ColorCorrector = &mc.ColorCorrector{
			Hdr10Metadata:        md,
			ColorSpaceConversion: mc.ColorSpaceConversionForceHdr10,
		}
	}

	return mvpp, nil
}

func unmute(n int, channel ...int64) []int64 {
	channel = append([]int64{}, channel...)
	channel[n] = 0
	return channel
}

func audioSplit(a mc.AudioDescription) (split []mc.AudioDescription) {
	if a.CodecSettings == nil ||
		a.CodecSettings.Codec != mc.AudioCodecWav ||
		a.CodecSettings.WavSettings == nil ||
		*a.CodecSettings.WavSettings.Channels < 2 {
		return []mc.AudioDescription{a}
	}

	n := int64(*a.CodecSettings.WavSettings.Channels)
	gain := make([]int64, n)
	*a.CodecSettings.WavSettings.Channels = 1

	for i := range gain {
		gain[i] = -60
	}
	for i := range gain {
		split = append(split, mc.AudioDescription{
			CodecSettings: a.CodecSettings,
			RemixSettings: &mc.RemixSettings{
				ChannelMapping: &mc.ChannelMapping{
					OutputChannels: []mc.OutputChannelMapping{{
						InputChannels: unmute(i, gain...),
					},
					}},
				ChannelsIn:  &n,
				ChannelsOut: aws.Int64(1),
			}})
	}
	return split
}

func audioPresetFrom(f job.File) (mc.AudioDescription, error) {
	mad := mc.AudioDescription{}

	if f.Audio.Normalize {
		mad.AudioNormalizationSettings = &mc.AudioNormalizationSettings{
			Algorithm:        mc.AudioNormalizationAlgorithmItuBs17703,
			AlgorithmControl: mc.AudioNormalizationAlgorithmControlCorrectAudio,
		}
	}

	codec := strings.ToLower(f.Audio.Codec)
	bitrate := int64(f.Audio.Bitrate)

	switch codec {
	case "pcm":
		mad.CodecSettings = &mc.AudioCodecSettings{
			Codec: mc.AudioCodecWav,
			WavSettings: &mc.WavSettings{
				BitDepth:   aws.Int64(24),
				Channels:   aws.Int64(2),
				SampleRate: aws.Int64(defaultAudioSampleRate),
				Format:     "RIFF",
			},
		}
	case "aac":
		mad.CodecSettings = &mc.AudioCodecSettings{
			Codec: mc.AudioCodecAac,
			AacSettings: &mc.AacSettings{
				SampleRate:      aws.Int64(defaultAudioSampleRate),
				Bitrate:         aws.Int64(bitrate),
				CodecProfile:    mc.AacCodecProfileLc,
				CodingMode:      mc.AacCodingModeCodingMode20,
				RateControlMode: mc.AacRateControlModeCbr,
			},
		}
	case "opus":
		mad.CodecSettings = &mc.AudioCodecSettings{
			Codec: mc.AudioCodecOpus,
			OpusSettings: &mc.OpusSettings{
				Channels:   aws.Int64(2),
				Bitrate:    aws.Int64(bitrate),
				SampleRate: aws.Int64(defaultAudioSampleRate),
			},
		}
	case "vorbis":
		mad.CodecSettings = &mc.AudioCodecSettings{
			Codec: mc.AudioCodecVorbis,
			VorbisSettings: &mc.VorbisSettings{
				Channels:   aws.Int64(2),
				SampleRate: aws.Int64(defaultAudioSampleRate),
				VbrQuality: aws.Int64(vbrLevel(bitrate)),
			},
		}
	default:
		return mc.AudioDescription{}, fmt.Errorf("%w: %q", ErrUnsupported, codec)
	}

	return mad, nil
}

func vbrLevel(bitrate int64) int64 {
	var level int64
	bKbps := bitrate / 1000

	switch {
	case bKbps == 0:
		level = 4
	case bKbps <= 128:
		level = (bKbps / 16) - 4
	case bKbps > 128 && bKbps <= 256:
		level = bKbps / 32
	case bKbps > 256:
		level = (bKbps / 64) + 4
	}

	if level < -1 {
		level = -1
	} else if level > 10 {
		level = 10
	}

	return level
}
