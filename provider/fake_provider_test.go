package provider

import (
	"context"

	"github.com/cbsinteractive/transcode-orchestrator/config"
	"github.com/cbsinteractive/transcode-orchestrator/db"
)

type fakeProvider struct {
	cap       Capabilities
	healthErr error
}

func (*fakeProvider) Transcode(context.Context, *db.Job) (*JobStatus, error) {
	return nil, nil
}

func (*fakeProvider) JobStatus(context.Context, *db.Job) (*JobStatus, error) {
	return nil, nil
}

func (*fakeProvider) CreatePreset(context.Context, db.Preset) (string, error) {
	return "", nil
}

func (*fakeProvider) GetPreset(context.Context, string) (interface{}, error) {
	return "", nil
}

func (*fakeProvider) DeletePreset(context.Context, string) error {
	return nil
}

func (*fakeProvider) CancelJob(context.Context, string) error {
	return nil
}

func (f *fakeProvider) Healthcheck() error {
	return f.healthErr
}

func (f *fakeProvider) Capabilities() Capabilities {
	return f.cap
}

func getFactory(fErr error, healthErr error, capabilities Capabilities) Factory {
	return func(*config.Config) (TranscodingProvider, error) {
		if fErr != nil {
			return nil, fErr
		}
		return &fakeProvider{healthErr: healthErr, cap: capabilities}, nil
	}
}
