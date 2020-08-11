package configuration

import (
	"fmt"

	"github.com/bitmovin/bitmovin-api-sdk-go"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/transcode-orchestrator/db"
	"github.com/cbsinteractive/transcode-orchestrator/provider/bitmovin/internal/configuration/codec"
	"github.com/pkg/errors"
)

// AV1 is a configuration service for content in av1
type AV1 struct {
	api  *bitmovin.BitmovinApi
	repo db.PresetSummaryRepository
}

// NewAV1 returns a service for managing av1 configurations
func NewAV1(api *bitmovin.BitmovinApi, repo db.PresetSummaryRepository) *AV1 {
	return &AV1{api: api, repo: repo}
}

// Create will create a new AV1 configuration based on a preset
func (c *AV1) Create(preset db.Preset) (db.PresetSummary, error) {
	vidCfgID, err := codec.NewAV1(c.api, preset)
	if err != nil {
		return db.PresetSummary{}, err
	}

	return db.PresetSummary{
		Name:          preset.Name,
		Container:     preset.Container,
		VideoCodec:    string(model.CodecConfigType_AV1),
		VideoConfigID: vidCfgID,
	}, nil
}

// Get retrieves a stored db.PresetSummary by its name
func (c *AV1) Get(presetName string) (db.PresetSummary, error) {
	return c.repo.GetPresetSummary(presetName)
}

// Delete removes the video configuration
func (c *AV1) Delete(presetName string) error {
	summary, err := c.Get(presetName)
	if err != nil {
		return err
	}

	_, err = c.api.Encoding.Configurations.Video.Av1.Delete(summary.VideoConfigID)
	if err != nil {
		return errors.Wrap(err, "removing the video config")
	}

	err = c.repo.DeletePresetSummary(presetName)
	if err != nil {
		return fmt.Errorf("deleting preset summary: %w", err)
	}

	return nil
}
