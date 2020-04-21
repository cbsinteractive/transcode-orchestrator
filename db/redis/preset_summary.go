package redis

import (
	"errors"

	"github.com/cbsinteractive/transcode-orchestrator/db"
	"github.com/cbsinteractive/transcode-orchestrator/db/redis/storage"
	"github.com/go-redis/redis"
)

const presetSummarySetKey = "presetsummaries"

func (r *redisRepository) CreatePresetSummary(summary *db.PresetSummary) error {
	if _, err := r.GetPresetSummary(summary.Name); err == nil {
		return db.ErrPresetSummaryAlreadyExists
	}

	return r.savePresetSummary(summary)
}

func (r *redisRepository) savePresetSummary(summary *db.PresetSummary) error {
	fields, err := r.storage.FieldMap(summary)
	if err != nil {
		return err
	}

	if summary.Name == "" {
		return errors.New("preset summary name missing")
	}

	presetSummaryKey := r.presetSummaryKey(summary.Name)

	return r.storage.RedisClient().Watch(func(tx *redis.Tx) error {
		err := tx.HMSet(presetSummaryKey, fields).Err()
		if err != nil {
			return err
		}
		return tx.SAdd(presetSummarySetKey, summary.Name).Err()
	}, presetSummaryKey)
}

func (r *redisRepository) GetPresetSummary(name string) (db.PresetSummary, error) {
	var summary db.PresetSummary

	err := r.storage.Load(r.presetSummaryKey(name), &summary)
	if err == storage.ErrNotFound {
		return db.PresetSummary{}, db.ErrPresetSummaryNotFound
	}

	return summary, err
}

func (r *redisRepository) DeletePresetSummary(name string) error {
	err := r.storage.Delete(r.presetSummaryKey(name))
	if err != nil {
		if err == storage.ErrNotFound {
			return db.ErrPresetSummaryNotFound
		}

		return err
	}

	r.storage.RedisClient().SRem(presetSummarySetKey, name)

	return nil
}

func (r *redisRepository) presetSummaryKey(name string) string {
	return "presetsummary:" + name
}
