package mediaconvert

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/cbsinteractive/transcode-orchestrator/db"
	"github.com/cbsinteractive/transcode-orchestrator/provider"
	"github.com/pkg/errors"
)

var timecodePositionMap = map[int]mediaconvert.TimecodeBurninPosition{
	0: mediaconvert.TimecodeBurninPositionTopCenter,
	1: mediaconvert.TimecodeBurninPositionTopLeft,
	2: mediaconvert.TimecodeBurninPositionTopRight,
	3: mediaconvert.TimecodeBurninPositionMiddleCenter,
	4: mediaconvert.TimecodeBurninPositionMiddleLeft,
	5: mediaconvert.TimecodeBurninPositionMiddleRight,
	6: mediaconvert.TimecodeBurninPositionBottomCenter,
	7: mediaconvert.TimecodeBurninPositionBottomLeft,
	8: mediaconvert.TimecodeBurninPositionBottomRight,
}

func outputFrom(preset db.Preset, sourceInfo db.File) (mediaconvert.Output, error) {
	container, err := containerFrom(preset.Container)
	if err != nil {
		return mediaconvert.Output{}, errors.Wrap(err, "mapping preset container to MediaConvert container")
	}

	var videoPreset *mediaconvert.VideoDescription
	if preset.Video != (db.VideoPreset{}) {
		videoPreset, err = videoPresetFrom(preset, sourceInfo)
		if err != nil {
			return mediaconvert.Output{}, errors.Wrap(err, "generating video preset")
		}
	}

	var audioPresets []mediaconvert.AudioDescription
	if preset.Audio != (db.AudioPreset{}) {
		audioPreset, err := audioPresetFrom(preset)
		if err != nil {
			return mediaconvert.Output{}, errors.Wrap(err, "generating audio preset")
		}
		if preset.Audio.DiscreteTracks {
			audioPresets = audioSplit(audioPreset)
		} else {
			audioPresets = append(audioPresets, audioPreset)
		}
	}

	output := mediaconvert.Output{
		ContainerSettings: containerSettingsFrom(container),
		VideoDescription:  videoPreset,
		AudioDescriptions: audioPresets,
	}

	return output, nil
}

func providerStatusFrom(status mediaconvert.JobStatus) provider.Status {
	switch status {
	case mediaconvert.JobStatusSubmitted:
		return provider.StatusQueued
	case mediaconvert.JobStatusProgressing:
		return provider.StatusStarted
	case mediaconvert.JobStatusComplete:
		return provider.StatusFinished
	case mediaconvert.JobStatusCanceled:
		return provider.StatusCanceled
	case mediaconvert.JobStatusError:
		return provider.StatusFailed
	default:
		return provider.StatusUnknown
	}
}

func containerFrom(container string) (mediaconvert.ContainerType, error) {
	container = strings.ToLower(container)
	switch container {
	case "mxf":
		return mediaconvert.ContainerTypeMxf, nil
	case "m3u8":
		return mediaconvert.ContainerTypeM3u8, nil
	case "cmaf":
		return mediaconvert.ContainerTypeCmfc, nil
	case "mp4":
		return mediaconvert.ContainerTypeMp4, nil
	case "mov":
		return mediaconvert.ContainerTypeMov, nil
	case "webm":
		return mediaconvert.ContainerTypeWebm, nil
	default:
		return "", fmt.Errorf("container %q not supported with mediaconvert", container)
	}
}

func containerSettingsFrom(container mediaconvert.ContainerType) *mediaconvert.ContainerSettings {
	cs := &mediaconvert.ContainerSettings{
		Container: container,
	}

	switch container {
	case mediaconvert.ContainerTypeMxf:
		// NOTE(as): AWS claims to auto-detect profile
	case mediaconvert.ContainerTypeMp4:
		cs.Mp4Settings = &mediaconvert.Mp4Settings{
			//ISO specification for base media file format
			Mp4MajorBrand: aws.String("isom"),
		}
	case mediaconvert.ContainerTypeMov:
		cs.MovSettings = &mediaconvert.MovSettings{
			ClapAtom:           mediaconvert.MovClapAtomExclude,
			CslgAtom:           mediaconvert.MovCslgAtomInclude,
			PaddingControl:     mediaconvert.MovPaddingControlOmneon,
			Reference:          mediaconvert.MovReferenceSelfContained,
			Mpeg2FourCCControl: mediaconvert.MovMpeg2FourCCControlMpeg,
		}
	}

	return cs
}

