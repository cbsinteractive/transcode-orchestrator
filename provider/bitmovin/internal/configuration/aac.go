package configuration

import (
	"github.com/bitmovin/bitmovin-api-sdk-go"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/video-transcoding-api/db"
	"github.com/cbsinteractive/video-transcoding-api/provider/bitmovin/internal/configuration/codec"
	"github.com/cbsinteractive/video-transcoding-api/provider/bitmovin/internal/types"
	"github.com/pkg/errors"
)

// AAC is a configuration service for content in this codec pair
type AAC struct {
	api *bitmovin.BitmovinApi
}

// NewAAC returns a service for managing AAC audio only configurations
func NewAAC(api *bitmovin.BitmovinApi) *AAC {
	return &AAC{api: api}
}

// Create will create a new AAC configuration based on a preset
func (c *AAC) Create(preset db.Preset) (string, error) {
	audCfgID, err := codec.NewAAC(c.api, preset.Audio.Bitrate, customDataWith("", "mp4"))

	if err != nil {
		return "", err
	}

	return audCfgID, nil
}

// Get retrieves audio with a presetID
// the function will return a boolean indicating whether the video
// configuration was found, a config object and an optional error
func (c *AAC) Get(cfgID string) (bool, Details, error) {
	audCfg, customData, err := c.audioConfigWithCustomDataFrom(cfgID)
	if err != nil {
		return false, Details{}, errors.Wrapf(err, "getting the audio configuration with ID %q", cfgID)
	}

	return true, Details{nil, audCfg, customData}, nil
}

// Delete removes the audio / video configurations
func (c *AAC) Delete(cfgID string) (found bool, e error) {
	customData, err := c.audioCustomDataFrom(cfgID)
	if err != nil {
		return found, err
	}

	audCfgID, err := AudCfgIDFrom(customData)
	if err != nil {
		return found, err
	}

	audCfg, err := c.api.Encoding.Configurations.Audio.Aac.Get(audCfgID)
	if err != nil {
		return found, errors.Wrap(err, "retrieving audio configuration")
	}
	found = true

	_, err = c.api.Encoding.Configurations.Audio.Aac.Delete(audCfg.Id)
	if err != nil {
		return found, errors.Wrap(err, "removing the audio config")
	}

	return found, nil
}

func (c *AAC) audioConfigWithCustomDataFrom(cfgID string) (*model.AacAudioConfiguration, types.CustomData, error) {
	audCfg, err := c.api.Encoding.Configurations.Audio.Aac.Get(cfgID)
	if err != nil {
		return nil, nil, errors.Wrap(err, "retrieving configuration with config ID")
	}

	data, err := c.audioCustomDataFrom(audCfg.Id)
	if err != nil {
		return nil, nil, err
	}

	return audCfg, data, nil
}

func (c *AAC) audioCustomDataFrom(cfgID string) (types.CustomData, error) {
	data, err := c.api.Encoding.Configurations.Audio.Aac.Customdata.Get(cfgID)
	if err != nil {
		return nil, errors.Wrap(err, "retrieving custom data with config ID")
	}

	return data.CustomData, nil
}
