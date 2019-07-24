package hybrik

import (
	"reflect"
	"testing"

	"github.com/NYTimes/video-transcoding-api/config"

	"github.com/google/go-cmp/cmp"

	"github.com/cbsinteractive/hybrik-sdk-go"

	"github.com/NYTimes/video-transcoding-api/db"
)

var (
	defaultPreset = db.Preset{
		Name:        "preset_name",
		Description: "test_desc",
		Container:   "mp4",
		RateControl: "VBR",
		TwoPass:     true,
		Video: db.VideoPreset{
			Profile:       "high",
			ProfileLevel:  "4.1",
			Width:         "300",
			Height:        "400",
			Codec:         "h264",
			Bitrate:       "400000",
			GopSize:       "120",
			InterlaceMode: "progressive",
		},
		Audio: db.AudioPreset{
			Codec:   "aac",
			Bitrate: "20000",
		},
	}
)

func TestHybrikProvider_hybrikPresetFrom(t *testing.T) {
	tests := []struct {
		name             string
		provider         *hybrikProvider
		preset           db.Preset
		wantHybrikPreset hybrik.Preset
		wantErr          bool
	}{
		{
			name: "a valid h264/aac mp4 preset results in the expected mediaconvert preset sent to the Hybrik API",
			provider: &hybrikProvider{
				config: &config.Hybrik{
					PresetPath: "some_preset_path",
				},
			},
			preset: defaultPreset,
			wantHybrikPreset: hybrik.Preset{
				Key:         defaultPreset.Name,
				Name:        defaultPreset.Name,
				Description: defaultPreset.Description,
				Kind:        "transcode",
				Path:        "some_preset_path",
				Payload: hybrik.PresetPayload{
					Targets: []hybrik.PresetTarget{
						{
							FilePattern: "",
							Container:   hybrik.TranscodeContainer{Kind: defaultPreset.Container},
							Video: hybrik.VideoTarget{
								Width:         intToPtr(300),
								Height:        intToPtr(400),
								Codec:         defaultPreset.Video.Codec,
								BitrateKb:     400,
								MaxGOPFrames:  120,
								Profile:       "high",
								Level:         "4.1",
								InterlaceMode: "progressive",
							},
							Audio: []hybrik.AudioTarget{
								{
									Codec:     defaultPreset.Audio.Codec,
									BitrateKb: 20,
								},
							},
							ExistingFiles: "replace",
							UID:           "target",
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			p, err := tt.provider.hybrikPresetFrom(tt.preset)
			if (err != nil) != tt.wantErr {
				t.Errorf("hybrikProvider.hybrikPresetFrom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if g, e := p, tt.wantHybrikPreset; !reflect.DeepEqual(g, e) {
				t.Fatalf("hybrikProvider.hybrikPresetFrom() wrong preset request\nWant %+v\nGot %+v\nDiff %s", e,
					g, cmp.Diff(e, g))
			}
		})
	}
}

func TestHybrikProvider_hybrikPresetFrom_fields(t *testing.T) {
	tests := []struct {
		name           string
		presetModifier func(preset db.Preset) db.Preset
		assertion      func(hybrik.Preset, *testing.T)
		wantErrMsg     string
	}{
		{
			name: "hevc/hdr10 presets are set correctly",
			presetModifier: func(p db.Preset) db.Preset {
				p.Video.Codec = "h265"
				p.Video.Profile = ""

				p.Video.HDR10Settings = db.HDR10Settings{
					Enabled:       true,
					MaxCLL:        10000,
					MaxFALL:       400,
					MasterDisplay: "G(13250,34500)B(7500,3000)R(34000,16000)WP(15635,16450)L(10000000,1)",
				}
				return p
			},
			assertion: func(input hybrik.Preset, t *testing.T) {
				firstTarget := input.Payload.Targets[0]

				tests := []struct {
					name      string
					got, want interface{}
				}{
					{
						name: "hdr10 master display",
						got:  firstTarget.Video.HDR10.MasterDisplay,
						want: "G(13250,34500)B(7500,3000)R(34000,16000)WP(15635,16450)L(10000000,1)",
					},
					{
						name: "hdr10 max cll",
						got:  firstTarget.Video.HDR10.MaxCLL,
						want: 10000,
					},
					{
						name: "hdr10 max fall",
						got:  firstTarget.Video.HDR10.MaxFALL,
						want: 400,
					},
					{
						name: "hdr10 color trc",
						got:  firstTarget.Video.ColorTRC,
						want: colorTRCSMPTE2084,
					},
					{
						name: "hdr10 color matrix",
						got:  firstTarget.Video.ColorMatrix,
						want: colorMatrixBT2020NC,
					},
					{
						name: "hdr10 color format",
						got:  firstTarget.Video.ChromaFormat,
						want: chromaFormatYUV420P10LE,
					},
					{
						name: "hdr10 color primaries",
						got:  firstTarget.Video.ColorPrimaries,
						want: colorPrimaryBT2020,
					},
					{
						name: "codec profile",
						got:  firstTarget.Video.Profile,
						want: "main10",
					},
					{
						name: "ffmpeg params",
						got:  firstTarget.Video.FFMPEGArgs,
						want: "-tag:v hvc1",
					},
				}

				for _, tt := range tests {
					tt := tt
					t.Run(tt.name, func(t *testing.T) {
						if g, e := tt.got, tt.want; !reflect.DeepEqual(g, e) {
							t.Fatalf("%s: got %q, expected %q", tt.name, g, e)
						}
					})
				}
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			p := &hybrikProvider{
				config: &config.Hybrik{
					PresetPath: "some_preset_path",
				},
			}
			gotPreset, err := p.hybrikPresetFrom(tt.presetModifier(defaultPreset))
			if err != nil && tt.wantErrMsg != err.Error() {
				t.Errorf("hybrikProvider.hybrikPresetFrom()error = %v, wantErr %q", err, tt.wantErrMsg)
				return
			}

			if tt.assertion != nil {
				tt.assertion(gotPreset, t)
			}
		})
	}
}

func intToPtr(i int) *int {
	return &i
}
