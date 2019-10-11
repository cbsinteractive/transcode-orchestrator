package hybrik

import "testing"

func Test_storageProviderFrom(t *testing.T) {
	tests := []struct {
		name, path   string
		wantProvider storageProvider
		wantErr      string
	}{
		{
			name:         "s3 schemes are identified correctly",
			path:         "s3://some-bucket/some-path",
			wantProvider: storageProviderS3,
		},
		{
			name:         "gcs schemes are identified correctly",
			path:         "gs://some-bucket/some-path",
			wantProvider: storageProviderGCS,
		},
		{
			name:         "http schemes are identified correctly",
			path:         "http://some-domain.com/some-path",
			wantProvider: storageProviderHTTP,
		},

		{
			name:         "https schemes are identified correctly",
			path:         "https://some-domain.com/some-path",
			wantProvider: storageProviderHTTP,
		},
		{
			name:    "unsupported schemes return a useful error",
			path:    "fakescheme://some-bucket/some-path",
			wantErr: `the scheme "fakescheme" is unsupported`,
		},
		{
			name:    "bad paths return a useful error",
			path:    "%fsdf://some-bucket/some-path",
			wantErr: `parse %fsdf://some-bucket/some-path: first path segment in URL cannot contain colon`,
		},
	}

	for _, tt := range tests {
		tt := tt
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
