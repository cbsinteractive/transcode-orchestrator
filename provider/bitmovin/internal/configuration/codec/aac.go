package codec

import (
	"fmt"

	"github.com/bitmovin/bitmovin-api-sdk-go"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/pkg/errors"
)

const defaultAACSampleRate = 48000

// NewAAC creates an AAC codec configuration and returns its ID
func NewAAC(api *bitmovin.BitmovinApi, bitrate int64) (string, error) {
	createCfg, err := aacConfigFrom(bitrate)
	if err != nil {
		return "", err
	}

	cfg, err := api.Encoding.Configurations.Audio.Aac.Create(createCfg)
	if err != nil {
		return "", errors.Wrap(err, "creating audio cfg")
	}

	return cfg.Id, nil
}

func aacConfigFrom(bitrate int64) (model.AacAudioConfiguration, error) {
	return model.AacAudioConfiguration{
		Name:    fmt.Sprintf("aac_%d_%d", bitrate, defaultAACSampleRate),
		Bitrate: &bitrate,
		Rate:    floatToPtr(defaultAACSampleRate),
	}, nil
}
