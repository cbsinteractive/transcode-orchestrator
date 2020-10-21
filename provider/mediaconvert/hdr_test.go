package mediaconvert

import (
	"reflect"
	"testing"
)

func Test_parseMasterDisplay(t *testing.T) {
	test := []struct {
		name        string
		encodedStr  string
		wantDisplay masterDisplay
		wantErr     bool
	}{
		{
			name:       "a well-formed encoded string is parsed correctly",
			encodedStr: "G(8500,39850)B(6550,2300)R(35400,14600)WP(15635,16450)L(100000000000,0)",
			wantDisplay: masterDisplay{
				greenPrimaryX: 8500,
				greenPrimaryY: 39850,
				bluePrimaryX:  6550,
				bluePrimaryY:  2300,
				redPrimaryX:   35400,
				redPrimaryY:   14600,
				whitePointX:   15635,
				whitePointY:   16450,
				maxLuminance:  100000000000,
				minLuminance:  0,
			},
		},
		{
			name:       "values are parsed correctly regardles of position in the encoded string",
			encodedStr: "B(6550,2300)G(8500,39850)L(100000000000,0)R(35400,14600)WP(15635,16450)",
			wantDisplay: masterDisplay{
				greenPrimaryX: 8500,
				greenPrimaryY: 39850,
				bluePrimaryX:  6550,
				bluePrimaryY:  2300,
				redPrimaryX:   35400,
				redPrimaryY:   14600,
				whitePointX:   15635,
				whitePointY:   16450,
				maxLuminance:  100000000000,
				minLuminance:  0,
			},
		},
		{
			name:       "incomplete values return an error",
			encodedStr: "B(6550,2300)G(8500,39850)L(100000000000,0)",
			wantErr:    true,
		},
		{
			name:       "invalid values return an error",
			encodedStr: "dsfdssdf",
			wantErr:    true,
		},
	}

	for _, tt := range test {

		t.Run(tt.name, func(t *testing.T) {
			display, err := parseMasterDisplay(tt.encodedStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseMasterDisplay() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if g, e := display, tt.wantDisplay; !reflect.DeepEqual(g, e) {
				t.Fatalf("parseMasterDisplay(): wrong display returned\nWant %+v\nGot %+v", e, g)
			}
		})
	}
}