func videoPresetFrom(preset db.Preset, sourceInfo db.File) (*mediaconvert.VideoDescription, error) {
	videoPreset := &mediaconvert.VideoDescription{
		ScalingBehavior:   mediaconvert.ScalingBehaviorDefault,
		TimecodeInsertion: mediaconvert.VideoTimecodeInsertionDisabled,
		AntiAlias:         mediaconvert.AntiAliasEnabled,
		RespondToAfd:      mediaconvert.RespondToAfdNone,
	}

	var (
		width, height int64
		err           error
		s             *mediaconvert.VideoCodecSettings
	)
	if preset.Video.Width != "" {
		width, err = strconv.ParseInt(preset.Video.Width, 10, 64)
		if err != nil {
			return nil, errors.Wrapf(err, "parsing video width %q to int64", preset.Video.Width)
		}
		videoPreset.Width = aws.Int64(width)
	}

	if preset.Video.Height != "" {
		height, err = strconv.ParseInt(preset.Video.Height, 10, 64)
		if err != nil {
			return nil, errors.Wrapf(err, "parsing video height %q to int64", preset.Video.Height)
		}
		videoPreset.Height = aws.Int64(height)
	}

	codec := strings.ToLower(preset.Video.Codec)
	switch codec {
	case "xdcam":
		s, err = mpeg2XDCAM.generate(preset)
		defer func() {
			videoPreset.VideoPreprocessors.Deinterlacer = nil
		}()
	case "h264":
		s, err = h264CodecSettingsFrom(preset)
	case "h265":
		s, err = h265CodecSettingsFrom(preset)
	case "vp8":
		s, err = vp8CodecSettingsFrom(preset)
	case "av1":
		s, err = av1CodecSettingsFrom(preset)
	default:
		return nil, fmt.Errorf("video codec %q is not yet supported with mediaconvert", codec)
	}
	if err != nil {
		return nil, fmt.Errorf("building %s codec settings: %w", codec, err)
	}
	videoPreset.CodecSettings = s

	videoPreprocessors, err := videoPreprocessorsFrom(preset.Video)
	if err != nil {
		return nil, errors.Wrap(err, "building videoPreprocessors")
	}
	videoPreset.VideoPreprocessors = videoPreprocessors

	// cropping
	if crop := preset.Video.Crop; !crop.Empty() && height > 0 && width > 0 {
		videoPreset.Crop = &mediaconvert.Rectangle{
			Height: aws.Int64(height - int64(crop.Top+crop.Bottom)),
			Width:  aws.Int64(width - int64(crop.Left+crop.Right)),
			X:      aws.Int64(int64(crop.Left)),
			Y:      aws.Int64(int64(crop.Top)),
		}
	}

	videoPreset = setter{dst: preset, src: sourceInfo}.ScanType(videoPreset)

	return videoPreset, nil
}

var (
	deinterlacerStandard = mediaconvert.Deinterlacer{
		Algorithm: mediaconvert.DeinterlaceAlgorithmInterpolate,
		Control:   mediaconvert.DeinterlacerControlNormal,
		Mode:      mediaconvert.DeinterlacerModeDeinterlace,
	}
	deinterlacerAdaptive = mediaconvert.Deinterlacer{
		Algorithm: mediaconvert.DeinterlaceAlgorithmInterpolate,
		Control:   mediaconvert.DeinterlacerControlNormal,
		Mode:      mediaconvert.DeinterlacerModeAdaptive,
	}
)

