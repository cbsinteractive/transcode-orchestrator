package provider

import (
	"context"

	"github.com/cbsinteractive/transcode-orchestrator/config"
	"github.com/cbsinteractive/transcode-orchestrator/job"
)

type fake struct {
	cap    Capabilities
	health error
}

func (fake) Create(context.Context, *job.Job) (*job.Status, error) { return nil, nil }
func (fake) Status(context.Context, *job.Job) (*job.Status, error) { return nil, nil }
func (fake) Cancel(context.Context, string) error                  { return nil }
func (f fake) Healthcheck() error                                  { return f.health }
func (f fake) Capabilities() Capabilities                          { return f.cap }

func getFactory(err error, health error, cap Capabilities) Factory {
	return func(*config.Config) (Provider, error) {
		if err != nil {
			return nil, err
		}
		return &fake{health: health, cap: cap}, nil
	}
}
