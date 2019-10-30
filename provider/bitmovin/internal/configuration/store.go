package configuration

import (
	"github.com/cbsinteractive/video-transcoding-api/db"
)

// Store is the interface for any underlying codec config services
type Store interface {
	Create(preset db.Preset) (string, error)
	Get(presetName string) (db.PresetSummary, error)
	Delete(presetName string) error
}
