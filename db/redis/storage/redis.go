package storage

import (
	"errors"

	"github.com/go-redis/redis"
)

// ErrNotFound is the error returned when the given key is not found.
var ErrNotFound = errors.New("not found")

// Storage is the basic type that provides methods for saving, listing and
// deleting types on Redis.
type Storage struct {
	config *Config
	client *redis.Client
}

// Config contains configuration for the Redis, in the standard proposed by
// Gizmo.
type Config struct {
	// Comma-separated list of sentinel servers.
	//
	// Example: 10.10.10.10:6379,10.10.10.1:6379,10.10.10.2:6379.
	SentinelAddrs      string `envconfig:"SENTINEL_ADDRS"`
	SentinelMasterName string `envconfig:"SENTINEL_MASTER_NAME"`

	RedisAddr          string `envconfig:"REDIS_ADDR" default:"127.0.0.1:6379"`
	Password           string `envconfig:"REDIS_PASSWORD"`
	PoolSize           int    `envconfig:"REDIS_POOL_SIZE"`
	PoolTimeout        int    `envconfig:"REDIS_POOL_TIMEOUT_SECONDS"`
	IdleTimeout        int    `envconfig:"REDIS_IDLE_TIMEOUT_SECONDS"`
	IdleCheckFrequency int    `envconfig:"REDIS_IDLE_CHECK_FREQUENCY_SECONDS"`
}

func NewStorage(cfg *Config) (*Storage, error)             { panic("not implemented") }
func (s *Storage) Save(key string, hash interface{}) error { panic("not implemented") }
func (s *Storage) Load(key string, out interface{}) error  { panic("not implemented") }
func (s *Storage) Delete(key string) error                 { panic("not implemented") }
