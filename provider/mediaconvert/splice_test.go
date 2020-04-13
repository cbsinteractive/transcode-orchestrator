package mediaconvert

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/cbsinteractive/pkg/timecode"
)

func TestSplice2Clippings(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input timecode.Splice
		want  []mediaconvert.InputClipping
	}{
		{"5-10s", timecode.Splice{{5, 10}}, makeIC("00:00:05:00", "00:00:10:00")},
	} {
		t.Run(tc.name, func(t *testing.T) {
			have := splice2clippings(tc.input, 0)
			want := tc.want
			if !reflect.DeepEqual(have, want) {
				t.Fatalf("have %v, want %v", have, want)
			}
		})
	}
}

func makeIC(s, e string) []mediaconvert.InputClipping {
	return []mediaconvert.InputClipping{{
		StartTimecode: &s,
		EndTimecode:   &e,
	}}
}
