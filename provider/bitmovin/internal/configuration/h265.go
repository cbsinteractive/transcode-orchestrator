package configuration

import (
	"fmt"

	"github.com/bitmovin/bitmovin-api-sdk-go"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/transcode-orchestrator/db"
	"github.com/cbsinteractive/transcode-orchestrator/provider/bitmovin/internal/configuration/codec"
	"github.com/pkg/errors"
)

// H265 is a configuration service for content in h.265
type H265 struct {
	api  *bitmovin.BitmovinApi
	repo db.PresetSummaryRepository
}

// NewH265 returns a service for managing h.265 configurations
func NewH265(api *bitmovin.BitmovinApi, repo db.PresetSummaryRepository) *H265 {
	return &H265{api: api, repo: repo}
}

// Create will create a new H265 configuration based on a preset
func (c *H265) Create(preset db.Preset) (string, error) {
	vidCfgID, err := codec.NewH265(c.api, preset)
	if err != nil {
		return "", err
	}

	err = c.repo.CreatePresetSummary(&db.PresetSummary{
		Name:          preset.Name,
		Container:     preset.Container,
		VideoCodec:    string(model.CodecConfigType_H265),
		VideoConfigID: vidCfgID,
	})
	if err != nil {
		return "", err
	}

	return preset.Name, nil
}

// Get retrieves a stored db.PresetSummary by its name
func (c *H265) Get(presetName string) (db.PresetSummary, error) {
	return c.repo.GetPresetSummary(presetName)
}

// Delete removes the video configuration
func (c *H265) Delete(presetName string) error {
	summary, err := c.Get(presetName)
	if err != nil {
		return err
	}

	_, err = c.api.Encoding.Configurations.Video.H265.Delete(summary.VideoConfigID)
	if err != nil {
		return errors.Wrap(err, "removing the video config")
	}

	err = c.repo.DeletePresetSummary(presetName)
	if err != nil {
		return fmt.Errorf("deleting preset summary: %w", err)
	}

	return nil
}
