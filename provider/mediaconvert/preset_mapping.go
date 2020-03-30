package mediaconvert

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/cbsinteractive/video-transcoding-api/db"
	"github.com/cbsinteractive/video-transcoding-api/provider"
	"github.com/pkg/errors"
)

func outputFrom(preset db.Preset) (mediaconvert.Output, error) {
	container, err := containerFrom(preset.Container)
	if err != nil {
		return mediaconvert.Output{}, errors.Wrap(err, "mapping preset container to MediaConvert container")
	}

	var videoPreset *mediaconvert.VideoDescription
	if preset.Video != (db.VideoPreset{}) {
		videoPreset, err = videoPresetFrom(preset)
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
		audioPresets = append(audioPresets, audioPreset)
	}

	output := mediaconvert.Output{
		ContainerSettings: &mediaconvert.ContainerSettings{
			Container: container,
		},
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
	case "m3u8":
		return mediaconvert.ContainerTypeM3u8, nil
	case "cmaf":
		return mediaconvert.ContainerTypeCmfc, nil
	case "mp4":
		return mediaconvert.ContainerTypeMp4, nil
	default:
		return "", fmt.Errorf("container %q not supported with mediaconvert", container)
	}
}

func videoPresetFrom(preset db.Preset) (*mediaconvert.VideoDescription, error) {
	videoPreset := mediaconvert.VideoDescription{
		ScalingBehavior:   mediaconvert.ScalingBehaviorDefault,
		TimecodeInsertion: mediaconvert.VideoTimecodeInsertionDisabled,
		AntiAlias:         mediaconvert.AntiAliasEnabled,
		RespondToAfd:      mediaconvert.RespondToAfdNone,
	}

	if preset.Video.Width != "" {
		width, err := strconv.ParseInt(preset.Video.Width, 10, 64)
		if err != nil {
			return nil, errors.Wrapf(err, "parsing video width %q to int64", preset.Video.Width)
		}
		videoPreset.Width = aws.Int64(width)
	}

	if preset.Video.Height != "" {
		height, err := strconv.ParseInt(preset.Video.Height, 10, 64)
		if err != nil {
			return nil, errors.Wrapf(err, "parsing video height %q to int64", preset.Video.Height)
		}
		videoPreset.Height = aws.Int64(height)
	}

	if hdr10 := preset.Video.HDR10Settings; hdr10.Enabled {
		mcHDR10Metadata := &mediaconvert.Hdr10Metadata{}
		if hdr10.MasterDisplay != "" {
			display, err := parseMasterDisplay(hdr10.MasterDisplay)
			if err != nil {
				return nil, errors.Wrap(err, "parsing master display string")
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

		videoPreset.VideoPreprocessors = &mediaconvert.VideoPreprocessor{
			ColorCorrector: &mediaconvert.ColorCorrector{
				Hdr10Metadata:        mcHDR10Metadata,
				ColorSpaceConversion: mediaconvert.ColorSpaceConversionForceHdr10,
			},
		}
	}

	codec := strings.ToLower(preset.Video.Codec)
	switch codec {
	case "h264":
		settings, err := h264CodecSettingsFrom(preset)
		if err != nil {
			return nil, errors.Wrap(err, "building h264 codec settings")
		}

		videoPreset.CodecSettings = settings
	case "h265":
		settings, err := h265CodecSettingsFrom(preset)
		if err != nil {
			return nil, errors.Wrap(err, "building h265 codec settings")
		}

		videoPreset.CodecSettings = settings
	case "av1":
		settings, err := av1CodecSettingsFrom(preset)
		if err != nil {
			return nil, errors.Wrap(err, "building av1 codec settings")
		}

		videoPreset.CodecSettings = settings
	default:
		return nil, fmt.Errorf("video codec %q is not yet supported with mediaconvert", codec)
	}

	if videoPreset.VideoPreprocessors == nil {
		videoPreset.VideoPreprocessors = &mediaconvert.VideoPreprocessor{}
	}

	videoPreset.VideoPreprocessors.Deinterlacer = &mediaconvert.Deinterlacer{
		Algorithm: mediaconvert.DeinterlaceAlgorithmInterpolate,
		Control:   mediaconvert.DeinterlacerControlNormal,
		Mode:      mediaconvert.DeinterlacerModeAdaptive,
	}

	return &videoPreset, nil
}

func audioPresetFrom(preset db.Preset) (mediaconvert.AudioDescription, error) {
	audioPreset := mediaconvert.AudioDescription{}

	codec := strings.ToLower(preset.Audio.Codec)
	switch codec {
	case "aac":
		bitrate, err := strconv.ParseInt(preset.Audio.Bitrate, 10, 64)
		if err != nil {
			return mediaconvert.AudioDescription{}, errors.Wrapf(err, "parsing audio bitrate %q to int64", preset.Audio.Bitrate)
		}

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
	default:
		return mediaconvert.AudioDescription{}, fmt.Errorf("audio codec %q is not yet supported with mediaconvert", codec)
	}

	return audioPreset, nil
}
