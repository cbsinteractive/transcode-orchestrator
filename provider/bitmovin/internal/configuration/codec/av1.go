package codec

import (
	"fmt"
	"strings"

	"github.com/bitmovin/bitmovin-api-sdk-go"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/transcode-orchestrator/db"
	"github.com/pkg/errors"
)

// NewAV1 creates a AV1 codec configuration and returns its ID
func NewAV1(api *bitmovin.BitmovinApi, preset db.Preset) (string, error) {
	newVidCfg, err := av1ConfigFrom(preset)
	if err != nil {
		return "", errors.Wrap(err, "creating av1 config object")
	}

	vidCfg, err := api.Encoding.Configurations.Video.Av1.Create(newVidCfg)
	if err != nil {
		return "", errors.Wrap(err, "creating av1 config with the API")
	}

	return vidCfg.Id, nil
}

func av1ConfigFrom(preset db.Preset) (model.Av1VideoConfiguration, error) {
	cfg := model.Av1VideoConfiguration{}

	cfg.Name = strings.ToLower(preset.Name)

	if n := int32(preset.Video.Width); n != 0 {
		cfg.Width = &n
	}
	if n := int32(preset.Video.Height); n != 0 {
		cfg.Height = &n
	}

	bitrate := int64(preset.Video.Bitrate)
	cfg.Bitrate = &bitrate

	gopSize := int32(preset.Video.GopSize)
	if gopSize != 0 {
		switch strings.ToLower(preset.Video.GopUnit) {
		case db.GopUnitFrames, "":
			cfg.MinGfInterval = &gopSize
			cfg.MaxGfInterval = &gopSize
		case db.GopUnitSeconds:
			return model.Av1VideoConfiguration{}, errors.New("Gop can only be expressed in frames in AV1")
		default:
			return model.Av1VideoConfiguration{}, fmt.Errorf("GopUnit %v not recognized", preset.Video.GopUnit)
		}
	}

	// Single-pass encoding throws an error
	cfg.EncodingMode = model.EncodingMode_TWO_PASS

	return cfg, nil
}
