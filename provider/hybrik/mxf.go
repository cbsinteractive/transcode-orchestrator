package hybrik

import (
	"fmt"

	"github.com/cbsinteractive/hybrik-sdk-go"
	"github.com/cbsinteractive/video-transcoding-api/db"
)

func modifyTranscodePayloadForMXFSources(payload hybrik.TranscodePayload, preset db.Preset) (hybrik.TranscodePayload, error) {
	transcodeTargets, ok := payload.Targets.([]hybrik.TranscodeTarget)
	if !ok {
		return hybrik.TranscodePayload{}, fmt.Errorf("targets are not TranscodeTargets: %v", payload.LocationTargetPayload.Targets)
	}

	modifiedTargets := []hybrik.TranscodeTarget{}
	for _, target := range transcodeTargets {
		if _, hdrEnabled := hdrTypeFromPreset(preset); hdrEnabled {
			// forcing this to two, mxf sources require two-pass
			// when processing sources for HDR output
			target.NumPasses = 2
		}

		if preset.Video.HDR10Settings.Enabled {
			// overriding this value to tell hybrik where to fetch the HDR metadata
			target.Video.HDR10.Source = "source_metadata"
		}
		modifiedTargets = append(modifiedTargets, target)
	}
	payload.Targets = modifiedTargets

	return payload, nil
}
