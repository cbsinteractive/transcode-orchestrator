package configuration

import (
	"fmt"

	"github.com/bitmovin/bitmovin-api-sdk-go"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/transcode-orchestrator/db"
	"github.com/cbsinteractive/transcode-orchestrator/provider/bitmovin/internal/configuration/codec"
	"github.com/pkg/errors"
)

// H265AAC is a configuration service for content in this codec pair
type H265AAC struct {
	api  *bitmovin.BitmovinApi
	repo db.PresetSummaryRepository
}

// NewH265AAC returns a service for managing h.265 / AAC configurations
func NewH265AAC(api *bitmovin.BitmovinApi, repo db.PresetSummaryRepository) *H265AAC {
	return &H265AAC{api: api, repo: repo}
}

// Create will create a new H265AAC configuration based on a preset
func (c *H265AAC) Create(preset db.Preset) (string, error) {
	audCfgID, err := codec.NewAAC(c.api, preset.Audio.Bitrate)
	if err != nil {
		return "", err
	}

	vidCfgID, err := codec.NewH265(c.api, preset)
	if err != nil {
		return "", err
	}

	err = c.repo.CreatePresetSummary(&db.PresetSummary{
		Name:          preset.Name,
		Container:     preset.Container,
		VideoCodec:    string(model.CodecConfigType_H265),
		VideoConfigID: vidCfgID,
		AudioCodec:    string(model.CodecConfigType_AAC),
		AudioConfigID: audCfgID,
	})
	if err != nil {
		return "", err
	}

	return preset.Name, nil
}

// Get retrieves a stored db.PresetSummary by its name
func (c *H265AAC) Get(presetName string) (db.PresetSummary, error) {
	return c.repo.GetPresetSummary(presetName)
}

// Delete removes the audio / video configurations
func (c *H265AAC) Delete(presetName string) error {
	summary, err := c.Get(presetName)
	if err != nil {
		return err
	}

	_, err = c.api.Encoding.Configurations.Audio.Aac.Delete(summary.AudioConfigID)
	if err != nil {
		return errors.Wrap(err, "removing the audio config")
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
