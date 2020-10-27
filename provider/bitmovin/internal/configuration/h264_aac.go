package configuration

import (
	"fmt"

	"github.com/bitmovin/bitmovin-api-sdk-go"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/transcode-orchestrator/db"
	"github.com/cbsinteractive/transcode-orchestrator/provider/bitmovin/internal/configuration/codec"
	"github.com/pkg/errors"
)

// H264AAC is a configuration service for content in this codec pair
type H264AAC struct {
	api  *bitmovin.BitmovinApi
	repo db.PresetSummaryRepository
}

// NewH264AAC returns a service for managing H264 / AAC configurations
func NewH264AAC(api *bitmovin.BitmovinApi, repo db.PresetSummaryRepository) *H264AAC {
	return &H264AAC{api: api, repo: repo}
}

// Create will create a new H264AAC configuration based on a preset
func (c *H264AAC) Create(preset db.Preset) (db.PresetSummary, error) {
	audCfgID, err := codec.NewAAC(c.api, int64(preset.Audio.Bitrate))
	if err != nil {
		return db.PresetSummary{}, err
	}

	vidCfgID, err := codec.NewH264(c.api, preset)
	if err != nil {
		return db.PresetSummary{}, err
	}

	return db.PresetSummary{
		Name:          preset.Name,
		Container:     preset.Container,
		VideoCodec:    string(model.CodecConfigType_H264),
		VideoConfigID: vidCfgID,
		AudioCodec:    string(model.CodecConfigType_AAC),
		AudioConfigID: audCfgID,
	}, nil
}

// Get retrieves a stored db.PresetSummary by its name
func (c *H264AAC) Get(presetName string) (db.PresetSummary, error) {
	return c.repo.GetPresetSummary(presetName)
}

// Delete removes the audio / video configurations
func (c *H264AAC) Delete(presetName string) error {
	summary, err := c.Get(presetName)
	if err != nil {
		return err
	}

	_, err = c.api.Encoding.Configurations.Audio.Aac.Delete(summary.AudioConfigID)
	if err != nil {
		return errors.Wrap(err, "removing the audio config")
	}

	_, err = c.api.Encoding.Configurations.Video.H264.Delete(summary.VideoConfigID)
	if err != nil {
		return errors.Wrap(err, "removing the video config")
	}

	err = c.repo.DeletePresetSummary(presetName)
	if err != nil {
		return fmt.Errorf("deleting preset summary: %w", err)
	}

	return nil
}
