package configuration

import (
	"github.com/cbsinteractive/transcode-orchestrator/db"
)

// Store is the interface for any underlying codec config services
type Store interface {
	Create(preset db.Preset) (db.PresetSummary, error)
	Get(presetName string) (db.PresetSummary, error)
	Delete(presetName string) error
}
