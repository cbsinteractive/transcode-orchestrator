package mediaconvert

import "testing"

func Test_vbrLevel(t *testing.T) {
	tests := []struct {
		name      string
		bitrate   int64
		wantLevel int64
	}{
		{
			name:      "NotSet",
			bitrate:   0,
			wantLevel: 4,
		},
		{
			name:      "45Kbps",
			bitrate:   45000,
			wantLevel: -1,
		},
		{
			name:      "128Kbps",
			bitrate:   128000,
			wantLevel: 4,
		},
		{
			name:      "196Kbps",
			bitrate:   196000,
			wantLevel: 6,
		},
		{
			name:      "256Kbps",
			bitrate:   256000,
			wantLevel: 8,
		},
		{
			name:      "500Kbps",
			bitrate:   500000,
			wantLevel: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := vbrLevel(tt.bitrate); got != tt.wantLevel {
				t.Errorf("vbrLevel() = %v, want %v", got, tt.wantLevel)
			}
		})
	}
}
