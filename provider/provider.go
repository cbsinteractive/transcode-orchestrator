package provider

import (
	"context"
	"errors"
	"fmt"
	"sort"

	"github.com/cbsinteractive/transcode-orchestrator/config"
	"github.com/cbsinteractive/transcode-orchestrator/job"
)

var providers = map[string]Factory{}

var (
	ErrRegistered = errors.New("provider is already registered")
	ErrNotFound   = errors.New("provider not found")
	ErrConfig     = errors.New("bad provider configuration")
	ErrPreset     = errors.New("preset not found in provider")
)

// Provider knows how to manage jobs for media transcoding
type Provider interface {
	Create(context.Context, *job.Job) (*job.Status, error)
	Status(context.Context, *job.Job) (*job.Status, error)
	Cancel(ctx context.Context, id string) error
	Healthcheck() error
	Capabilities() Capabilities
}

// Factory is the function responsible for creating the instance of a
// provider.
type Factory func(cfg *config.Config) (Provider, error)

// InvalidConfigError is returned if a provider could not be configured properly
type InvalidConfigError string

// JobNotFoundError is returned if a job with a given id could not be found by the provider
type JobNotFoundError struct {
	ID string
}

func (err InvalidConfigError) Error() string {
	return string(err)
}

func (err JobNotFoundError) Error() string {
	return fmt.Sprintf("could not found job with id: %s", err.ID)
}

// Register register a new provider in the internal list of providers.
func Register(name string, provider Factory) error {
	if _, ok := providers[name]; ok {
		return ErrRegistered
	}
	providers[name] = provider
	return nil
}

// GetProviderFactory looks up the list of registered providers and returns the
// factory function for the given provider name, if it's available.
func GetFactory(name string) (Factory, error) {
	factory, ok := providers[name]
	if !ok {
		return nil, ErrNotFound
	}
	return factory, nil
}

// List returns the list of currently registered providers,
// alphabetically ordered.
func List(c *config.Config) []string {
	providerNames := make([]string, 0, len(providers))
	for name, factory := range providers {
		if _, err := factory(c); err == nil {
			providerNames = append(providerNames, name)
		}
	}
	sort.Strings(providerNames)
	return providerNames
}

// Describe describes the given provider. It includes information about
// the provider's capabilities and its current health state.
func Describe(name string, c *config.Config) (*Description, error) {
	factory, err := GetFactory(name)
	if err != nil {
		return nil, err
	}
	description := Description{Name: name}
	provider, err := factory(c)
	if err != nil {
		return &description, nil
	}
	description.Enabled = true
	description.Capabilities = provider.Capabilities()
	description.Health = Health{OK: true}
	if err = provider.Healthcheck(); err != nil {
		description.Health = Health{OK: false, Message: err.Error()}
	}
	return &description, nil
}
