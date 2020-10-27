package codec

import (
	"fmt"

	"github.com/bitmovin/bitmovin-api-sdk-go"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/pkg/errors"
)

const defaultOpusSampleRate = 48000

// NewOpus creates an Opus codec configuration and returns its ID
func NewOpus(api *bitmovin.BitmovinApi, bitrate int64) (string, error) {
	cfg, err := api.Encoding.Configurations.Audio.Opus.Create(model.OpusAudioConfiguration{
		Name:    fmt.Sprintf("opus_%s_%d", bitrate, defaultOpusSampleRate),
		Bitrate: &bitrate,
		Rate:    floatToPtr(defaultOpusSampleRate),
	})
	if err != nil {
		return "", errors.Wrap(err, "creating audio cfg")
	}
	return cfg.Id, nil
}
