package codec

import (
	"fmt"
	"strings"

	"github.com/bitmovin/bitmovin-api-sdk-go"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/video-transcoding-api/db"
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

	presetWidth := preset.Video.Width
	if presetWidth != "" {
		width, err := dimensionFrom(presetWidth)
		if err != nil {
			return model.Av1VideoConfiguration{}, err
		}
		cfg.Width = width
	}

	presetHeight := preset.Video.Height
	if presetHeight != "" {
		height, err := dimensionFrom(presetHeight)
		if err != nil {
			return model.Av1VideoConfiguration{}, err
		}
		cfg.Height = height
	}

	bitrate, err := bitrateFrom(preset.Video.Bitrate)
	if err != nil {
		return model.Av1VideoConfiguration{}, err
	}
	cfg.Bitrate = bitrate

	presetGOPSize := preset.Video.GopSize
	if presetGOPSize != "" {
		switch strings.ToLower(preset.Video.GopUnit) {
		case db.GopUnitFrames, "":
			gopSize, err := gopSizeFrom(presetGOPSize)
			if err != nil {
				return model.Av1VideoConfiguration{}, err
			}
			cfg.MinGfInterval = gopSize
			cfg.MaxGfInterval = gopSize
		case db.GopUnitSeconds:
			return model.Av1VideoConfiguration{}, errors.New("Gop can only be expressed in frames in AV1")
		default:
			return model.Av1VideoConfiguration{}, fmt.Errorf("GopUnit %v not recognized", preset.Video.GopUnit)
		}
	}

	cfg.EncodingMode = model.EncodingMode_SINGLE_PASS
	if preset.TwoPass {
		cfg.EncodingMode = model.EncodingMode_TWO_PASS
	}

	return cfg, nil
}
