package hybrik

import "testing"

func TestStorage(t *testing.T) {
	tests := []struct {
		name, path   string
		wantProvider storageProvider
		wantErr      string
	}{
		{"s3", "s3://some-bucket/some-path", storageProviderS3, ""},
		{"gcs", "gs://some-bucket/some-path", storageProviderGCS, ""},
		{"http", "http://some-domain.com/some-path", storageProviderHTTP, ""},
		{"https", "https://some-domain.com/some-path", storageProviderHTTP, ""},
		{"unsupported", "fakescheme://some-bucket/some-path", "", `the scheme "fakescheme" is unsupported`},
		{"bad", "%fsdf://some-bucket/some-path", "", `parse "%fsdf://some-bucket/some-path": first path segment in URL cannot contain colon`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prov, err := storageProviderFrom(tt.path)
			if shouldReturn := assertWantErr(err, tt.wantErr, "storageProviderFrom()", t); shouldReturn {
				return
			}

			if g, e := prov, tt.wantProvider; g != e {
				t.Fatalf("storageProviderFrom() wrong provider: got %q, expected %q", g, e)
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
