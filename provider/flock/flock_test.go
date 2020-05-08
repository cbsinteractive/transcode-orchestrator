package flock

import (
	"context"
	"net/http"
	"testing"

	"github.com/cbsinteractive/transcode-orchestrator/config"
)

func TestFlock_CancelJob(t *testing.T) {
	tests := []struct {
		name         string
		cfg          config.Flock
		providerID   string
		reqAssertion func(*testing.T, *http.Request)
		response     http.Response
		expectErr    string
	}{
		{
			name:       "the correct url is requested to flock on cancel",
			providerID: "some-id",
			cfg:        config.Flock{Endpoint: "http://flock.com"},
			reqAssertion: func(t *testing.T, r *http.Request) {
				wantURL := "http://flock.com/api/v1/jobs/some-id"
				if g, e := r.URL.String(), wantURL; g != e {
					t.Errorf("CancelJob() wrong url requested, got %q, expected %q", g, e)
				}
			},
		},
		{
			name: "the credential is added to the request to flock on cancel",
			cfg:  config.Flock{Credential: "some-credential"},
			reqAssertion: func(t *testing.T, r *http.Request) {
				wantCredential := "some-credential"
				if g, e := r.Header.Get("Authorization"), wantCredential; g != e {
					t.Errorf("CancelJob() wrong credential sent, got %q, expected %q", g, e)
				}
			},
		},
		{
			name:      "if the flock endpoint is malformed, a useful error is returned",
			cfg:       config.Flock{Endpoint: ":::"},
			expectErr: `parse ":::/api/v1/jobs/": missing protocol scheme`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockTransport := &mockRoundTripper{returnsResp: tt.response}
			client := &http.Client{Transport: mockTransport}

			err := (&flock{cfg: &tt.cfg, client: client}).CancelJob(context.Background(), tt.providerID)
			if err != nil {
				if g, e := err.Error(), tt.expectErr; g != e {
					t.Errorf("CancelJob() wrong error returned, got: %v, want: %v", g, e)
				}
			} else if tt.expectErr != "" {
				t.Error("CancelJob() expected an error, got nil")
			}

			if tt.reqAssertion != nil {
				tt.reqAssertion(t, mockTransport.calledWithReq)
			}
		})
	}
}

type mockRoundTripper struct {
	calledWithReq *http.Request
	returnsResp   http.Response
	returnsErr    error
}

func (rt *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	rt.calledWithReq = req
	return &rt.returnsResp, rt.returnsErr
}
