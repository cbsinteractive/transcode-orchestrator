package mediaconvert

import (
	"reflect"
	"testing"

	mc "github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/cbsinteractive/transcode-orchestrator/config"
	"github.com/cbsinteractive/transcode-orchestrator/job"
)

func TestHDRMasterDisplay(t *testing.T) {
	test := []struct {
		name        string
		encodedStr  string
		wantDisplay masterDisplay
		wantErr     bool
	}{
		{
			name:       "GBRWPL",
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
			name:       "BGLRWP",
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
			name:       "Incomplete",
			encodedStr: "B(6550,2300)G(8500,39850)L(100000000000,0)",
			wantErr:    true,
		},
		{
			name:       "Invalid",
			encodedStr: "dsfdssdf",
			wantErr:    true,
		},
	}

	for _, tt := range test {

		t.Run(tt.name, func(t *testing.T) {
			display, err := parseMasterDisplay(tt.encodedStr)
			if (err != nil) != tt.wantErr {
				t.Fatalf("parseMasterDisplay() error = %v, wantErr %v", err, tt.wantErr)
			}
			if g, e := display, tt.wantDisplay; !reflect.DeepEqual(g, e) {
				t.Fatalf("parseMasterDisplay(): wrong display returned\nWant %+v\nGot %+v", e, g)
			}
		})
	}
}

func TestHDRRequest(t *testing.T) {
	i := aws.Int64
	display := "G(8500,39850)B(6550,2300)R(35400,14600)WP(15635,16450)L(100000000000,0)"
	want := &mc.Hdr10Metadata{
		GreenPrimaryX: i(8500), GreenPrimaryY: i(39850),
		BluePrimaryX: i(6550), BluePrimaryY: i(2300),
		RedPrimaryX: i(35400), RedPrimaryY: i(14600),
		WhitePointX: i(15635), WhitePointY: i(16450),
		MinLuminance: i(0), MaxLuminance: i(100000000000),
		MaxContentLightLevel:      i(10000),
		MaxFrameAverageLightLevel: i(400),
	}

	p := defaultPreset
	p.Video.HDR10.Enabled = true
	p.Video.HDR10.MaxCLL = 10000
	p.Video.HDR10.MaxFALL = 400
	p.Video.HDR10.MasterDisplay = display

	d := &driver{cfg: config.MediaConvert{Destination: "s3://some_dest"}}
	req, err := d.createRequest(nil, &job.Job{
		Outputs: []job.TranscodeOutput{{Preset: p}},
	})
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	cc := req.Settings.OutputGroups[0].Outputs[0].VideoDescription.VideoPreprocessors.ColorCorrector
	if g, e := cc.ColorSpaceConversion, mc.ColorSpaceConversionForceHdr10; g != e {
		t.Fatalf("force hdr10: have %v, want %v", g, e)
	}

	if g, e := cc.Hdr10Metadata, want; !reflect.DeepEqual(g, e) {
		t.Fatalf("metadata: have %v, want %v", g, e)
	}
}