type setter struct {
	dst db.Preset
	src db.File
}

func (s setter) ScanType(v *mediaconvert.VideoDescription) *mediaconvert.VideoDescription {
	const (
		// constants have same value for src/dst, but different types...
		progressive = string(db.ScanTypeProgressive)
		interlaced  = string(db.ScanTypeInterlaced)
	)
	if v == nil {
		v = &mediaconvert.VideoDescription{}
	}
	if v.VideoPreprocessors == nil {
		v.VideoPreprocessors = &mediaconvert.VideoPreprocessor{}
	}

	switch s.dst.Video.InterlaceMode {
	case interlaced:
		switch string(s.src.ScanType) {
		case progressive:
		case interlaced:
		default:
		}
	case progressive:
		fallthrough
	default: // progressive
		switch string(s.src.ScanType) {
		case progressive:
		case interlaced:
			v.VideoPreprocessors.Deinterlacer = &deinterlacerStandard
		default:
			v.VideoPreprocessors.Deinterlacer = &deinterlacerAdaptive
		}
	}
	return v
}

func videoPreprocessorsFrom(videoPreset db.VideoPreset) (*mediaconvert.VideoPreprocessor, error) {
	videoPreprocessor := &mediaconvert.VideoPreprocessor{}

	if videoPreset.Overlays != nil && videoPreset.Overlays.TimecodeBurnin != nil {
		if tcBurnin := videoPreset.Overlays.TimecodeBurnin; tcBurnin.Enabled {
			videoPreprocessor.TimecodeBurnin = &mediaconvert.TimecodeBurnin{
				Prefix:   &tcBurnin.Prefix,
				FontSize: aws.Int64(int64(tcBurnin.FontSize)),
				Position: timecodePositionMap[tcBurnin.Position],
			}
		}
	}

	if hdr10 := videoPreset.HDR10Settings; hdr10.Enabled {
		mcHDR10Metadata := &mediaconvert.Hdr10Metadata{}
		if hdr10.MasterDisplay != "" {
			display, err := parseMasterDisplay(hdr10.MasterDisplay)
			if err != nil {
				return videoPreprocessor, errors.Wrap(err, "parsing master display string")
			}
			mcHDR10Metadata.BluePrimaryX = aws.Int64(display.bluePrimaryX)
			mcHDR10Metadata.BluePrimaryY = aws.Int64(display.bluePrimaryY)
			mcHDR10Metadata.GreenPrimaryX = aws.Int64(display.greenPrimaryX)
			mcHDR10Metadata.GreenPrimaryY = aws.Int64(display.greenPrimaryY)
			mcHDR10Metadata.RedPrimaryX = aws.Int64(display.redPrimaryX)
			mcHDR10Metadata.RedPrimaryY = aws.Int64(display.redPrimaryY)
			mcHDR10Metadata.WhitePointX = aws.Int64(display.whitePointX)
			mcHDR10Metadata.WhitePointY = aws.Int64(display.whitePointY)
			mcHDR10Metadata.MaxLuminance = aws.Int64(display.maxLuminance)
			mcHDR10Metadata.MinLuminance = aws.Int64(display.minLuminance)
		}

		if hdr10.MaxCLL != 0 {
			mcHDR10Metadata.MaxContentLightLevel = aws.Int64(int64(hdr10.MaxCLL))
		}

		if hdr10.MaxFALL != 0 {
			mcHDR10Metadata.MaxFrameAverageLightLevel = aws.Int64(int64(hdr10.MaxFALL))
		}

		videoPreprocessor.ColorCorrector = &mediaconvert.ColorCorrector{
			Hdr10Metadata:        mcHDR10Metadata,
			ColorSpaceConversion: mediaconvert.ColorSpaceConversionForceHdr10,
		}
	}

	return videoPreprocessor, nil
}

func unmute(n int, channel ...int64) []int64 {
	channel = append([]int64{}, channel...)
	channel[n] = 0
	return channel
}

