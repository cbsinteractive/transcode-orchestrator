package configuration

import (
	"fmt"

	"github.com/bitmovin/bitmovin-api-sdk-go"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/video-transcoding-api/db"
	"github.com/cbsinteractive/video-transcoding-api/provider/bitmovin/internal/configuration/codec"
	"github.com/pkg/errors"
)

// AAC is a configuration service for content using only this codec
type AAC struct {
	api  *bitmovin.BitmovinApi
	repo db.PresetSummaryRepository
}

// NewAAC returns a service for managing AAC audio only configurations
func NewAAC(api *bitmovin.BitmovinApi, repo db.PresetSummaryRepository) *AAC {
	return &AAC{api: api, repo: repo}
}

// Create will create a new AAC configuration based on a preset
func (c *AAC) Create(preset db.Preset) (string, error) {
	audCfgID, err := codec.NewAAC(c.api, preset.Audio.Bitrate)
	if err != nil {
		return "", err
	}

	err = c.repo.CreatePresetSummary(&db.PresetSummary{
		Name:          preset.Name,
		Container:     preset.Container,
		AudioCodec:    string(model.CodecConfigType_AAC),
		AudioConfigID: audCfgID,
	})
	if err != nil {
		return "", err
	}

	return preset.Name, nil
}

// Get retrieves audio with a presetID
// the function will return a boolean indicating whether the audio
// configuration was found, a config object and an optional error
func (c *AAC) Get(presetName string) (db.PresetSummary, error) {
	return c.repo.GetPresetSummary(presetName)
}

// Delete removes the audio configurations
func (c *AAC) Delete(presetName string) error {
	summary, err := c.Get(presetName)
	if err != nil {
		return err
	}

	_, err = c.api.Encoding.Configurations.Audio.Aac.Delete(summary.AudioConfigID)
	if err != nil {
		return errors.Wrap(err, "removing the audio config")
	}

	err = c.repo.DeletePresetSummary(presetName)
	if err != nil {
		return fmt.Errorf("deleting preset summary: %w", err)
	}

	return nil
}
