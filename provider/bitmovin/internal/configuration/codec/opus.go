package codec

import (
	"fmt"
	"strconv"

	"github.com/bitmovin/bitmovin-api-sdk-go"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/pkg/errors"
)

const defaultOpusSampleRate = 48000

// NewOpus creates an Opus codec configuration and returns its ID
func NewOpus(api *bitmovin.BitmovinApi, bitrate string) (string, error) {
	createCfg, err := opusConfigFrom(bitrate)
	if err != nil {
		return "", err
	}

	cfg, err := api.Encoding.Configurations.Audio.Opus.Create(createCfg)
	if err != nil {
		return "", errors.Wrap(err, "creating audio cfg")
	}

	return cfg.Id, nil
}

func opusConfigFrom(bitrate string) (model.OpusAudioConfiguration, error) {
	convertedBitrate, err := strconv.ParseInt(bitrate, 10, 64)
	if err != nil {
		return model.OpusAudioConfiguration{}, errors.Wrapf(err, "parsing audio bitrate %q to int64", bitrate)
	}

	return model.OpusAudioConfiguration{
		Name:    fmt.Sprintf("opus_%s_%d", bitrate, defaultOpusSampleRate),
		Bitrate: &convertedBitrate,
		Rate:    floatToPtr(defaultOpusSampleRate),
	}, nil
}
