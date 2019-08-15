package hybrik

import (
	hwrapper "github.com/cbsinteractive/hybrik-sdk-go"
	"github.com/cbsinteractive/video-transcoding-api/db"
)

func modifyPresetForMXFSources(hybrikPreset hwrapper.Preset, preset db.Preset) (hwrapper.Preset, error) {
	modifiedTargets := []hwrapper.PresetTarget{}
	for _, target := range hybrikPreset.Payload.Targets {
		if _, hdrEnabled := hdrTypeFromPreset(preset); hdrEnabled {
			// forcing this to two, mxf sources require two-pass
			// when processing sources for HDR output
			target.NumPasses = 2
		}

		if preset.Video.HDR10Settings.Enabled {
			// overriding this value to tell hybrik where to fetch the HDR metadata
			target.Video.HDR10.Source = "source_metadata"

			// resetting the ffmpeg args to remove codec tags that break hybrik during
			// hdr10 h265 jobs
			target.Video.FFMPEGArgs = ""
		}
		modifiedTargets = append(modifiedTargets, target)
	}
	hybrikPreset.Payload.Targets = modifiedTargets

	return hybrikPreset, nil
}
