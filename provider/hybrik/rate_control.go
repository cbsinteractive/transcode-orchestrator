package hybrik

import (
	"fmt"
	"strings"

	hwrapper "github.com/cbsinteractive/hybrik-sdk-go"
	"github.com/cbsinteractive/video-transcoding-api/db"
)

type rateControlMode = string

const (
	rateControlModeVBR = "vbr"
	rateControlModeCBR = "cbr"

	vbrVariabilityPercent = 10
)

var supportedRateControlModes = map[rateControlMode]struct{}{
	rateControlModeCBR: {},
	rateControlModeVBR: {},
}

func enrichPresetWithRateControl(hybrikPreset hwrapper.Preset, preset db.Preset) (hwrapper.Preset, error) {
	mode := strings.ToLower(preset.RateControl)

	_, found := supportedRateControlModes[mode]
	if !found {
		return hwrapper.Preset{}, fmt.Errorf("rate control mode %q is not supported in hybrik, the currently "+
			"supported modes are %v", mode, supportedRateControlModes)
	}

	for idx, target := range hybrikPreset.Payload.Targets {
		target.Video.BitrateMode = mode

		// in the case of vbr we constrain the min/max bitrate based on a hardcoded variability percent
		if mode == rateControlModeVBR {
			target.Video.MaxBitrateKb = percentTargetOf(target.Video.BitrateKb, vbrVariabilityPercent)
			target.Video.MinBitrateKb = percentTargetOf(target.Video.BitrateKb, -vbrVariabilityPercent)
		}

		// set the enriched target back onto the preset
		hybrikPreset.Payload.Targets[idx] = target
	}

	return hybrikPreset, nil
}

func percentTargetOf(bitrate, percent int) int {
	return bitrate * (100 + percent) / 100
}
