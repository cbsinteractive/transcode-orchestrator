package tracing

import (
	"context"
	"net/http"
)

type nopCloser struct{}

func (nopCloser) Close(error) {}

type NoopTracer struct{}

func (NoopTracer) Init() error                        { return nil }
func (NoopTracer) Client(c *http.Client) *http.Client { return c }
func (NoopTracer) BeginSubsegment(ctx context.Context, name string) interface{ Close(error) } {
	return nopCloser{}
}
func (NoopTracer) Handle(_ interface{ Name(host string) string }, h http.Handler) http.Handler {
	return h
}
