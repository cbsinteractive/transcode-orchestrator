package hybrik

import (
	"testing"

	"github.com/cbsinteractive/transcode-orchestrator/job"
)

func TestStorage(t *testing.T) {
	tests := []struct {
		name, path, want string
	}{
		{"s3", "s3://some-bucket/some-path", "s3"},
		{"gcs", "gs://some-bucket/some-path", "gcs"},
		{"http", "http://some-domain.com/some-path", "http"},
		{"https", "https://some-domain.com/some-path", "https"},
		// {"unsupported", "fakescheme://some-bucket/some-path", "", `the scheme "fakescheme" is unsupported`},
		{"bad", "%fsdf://some-bucket/some-path", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := job.File{Name: tt.path}.Provider()
			if h != tt.want {
				t.Fatalf("wrong provider: have %q, want %q", h, tt.want)
			}
		})
	}
}

func assertWantErr(err error, wantErr, caller string, t *testing.T) bool {
	if err != nil {
		if wantErr != err.Error() {
			t.Errorf("%s error = %v, wantErr %q", caller, err, wantErr)
		}

		return true
	} else if wantErr != "" {
		t.Errorf("%s expected error %q, did not receive an error", caller, wantErr)
		return true
	}

	return false
}
