package configuration

import (
	"fmt"

	"github.com/bitmovin/bitmovin-api-sdk-go"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/transcode-orchestrator/db"
	"github.com/cbsinteractive/transcode-orchestrator/provider/bitmovin/internal/configuration/codec"
	"github.com/pkg/errors"
)

// Opus is a configuration service for content using only this codec
type Opus struct {
	api  *bitmovin.BitmovinApi
	repo db.PresetSummaryRepository
}

// NewOpus returns a service for managing Opus audio only configurations
func NewOpus(api *bitmovin.BitmovinApi, repo db.PresetSummaryRepository) *Opus {
	return &Opus{api: api, repo: repo}
}

// Create will create a new Opus configuration based on a preset
func (c *Opus) Create(preset db.Preset) (db.PresetSummary, error) {
	audCfgID, err := codec.NewOpus(c.api, preset.Audio.Bitrate)
	if err != nil {
		return db.PresetSummary{}, err
	}

	return db.PresetSummary{
		Name:          preset.Name,
		Container:     preset.Container,
		AudioCodec:    string(model.CodecConfigType_OPUS),
		AudioConfigID: audCfgID,
	}, nil
}

// Get retrieves a stored db.PresetSummary by its name
func (c *Opus) Get(presetName string) (db.PresetSummary, error) {
	return c.repo.GetPresetSummary(presetName)
}

// Delete removes the audio configurations
func (c *Opus) Delete(presetName string) error {
	summary, err := c.Get(presetName)
	if err != nil {
		return err
	}

	_, err = c.api.Encoding.Configurations.Audio.Opus.Delete(summary.AudioConfigID)
	if err != nil {
		return errors.Wrap(err, "removing the audio config")
	}

	err = c.repo.DeletePresetSummary(presetName)
	if err != nil {
		return fmt.Errorf("deleting preset summary: %w", err)
	}

	return nil
}
