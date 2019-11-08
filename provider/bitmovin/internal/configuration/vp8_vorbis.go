package configuration

import (
	"fmt"

	"github.com/bitmovin/bitmovin-api-sdk-go"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/video-transcoding-api/db"
	"github.com/cbsinteractive/video-transcoding-api/provider/bitmovin/internal/configuration/codec"
	"github.com/pkg/errors"
)

// VP8Vorbis is a configuration service for content in this codec pair
type VP8Vorbis struct {
	api  *bitmovin.BitmovinApi
	repo db.PresetSummaryRepository
}

// NewVP8Vorbis returns a service for managing VP8 / Vorbis configurations
func NewVP8Vorbis(api *bitmovin.BitmovinApi, repo db.PresetSummaryRepository) *VP8Vorbis {
	return &VP8Vorbis{api: api, repo: repo}
}

// Create will create a new VP8 configuration based on a preset
func (c *VP8Vorbis) Create(preset db.Preset) (string, error) {
	audCfgID, err := codec.NewVorbis(c.api, preset.Audio.Bitrate)
	if err != nil {
		return "", err
	}

	vidCfgID, err := codec.NewVP8(c.api, preset)
	if err != nil {
		return "", err
	}

	err = c.repo.CreatePresetSummary(&db.PresetSummary{
		Name:          preset.Name,
		Container:     preset.Container,
		VideoCodec:    string(model.CodecConfigType_VP8),
		VideoConfigID: vidCfgID,
		AudioCodec:    string(model.CodecConfigType_VORBIS),
		AudioConfigID: audCfgID,
	})
	if err != nil {
		return "", err
	}

	return preset.Name, nil
}

// Get retrieves a stored db.PresetSummary by its name
func (c *VP8Vorbis) Get(presetName string) (db.PresetSummary, error) {
	return c.repo.GetPresetSummary(presetName)
}

// Delete removes the audio / video configurations
func (c *VP8Vorbis) Delete(presetName string) error {
	summary, err := c.Get(presetName)
	if err != nil {
		return err
	}

	_, err = c.api.Encoding.Configurations.Audio.Vorbis.Delete(summary.AudioConfigID)
	if err != nil {
		return errors.Wrap(err, "removing the audio config")
	}

	_, err = c.api.Encoding.Configurations.Video.Vp8.Delete(summary.VideoConfigID)
	if err != nil {
		return errors.Wrap(err, "removing the video config")
	}

	err = c.repo.DeletePresetSummary(presetName)
	if err != nil {
		return fmt.Errorf("deleting preset summary: %w", err)
	}

	return nil
}
