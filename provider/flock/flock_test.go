package flock

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
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
			name: "if the backend returns a non 2xx status code, a useful error is returned",
			response: http.Response{
				StatusCode: 500,
				Body:       ioutil.NopCloser(strings.NewReader("oops something went wrong")),
			},
			expectErr: "received non 2xx status code, got 500 with body: oops something went wrong",
		},
		{
			name: "if we fail to read the resp body, a useful error is returned",
			response: http.Response{
				Body: ioutil.NopCloser(errReader{}),
			},
			expectErr: "reading resp body: error forced by mock reader",
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

			provider := &flock{
				cfg:    &tt.cfg,
				client: &http.Client{Transport: mockTransport},
			}

			err := provider.CancelJob(context.Background(), tt.providerID)
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

	if rt.returnsResp.Body == nil {
		rt.returnsResp.Body = ioutil.NopCloser(strings.NewReader(""))
	}

	return &rt.returnsResp, rt.returnsErr
}

type errReader struct{}

func (errReader) Read([]byte) (int, error) {
	return 0, errors.New("error forced by mock reader")
}
