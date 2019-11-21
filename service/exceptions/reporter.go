package exceptions

import (
	"time"

	"github.com/getsentry/sentry-go"
)

const defaultFlushTimeout = time.Second * 5

// Reporter sends exceptions to an external source
type Reporter interface {
	ReportException(err error)
}

// NoopReporter is a no-op exception reporter
type NoopReporter struct{}

// ReportException does nothing
func (r *NoopReporter) ReportException(_ error) {}

// SentryReporter is an ErrorReporter that sends error information to Sentry
type SentryReporter struct{}

// NewSentryReporter creates and returns an instance of SentryReporter
func NewSentryReporter(dsn, env string) (*SentryReporter, error) {
	err := sentry.Init(sentry.ClientOptions{Dsn: dsn, Environment: env})
	if err != nil {
		return nil, err
	}

	return &SentryReporter{}, nil
}

// ReportException will send errors to Sentry
func (r *SentryReporter) ReportException(err error) {
	sentry.CaptureException(err)
	sentry.Flush(defaultFlushTimeout)
}
