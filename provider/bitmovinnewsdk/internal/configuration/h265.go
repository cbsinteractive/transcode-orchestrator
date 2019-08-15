package configuration

import (
	"github.com/bitmovin/bitmovin-api-sdk-go"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/video-transcoding-api/db"
	"github.com/cbsinteractive/video-transcoding-api/provider/bitmovinnewsdk/internal/configuration/codec"
	"github.com/cbsinteractive/video-transcoding-api/provider/bitmovinnewsdk/internal/types"
	"github.com/pkg/errors"
)

// H265 is a configuration service for content in h.265
type H265 struct {
	api *bitmovin.BitmovinApi
}

// NewH265 returns a service for managing h.265 configurations
func NewH265(api *bitmovin.BitmovinApi) *H265 {
	return &H265{api: api}
}

// Create will create a new H265 configuration based on a preset
func (c *H265) Create(preset db.Preset) (string, error) {
	vidCfgID, err := codec.NewH265(c.api, preset, customDataWith("", preset.Container))
	if err != nil {
		return "", err
	}

	return vidCfgID, nil
}

// Get retrieves video configuration with a presetID
// the function will return a boolean indicating whether the video
// configuration was found, a config object and an optional error
func (c *H265) Get(cfgID string) (bool, Details, error) {
	vidCfg, customData, err := c.vidConfigWithCustomDataFrom(cfgID)
	if err != nil {
		return false, Details{}, err
	}

	return true, Details{Video: vidCfg, CustomData: customData}, nil
}

// Delete removes the video configuration
func (c *H265) Delete(cfgID string) (found bool, e error) {
	_, err := c.api.Encoding.Configurations.Video.H265.Delete(cfgID)
	if err != nil {
		return found, errors.Wrap(err, "removing the video config")
	}

	return found, nil
}

func (c *H265) vidConfigWithCustomDataFrom(cfgID string) (*model.H265VideoConfiguration, types.CustomData, error) {
	vidCfg, err := c.api.Encoding.Configurations.Video.H265.Get(cfgID)
	if err != nil {
		return nil, nil, errors.Wrap(err, "retrieving configuration with config ID")
	}

	data, err := c.vidCustomDataFrom(vidCfg.Id)
	if err != nil {
		return nil, nil, err
	}

	return vidCfg, data, nil
}

func (c *H265) vidCustomDataFrom(cfgID string) (types.CustomData, error) {
	data, err := c.api.Encoding.Configurations.Video.H265.Customdata.Get(cfgID)
	if err != nil {
		return nil, errors.Wrap(err, "retrieving custom data with config ID")
	}

	return data.CustomData, nil
}
