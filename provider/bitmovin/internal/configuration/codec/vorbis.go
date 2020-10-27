package codec

import (
	"fmt"

	"github.com/bitmovin/bitmovin-api-sdk-go"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/pkg/errors"
)

const defaultVorbisSampleRate = 48000

// NewVorbis creates a Vorbis codec configuration and returns its ID
func NewVorbis(api *bitmovin.BitmovinApi, bitrate int64) (string, error) {
	cfg, err := api.Encoding.Configurations.Audio.Vorbis.Create(model.VorbisAudioConfiguration{
		Name:    fmt.Sprintf("vorbis_%d_%d", bitrate, defaultVorbisSampleRate),
		Bitrate: &bitrate,
		Rate:    floatToPtr(defaultVorbisSampleRate),
	})
	if err != nil {
		return "", errors.Wrap(err, "creating audio config")
	}
	return cfg.Id, nil
}