func audioSplit(a mediaconvert.AudioDescription) (split []mediaconvert.AudioDescription) {
	if a.CodecSettings == nil ||
		a.CodecSettings.Codec != mediaconvert.AudioCodecWav ||
		a.CodecSettings.WavSettings == nil ||
		*a.CodecSettings.WavSettings.Channels < 2 {
		return []mediaconvert.AudioDescription{a}
	}

	n := int64(*a.CodecSettings.WavSettings.Channels)
	gain := make([]int64, n)
	*a.CodecSettings.WavSettings.Channels = 1

	for i := range gain {
		gain[i] = -60
	}
	for i := range gain {
		split = append(split, mediaconvert.AudioDescription{
			CodecSettings: a.CodecSettings,
			RemixSettings: &mediaconvert.RemixSettings{
				ChannelMapping: &mediaconvert.ChannelMapping{
					OutputChannels: []mediaconvert.OutputChannelMapping{{
						InputChannels: unmute(i, gain...),
					},
					}},
				ChannelsIn:  &n,
				ChannelsOut: aws.Int64(1),
			}})
	}
	return split
}

func audioPresetFrom(preset db.Preset) (mediaconvert.AudioDescription, error) {
	audioPreset := mediaconvert.AudioDescription{}

	if preset.Audio.Normalization {
		audioPreset.AudioNormalizationSettings = &mediaconvert.AudioNormalizationSettings{
			Algorithm:        mediaconvert.AudioNormalizationAlgorithmItuBs17703,
			AlgorithmControl: mediaconvert.AudioNormalizationAlgorithmControlCorrectAudio,
		}
	}

	codec := strings.ToLower(preset.Audio.Codec)
	bitrate, err := strconv.ParseInt(preset.Audio.Bitrate, 10, 64)
	if err != nil && codec != "pcm" {
		return mediaconvert.AudioDescription{}, errors.Wrapf(err, "parsing audio bitrate %q to int64", preset.Audio.Bitrate)
	}

	switch codec {
	case "pcm":
		audioPreset.CodecSettings = &mediaconvert.AudioCodecSettings{
			Codec: mediaconvert.AudioCodecWav,
			WavSettings: &mediaconvert.WavSettings{
				BitDepth:   aws.Int64(24),
				Channels:   aws.Int64(2),
				SampleRate: aws.Int64(defaultAudioSampleRate),
				Format:     "RIFF",
			},
		}
	case "aac":
		audioPreset.CodecSettings = &mediaconvert.AudioCodecSettings{
			Codec: mediaconvert.AudioCodecAac,
			AacSettings: &mediaconvert.AacSettings{
				SampleRate:      aws.Int64(defaultAudioSampleRate),
				Bitrate:         aws.Int64(bitrate),
				CodecProfile:    mediaconvert.AacCodecProfileLc,
				CodingMode:      mediaconvert.AacCodingModeCodingMode20,
				RateControlMode: mediaconvert.AacRateControlModeCbr,
			},
		}
	case "opus":
		audioPreset.CodecSettings = &mediaconvert.AudioCodecSettings{
			Codec: mediaconvert.AudioCodecOpus,
			OpusSettings: &mediaconvert.OpusSettings{
				Channels:   aws.Int64(2),
				Bitrate:    aws.Int64(bitrate),
				SampleRate: aws.Int64(defaultAudioSampleRate),
			},
		}
	case "vorbis":
		audioPreset.CodecSettings = &mediaconvert.AudioCodecSettings{
			Codec: mediaconvert.AudioCodecVorbis,
			VorbisSettings: &mediaconvert.VorbisSettings{
				Channels:   aws.Int64(2),
				SampleRate: aws.Int64(defaultAudioSampleRate),
				VbrQuality: aws.Int64(vbrLevel(bitrate)),
			},
		}
	default:
		return mediaconvert.AudioDescription{}, fmt.Errorf("audio codec %q is not yet supported with mediaconvert", codec)
	}

	return audioPreset, nil
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
